package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, "OK", map[string]interface{}{}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, message, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, "OK", data, c)
}

func OkWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(http.StatusOK, message, data, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusInternalServerError, "SERVER ERROR", map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusInternalServerError, message, map[string]interface{}{}, c)
}

func FailWithCode(code int, message string, c *gin.Context) {
	Result(code, message, map[string]interface{}{}, c)
}

func FailWithData(message string, data interface{}, c *gin.Context) {
	Result(http.StatusInternalServerError, message, data, c)
}

func FailWithDetailed(code int, message string, data interface{}, c *gin.Context) {
	Result(code, message, data, c)
}
