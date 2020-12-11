package view

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	css = `
	.grid {
		display: grid;
		gap: 10px;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
	}
	.item {
		border-radius: 10px;
		background: #edf;
		padding: 15px;
		text-align: center;
	}
	img {
		max-width: 100%;
		height: auto;
	}
	video {
		width: 300px;
		height: auto;
	}
	object {
		width: 180px;
		height: auto;
	}
	p {
		text-align: left;
	}
	.linkbox {
		position: relative;
	}
	.linkbox a {
		position: absolute;
		top: 0;
		left: 0;
		height:100%;
		width: 100%;
	}
	`
)

//Template html
func Template(searchPath string, fInfos []os.FileInfo) string {
	title := "FileGEA"

	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.IsDir() {
			// ディレクトリ
			basePath := "/filegea"
			path := filepath.Join(basePath, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<a href="` + path + `"></a>
			<p>` + finfo.Name() + `</p>
			</div>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル
			basePath := "/Data"
			path := filepath.Join(basePath, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item">
			<video src="` + path + `" controls playline></video>
			<p>` + finfo.Name() + `</p>
			</div>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".pdf") {
			// PDF ファイル
			basePath := "/Data"
			path := filepath.Join(basePath, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<a href="` + path + `"></a>
			<p>` + finfo.Name() + `</p>
			</div>
			`)

		} else {
			// 画像ファイル
			basePath := "/Data"
			path := filepath.Join(basePath, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<img src="` + path + `" />
			<a href="` + path + `"></a>
			<p>` + finfo.Name() + `</p>
			</div>
			`)
		}
	}

	html := `
	<!DOCTYPE html>
	<html lang="ja">
	<head>
		<meta charset="UTF-8">
		<title>` + title + `</title>
		<style>` + css + `</style>
	</head>
	<body>
		<div class="grid">` + items.String() + `</div>
	</body>
	</html>
	`

	return html
}
