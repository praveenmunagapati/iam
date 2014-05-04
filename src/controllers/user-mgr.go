package controllers

import (
    "../../deps/lessgo/data/rdc"
    "../../deps/lessgo/pagelet"
    "../../deps/lessgo/pass"
    "../../deps/lessgo/utils"
    "../reg/signup"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

var (
    userMgrPasswdHidden = "************"
)

type RoleEntry struct {
    Rid, Name, Checked string
}

type UserMgr struct {
    *pagelet.Controller
}

func (c UserMgr) IndexAction() {

    if !c.Session.AccessAllowed("user.admin") {
        c.RenderError(401, "Access Denied")
        return
    }
}

func (c UserMgr) ListAction() {

    if !c.Session.AccessAllowed("user.admin") {
        c.RenderError(200, "Access Denied")
        return
    }

    dcn, err := rdc.InstancePull("def")
    if err != nil {
        return
    }

    rdict := map[string]string{}
    q := rdc.NewQuerySet().From("ids_role").Limit(100)
    rsr, err := dcn.Query(q)
    if err == nil && len(rsr) > 0 {
        //c.ViewData["roles"] = rsr
        for _, v := range rsr {
            rdict[fmt.Sprintf("%v", v["rid"])] = v["name"].(string)
        }
    }

    // filter: query_text
    q = rdc.NewQuerySet().From("ids_login").Limit(20)
    if query_text := c.Params.Get("query_text"); query_text != "" {
        q.Where.And("name.like", "%"+query_text+"%").
            Or("uname.like", "%"+query_text+"%").
            Or("email.like", "%"+query_text+"%")
        c.ViewData["query_text"] = query_text
    }

    rsl, err := dcn.Query(q)
    if err == nil && len(rsl) > 0 {

        for k, v := range rsl {

            rids := strings.Split(v["roles"].(string), ",")
            for rk, rv := range rids {

                rname, ok := rdict[rv]
                if !ok {
                    continue
                }

                rids[rk] = rname
            }

            rsl[k]["roles_display"] = rids
        }

        c.ViewData["list"] = rsl
    }

    c.ViewData["query_role"] = c.Params.Get("query_role")
}

func (c UserMgr) EditAction() {

    if !c.Session.AccessAllowed("user.admin") {
        c.RenderError(200, "Access Denied")
        return
    }

    dcn, err := rdc.InstancePull("def")
    if err != nil {
        c.RenderError(500, http.StatusText(500))
        return
    }

    //
    roles := []RoleEntry{}
    q := rdc.NewQuerySet().From("ids_role").Limit(100)
    q.Where.And("status", 1)
    rsr, err := dcn.Query(q)
    if err == nil && len(rsr) > 0 {

        for _, v := range rsr {
            roles = append(roles, RoleEntry{
                fmt.Sprintf("%v", v["rid"]),
                fmt.Sprintf("%v", v["name"]),
                ""})
        }
    }

    if c.Params.Get("uid") != "" {

        q := rdc.NewQuerySet().From("ids_login").Limit(1)
        q.Where.And("uid", c.Params.Get("uid"))
        rslogin, err := dcn.Query(q)
        if err != nil || len(rslogin) == 0 {
            c.RenderError(400, http.StatusText(400))
            return
        }

        rls := strings.Split(rslogin[0]["roles"].(string), ",")
        for _, v := range rls {
            for k2, v2 := range roles {
                if v2.Rid == v {
                    roles[k2].Checked = "1"
                    break
                }
            }
        }

        c.ViewData["uid"] = c.Params.Get("uid")
        c.ViewData["uname"] = rslogin[0]["uname"]
        c.ViewData["email"] = rslogin[0]["email"]
        c.ViewData["passwd"] = userMgrPasswdHidden
        c.ViewData["name"] = rslogin[0]["name"]

        q.From("ids_profile")
        rsprofile, err := dcn.Query(q)
        if err == nil && len(rsprofile) == 1 {
            c.ViewData["birthday"] = rsprofile[0]["birthday"]
            c.ViewData["aboutme"] = rsprofile[0]["aboutme"]
        }

        c.ViewData["panel_title"] = "Edit Account"
        c.ViewData["uid"] = c.Params.Get("uid")
    } else {

        c.ViewData["panel_title"] = "New Account"
        c.ViewData["uid"] = ""
    }

    c.ViewData["roles"] = roles
}

func (c UserMgr) SaveAction() {

    c.AutoRender = false

    var rsp ResponseJson
    rsp.ApiVersion = apiVersion
    rsp.Status = 400
    rsp.Message = "Bad Request"

    defer func() {
        if rspj, err := utils.JsonEncode(rsp); err == nil {
            io.WriteString(c.Response.Out, rspj)
        }
    }()

    if err := signup.Validate(c.Params); err != nil {
        rsp.Message = err.Error()
        return
    }

    dcn, err := rdc.InstancePull("def")
    if err != nil {
        rsp.Message = "Internal Server Error"
        return
    }

    q := rdc.NewQuerySet().From("ids_login").Limit(1)

    isNew := true
    loginset := map[string]interface{}{}

    if c.Params.Get("uid") != "" {

        q.Where.And("uid", c.Params.Get("uid"))

        rslogin, err := dcn.Query(q)
        if err != nil || len(rslogin) == 0 {
            c.RenderError(400, http.StatusText(400))
            return
        }

        isNew = false
    }

    //
    q = rdc.NewQuerySet().From("ids_login").Limit(1)
    q.Where.And("email", c.Params.Get("email"))
    rsu, err := dcn.Query(q)
    if err == nil && len(rsu) == 1 {

        if isNew || fmt.Sprintf("%v", rsu[0]["uid"]) != c.Params.Get("uid") {
            rsp.Message = "The `Email` already exists, please choose another one"
            return
        }

    } else {
        loginset["email"] = c.Params.Get("email")
    }

    //
    q = rdc.NewQuerySet().From("ids_login").Limit(1)
    q.Where.And("uname", c.Params.Get("uname"))
    rsu, err = dcn.Query(q)
    if err == nil && len(rsu) == 1 {

        if isNew || fmt.Sprintf("%v", rsu[0]["uid"]) != c.Params.Get("uid") {
            rsp.Message = "The `Username` already exists, please choose another one"
            return
        }

    } else {
        loginset["uname"] = c.Params.Get("uname")
    }

    if c.Params.Get("passwd") != userMgrPasswdHidden {

        pass, err := pass.HashDefault(c.Params.Get("passwd"))
        if err != nil {
            return
        }
        loginset["pass"] = pass
    }

    if isNew {
        loginset["created"] = rdc.TimeNow("datetime")
        loginset["status"] = 1
        loginset["timezone"] = "UTC"
    }
    loginset["updated"] = rdc.TimeNow("datetime")
    loginset["name"] = c.Params.Get("name")
    loginset["roles"] = strings.Join(c.Params.Values["roles"], ",")

    frupd := rdc.NewFilter()

    if isNew {
        rst, err := dcn.Insert("ids_login", loginset)
        if err != nil {
            rsp.Status = 500
            rsp.Message = "Can not write to database"
            return
        }

        lastid, err := rst.LastInsertId()
        if err != nil || lastid == 0 {
            rsp.Status = 500
            rsp.Message = "Can not write to database"
            return
        }

        c.Params.Set("uid", fmt.Sprintf("%v", lastid))

    } else {

        frupd.And("uid", c.Params.Get("uid"))
        if _, err := dcn.Update("ids_login", loginset, frupd); err != nil {
            rsp.Status = 500
            rsp.Message = "Can not write to database"
            return
        }
    }

    if _, err := time.Parse("2006-01-02", c.Params.Get("birthday")); err != nil {
        c.Params.Set("birthday", "0000-00-00")
    }

    profile := map[string]interface{}{
        "birthday": c.Params.Get("birthday"),
        "aboutme":  c.Params.Get("aboutme"),
        "updated":  rdc.TimeNow("datetime"),
    }
    if isNew {
        profile["uid"] = c.Params.Get("uid")
        profile["gender"] = 0
        profile["created"] = rdc.TimeNow("datetime")

        dcn.Insert("ids_profile", profile)
    } else {
        dcn.Update("ids_profile", profile, frupd)
    }

    rsp.Status = 200
    rsp.Message = ""
}
