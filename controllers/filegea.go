package controllers

import (
	"filegea/internal/view"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//FileGeaController filegea controller
type FileGeaController struct{}

//NewFileGeaController filegea controller
func NewFileGeaController() *FileGeaController {
	return new(FileGeaController)
}

//Index index page
func (fgc *FileGeaController) Index(c *gin.Context) {

	files, _ := ioutil.ReadDir("./Trash/picture")

	c.Writer.WriteString(view.Template(files))
}
