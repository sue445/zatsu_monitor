package main

import (
	"flag"
	"fmt"
)

func main() {
	configFile := *flag.String("config", "", "Path to config file")
	dataDir := *flag.String("data", "", "Path to data dir for cache")
	configName := *flag.String("name", "", "Name for url checking (default: all)")
	flag.Parse()

	fmt.Printf("configFile=%v\n", configFile)
	fmt.Printf("dataDir=%v\n", dataDir)
	fmt.Printf("configName=%v\n", configName)
}
