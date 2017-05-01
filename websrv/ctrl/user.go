// Copyright 2014 lessos Authors, All rights reserved.
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

package ctrl

import (
	"code.hooto.com/lessos/iam/config"
	"code.hooto.com/lessos/iam/iamclient"
	"github.com/lessos/lessgo/httpsrv"
)

type User struct {
	*httpsrv.Controller
}

func (c User) PanelInfoAction() {

	rsp := map[string]interface{}{}
	//
	nav := []map[string]string{
		{"path": "my-app/index", "title": "My Authorized Apps"},
	}

	if iamclient.SessionAccessAllowed(c.Session, "user.admin", config.Config.InstanceID) {
		nav = append(nav, map[string]string{
			"path":  "user-mgr/index",
			"title": "Accounts",
		})
	}

	if iamclient.SessionAccessAllowed(c.Session, "sys.admin", config.Config.InstanceID) {

		nav = append(nav, map[string]string{
			"path":  "app-mgr/index",
			"title": "Authorized Apps",
		})
		nav = append(nav, map[string]string{
			"path":  "sys-mgr/index",
			"title": "Sys Settings",
		})
	}

	rsp["topnav"] = nav
	rsp["webui_banner_title"] = config.Config.WebUiBannerTitle

	c.RenderJson(rsp)
}
