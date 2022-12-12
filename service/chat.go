package service

import (
	//"fmt"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
	//database "webgoapp.go/db"
)

func OpenchatForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "public.navbar.html", nil)
}
