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
	r.GET("/uploaddir/*path", filegea.UploadFromDir)
	r.GET("/delete/*path", filegea.DeleteForm)
	r.GET("/download/*path", filegea.DownloadForm)

	r.POST("/upload/*path", filegea.Upload)
	r.POST("/delete/*path", filegea.Delete)
	r.POST("/download/*path", filegea.Download)

	r.StaticFS("/Data", http.Dir(filegea.Conf.DataPath))

	return r
}
