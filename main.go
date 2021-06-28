package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Identifier string

type global struct {
	asyncapi     string                // validate format #digit#.#digit#.#digit#
	id           Identifier            // not required
	info         Info                  // required
	servers      Servers               // not required
	channels     Channels              // required
	components   Components            // not required
	tags         Tags                  // not required
	externalDocs ExternalDocumentation // not required
}

type Info struct {
	title          string // required
	version        string // required
	description    string
	termsOfService string
	contact        Contact
	license        Licence
}

type Servers struct {
	servers map[string]Server
}

type Server struct {
	url             string // required
	protocol        string // required
	protocolVersion string
	description     string
	variables       map[string]ServerVariable
	security        []SecurityRequirement
	bindings        ServerBindings
}

type ServerVariable struct {
	enum        []string
	myDefault   string
	description string
	exemples    []string
}

type SecurityRequirement struct {
	name map[string][]string
}

type ServerBindings struct {
}

type Channels struct {
	channels map[string]Channel
}

type Channel struct {
	ref         string
	description string
	subscribe   Operation
	publish     Operation
	parameters  Parameters
	bindings    ChannelBindings
}

type Components struct {
}

type Tags struct {
}

type ExternalDocumentation struct {
}

type Contact struct {
	name  string
	url   string
	email string
}

type Licence struct {
	name string
	url  string
}

func readYamlFile() map[string]interface{} {

	mapAsyncApi := make(map[string]interface{})

	yamlFile, err := ioutil.ReadFile("./test/simple.yaml")
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

	fmt.Println("First level items")

	for k := range mapResult {
		fmt.Println(k)
	}

}
