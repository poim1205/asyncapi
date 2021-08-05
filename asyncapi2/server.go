package asyncapi2

import (
	"fmt"
)

type Servers map[string]*Server

func NewServers() Servers {
	return make(Servers)
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
	default:
	}
	return s
}

type Server struct {
	Url             string `validate:"required,url"`
	Protocol        string `validate:"required"`
	ProtocolVersion string
	Description     string
	Variables       map[string]*ServerVariable
	Security        map[string][]string
	Bindings        ServerBindings
}

func NewServer() *Server {
	return &Server{}
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
					msv := make(map[string]*ServerVariable)
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
			if keyString == "security" {
				switch securityVal := val.(type) {
				case map[interface{}]interface{}:
					mapSecVal := make(map[string][]string)
					for varKey, varVal := range securityVal {
						keySerVar := fmt.Sprintf("%v", varKey)

						_, Ok := mapSecVal[keySerVar]
						if !Ok {
							sliceListSec := make([]string, 0)
							switch listSec := varVal.(type) {
							case []interface{}:
								for _, strSecVal := range listSec {
									sliceListSec = append(sliceListSec, fmt.Sprintf("%v", strSecVal))
								}
							default:
							}
							mapSecVal[keySerVar] = sliceListSec
						}
					}
					value.Security = mapSecVal
				default:
				}
			}
			if keyString == "bindings" {
				newServerBindings := NewServerBindings()

				value.Bindings = newServerBindings.SetValues(val)
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
