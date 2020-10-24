// Copyright 2020 Elton Zheng
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

package sms

import (
	"ucenter/app/services/sms"

	"github.com/apptut/rsp"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
)

// Send 发送验证码
func Send(ctx *gin.Context) {
	// 获取用户IP
	// ip := ctx.ClientIP()

	// 传入参数验证
	mobile, err := validSendParams(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 发送验证码
	err = sms.PostCode(mobile)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	rsp.JsonOk(ctx)

}

func validSendParams(ctx *gin.Context) (string, *rsp.Error) {
	// 获取GET请求参数
	mobile := ctx.Query("mobile")

	// 验证手机号是否合法
	_, err := validator.New(map[string][]string{
		// 数据，传入的是[]string，如{"18622222","19122232323"}
		"mobile": {mobile},
	}, map[string]string{
		// 验证规则
		"mobile": "mobile",
	}, map[string]string{
		// 错误提示
		"mobile": "手机号码格式不正确",
	})

	if err != nil {
		return "", rsp.NewErr(err)
	}

	return mobile, nil
}
