package main

import "filegea/server"

func main() {
	port := "1270"
	server.Init().Run(port)

	//files, _ := ioutil.ReadDir("./Trash")
	//for _, file := range files {
	//	fmt.Println(file.Name())
	//}

}
