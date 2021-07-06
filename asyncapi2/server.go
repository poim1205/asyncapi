package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Servers map[string]*Server

func NewServers() Servers {
	return make(Servers, 1)
}

func (s Servers) SetValues(v interface{}) Servers {
	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			_, Ok := s[keyString]
			if !Ok {
				ns := NewServer()

				s[keyString] = ns.SetValues(val)
			}
		}
	case []interface{}:
		for _, val := range arrayVal {
			ns := NewServer()
			s["not"] = ns.SetValues(val)
		}
	default:
	}
	return s
}

type Server struct {
	Url             string // required
	Protocol        string // required
	ProtocolVersion string
	Description     string
	Variables       map[string]*ServerVariable
	Security        []*SecurityRequirement
	Bindings        ServerBindings
}

func NewServer() *Server {
	return &Server{}
}

func (value *Server) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Server) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

func (value *Server) SetValues(v interface{}) *Server {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "url" {
				value.Url = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}
			if keyString == "protocol" {
				value.Protocol = fmt.Sprintf("%v", val)
			}
			if keyString == "protocolVersion" {
				value.ProtocolVersion = fmt.Sprintf("%v", val)
			}
		}
	default:
	}
	return value
}

type ServerVariable struct {
	Enum        []string
	Default     string
	Description string
	//Exemples    []string
}

func NewServerVariable() *ServerVariable {
	return &ServerVariable{}
}

func (value *ServerVariable) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *ServerVariable) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type SecurityRequirement struct {
	Name map[string][]string
}

func (value *SecurityRequirement) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *SecurityRequirement) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type ServerBindings map[string]ServerBinding

type ServerBinding struct {
}
