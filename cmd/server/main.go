package main

import (
	"fmt"
	"os"
)

func main() {
	_ = os.Setenv("DEBUG", "true")
	fmt.Println("启动！")
}
