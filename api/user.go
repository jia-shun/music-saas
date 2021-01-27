package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"music-saas/global"
	"music-saas/middleware"
	"music-saas/model"
	"music-saas/model/request"
	"music-saas/model/response"
	"music-saas/service"
	"music-saas/utils"
	"net/http"
	"time"
)

func Login(ctx *gin.Context) {
	var R request.Login
	_ = ctx.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	U := model.User{Username: R.Username, Password: R.Password}
	if user, err := service.Login(U); err != nil {
		errMsg := "Login failed: The username does not exist or the password is incorrect"
		global.LOG.Error(errMsg)
		response.FailWithDetailed(http.StatusBadRequest, errMsg, user, ctx)
	} else {
		issueToken(user, ctx)
	}
}

func Logout(ctx *gin.Context) {
	response.OkWithMessage("Logout success", ctx)
}

func GetInfo(ctx *gin.Context) {
	if claims, exists := ctx.Get("claims"); !exists {
		global.LOG.Error("get user id from context failed")
		response.FailWithCode(http.StatusUnauthorized, "Get user info failed: unauthorized", ctx)
	} else {
		user := claims.(*request.CustomClaims)
		userReturn, err := service.FindUserById(user.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), ctx)
		}
		response.OkWithData(userReturn, ctx)
	}
}

func Register(ctx *gin.Context) {
	var R request.Register
	_ = ctx.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	u := &model.User{Username: R.Username, NickName: R.NickName, Password: R.Password, Avatar: R.Avatar, Phone: R.Phone,
		Email: R.Email, Sex: R.Sex, Age: R.Age, Status: true}
	userReturn, err := service.Register(*u)
	if err != nil {
		global.LOG.Error("Register failed", zap.Any("err", err))
		response.FailWithDetailed(http.StatusBadRequest, "Register failed: "+err.Error(), response.SysUserResponse{User: userReturn}, ctx)
	} else {
		response.OkWithDetailed("Register success", response.SysUserResponse{User: userReturn}, ctx)
	}
	return
}

func issueToken(user model.User, ctx *gin.Context) {
	j := middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)}
	claims := request.CustomClaims{
		ID:         user.ID,
		Username:   user.Username,
		BufferTime: global.CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime,
			Issuer:    "coolMusic",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("Get token failed", zap.Any("err", err))
		response.FailWithMessage("Get token failed", ctx)
		return
	}
	response.OkWithDetailed("Login success", response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, ctx)
	return
}
