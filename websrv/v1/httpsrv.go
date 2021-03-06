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

package v1

import (
	"github.com/hooto/httpsrv"
)

func NewModule() httpsrv.Module {

	module := httpsrv.NewModule("iam_api")

	module.ControllerRegister(new(Service))

	module.ControllerRegister(new(User))
	module.ControllerRegister(new(Account))
	module.ControllerRegister(new(AccountCharge))
	module.ControllerRegister(new(App))

	module.ControllerRegister(new(AppAuth))
	module.ControllerRegister(new(AccessKey))
	module.ControllerRegister(new(Status))

	module.ControllerRegister(new(UserMgr))
	module.ControllerRegister(new(AppMgr))
	module.ControllerRegister(new(AccountMgr))
	module.ControllerRegister(new(SysConfig))

	return module
}
