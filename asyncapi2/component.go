package asyncapi2

import "gopkg.in/yaml.v2"

type Components struct {
	schemas           map[string]Schema
	messages          map[string]Message
	securitySchemes   map[string]SecurityScheme
	parameters        map[string]Parameter
	correlationIds    map[string]CorrelationID
	operationTraits   map[string]OperationTrait
	messageTraits     map[string]MessageTrait
	channelBindings   map[string]ChannelBindings
	operationBindings map[string]OperationBinding
	messageBindings   map[string]MessageBinding
}

func (value *Components) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *Components) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type SecurityScheme struct {
	_type            string
	description      string
	name             string
	in               string
	scheme           string
	bearerFormat     string
	flows            OAuthFlows
	openIdConnectUrl string
}

func (value *SecurityScheme) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *SecurityScheme) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type OAuthFlows struct {
	implicit          OAuthFlow
	password          OAuthFlow
	clientCredentials OAuthFlow
	authorizationCode OAuthFlow
}

func (value *OAuthFlows) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *OAuthFlows) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}

type OAuthFlow struct {
	authorizationUrl string
	tokenUrl         string
	refreshUrl       string
	scopes           map[string]string
}

func (value *OAuthFlow) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(value)
}

func (value *OAuthFlow) UnmarshalYAML(data []byte) error {
	return yaml.Unmarshal(data, value)
}
