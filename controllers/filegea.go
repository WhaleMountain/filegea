package controllers

import (
	"filegea/internal/fileope"
	"filegea/internal/view"
	"io/ioutil"
	"log"
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

	c.Writer.WriteString(view.Index(searchPATH, files))
}

//UploadFrom upload form
func (fgc *FileGeaController) UploadFrom(c *gin.Context) {
	savePath := c.Param("path")
	c.Writer.WriteString(view.Upload(savePath))
}

//Upload upload file
func (fgc *FileGeaController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("form file error %s\n", err)
		c.Redirect(http.StatusMovedPermanently, "/filegea")
	}

	savePath := c.Param("path")
	if err := fileope.Save(file, savePath); err != nil {
		log.Printf("form file error %s\n", err)
	}

	c.Redirect(http.StatusMovedPermanently, "/filegea")
}

//Redirect / -> /filegea
func (fgc *FileGeaController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/filegea")
}
