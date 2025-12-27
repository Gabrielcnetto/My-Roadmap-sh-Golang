package main

import (
	"fmt"
	"sync"
)

//testando o once, que ainda nao conhecia

type Configuration struct {
	Port    int
	IsDebug bool
}

var (
	serverConfig Configuration
	once         sync.Once //sync.once
)

func initializeServerConfig() {
	fmt.Println("Initializing server configuration...")
	serverConfig = Configuration{
		Port:    8080,
		IsDebug: false,
	}
}

// getServerConfig returns the server configuration, initializing it if not done already.
func getServerConfig() Configuration {
	once.Do(initializeServerConfig)
	return serverConfig
}

/*
* logica do once é, evitar chamar a configuracao do server mais que uma vez, tipo o redis por exemplo
 */
func main() {

	config1 := getServerConfig()
	config2 := getServerConfig()

	fmt.Println("Config1:", config1)
	fmt.Println("Config2:", config2)
}
