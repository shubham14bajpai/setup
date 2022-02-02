package main

import (
	"fmt"

	config "github.com/shubham14bajpai/setup/pkg"
)

func main() {
	config.Install()
	fmt.Println("hello!")
}
