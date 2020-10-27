package user

import (
	"ucenter/app/services/reg"

	"github.com/apptut/rsp"
	"github.com/gin-gonic/gin"
)

// Reg 注册用户信息
func Reg(ctx *gin.Context) {
	// 参数验证
	err := reg.Valid(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 验证验证码
	err = reg.CheckCode(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 验证用户信息
	err = reg.CheckUser(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 保存数据
	err = reg.SaveData(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}
	rsp.JsonOk(ctx)
}
