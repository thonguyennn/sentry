package main

import (
	"bot_tele_test/logger"
	"fmt"
)

func main() {
	defer logger.SentryRecover()
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[3])
}
