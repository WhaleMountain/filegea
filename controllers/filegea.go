package controllers

import (
	"filegea/internal/view"
	"io/ioutil"
	"net/http"
	"path/filepath"

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
	basePath := "./Data"
	searchPATH := c.Param("path")

	path := filepath.Join(basePath, searchPATH)

	files, _ := ioutil.ReadDir(path)

	c.Writer.WriteString(view.Template(searchPATH, files))
}

//Redirect / -> /filegea
func (fgc *FileGeaController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/filegea")
}
