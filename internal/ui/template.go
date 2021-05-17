package ui

// Form HTML and CSS.
// https://codepen.io/TheLukasWeb/pen/qlGDa

// Modal THML and CSS
// https://dev.to/mandrewdarts/css-only-modal-using-target-5ac7

//CSS
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
)

//HTML
const (
	html = `
	<!DOCTYPE html>
	<html lang="ja">
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		<style>` + css + `</style>
		<header>
		<h1>
			<a href="/">FILEGEA</a>
		</h1>
		<nav class="pc-navi">
			<ul>
				<li><a href="/filegea"><h2>HOME</h2></a></li>
				<li><a href="/upload{{.NaviPath}}"><h2>UPLOAD</h2></a></li>
				<li><a href="/uploaddir{{.NaviPath}}"><h2>UPLOAD DIRECTORY</h2></a></li>
				<li><a href="/download{{.NaviPath}}"><h2>DOWNLOAD</h2></a></li>
				<li><a href="/delete{{.NaviPath}}"><h2>DELETE</h2></a></li>
				{{ if .Upload }}
				<li><button type="submit" class="upbtn" form="upload">UPLOAD</button></li>
				{{ else if .UploadDir }}
				<li><button type="submit" class="upbtn" form="upload">UPLOAD DIR</button></li>
				{{ else if .Delete }}
				<li><button onclick="location.href='#warning'" class="delbtn" >DELETE</button></li>
				{{ else if .Download }}
				<li><button type="submit" class="dlbtn" form="download">DOWNLOAD</button></li>
				{{ end }}

			</ul>
		</nav>
		</header>
	</head>
	<body>
		{{ if .Upload }}
		<form action="/upload{{.SavePath}}" method="POST" id="upload" class="upload" enctype="multipart/form-data">
			<input type="file" class="file" name="file" multiple >
			<p>{{.Status}}</p>
		</form>

		{{ else if .UploadDir }}
		<form action="/upload{{.SavePath}}" method="POST" id="upload" class="upload" enctype="multipart/form-data">
			<input type="file" class="file" name="file" webkitdirectory mozdirectory>
			<p>{{.Status}}</p>
		</form>

		{{ else if .Delete }}
		<form action="/delete{{.DeletePath}}" id="delete" class="delete" method="POST">
			<div class="grid">{{.Items}}</div>
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

		{{ else if .Download }}
		<form action="/download{{.DownloadPath}}" id="download" class="download" method="POST">
			<div class="grid">{{.Items}}</div>
		</form>

		{{ else }}
		<div class="grid">{{.Items}}</div>
		{{ end }}
	</body>
	</html>
	`
)

//DIV
const (
	div = `
	{{ if eq .Ctype "dir" }}	
		{{ if .CheckBox }}
			<label>
				<div class="item linkbox dir">
					<input type="checkbox" name="target" value="{{.HostPath}}" >
					<p>{{.Name}}</p>
				</div>
			</label>

		{{ else }}
			<div class="item linkbox dir">
				<a href="{{.Path}}"></a>
				<p>{{.Name}}</p>
			</div>
		{{ end }}
	

	{{ else if eq .Ctype "video" }}
		{{ if .CheckBox }}
			<label>
				<div class="item">
					<input type="checkbox" name="target" value="{{.HostPath}}" >
					<video src="{{.Path}}" controls playline></video>
					<p>{{.Name}}</p>
				</div>
			</label>
		
		{{ else }}
			<div class="item">
				<video src="{{.Path}}" controls playline></video>
			</div>
		{{ end }}


	{{ else if eq .Ctype "img" }}
		{{ if .CheckBox }}
			<label>
				<div class="item linkbox">
					<input type="checkbox" name="target" value="{{.HostPath}}" >
					<img src="{{.Path}}" />
					<p>{{.Name}}</p>
				</div>
			</label>
		{{ else }}
			<div class="item linkbox">
				<img src="{{.Path}}" />
				<a href="{{.Path}}"></a>
			</div>
		{{ end }}


	{{ else if eq .Ctype "other" }}
		{{ if .CheckBox }}
			<label>
				<div class="item linkbox">
					<input type="checkbox" name="target" value="{{.HostPath}}" >
					<p>{{.Name}}</p>
				</div>
			</label>

		{{ else }}
			<div class="item linkbox">
				<p>{{.Name}}</p>
				<a href="{{.Path}}"></a>
			</div>
		{{ end }}
	{{ end }}
	`
)