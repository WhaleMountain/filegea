package server

import (
	"github.com/gin-gonic/gin"
)

//FileGea file gea
type FileGea struct {
	route *gin.Engine
}

//Init route init
func Init() *FileGea {
	r := router()

	return &FileGea{
		route: r,
	}
}

//Run run gin server
func (fg *FileGea) Run(port string) {
	fg.route.Run(":"+port)
}
