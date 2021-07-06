package main

import (
	"fmt"
	"io/ioutil"

	"github.com/poim1205/asyncapi/asyncapi2"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func readYamlFile() map[string]interface{} {

	mapAsyncApi := make(map[string]interface{})

	yamlFile, err := ioutil.ReadFile("./test/info-servers-variables.yaml")
	if err != nil {
		logrus.Fatalf("Error while reading content of yaml file : #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, mapAsyncApi)
	if err != nil {
		logrus.Fatalf("Unmarshal: %v", err)
	}

	return mapAsyncApi
}

func main() {

	mapResult := readYamlFile()

	async := asyncapi2.NewAsyncAPI()

	for k, v := range mapResult {

		if k == "asyncapi" {
			fmt.Printf("Assigning AsyncAPI value %v\n", v)
			async.Asyncapi = fmt.Sprintf("%v", v)
		}

		if k == "info" {
			fmt.Printf("Assigning Info values %v\n", v)

			i := asyncapi2.NewInfo()

			async.Info = i.SetValues(v) // async.Info.Title
		}

		if k == "servers" {
			fmt.Printf("Assigning Servers values %v\n", v)

			s := asyncapi2.NewServers()
			async.Servers = s.SetValues(v)
			// map[interface {}]interface {}
			// []interface {}
		}

		if k == "channels" {
			fmt.Printf("Assigning Channels values %v\n", v)
		}

		if k == "components" {
			fmt.Printf("Assigning Components values %v\n", v)
		}
	}
	//fmt.Printf("%s", async.Info.String())

	fmt.Printf("%v", async.Servers["not"])
}
