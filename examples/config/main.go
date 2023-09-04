package main

import (
	"fmt"
	"ssikr/config"
)

func main() {
	fmt.Println("Config registrar address: ", config.SystemConfig.RegistrarAddr)
	fmt.Println("Config resolver address: ", config.SystemConfig.ResolverAddr)
}
