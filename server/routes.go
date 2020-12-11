package server

import (
	"filegea/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	filegea := controllers.NewFileGeaController()

	r.GET("/", filegea.Redirect)
	r.GET("/filegea/*path", filegea.Index)
	r.GET("/upload/*path", filegea.UploadFrom)

	r.POST("/upload/*path", filegea.Upload)

	r.StaticFS("/Data", http.Dir("./Data"))

	return r
}
