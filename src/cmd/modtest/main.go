package main

import "fmt"
import "myapp/mymath"

func main() {
	# modファイルの生成
	# go mod init myapp
	fmt.Println(mymath.Add(100, 200))
}
