package asyncapi2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Servers interface {
	PrintServers()
}

type MapServers map[string]*Server

type SliceServers []*Server

func NewServers(v interface{}) Servers {
	var s Servers
	switch arrayVal := v.(type) {
	case map[interface{}]interface{}:
		ms := make(MapServers, 1)
		for key, val := range arrayVal {
			keyString := fmt.Sprintf("%v", key)
			_, Ok := ms[keyString]
			if !Ok {
				ns := NewServer()

				ms[keyString] = ns.SetValues(val)
			}
		}
		s = ms
	case []interface{}:
		ss := make(SliceServers, 0, 1)
		for _, val := range arrayVal {
			ns := NewServer()
			ss = append(ss, ns.SetValues(val))
		}
		s = ss
	default:
	}
	return s
}

func (ms MapServers) PrintServers() {
	fmt.Println("Printing mapservers")
	for k, v := range ms {
		fmt.Printf("Servers.%s.url : %s\n", k, v.Url)
		fmt.Printf("Servers.%s.description : %s\n", k, v.Description)
		fmt.Printf("Servers.%s.protocol : %s\n", k, v.Protocol)
		fmt.Printf("Servers.%s.protocolVersion : %s\n", k, v.ProtocolVersion)
	}
}

func (ss SliceServers) PrintServers() {
	fmt.Println("Printing sliceservers")

	for _, v := range ss {
		fmt.Printf("Servers.url : %s\n", v.Url)
		fmt.Printf("Servers.description : %s\n", v.Description)
		fmt.Printf("Servers.protocol : %s\n", v.Protocol)
		fmt.Printf("Servers.protocolVersion : %s\n", v.ProtocolVersion)
		for svk, svv := range v.Variables {
			for i, enu := range svv.Enum {
				fmt.Printf("Servers.variables.%s.enum::val%d : %s\n", svk, i+1, enu)
			}
			fmt.Printf("Servers.variables.%s.default : %s\n", svk, svv.Default)
			fmt.Printf("Servers.variables.%s.description : %s\n", svk, svv.Description)
		}
	}
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
			if keyString == "variables" {
				switch variablesVal := val.(type) {
				case map[interface{}]interface{}:
					msv := make(map[string]*ServerVariable, 1)
					for varKey, varVal := range variablesVal {
						keySerVar := fmt.Sprintf("%v", varKey)
						nsv := NewServerVariable()

						_, Ok := msv[keySerVar]
						if !Ok {
							msv[keySerVar] = nsv.SetValues(varVal)
						}
					}
					value.Variables = msv
				default:
				}
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

func (value *ServerVariable) SetValues(v interface{}) *ServerVariable {
	switch mapVal := v.(type) {
	case map[interface{}]interface{}:

		for key, val := range mapVal {
			keyString := fmt.Sprintf("%v", key)
			if keyString == "enum" {
				switch valSlice := val.(type) {
				case []interface{}:
					sliceEnum := make([]string, 0, 1)
					for _, valEnum := range valSlice {
						sliceEnum = append(sliceEnum, fmt.Sprintf("%v", valEnum))
					}
					value.Enum = sliceEnum
				default:
				}
			}
			if keyString == "default" {
				value.Default = fmt.Sprintf("%v", val)
			}
			if keyString == "description" {
				value.Description = fmt.Sprintf("%v", val)
			}
		}
	default:
	}

	return value
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
