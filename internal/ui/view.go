package ui

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Form HTML and CSS.
// https://codepen.io/TheLukasWeb/pen/qlGDa

// Modal THML and CSS
// https://dev.to/mandrewdarts/css-only-modal-using-target-5ac7

const (
	css = `
	.delete{
		position: relative;
		z-index: 20;
	}
	.upload{
		position: absolute;
		top: 40%;
		left: 20%;
		margin-top: -100px;
		margin-left: -100px;
		width: 850px;
		height: 500px;
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
	.upbtn{
		font-weight: bold;
		color: #4b4b4b;
		background: #a6deab;
		width: 100px;
		height: 50px;
		box-shadow: 4px 4px;
	}
	.upbtn:hover {
		background-color: #9bcd87;
	}
	.delete {
		text-align: center;
	}
	.delbtn{
		font-weight: bold;
		color: #4b4b4b;
		background: #a6deab;
		width: 100px;
		height: 50px;
		box-shadow: 4px 4px;
	}
	.delbtn:hover {
		background-color: #9bcd87;
	}
	.download {
		text-align: center;
	}
	.dlbtn{
		font-weight: bold;
		color: #4b4b4b;
		background: #a6deab;
		width: 100px;
		height: 50px;
		box-shadow: 4px 4px;
	}
	.dlbtn:hover {
		background-color: #9bcd87;
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
	.item:hover {
		background-color: #91a9b8;
	}
	.dir {
		background: #e7dfa0;
	}
	.dir:hover {
		background-color: #d3c77d;
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
	li a:hover {
		color: #b9c8d5;
		background-color:#999;
	}
	nav {
		margin: auto auto auto 0;
	}
	/* Greeting Modal Container */
	#warning {
	  visibility: hidden;
	  opacity: 0;
	  transition: all .5s cubic-bezier(0.075, 0.82, 0.165, 1);
	}
	
	/* Greeting Modal Container - when open */
	#warning:target {
	  visibility: visible;
	  opacity: 1;
	}
	
	/* Greeting Modal */
	#warning .modal {
	  opacity: 0;
	  transform: translateY(-1rem);
	  transition: all .3s cubic-bezier(0.075, 0.82, 0.165, 1);
	  transition-delay: .2s;
	}
	
	/* Greeting Modal - when open */
	#warning:target .modal {
	transform: translateY(0);
	opacity: 1;
	}
	
	/* Modal Container Styles */
	.modal-container {
	  position: fixed;
	  z-index: 30;
	  top: 0;
	  left: 0;
	  right: 0;
	  bottom: 0;
	  display: flex;
	  justify-content: center;
	  align-items: center;
	}
	
	/* Modal Background Styles */
	.modal-bg {
	  position: fixed;
	  top: 0;
	  left: 0;
	  right: 0;
	  bottom: 0;
	  background-color: rgba(0, 0, 0, .2);
	}
	
	/* Modal Body Styles */
	.modal {
	  z-index: 40;
	  background-color: white;
	  width: 80%;
	  max-width: 500px;
	  padding: 1rem;
	  border-radius: 8px;
	}
	.modal h3 {
		color: red;
	}
	.warbtn{
		font-weight: bold;
		color: #4b4b4b;
		background: #a6deab;
		width: 150px;
		height: 50px;
		box-shadow: 4px 4px;
		display: block;
		margin: 0 0 0 auto;
	}
	.warbtn:hover {
		background-color: #9bcd87;
	}
	`

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
			<div class="item linkbox dir">
			<a href="` + path + `"></a>
			<p>` + finfo.Name() + `</p>
			</div>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") || strings.HasSuffix(finfo.Name(), ".MP4") {
			// 動画ファイル
			path := filepath.Join(staticFS, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item">
			<video src="` + path + `" controls playline></video>
			<!-- <p>` + finfo.Name() + `</p> -->
			</div>
			`)

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル
			path := filepath.Join(staticFS, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
			<img src="` + path + `" />
			<a href="` + path + `"></a>
			<!-- <p>` + finfo.Name() + `</p> -->
			</div>
			`)

		} else {
			//ディレクトリ, 画像, 動画 以外
			path := filepath.Join(staticFS, searchPath, finfo.Name())

			items.WriteString(`
			<div class="item linkbox">
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
				<li><a href="/upload` + searchPath + `"><h2>UPLOAD</h2></a></li>
				<li><a href="/uploaddir` + searchPath + `"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download` + searchPath + `"><h2>DOWNLOAD</h2></a></li>
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
func Upload(savePath, status string) string {
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
				<li><a href="/uploaddir` + savePath + `"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download` + savePath + `"><h2>DOWNLOAD</h2></a></li>
				<li><a href="/delete` + savePath + `"><h2>DELETE</h2></a></li>
				<li><button type="submit" class="upbtn" form="upload">UPLOAD</button></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/upload` + savePath + `" method="POST" id="upload" class="upload" enctype="multipart/form-data">
			<input type="file" class="file" name="file" multiple >
			<p>` + status + `</p>
		</form>
	</body>
	</html>
	`

	return html
}

//UploadDir upload html
func UploadDir(savePath, status string) string {
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
				<li><a href="/uploaddir` + savePath + `"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download` + savePath + `"><h2>DOWNLOAD</h2></a></li>
				<li><a href="/delete` + savePath + `"><h2>DELETE</h2></a></li>
				<li><button type="submit" class="upbtn" form="upload">UPLOAD DIR</button></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/upload` + savePath + `" method="POST" id="upload" class="upload" enctype="multipart/form-data">
			<input type="file" class="file" name="file" webkitdirectory mozdirectory>
			<p>` + status + `</p>
		</form>
	</body>
	</html>
	`

	return html
}

//Delete delete html
func Delete(searchPath string, fInfos []os.FileInfo) string {
	title := "FileGEA Delete"

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
			<div class="item linkbox dir">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル

			items.WriteString(`
			<label>
			<div class="item">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<video src="` + fsPath + `" controls playline></video>
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<img src="` + fsPath + `" />
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else {
			//ディレクトリ, 画像, 動画 以外

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="target" value="` + hostPath + `" >
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
				<li><a href="/uploaddir` + searchPath + `"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download` + searchPath + `"><h2>DOWNLOAD</h2></a></li>
				<li><a href="/delete` + searchPath + `"><h2>DELETE</h2></a></li>
				<li><button onclick="location.href='#warning'" class="delbtn" >DELETE</button></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/delete` + searchPath + `" id="delete" class="delete" method="POST">
			<div class="grid">` + items.String() + `</div>
		</form>
		<!-- Modal container -->
		<div 
		  class="modal-container" id="warning">
		  <!-- Modal  -->
		  <div class="modal">
			<h3>Warning</h3>
			<p>Do you really want to delete this.</p>
			<button type="submit" class="warbtn" form="delete">Yes I really mean it</button>
		  </div>
		  <!-- Background, click to close -->
		  <a href="#" class="modal-bg"></a>
		</div>	  
	</body>
	</html>
	`

	return html
}

//Download download html
func Download(searchPath string, fInfos []os.FileInfo) string {
	title := "FileGEA Download"

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
			<div class="item linkbox dir">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else if strings.HasSuffix(finfo.Name(), ".mp4") {
			// 動画ファイル

			items.WriteString(`
			<label>
			<div class="item">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<video src="` + fsPath + `" controls playline></video>
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else if re.MatchString(finfo.Name()) {
			// 画像ファイル

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="target" value="` + hostPath + `" >
			<img src="` + fsPath + `" />
			<p>` + finfo.Name() + `</p>
			</div>
			</label>
			`)

		} else {
			//ディレクトリ, 画像, 動画 以外

			items.WriteString(`
			<label>
			<div class="item linkbox">
			<input type="checkbox" name="target" value="` + hostPath + `" >
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
				<li><a href="/uploaddir` + searchPath + `"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download` + searchPath + `"><h2>DOWNLOAD</h2></a></li>
				<li><a href="/delete` + searchPath + `"><h2>DELETE</h2></a></li>
				<li><button type="submit" class="dlbtn" form="download">DOWNLOAD</button></li>
			</ul>
		</nav>
		</header>
	</head>
	<body>
		<form action="/download` + searchPath + `" id="download" class="download" method="POST">
			<div class="grid">` + items.String() + `</div>
		</form>
	</body>
	</html>
	`

	return html
}
