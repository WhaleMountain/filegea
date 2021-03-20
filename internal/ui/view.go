package ui

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const (
	staticFS = "/Data"
	uri      = "/filegea"
)

var (
	imagePrefix = []string{
		"png", "PNG",
		"jpeg", "JPG", "JPEG", "jpg",
		"gif", "svg",
	}

	re = regexp.MustCompile(strings.Join(imagePrefix, "|"))
)

type FileGeaTemplate struct {
	Title        string
	NaviPath     string
	SavePath     string
	DeletePath   string
	DownloadPath string
	Status       string
	Items        string
	Upload       bool
	UploadDir    bool
	Delete       bool
	Download     bool
}

type ContentTemplate struct {
	Ctype    string
	Path     string
	HostPath string
	Name     string
	CheckBox bool
}

func createHtml(fgTmpl FileGeaTemplate) string {
	var resHtml strings.Builder
	htmlTmpl, _ := template.New("html").Parse(html)

	if err := htmlTmpl.Execute(&resHtml, fgTmpl); err != nil {
		log.Print(err)
	}

	return resHtml.String()
}

func newDiv(content ContentTemplate) string {
	var contentBuf strings.Builder
	divTmpl, _ := template.New("div").Parse(div)

	if err := divTmpl.Execute(&contentBuf, content); err != nil {
		log.Print(err)
	}

	return contentBuf.String()
}

//Index index html
func Index(searchPath string, fInfos []os.FileInfo) string {
	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.Name() == ".DS_Store" {
			continue
		}

		var ctype, name, path string
		if finfo.IsDir() {
			// ディレクトリ
			ctype = "dir"
			path = filepath.Join(uri, searchPath, finfo.Name())
			name = finfo.Name()

		} else if strings.HasSuffix(finfo.Name(), ".mp4") || strings.HasSuffix(finfo.Name(), ".MP4") {
			// 動画ファイル
			ctype = "video"
			path = filepath.Join(staticFS, searchPath, finfo.Name())
			name = ""

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル
			ctype = "img"
			path = filepath.Join(staticFS, searchPath, finfo.Name())
			name = ""

		} else {
			//ディレクトリ, 画像, 動画 以外
			ctype = "other"
			path = filepath.Join(staticFS, searchPath, finfo.Name())
			name = finfo.Name()
		}

		content := ContentTemplate{
			Ctype: ctype,
			Path:  path,
			Name:  name,
		}
		items.WriteString(newDiv(content))

	}

	fgTmpl := FileGeaTemplate{
		Title:    "FileGEA",
		NaviPath: searchPath,
		Items:    items.String(),
	}

	return createHtml(fgTmpl)
}

//Upload upload html
func Upload(savePath, status string) string {
	fgTmpl := FileGeaTemplate{
		Title:    "Upload",
		NaviPath: savePath,
		SavePath: savePath,
		Status:   status,
		Upload:   true,
	}

	return createHtml(fgTmpl)
}

//UploadDir upload html
func UploadDir(savePath, status string) string {
	fgTmpl := FileGeaTemplate{
		Title:     "UploadDir",
		NaviPath:  savePath,
		SavePath:  savePath,
		Status:    status,
		UploadDir: true,
	}

	return createHtml(fgTmpl)
}

//Delete delete html
func Delete(searchPath string, fInfos []os.FileInfo) string {
	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.Name() == ".DS_Store" {
			continue
		}
		var ctype string
		if finfo.IsDir() {
			// ディレクトリ
			ctype = "dir"
		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル
			ctype = "video"

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル
			ctype = "img"

		} else {
			//ディレクトリ, 画像, 動画 以外
			ctype = "other"
		}

		content := ContentTemplate{
			Ctype:    ctype,
			Path:     filepath.Join(staticFS, searchPath, finfo.Name()),
			HostPath: filepath.Join(searchPath, finfo.Name()),
			Name:     finfo.Name(),
			CheckBox: true,
		}
		items.WriteString(newDiv(content))
	}

	fgTmpl := FileGeaTemplate{
		Title:      "Delete",
		NaviPath:   searchPath,
		DeletePath: searchPath,
		Items:      items.String(),
		Delete:     true,
	}

	return createHtml(fgTmpl)
}

//Download download html
func Download(searchPath string, fInfos []os.FileInfo) string {
	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.Name() == ".DS_Store" {
			continue
		}
		var ctype string
		if finfo.IsDir() {
			// ディレクトリ
			ctype = "dir"
		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル
			ctype = "video"

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル
			ctype = "img"

		} else {
			//ディレクトリ, 画像, 動画 以外
			ctype = "other"
		}

		content := ContentTemplate{
			Ctype:    ctype,
			Path:     filepath.Join(staticFS, searchPath, finfo.Name()),
			HostPath: filepath.Join(searchPath, finfo.Name()),
			Name:     finfo.Name(),
			CheckBox: true,
		}
		items.WriteString(newDiv(content))
	}

	fgTmpl := FileGeaTemplate{
		Title:        "Download",
		NaviPath:     searchPath,
		DownloadPath: searchPath,
		Items:        items.String(),
		Download:     true,
	}

	return createHtml(fgTmpl)
}
