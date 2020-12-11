package server

import (
	"filegea/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	filegea := controllers.NewFileGeaController()

	r.GET("/filegea", filegea.Index)

	r.StaticFS("/Trash", http.Dir("./Trash/picture"))

	return r
}
