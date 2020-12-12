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
	r.GET("/delete/*path", filegea.DeleteForm)

	r.POST("/upload/*path", filegea.Upload)
	r.POST("/delete/*path", filegea.Delete)

	r.StaticFS("/Data", http.Dir(filegea.Conf.DataPath))

	return r
}
