// Copyright 2014 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iamapi

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/lessos/lessgo/crypto/idhash"
	"github.com/lessos/lessgo/encoding/json"
	"github.com/lessos/lessgo/types"
)

const (
	AccessTokenKey = "_iam_at"
	UserKeyDefault = "std"

	ErrCodeAccessDenied    = "AccessDenied"
	ErrCodeUnauthorized    = "Unauthorized" // Need to login and fetch a new access-token
	ErrCodeInvalidArgument = "InvalidArgument"
	ErrCodeUnavailable     = "Unavailable"
	ErrCodeInternalError   = "InternalError"
	ErrCodeNotFound        = "NotFound"
	ErrCodeAccChargeOut    = "AccChargeOut"
)

var (
	UserNameRe2            = regexp.MustCompile("^[a-z]{1}[a-z0-9]{3,29}$")
	UserEmailRe2           = regexp.MustCompile("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,10})$")
	accessTokenFrontendRe2 = regexp.MustCompile("^[a-z0-9]{4,30}\\/[a-f0-9]{20,40}$")
)

func UserIdBytes(name string) []byte {
	return idhash.Hash([]byte(name), 4)
}

func UserId(name string) string {
	return hex.EncodeToString(UserIdBytes(name))
}

type UserSession struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token,omitempty"`
	UserName     string            `json:"username"`
	DisplayName  string            `json:"display_name,omitempty"`
	Roles        types.ArrayUint32 `json:"roles,omitempty"`
	Groups       types.ArrayUint32 `json:"groups,omitempty"`
	ClientAddr   string            `json:"client_addr,omitempty"`
	Created      types.MetaTime    `json:"created"`
	Expired      types.MetaTime    `json:"expired"`
}

func (s *UserSession) IsLogin() bool {
	return (s.AccessToken != "")
}

func (s *UserSession) FullToken() string {
	return s.UserName + "/" + s.AccessToken
}

func (s *UserSession) UserId() string {
	return UserId(s.UserName)
}

type UserAccessEntry struct {
	types.TypeMeta `json:",inline"`
	AccessToken    string `json:"access_token"`
	Privilege      string `json:"privilege"`
	InstanceID     string `json:"instance_id,omitempty"`
}

type User struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Email       string            `json:"email,omitempty"`
	DisplayName string            `json:"display_name,omitempty"`
	Keys        types.KvPairs     `json:"keys,omitempty"`
	Roles       types.ArrayUint32 `json:"roles,omitempty"`
	Groups      types.ArrayUint32 `json:"groups,omitempty"`
	Status      uint8             `json:"status"`
	Created     types.MetaTime    `json:"created"`
	Updated     types.MetaTime    `json:"updated"`
}

type UserEntry struct {
	types.TypeMeta `json:",inline"`
	Login          User         `json:"login,omitempty"`
	Profile        *UserProfile `json:"profile,omitempty"`
}

type UserList struct {
	types.TypeMeta `json:",inline"`
	Meta           types.ListMeta `json:"meta,omitempty"`
	Items          []User         `json:"items,omitempty"`
}

type UserProfile struct {
	Login       *User          `json:"login,omitempty"`
	Gender      uint8          `json:"gender,omitempty"`
	Photo       string         `json:"photo,omitempty"`
	PhotoSource string         `json:"photo_source,omitempty"`
	Birthday    string         `json:"birthday,omitempty"`
	About       string         `json:"about,omitempty"`
	Updated     types.MetaTime `json:"updated,omitempty"`
}

type UserPasswordSet struct {
	types.TypeMeta  `json:",inline"`
	CurrentPassword string `json:"current_password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
}

type UserPasswordReset struct {
	types.TypeMeta `json:",inline"`
	Id             string `json:"id,omitempty"`
	UserName       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Expired        string `json:"expired,omitempty"`
}

type UserEmailSet struct {
	types.TypeMeta `json:",inline"`
	Auth           string `json:"auth,omitempty"`
	Email          string `json:"email,omitempty"`
}

type UserPhotoSet struct {
	types.TypeMeta `json:",inline"`
	Name           string `json:"name,omitempty"`
	Size           int    `json:"size,omitempty"`
	Data           string `json:"data,omitempty"`
}

type UserRole struct {
	Id         uint32         `json:"id"`
	Name       string         `json:"name"`
	User       string         `json:"user,omitempty"`
	Status     uint8          `json:"status,omitempty"`
	Desc       string         `json:"desc,omitempty"`
	Privileges []string       `json:"privileges,omitempty"`
	Created    types.MetaTime `json:"created,omitempty"`
	Updated    types.MetaTime `json:"updated,omitempty"`
}

type UserRoleList struct {
	types.TypeMeta `json:",inline"`
	Items          []UserRole `json:"items,omitempty"`
}

type UserPrivilege struct {
	types.TypeMeta `json:",inline"`
	Meta           types.InnerObjectMeta `json:"meta,omitempty"`
	Instance       string                `json:"instance"`
	Desc           string                `json:"desc,omitempty"`
}

type UserPrivilegeList struct {
	types.TypeMeta `json:",inline"`
	Items          []UserPrivilege `json:"items,omitempty"`
}

type AppPrivilege struct {
	// ID        uint32   `json:"id,omitempty"`
	Privilege string            `json:"privilege"`
	Desc      string            `json:"desc,omitempty"`
	Roles     types.ArrayUint32 `json:"roles,omitempty"`
	// ExtRoles  types.ArrayUint32 `json:"extroles,omitempty"`
}

type AppInstance struct {
	Meta       types.InnerObjectMeta `json:"meta,omitempty"`
	AppID      string                `json:"app_id,omitempty"`
	AppTitle   string                `json:"app_title,omitempty"`
	Version    string                `json:"version,omitempty"`
	Status     uint8                 `json:"status,omitempty"`
	Url        string                `json:"url,omitempty"`
	Privileges []AppPrivilege        `json:"privileges,omitempty"`
	SecretKey  string                `json:"secret_key,omitempty"`
}

type AppInstanceList struct {
	types.TypeMeta `json:",inline"`
	Meta           types.ListMeta `json:"meta,omitempty"`
	Items          []AppInstance  `json:"items,omitempty"`
}

type AppAuthInfo struct {
	types.TypeMeta `json:",inline"`
	InstanceID     string `json:"instance_id"`
	AppID          string `json:"app_id"`
	// Version        string `json:"version,omitempty"`
}

type AppInstanceRegister struct {
	types.TypeMeta `json:",inline"`
	AccessToken    string      `json:"access_token,omitempty"`
	Instance       AppInstance `json:"instance"`
}

type SysConfigList struct {
	types.TypeMeta `json:",inline"`
	Items          types.Labels `json:"items,omitempty"`
}

type SysConfigMailer struct {
	SmtpHost string `json:"smtp_host"`
	SmtpPort string `json:"smtp_port"`
	SmtpUser string `json:"smtp_user"`
	SmtpPass string `json:"smtp_pass"`
}

type AccessTokenFrontend string

func NewAccessTokenFrontend(username, tk string) AccessTokenFrontend {
	return AccessTokenFrontend(username + "/" + tk)
}

func (t *AccessTokenFrontend) Valid() bool {
	return accessTokenFrontendRe2.MatchString(string(*t))
}

func (t *AccessTokenFrontend) SessionPath() string {
	return string(*t)
}

func (t *AccessTokenFrontend) String() string {
	return string(*t)
}

func (t *AccessTokenFrontend) User() string {
	if n := strings.Index(string(*t), "/"); n > 0 {
		return string(*t)[:n]
	}
	return ""
}

func (t *AccessTokenFrontend) Id() string {
	if n := strings.Index(string(*t), "/"); n > 0 {
		return string(*t)[n+1:]
	}
	return ""
}

type ServiceLoginAuth struct {
	types.TypeMeta `json:",inline"`
	RedirectUri    string `json:"redirect_uri,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
}

type ServiceRedirectToken struct {
	RedirectUri string `json:"uri,omitempty"`
	State       string `json:"state,omitempty"`
	ClientId    string `json:"cid,omitempty"`
	Persistent  int    `json:"pt,omitempty"`
}

func ServiceRedirectTokenValid(tokenstr string) bool {
	if _, err := base64.StdEncoding.DecodeString(tokenstr); err == nil {
		return true
	}
	return false
}

func (s *ServiceRedirectToken) Encode() string {

	if len(s.RedirectUri) > 200 {
		s.RedirectUri = s.RedirectUri[:200]
	}
	if len(s.State) > 100 {
		s.State = s.State[:100]
	}
	if len(s.ClientId) > 40 {
		s.ClientId = s.ClientId[:40]
	}

	js, _ := json.Encode(s, "")
	return base64.StdEncoding.EncodeToString(js)
}

func ServiceRedirectTokenDecode(tokenstr string) ServiceRedirectToken {
	var token ServiceRedirectToken
	if jsb, err := base64.StdEncoding.DecodeString(tokenstr); err == nil {
		json.Decode(jsb, &token)
	}
	return token
}

func Uint32ToBytes(v uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, v)
	return bs
}

func BytesToHexString(bs []byte) string {
	return hex.EncodeToString(bs)
}

func HexStringToBytes(s string) []byte {
	dec, err := hex.DecodeString(s)
	if err != nil {
		return []byte{}
	}
	return dec
}
