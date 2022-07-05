package main

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/config"
)

func main() {
	config, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Println("❌ Failed to load configuration:", err)
		return
	}

	fmt.Println("✅ Loaded configuration")
	fmt.Println("\nPhilips Hue bridges:")
	for i, bridge := range config.Bridges {
		fmt.Printf("%d. %s -- %s via %s\n", i+1, bridge.Name, bridge.IPAddress, bridge.Username)
	}
}
