// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package conf

import "time"

type User struct {
	UserId                          string       `json:"userId"`
	UserName                        string       `json:"userName"`
	UserAvatarURL                   string       `json:"userAvatarURL"`
	UserHomeBImgURL                 string       `json:"userHomeBImgURL"`
	UserTitles                      []*UserTitle `json:"userTitles"`
	UserIntro                       string       `json:"userIntro"`
	UserNickname                    string       `json:"userNickname"`
	UserCreateTime                  string       `json:"userCreateTime"`
	UserSiYuanProExpireTime         float64      `json:"userSiYuanProExpireTime"`
	UserToken                       string       `json:"userToken"`
	UserTokenExpireTime             string       `json:"userTokenExpireTime"`
	UserSiYuanRepoSize              float64      `json:"userSiYuanRepoSize"`
	UserSiYuanPointExchangeRepoSize float64      `json:"userSiYuanPointExchangeRepoSize"`
	UserSiYuanAssetSize             float64      `json:"userSiYuanAssetSize"`
	UserTrafficUpload               float64      `json:"userTrafficUpload"`
	UserTrafficDownload             float64      `json:"userTrafficDownload"`
	UserTrafficAPIGet               float64      `json:"userTrafficAPIGet"`
	UserTrafficAPIPut               float64      `json:"userTrafficAPIPut"`
	UserTrafficTime                 float64      `json:"userTrafficTime"`
	UserSiYuanSubscriptionPlan      float64      `json:"userSiYuanSubscriptionPlan"`   // -1：未订阅，0：标准订阅，1：教育订阅，2：试用
	UserSiYuanSubscriptionStatus    float64      `json:"userSiYuanSubscriptionStatus"` // -1：未订阅，0：订阅可用，1：订阅封禁，2：订阅过期
	UserSiYuanSubscriptionType      float64      `json:"userSiYuanSubscriptionType"`   // 0 年付；1 终生；2 月付
	UserSiYuanOneTimePayStatus      float64      `json:"userSiYuanOneTimePayStatus"`   // 0 未付费；1 已付费
}

type UserTitle struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Icon string `json:"icon"`
}

func (user *User) GetCloudRepoAvailableSize() int64 {
	return int64(user.UserSiYuanRepoSize - user.UserSiYuanAssetSize)
}

func (user *User) hasSubscribed() bool {
	return 0 == user.UserSiYuanSubscriptionStatus &&
		(0 == user.UserSiYuanSubscriptionPlan || 1 == user.UserSiYuanSubscriptionPlan) &&
		1 == user.UserSiYuanOneTimePayStatus
}

func (user *User) ToSubscriptionUser() {
	if user.hasSubscribed() {
		return
	}

	user.UserSiYuanProExpireTime = float64(time.Now().UnixMilli() + (365 * 24 * 3600 * 1000 / 2))
	user.UserSiYuanSubscriptionPlan = 0
	user.UserSiYuanSubscriptionStatus = 0
	user.UserSiYuanSubscriptionType = 0
	user.UserSiYuanOneTimePayStatus = 1
}
