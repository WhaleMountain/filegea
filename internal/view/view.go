package view

import (
	"os"
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
	p {
		text-align: left;
	}
	`
)

//Template html
func Template(fInfos []os.FileInfo) string {
	title := "FileGEA"

	var items strings.Builder
	for _, finfo := range fInfos {
		items.WriteString(`
		<div class="item">
		<img src="` + "/Trash/"+ finfo.Name() + `" alt="" />
		<p>
	  		abcd.
		</p>
		</div>
		`)
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
