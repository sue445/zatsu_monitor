package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	// Version represents app version (injected from ldflags)
	Version string

	// Revision represents app revision (injected from ldflags)
	Revision string
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
	fmt.Printf("zatsu_monitor v%s, build %s\n", Version, Revision)
}

func perform(name string, values map[string]string) {
	notifierType := values["type"]

	var notifier Notifier

	switch notifierType {
	case "chatwork":
		notifier = NewChatworkNotifier(values["api_token"], values["room_id"])
	case "slack":
		notifier = NewSlackNotifier(values["api_token"], values["user_name"], values["channel"])
	default:
		panic(fmt.Sprintf("Unknown type: %s in %s", notifierType, configFile))
	}

	// If it does not exist even one expected key, skip
	for _, expectedKey := range notifier.ExpectedKeys() {
		if _, ok := values[expectedKey]; !ok {
			return
		}
	}

	checkURL := values["check_url"]

	start := time.Now()
	currentStatusCode, httpError := GetStatusCode(checkURL)
	end := time.Now()
	responseTime := (end.Sub(start)).Seconds()

	fmt.Printf("time:%v\tcheck_url:%s\tstatus:%d\tresponse_time:%f\terror:%v\n", time.Now(), checkURL, currentStatusCode, responseTime, httpError)

	store := NewStatusStore(dataDir)
	beforeStatusCode, err := store.GetDbStatus(name)

	if err != nil {
		panic(err)
	}

	err = store.SaveDbStatus(name, currentStatusCode)

	if err != nil {
		panic(err)
	}

	onlyCheckOnTheOrderOf100 := false
	if values["check_only_top_of_status_code"] == "true" {
		onlyCheckOnTheOrderOf100 = true
	}

	if isNotify(beforeStatusCode, currentStatusCode, onlyCheckOnTheOrderOf100) {
		// When status code changes from the previous, notify
		param := PostStatusParam{
			CheckURL:          checkURL,
			BeforeStatusCode:  beforeStatusCode,
			CurrentStatusCode: currentStatusCode,
			HTTPError:         httpError,
			ResponseTime:      responseTime,
		}
		notifier.PostStatus(&param)
	}
}

func isNotify(beforeStatusCode int, currentStatusCode int, checkOnlyTopOfStatusCode bool) bool {
	if beforeStatusCode == NOT_FOUND_KEY {
		return false
	}

	if checkOnlyTopOfStatusCode {
		if beforeStatusCode/100 == currentStatusCode/100 {
			return false
		}

	} else {
		if beforeStatusCode == currentStatusCode {
			return false
		}
	}

	return true
}
