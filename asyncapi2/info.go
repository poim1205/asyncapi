package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Info struct {
	Title          string // required
	Version        string // required
	Description    string
	TermsOfService string
	Contact        *Contact
	License        *License
}

func NewInfo() *Info {
	return &Info{}
}

func (value *Info) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Info) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *Info) SetValues(v interface{}) *Info {

	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "title" {
				value.Title = fmt.Sprintf("%v", val)
			}
			if keyString == "version" {
				value.Version = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}
			if keyString == "termsOfService" {
				value.TermsOfService = fmt.Sprintf("%v", val)
			}
			if keyString == "contact" {
				contact := NewContact()
				value.Contact = contact.SetValues(val)
			}
			if keyString == "license" {
				license := NewLicense()
				value.License = license.SetValues(val)
			}
		}
	default:

	}
	return value
}

func (value *Info) String() string {
	return fmt.Sprintf("Info.Title : %s\n Info.Description : %s\n Info.Version : %s\n Info.TermsOfService : %s\n %s %s", value.Title, value.Description, value.Version, value.TermsOfService, value.Contact.String(), value.License.String())
}

type Contact struct {
	Name  string
	Url   string
	Email string
}

func NewContact() *Contact {
	return &Contact{}
}

func (value *Contact) SetValues(v interface{}) *Contact {

	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "name" {
				value.Name = fmt.Sprintf("%v", val)
			}
			if keyString == "url" {
				value.Url = fmt.Sprintf("%v", val)
			}
			if keyString == "email" {
				value.Email = fmt.Sprintf("%v", val)
			}
		}
	default:

	}

	return value
}

func (value *Contact) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Contact) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *Contact) String() string {
	return fmt.Sprintf("Info.Contact.Name : %s\n Info.Contact.Url : %s\n Info.Contact.Email : %s\n", value.Name, value.Url, value.Email)
}

type License struct {
	Name string
	Url  string
}

func NewLicense() *License {
	return &License{}
}

func (value *License) SetValues(v interface{}) *License {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "name" {
				value.Name = fmt.Sprintf("%v", val)
			}
			if keyString == "url" {
				value.Url = fmt.Sprintf("%v", val)
			}
		}
	default:

	}

	return value
}

func (value *License) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *License) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *License) String() string {
	return fmt.Sprintf("Info.License.Name : %s\n Info.License.Url : %s\n", value.Name, value.Url)
}
