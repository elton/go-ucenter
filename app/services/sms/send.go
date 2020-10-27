/* Package sms 短信发送服务
 *
 * @Author: Elton Zheng
 * @Date: 2020-10-26 13:51:27
 * @Last Modified by: Elton Zheng
 * @Last Modified time: 2020-10-26 13:54:08
 */

package sms

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"ucenter/app"
	"ucenter/app/utils/crypto"

	"github.com/apptut/rsp"
	"github.com/go-redis/redis/v8"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

var ctx = context.Background()
var redisClinet *redis.Client
var err error

// UpdateSMSCache 更新SMS发送记录缓存
func UpdateSMSCache(ip string, mobile string) error {
	redisKey := "sms:" + crypto.MD5(ip+mobile)

	if err = app.GetRedis().Set(ctx, redisKey, true, time.Minute*1).Err(); err != nil {
		return errors.New("添加失败")
	}

	return nil
}

// SendEnable 验证发送频率
func SendEnable(ip string, mobile string) bool {
	redisKey := "sms:" + crypto.MD5(ip+mobile)
	if err = app.GetRedis().Get(ctx, redisKey).Err(); err != nil {
		return err == redis.Nil
	}
	return false
}

// PostCode 发送验证码
func PostCode(mobile string) *rsp.Error {
	code := generateCode()

	// 云片短信平台 API-KEY
	client := ypclnt.New("ee89935b41b8d4262b4c56e2594dfd49")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = mobile
	param[ypclnt.TEXT] = "【郑振宇】您的验证码是" + code

	r := client.Sms().SingleSend(param)
	if r.Code != ypclnt.SUCC {
		return rsp.NewErrMsg(r.Msg)
	}
	return nil
}

// 生成验证码
func generateCode() string {
	code := ""
	// 以当前时间作为伪随机种子
	rand.Seed(time.Now().Unix())

	// 生成一个[0,10)范围的数字，循环4次，生成长度为4位的字符串
	for i := 0; i < 4; i++ {
		// strconv是golang用来做数据类型转换的一个库。
		// Itoa 将 int 转化为 string
		// rand.Intn(n int) int，返回一个大于等于0小于n的正整数 [0,n)。
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
