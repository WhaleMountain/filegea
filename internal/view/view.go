package view

import (
	"os"
	"path/filepath"
	"strings"
)

// Form HTML and CSS. Thank you.
// https://codepen.io/TheLukasWeb/pen/qlGDa

const (
	css = `
	.delete{
		position: relative;
		z-index: 20;
	}
	.upload{
		position: absolute;
		top: 40%;
		left: 50%;
		margin-top: -100px;
		margin-left: -250px;
		width: 500px;
		height: 200px;
		border: 4px dashed #b9c8d5;
	}
	.upload p{
		width: 100%;
		height: 100%;
		text-align: center;
		line-height: 170px;
		color: #4b4b4b;
	}
	.upload .file{
		position: absolute;
		margin: 0;
		padding: 0;
		width: 100%;
		height: 100%;
		outline: none;
		opacity: 0;
	}
	.upload button{
		margin: 0;
		color: #4b4b4b;
		background: #b9c8d5;
		width: 508px;
		height: 35px;
		margin-top: -20px;
		margin-left: -4px;
		transition: all .2s ease;
		outline: none;
		box-shadow: 4px 4px;
	}
	.delete {
		text-align: center;
	}
	.delete button{
		color: #fff;
		background: #9179a4;
		width: 600px;
		height: 50px;
		box-shadow: 4px 4px;
	}
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
		word-wrap: break-word;
	}
	img {
		max-width: 100%;
		height: auto;
	}
	video {
		width: 100%;
		height: auto;
	}
	div p {
		text-align: left;
		font-size: large;
		font-weight: bold;
		color: #4b4b4b;
		text-align: center;
	}
	.linkbox {
		position: relative;
		z-index 10;
	}
	.linkbox a {
		position: absolute;
		top: 0;
		left: 0;
		height: 100%;
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

	staticFS = "/Data"
	uri      = "/filegea"
)

//Index index html
func Index(searchPath string, fInfos []os.FileInfo) string {
	title := "FileGEA"

	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.Name() == ".DS_Store" {
			continue
		}
		if finfo.IsDir() {
			// ディレクトリ
			path := filepath.Join(uri, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<a href="` + path + `"></a>
			<p>` + finfo.Name() + `</p>
			</div>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル
			path := filepath.Join(staticFS, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item">
			<video src="` + path + `" controls playline></video>
			<!-- <p>` + finfo.Name() + `</p> -->
			</div>
			`)

		} else {
			// 画像ファイル
			path := filepath.Join(staticFS, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<img src="` + path + `" />
			<a href="` + path + `"></a>
			<!-- <p>` + finfo.Name() + `</p> -->
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
				<li><a href="/upload` + searchPath + `"><h2>UPLOAD</h2></a></li>
				<li><a href="/delete` + searchPath + `"><h2>DELETE</h2></a></li>
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

//Upload upload html
func Upload(savePath string) string {
	title := "FileGEA Upload"

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
				<li><a href="/upload` + savePath + `"><h2>UPLOAD</h2></a></li>
				<li><a href="/delete` + savePath + `"><h2>DELETE</h2></a></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/upload` + savePath + `" method="POST" class="upload" enctype="multipart/form-data">
			<input type="file" class="file" name="file" multiple>
			<p>Drag and Drop or Click in this area.</p>
			<button type="submit">Upload</button>
  		</form>
	</body>
	</html>
	`

	return html
}

//Delete delete html
func Delete(searchPath string, fInfos []os.FileInfo) string {
	title := "FileGEA"

	var items strings.Builder
	for _, finfo := range fInfos {
		if finfo.Name() == ".DS_Store" {
			continue
		}

		fsPath := filepath.Join(staticFS, searchPath, finfo.Name())
		hostPath := filepath.Join(searchPath, finfo.Name())
		if finfo.IsDir() {
			// ディレクトリ

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="dir" value="` + hostPath + `" >
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル

			items.WriteString(`
			<label>
			<div class="item">
			<input type="checkbox" name="video" value="` + hostPath + `" >
			<video src="` + fsPath + `" controls playline></video>
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else {
			// 画像ファイル

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="img" value="` + hostPath + `" >
			<img src="` + fsPath + `" />
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
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
				<li><a href="/upload` + searchPath + `"><h2>UPLOAD</h2></a></li>
				<li><a href="/delete` + searchPath + `"><h2>DELETE</h2></a></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/delete` + searchPath + `" class="delete" method="POST">
			<div class="grid">` + items.String() + `</div>
			<button type="submit">Delete</button>
		</form>
	</body>
	</html>
	`

	return html
}
