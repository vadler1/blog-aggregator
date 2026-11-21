package main

import (
	"fmt"

	"github.com/vadler1/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	err = cfg.SetUser("Valentin")
	if err != nil {
		fmt.Println("Error setting user:", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading config after setting user:", err)
		return
	}

	fmt.Println("Current user:", cfg.CurrentUserName)
	fmt.Println("DB URL:", cfg.DBUrl)
}
