package main

import (
	"github.com/gin-gonic/gin"
	"mygo/myapp/controller"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob(filepath.Clean(`static/*.html`))
	r.Static("static", "static")

	r.GET("/notes", controller.ListNotes)
	r.GET("/test", controller.GetTest)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
