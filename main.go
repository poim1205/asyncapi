package main

import (
	"fmt"

	"github.com/poim1205/asyncapi/asyncapi2"
	"github.com/sirupsen/logrus"
)

func main() {
	loader := asyncapi2.NewLoader()

	async, err := loader.LoadFromFile("./test/streetlights-mqtt.yml")
	if err != nil {
		logrus.Fatalf("Error loading file %s, with the following error: %v", "./test/streetlights-mqtt.yml\n", err)
	}

	fmt.Printf("Info.description: %s\n", async.Info.Description)
}
