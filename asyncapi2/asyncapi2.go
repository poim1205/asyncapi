package asyncapi2

type T struct {
	Asyncapi           string `validate:"required,version"`
	ID                 string
	Info               *Info `validate:"required"`
	Servers            Servers
	DefaultContentType string
	Channels           Channels `validate:"required"`
	Components         *Components
	Tags               Tags
	ExternalDocs       *ExternalDocs
}

func NewAsyncAPI() *T {
	return &T{}
}

func (t *T) Validate() error {
	return ValidateStruct(t)
}
