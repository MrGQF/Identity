package http

import (
	"fmt"
	h "gitee.com/QunXiongZhuLu/KongMing/http"
	"gitee.com/QunXiongZhuLu/KongMing/log"
	"gitee.com/QunXiongZhuLu/KongMing/tars-protocol/account"
	bm "gitee.com/QunXiongZhuLu/kratos/pkg/net/http/blademaster"
	"identify/app/model/param"

	"time"
)

func GetUserInfoByToken(c *bm.Context) {
	var (
		req = new(param.GetUserInfoByTokenReq)
		err error
	)

	// 调用记录
	defer func(name string, tw time.Time) {
		tc := time.Since(tw)
		log.TLog.Warn(fmt.Sprintf("LogRec Method:%v, timecost:%v, req:%v, err:%v", name, tc, req, err))
	}("Http ValidateToken", time.Now())

	err = c.Bind(req)
	if err != nil {
		return
	}

	// 校验token
	common := h.GetCommon(c)
	c.JSON(svc.GetUserInfoByToken(req, common))
	return

}

func RefreshToken(c *bm.Context) {
	var (
		req = new(param.RefreshTokenReq)
		res = new(param.RefreshTokenRes)
		err error
	)

	// 调用记录
	defer func(name string, tw time.Time) {
		tc := time.Since(tw)
		log.TLog.Warn(fmt.Sprintf("LogRec Method:%v, timecost:%v, req:%v, res:%v, err:%v", name, tc, req, res, err))
	}("Http RefreshToken", time.Now())

	err = c.Bind(req)
	if err != nil {
		return
	}

	res, err = svc.RefreshToken(c, req)
	c.JSON(res, err)
}

func Logout(c *bm.Context) {
	var (
		req = new(param.LogoutReq)
		err error
	)

	// 调用记录
	defer func(name string, tw time.Time) {
		tc := time.Since(tw)
		log.TLog.Warn(fmt.Sprintf("LogRec Method:%v, timecost:%v, req:%v, err:%v", name, tc, req, err))
	}("Http Logout", time.Now())

	err = c.Bind(req)
	if err != nil {
		return
	}

	c.JSON(nil, svc.Logout(c, req))
}

func GetPCPassport(c *bm.Context) {
	var (
		req = new(param.GetPassportReq)
		err error
		res string
	)

	err = c.Bind(req)
	if err != nil {
		return
	}

	res, err = svc.GetPCPassport(c, req)
	c.JSON(res, err)
	return
}

func GetCookie(c *bm.Context) {
	var (
		req = new(param.GetCookieReq)
		res = new(account.GetCookieRes)
		err error
	)

	err = c.Bind(req)
	if err != nil {
		return
	}

	err = svc.GetCookie(c, req, res)
	c.JSON(res, err)
	return
}
