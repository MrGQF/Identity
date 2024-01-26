package service

import (
	"encoding/json"
	"gitee.com/QunXiongZhuLu/KongMing/config"
	base "gitee.com/QunXiongZhuLu/KongMing/error"
	error2 "gitee.com/QunXiongZhuLu/KongMing/error"
	ac "gitee.com/QunXiongZhuLu/KongMing/tars-protocol/account"
	it "gitee.com/QunXiongZhuLu/KongMing/tars-protocol/identify"
	"gitee.com/QunXiongZhuLu/KongMing/util"
	"gitee.com/QunXiongZhuLu/kratos/pkg/ecode"
	"github.com/TarsCloud/TarsGo/tars"
)

var (
	Comm                  *tars.Communicator
	CertificateServerAddr = config.CertificateServerAddr
	UserBaseServerAddr    = config.UserBaseServerAddr
)

func init() {
	Comm = tars.NewCommunicator()
}

func (s *Service) TarsGetAppSecret(appId string, appKey string) (secret string, err error) {
	var (
		obj       = CertificateServerAddr
		appSecret = new(string)
		ret       int32
	)

	app := new(it.Token)
	Comm.StringToProxy(obj, app)

	ret, err = app.GetAppKeySecret(appId, appKey, appSecret)
	if err != nil {
		err = error2.TarsGetAppSecretError
		return
	}

	if ret == config.AppSecretNotFound {
		err = error2.TarsAppSecretNotFound
		return
	}

	if *appSecret == "" {
		err = error2.TarsAppSecretIsEmpty
		return
	}

	return *appSecret, nil
}

func (s *Service) TarsCheckToken(appId, token, version string) (cert *it.CertificationInfo, err error) {
	var (
		ret int32
		obj = CertificateServerAddr
	)

	app := new(it.Token)
	Comm.StringToProxy(obj, app)
	cert = new(it.CertificationInfo)
	ret, err = app.Validate(appId, token, version, cert)
	if err != nil {
		return nil, error2.TarsCheckTokenError
	}

	if ret == config.TokenExpire {
		return nil, error2.TarsTokenExpired
	}

	if cert.Thisid == "" || cert.Extern == "" {
		return nil, error2.TarsTokenShallNotValid
	}

	return cert, err
}

func (s *Service) TarsRefreshToken(appId string, refreshToken string) (tokenInfo *it.TokenInfo, err error) {
	var (
		obj = CertificateServerAddr
		ret int32
	)

	app := new(it.Token)
	Comm.StringToProxy(obj, app)

	tokenInfo = new(it.TokenInfo)
	ret, err = app.RefreshToken(appId, refreshToken, tokenInfo)
	if err != nil {
		return nil, error2.TarsRefreshTokenError
	}

	if ret == config.RefreshTokenNotFound {
		return nil, error2.TarsRefreshTokenNotFound
	}

	if tokenInfo.RefreshToken == "" || tokenInfo.Refresh_expire == "" || tokenInfo.AccessToken == "" || tokenInfo.Access_expire == "" {
		return nil, error2.TarsTokenInfoNotValid
	}

	return tokenInfo, nil
}

func (s *Service) TarsClearToken(appId, refreshToken string) (err error) {
	var (
		obj = CertificateServerAddr
	)

	app := new(it.Token)
	Comm.StringToProxy(obj, app)
	_, err = app.ClearToken(appId, refreshToken)
	if err != nil {
		return error2.TarsClearTokenError
	}

	return nil
}

func (s *Service) TarsGetCookie(userId int32, signValid string, res *ac.GetCookieRes) (err error) {

	var (
		app       = new(ac.Accbase)
		errMsg    []byte
		errMsgStr string
		req       = &ac.GetCookieReq{
			UserId:    userId,
			SignValid: signValid,
		}
		result = new(ac.Result)
	)

	Comm.StringToProxy(UserBaseServerAddr, app)
	app.TarsSetTimeout(5000)
	if _, err = app.GetCookie(req, res, result); err != nil {
		return
	}

	if result.Code == 0 {
		return
	}

	errMsg, err = json.Marshal(result.Stacktrace)
	errMsgStr = *util.BytesToStr(errMsg)
	if result.Code == -1 {
		err = ecode.Error(ecode.Code(base.TarsGetCookieError.Code()), errMsgStr)
		return
	}
	err = ecode.Error(ecode.Code(result.Code), errMsgStr)

	return
}

func (s *Service) TarsGetPCPassport(req *ac.PCPassport, res *ac.GetPassportRes) (err error) {

	var (
		app       = new(ac.Accbase)
		errMsg    []byte
		errMsgStr string
		result    = new(ac.Result)
		mobileReq = new(ac.MobilePassport)
	)

	Comm.StringToProxy(UserBaseServerAddr, app)
	app.TarsSetTimeout(1000)
	if _, err = app.GetPassport(1, req, mobileReq, res, result); err != nil {
		return
	}

	if result.Code == 0 {
		return
	}

	errMsg, err = json.Marshal(result.Stacktrace)
	errMsgStr = *util.BytesToStr(errMsg)
	if result.Code == -1 {
		err = ecode.Error(ecode.Code(base.TarsGetPCPassportError.Code()), errMsgStr)
		return
	}
	err = ecode.Error(ecode.Code(result.Code), errMsgStr)

	return
}
