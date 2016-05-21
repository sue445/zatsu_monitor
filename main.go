package main

import (
	"flag"
	"fmt"
	"log"
)

var configFile, dataDir string
var isPrintVersion bool

func init() {
	flag.StringVar(&configFile, "config", "", "Path to config file")
	flag.StringVar(&dataDir, "data", "", "Path to data dir for cache")
	flag.BoolVar(&isPrintVersion, "version", false, "Whether showing version")

	flag.Parse()
}

func main() {
	if isPrintVersion {
		printVersion()
		return
	}

	if len(configFile) == 0 || len(dataDir) == 0 {
		flag.PrintDefaults()
		return
	}

	config, err := LoadConfigFromFile(configFile)

	if err != nil {
		panic(err)
	}

	for name, values := range config {
		perform(name, values)
	}
}

func printVersion() {
	fmt.Printf("zatsu_monitor v%s\n", VERSION)
}

func perform(name string, values map[string]string) {
	notifier_type := values["type"]

	var notifier Notifier

	switch notifier_type {
	case "chatwork":
		notifier = NewChatworkNotifier(values["api_token"], values["room_id"])
	case "slack":
		notifier = NewSlackNotifier(values["api_token"], values["user_name"], values["channel"])
	default:
		panic(fmt.Sprintf("Unknown type: %s in %s", notifier_type, configFile))
	}

	// If it does not exist even one expected key, skip
	for _, expectedKey := range notifier.ExpectedKeys() {
		if _, ok := values[expectedKey]; !ok {
			return
		}
	}

	checkUrl := values["check_url"]
	currentStatusCode, err := HttpStatusCode(checkUrl)
	log.Printf("%s [status %d]\n", checkUrl, currentStatusCode)

	if err != nil {
		panic(err)
	}

	z := NewZatsuMonitor(dataDir)
	beforeStatusCode, err := z.GetDbStatus(name)

	if err != nil {
		panic(err)
	}

	err = z.SaveDbStatus(name, currentStatusCode)

	if err != nil {
		panic(err)
	}

	if beforeStatusCode > 0 && currentStatusCode > 0 && beforeStatusCode != currentStatusCode {
		// When status code changes from the previous, notify
		notifier.PostStatus(checkUrl, beforeStatusCode, currentStatusCode)
	}
}
