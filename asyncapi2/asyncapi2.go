package asyncapi2

import "gopkg.in/yaml.v2"

type Identifier string

type T struct {
	Asyncapi     string       // validate format #digit#.#digit#.#digit#
	ID           Identifier   // not required
	Info         *Info        // required
	Servers      Servers      // not required
	Channels     Channels     // required
	Components   Components   // not required
	Tags         Tags         // not required
	ExternalDocs ExternalDocs // not required
}

func NewAsyncAPI() *T {
	return &T{}
}

func (doc *T) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(doc)
}

func (doc *T) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, doc)
}
