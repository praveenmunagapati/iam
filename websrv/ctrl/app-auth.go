package ctrl

import (
	"github.com/hooto/httpsrv"
	// "github.com/lessos/lessgo/utils"
)

type AppAuth struct {
	*httpsrv.Controller
}

// func (c AppAuth) RegisterAction() {

// 	set := iamapi.AppInstanceRegister{}

// 	defer c.RenderJson(&set)

// 	if err := c.Request.JsonDecode(&set); err != nil {
// 		set.Error = &types.ErrorMeta{iamapi.ErrCodeInvalidArgument, "Bad Argument"}
// 	}

// 	if set.Meta.ID == "" {

// 		set.Error = &types.ErrorMeta{iamapi.ErrCodeInvalidArgument, "Bad Argument"}

// 	} else {

// 	}

// 	// if !c.Session.AccessAllowed("sys.admin") {
// 	//     set.Error = &types.ErrorMeta{iamapi.ErrCodeAccessDenied, "Unauthorized"}
// 	//     return
// 	// }

// 	// sess, err := c.Session.SessionFetch()

// 	var prevVersion uint64
// 	var prev iamapi.AppInstance

// 	if obj := store.PvGet(btapi.ObjectPut{
// 		Meta: btapi.ObjectMeta{
// 			Path: "/app-instance/" + set.Meta.ID,
// 		},
// 	}); obj.Error == nil {
// 		obj.JsonDecode(&prev)
// 		prevVersion = obj.Meta().Version
// 	}

// 	if prev.Meta.ID == "" {

// 		set.Meta.Created = utilx.TimeNow("datetime")
// 		set.Meta.Updated = utilx.TimeNow("datetime")
// 		set.Status = 1
// 		set.Meta.UserID = ""

// 	} else {

// 		set.Meta.Created = prev.Meta.Created
// 		set.Meta.UserID = prev.Meta.UserID
// 		set.Status = prev.Status
// 	}

// 	setjs, _ := utils.JsonEncode(set)

// 	if obj := store.PvPut(btapi.ObjectPut{
// 		Meta: btapi.ObjectMeta{
// 			Path: "/app-instance/" + set.Meta.ID,
// 		},
// 		Data:        setjs,
// 		PrevVersion: prevVersion,
// 	}); obj.Error != nil {
// 		set.Error = &types.ErrorMeta{iamapi.ErrCodeInternalError, obj.Error.Message}
// 		return
// 	}

// 	//
// 	// q = base.NewQuerySet().From("iam_privilege").Limit(1000)
// 	// q.Where.And("instance", req.Data.InstanceId)
// 	// rs, err = dcn.Base.Query(q)
// 	// if err != nil {
// 	//  rsp.Message = "Internal Server Error"
// 	//  return
// 	// }

// 	// for _, prePriv := range rs {

// 	//  isExist := false
// 	//  for _, curPrev := range req.Data.Privileges {

// 	//      if prePriv.Field("privilege").String() == curPrev.Key {
// 	//          isExist = true
// 	//          break
// 	//      }
// 	//  }

// 	//  if !isExist {
// 	//      frupd := base.NewFilter()
// 	//      frupd.And("instance", req.Data.InstanceId).And("privilege", prePriv.Field("privilege").String())
// 	//      dcn.Base.Delete("iam_privilege", frupd)
// 	//  }
// 	// }

// 	// for _, curPrev := range req.Data.Privileges {

// 	//  isExist := false

// 	//  for _, prePriv := range rs {

// 	//      if prePriv.Field("privilege").String() == curPrev.Key {
// 	//          isExist = true
// 	//          break
// 	//      }
// 	//  }

// 	//  if !isExist {
// 	//      item := map[string]interface{}{
// 	//          "instance":  req.Data.InstanceId,
// 	//          "uid":       sess.UserID,
// 	//          "privilege": curPrev.Key,
// 	//          "desc":      curPrev.Desc,
// 	//          "created":   base.TimeNow("datetime"),
// 	//      }

// 	//      if _, err := dcn.Base.Insert("iam_privilege", item); err != nil {
// 	//          rsp.Status = 500
// 	//          rsp.Message = "Can not write to database" + err.Error()
// 	//          return
// 	//      }
// 	//  }
// 	// }

// 	set.Kind = "AppInstance"
// }
