package asyncapi2

import "fmt"

type Components struct {
	Schemas           map[string]*Schema
	Messages          map[string]*Message
	SecuritySchemes   map[string]*SecurityScheme
	Parameters        map[string]*Parameter
	CorrelationIds    map[string]*CorrelationID
	OperationTraits   map[string]*OperationTrait
	MessageTraits     map[string]*MessageTrait
	OperationBindings OperationBindings
	MessageBindings   MessageBindings
}

func NewComponents() *Components {
	return &Components{}
}

func (c *Components) SetValues(v interface{}) *Components {
	switch mapComponents := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapComponents {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "schemas" {
				switch mapSchemas := val.(type) {
				case map[interface{}]interface{}:
					schemaMap := make(map[string]*Schema, 0)
					for key, mapval := range mapSchemas {
						keySchemaStr := fmt.Sprintf("%v", key)
						_, Ok := schemaMap[keySchemaStr]
						if !Ok {
							sch := NewSchema()
							schemaMap[keySchemaStr] = sch.SetValues(mapval)
						}
					}
					c.Schemas = schemaMap
				default:
				}
			}
			if keyStr == "messages" {
				switch mapMessages := val.(type) {
				case map[interface{}]interface{}:
					messMap := make(map[string]*Message, 0)
					for key, mapval := range mapMessages {
						keyMessageStr := fmt.Sprintf("%v", key)
						_, Ok := messMap[keyMessageStr]
						if !Ok {
							mess := NewMessage()
							messMap[keyMessageStr] = mess.SetValues(mapval)
						}
					}
					c.Messages = messMap
				default:
				}
			}
			// TODO: fix security schemes
			// if keyStr == "securitySchemes" {
			// 	switch mapSecuritySchemes := val.(type) {
			// 	case map[interface{}]interface{}:
			// 		secSchemeMap := make(map[string]*SecurityScheme)
			// 		for key, mapval := range mapSecuritySchemes {
			// 			keySecSchemeStr := fmt.Sprintf("%v", key)
			// 			_, Ok := secSchemeMap[keySecSchemeStr]
			// 			if !Ok {
			// 				secScheme := NewSecurityScheme()
			// 				secSchemeMap[keySecSchemeStr] = secScheme.SetValues(mapval)
			// 			}
			// 		}
			// 		c.SecuritySchemes = secSchemeMap
			// 	default:
			// 	}
			// }
			if keyStr == "parameters" {
				switch mapParameters := val.(type) {
				case map[interface{}]interface{}:
					paramMap := make(map[string]*Parameter)
					for key, mapval := range mapParameters {
						keyParameterStr := fmt.Sprintf("%v", key)
						_, Ok := paramMap[keyParameterStr]
						if !Ok {
							paramMap[keyParameterStr] = NewParameter(mapval)
						}
					}
					c.Parameters = paramMap
				default:
				}
			}
			if keyStr == "correlationIds" {
				switch mapCorrelIds := val.(type) {
				case map[interface{}]interface{}:
					correlMap := make(map[string]*CorrelationID)
					for key, mapval := range mapCorrelIds {
						keyCorrelIdStr := fmt.Sprintf("%v", key)
						_, Ok := correlMap[keyCorrelIdStr]
						if !Ok {
							corellId := NewCorrelationId()
							correlMap[keyCorrelIdStr] = corellId.SetValues(mapval)
						}
					}
					c.CorrelationIds = correlMap
				default:
				}
			}
			if keyStr == "operationTraits" {
				switch mapOperationTraits := val.(type) {
				case map[interface{}]interface{}:
					opTraitsMap := make(map[string]*OperationTrait)
					for key, mapval := range mapOperationTraits {
						keyOperationTraitsStr := fmt.Sprintf("%v", key)
						_, Ok := opTraitsMap[keyOperationTraitsStr]
						if !Ok {
							operationTrait := NewOperationTrait()
							opTraitsMap[keyOperationTraitsStr] = operationTrait.SetValues(mapval)
						}
					}

					c.OperationTraits = opTraitsMap
				default:
				}
			}
			if keyStr == "messageTraits" {
				switch mapMessageTraits := val.(type) {
				case map[interface{}]interface{}:
					messTraitsMap := make(map[string]*MessageTrait)
					for key, mapval := range mapMessageTraits {
						keyMessageTraitsStr := fmt.Sprintf("%v", key)
						_, Ok := messTraitsMap[keyMessageTraitsStr]
						if !Ok {
							messageTrait := NewMessageTrait()
							messTraitsMap[keyMessageTraitsStr] = messageTrait.SetValues(mapval)
						}
					}
					c.MessageTraits = messTraitsMap
				default:
				}
			}
			if keyStr == "operationBindings" {
				newOperationBindings := NewOperationBindings()
				c.OperationBindings = newOperationBindings.SetValues(val)
			}
			if keyStr == "messageBindings" {
				newMessBindings := NewMessageBindings()
				c.MessageBindings = newMessBindings.SetValues(val)
			}
		}

	default:
	}
	return c
}

type SecurityScheme struct {
	Type             string
	Description      string
	Name             string
	In               string
	Scheme           string
	BearerFormat     string
	Flows            *OAuthFlows
	OpenIdConnectUrl string
}

func NewSecurityScheme() *SecurityScheme {
	return &SecurityScheme{}
}

func (ss *SecurityScheme) SetValues(v interface{}) *SecurityScheme {
	switch mapOpBind := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapOpBind {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "type" {
				ss.Type = fmt.Sprintf("%v", val)
			}
			if keyStr == "description" {
				ss.Description = fmt.Sprintf("%v", val)
			}
			if keyStr == "name" {
				ss.Name = fmt.Sprintf("%v", val)
			}
			if keyStr == "in" {
				ss.In = fmt.Sprintf("%v", val)
			}
			if keyStr == "scheme" {
				ss.Scheme = fmt.Sprintf("%v", val)
			}
			if keyStr == "bearerFormat" {
				ss.BearerFormat = fmt.Sprintf("%v", val)
			}
			if keyStr == "flows" {
				newFlows := NewOAuthFlows()
				ss.Flows = newFlows.SetValues(val)
			}
			if keyStr == "openIdConnectUrl" {
				ss.OpenIdConnectUrl = fmt.Sprintf("%v", val)
			}
		}

	default:
	}

	return ss
}

type OAuthFlows struct {
	Implicit          *OAuthFlow
	Password          *OAuthFlow
	ClientCredentials *OAuthFlow
	AuthorizationCode *OAuthFlow
}

func NewOAuthFlows() *OAuthFlows {
	return &OAuthFlows{}
}

func (flows *OAuthFlows) SetValues(v interface{}) *OAuthFlows {
	switch mapFlows := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapFlows {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "implicit" {
				newFlow := NewOAuthFlow()
				flows.Implicit = newFlow.SetValues(val)
			}
			if keyStr == "password" {
				newFlow := NewOAuthFlow()
				flows.Password = newFlow.SetValues(val)
			}
			if keyStr == "clientCredentials" {
				newFlow := NewOAuthFlow()
				flows.ClientCredentials = newFlow.SetValues(val)
			}
			if keyStr == "authorizationCode" {
				newFlow := NewOAuthFlow()
				flows.AuthorizationCode = newFlow.SetValues(val)
			}
		}
	default:
	}
	return flows
}

type OAuthFlow struct {
	AuthorizationUrl string
	TokenUrl         string
	RefreshUrl       string
	Scopes           map[string]string
}

func NewOAuthFlow() *OAuthFlow {
	return &OAuthFlow{}
}

func (flow *OAuthFlow) SetValues(v interface{}) *OAuthFlow {
	switch mapFlow := v.(type) {
	case map[interface{}]interface{}:
		for key, val := range mapFlow {
			keyStr := fmt.Sprintf("%v", key)
			if keyStr == "authorizationUrl" {
				flow.AuthorizationUrl = fmt.Sprintf("%v", val)
			}
			if keyStr == "tokenUrl" {
				flow.TokenUrl = fmt.Sprintf("%v", val)
			}
			if keyStr == "refreshUrl" {
				flow.RefreshUrl = fmt.Sprintf("%v", val)
			}
			if keyStr == "scopes" {
				switch mapScope := val.(type) {
				case map[interface{}]interface{}:
					for keyScope, valScope := range mapScope {
						keyScopeStr := fmt.Sprintf("%v", keyScope)

						_, Ok := flow.Scopes[keyScopeStr]

						if !Ok {
							flow.Scopes[keyScopeStr] = fmt.Sprintf("%v", valScope)
						}
					}
				}
			}
		}

	default:
	}
	return flow
}
