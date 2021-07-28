package asyncapi2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadFromFile(fullFilePath string) (*T, error) {

	byteSlice, err := ioutil.ReadFile(fullFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error loading the file: %s", fullFilePath)
	}

	mapAsyncApi := make(map[string]interface{})
	err = yaml.Unmarshal(byteSlice, mapAsyncApi)
	if err != nil {
		err = json.Unmarshal(byteSlice, &mapAsyncApi)
		if err != nil {
			return nil, fmt.Errorf("Error un-marshalling the file: %s", fullFilePath)
		}
	}

	async := loadAsyncApiData(mapAsyncApi)

	return async, nil
}

func loadAsyncApiData(mapAsyncApi map[string]interface{}) *T {

	async := NewAsyncAPI()

	for k, v := range mapAsyncApi {

		if k == "asyncapi" {
			async.Asyncapi = fmt.Sprintf("%v", v)
		}

		if k == "info" {
			i := NewInfo()
			async.Info = i.SetValues(v)
		}

		if k == "servers" {
			s := NewServers()
			async.Servers = s.SetValues(v)
		}

		if k == "defaultContentType" {
			async.DefaultContentType = fmt.Sprintf("%v", v)
		}

		if k == "channels" {
			c := NewChannels()
			async.Channels = c.SetValues(v)
		}

		if k == "components" {
			comp := NewComponents()
			async.Components = comp.SetValues(v)
		}

		if k == "tags" {
			tags := NewTags()
			async.Tags = tags.SetValues(v)
		}

		if k == "externalDocs" {
			newExternalDocs := NewExternalDocs()
			async.ExternalDocs = newExternalDocs.SetValues(v)
		}
	}

	return async
}
