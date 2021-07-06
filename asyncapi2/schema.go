package asyncapi2

type Schema struct {
	SchemaUrl  string `json:"$schema"`
	ID         string `json:"$id"`
	Vocabulary struct {
		HTTPSJSONSchemaOrgDraft202012VocabCore             bool `json:"https://json-schema.org/draft/2020-12/vocab/core"`
		HTTPSJSONSchemaOrgDraft202012VocabApplicator       bool `json:"https://json-schema.org/draft/2020-12/vocab/applicator"`
		HTTPSJSONSchemaOrgDraft202012VocabUnevaluated      bool `json:"https://json-schema.org/draft/2020-12/vocab/unevaluated"`
		HTTPSJSONSchemaOrgDraft202012VocabValidation       bool `json:"https://json-schema.org/draft/2020-12/vocab/validation"`
		HTTPSJSONSchemaOrgDraft202012VocabMetaData         bool `json:"https://json-schema.org/draft/2020-12/vocab/meta-data"`
		HTTPSJSONSchemaOrgDraft202012VocabFormatAnnotation bool `json:"https://json-schema.org/draft/2020-12/vocab/format-annotation"`
		HTTPSJSONSchemaOrgDraft202012VocabContent          bool `json:"https://json-schema.org/draft/2020-12/vocab/content"`
	} `json:"$vocabulary"`
	DynamicAnchor string `json:"$dynamicAnchor"`
	Title         string `json:"title"`
	AllOf         []struct {
		Ref string `json:"$ref"`
	} `json:"allOf"`
	Type       []string `json:"type"`
	Comment    string   `json:"$comment"`
	Properties struct {
		Definitions struct {
			Comment              string `json:"$comment"`
			Type                 string `json:"type"`
			AdditionalProperties struct {
				DynamicRef string `json:"$dynamicRef"`
			} `json:"additionalProperties"`
			Deprecated bool `json:"deprecated"`
			Default    struct {
			} `json:"default"`
		} `json:"definitions"`
		Dependencies struct {
			Comment              string `json:"$comment"`
			Type                 string `json:"type"`
			AdditionalProperties struct {
				AnyOf []struct {
					DynamicRef string `json:"$dynamicRef,omitempty"`
					Ref        string `json:"$ref,omitempty"`
				} `json:"anyOf"`
			} `json:"additionalProperties"`
			Deprecated bool `json:"deprecated"`
			Default    struct {
			} `json:"default"`
		} `json:"dependencies"`
		RecursiveAnchor struct {
			Comment    string `json:"$comment"`
			Ref        string `json:"$ref"`
			Deprecated bool   `json:"deprecated"`
		} `json:"$recursiveAnchor"`
		RecursiveRef struct {
			Comment    string `json:"$comment"`
			Ref        string `json:"$ref"`
			Deprecated bool   `json:"deprecated"`
		} `json:"$recursiveRef"`
	} `json:"properties"`
	discriminator string
	externalDocs  ExternalDocs
	deprecated    bool
}
