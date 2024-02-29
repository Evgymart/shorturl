package main

import (
	"fmt"
	"main/internal/config"
)

func main() {
	conf := config.MustLoad()
	fmt.Println(conf)
}
