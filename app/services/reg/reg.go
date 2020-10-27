package reg

import (
	"context"
	"ucenter/app"
	"ucenter/app/models"

	"github.com/apptut/rsp"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SaveData 保存表单数据
func SaveData(ctx *gin.Context) *rsp.Error {
	password, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")), 10)
	model := models.User{
		UserName: ctx.PostForm("username"),
		Mobile:   ctx.PostForm("mobile"),
		Password: string(password),
	}
	err = app.GetDB().Create(&model).Error
	if err != nil {
		return rsp.NewErr(err)
	}
	return nil
}

// CheckUser 检测用户是否已经注册
func CheckUser(ctx *gin.Context) *rsp.Error {
	username := ctx.PostForm("username")
	mobile := ctx.PostForm("mobile")

	var model models.User
	err := app.GetDB().Find(&model, "username=? or mobile=?", username, mobile).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil
	}
	return rsp.NewErrMsg("User already exists.")
}

// CheckCode 手机验证码核对
func CheckCode(ctx *gin.Context) *rsp.Error {
	ctx := context.Background()
	model := ctx.PostForm("model")
	code := ctx.PostForm("code")

	cachedCode, err := app.GetRedis().Get(ctx, "reg:"+mobile).Result()
	if err != nil {
		return rsp.NewErr(err)
	}

	if code != cachedCode {
		return rsp.NewErrMsg("验证码不正确")
	}
	return nil
}

// Valid 验证表单
func Valid(ctx *gin.Context) *rsp.Error {
	_, err := validator.New(map[string][]string{
		"username": {ctx.PostForm("username")},
		"mobile":   {ctx.PostForm("mobile")},
		"password": {ctx.PostForm("password")},
		"code":     {ctx.PostForm("code")},
	}, map[string]string{
		"username": "regex:^[\\w_]{6,20}$",
		"mobile":   "mobile",
		"password": "regex:^[\\S]{6,20}$",
		"code":     "regex:^[0-9]{4}$",
	})

	if err != nil {
		return rsp.NewErr(err)
	}
	return nil
}
