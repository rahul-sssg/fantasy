package handler

import (
	usqtemplate "fantasy/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsqHandler(ctx *gin.Context) {
	usqtemplate.GenerateHtml()

	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/login.html")
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/indexTemplateOutput.html")
}

func UsqHandlerErr(ctx *gin.Context) {
	usqtemplate.GenerateHtml()
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/error.html")
}

func Login(ctx *gin.Context) {
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/login.html")
}

func LoginValidate(ctx *gin.Context) {
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/loginSuccess.html")
}

func Glogin(ctx *gin.Context) {
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/glogin.html")
}

func Applelogin(ctx *gin.Context) {
	ctx.File("/Users/Rahulkumar/Downloads/sssg/fantasy/template/applelogin.html")
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	c.String(http.StatusOK, "Cookie value: %s", cookie)
}
