package controller

import (
	"github.com/gin-gonic/gin"
	"mygo/myapp/model"
	"net/http"
)

func ListNotes(ctx *gin.Context) {
	notes, err := model.ListNotes()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "list notes failed")
		return
	}
	ctx.HTML(http.StatusOK, "list_notes.html", notes)
}

func GetTest(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "test.html", nil)
}
