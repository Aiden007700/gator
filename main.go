package main

import (
	"fmt"
	"github.com/aiden007700/gator/internal/config"
	"os"
)

func main() {
	config := readConfigAAndReturn()
	config.SetUser("Aiden")
	config = readConfigAAndReturn()

}

func readConfigAAndReturn() config.Config {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(config)
	return config
}
