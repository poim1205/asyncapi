package asyncapi2

const (
	INDENT = "  "
)

type T struct {
	Asyncapi           string        // validate format #digit#.#digit#.#digit#
	ID                 string        // not required
	Info               *Info         // required
	Servers            Servers       // not required
	DefaultContentType string        // list of content type
	Channels           Channels      // required
	Components         *Components   // not required
	Tags               Tags          // not required
	ExternalDocs       *ExternalDocs // not required
}

func NewAsyncAPI() *T {
	return &T{}
}

// TODO: implement print T struct
