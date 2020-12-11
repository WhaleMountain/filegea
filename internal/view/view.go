package view

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	css = `
	body {
		margin: 20px;
		padding: -10px;
		background-color: #e6e6e6;
	}
	.grid {
		display: grid;
		padding: 10px;
		gap: 10px;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
	}
	.item {
		border-radius: 10px;
		background: #b9c8d5;
		padding: 15px;
		text-align: center;
		box-shadow: 4px 4px;
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
		font-size: large;
		font-weight: bold;
		color: #4b4b4b;
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
	header {
		padding: 30px 4% 5px;
		top: 0;
		width: 100%;
		background-color: #fff;
		display: flex;
		align-items: center;
	}
	a {	
		text-decoration: none;
		color: #4b4b4b;
	}
	ul {
		list-style: none;
		margin: 0;
		display: flex;
	}
	li {
		margin: 0 0 0 15px;
		font-size: 14px;
	}
	nav {
		margin: auto auto auto 0;
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
		<header>
		<h1>
			<a href="/">FILEGEA</a>
		</h1>
		<nav class="pc-navi">
			<ul>
				<li><a href="/filegea"><h2>HOME</h2></a></li>
				<li><a href="#"><h2>UPLOAD</h2></a></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<div class="grid">` + items.String() + `</div>
	</body>
	</html>
	`

	return html
}
