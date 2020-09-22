package validation

type Request struct {
	WebSite string `from:"query" json:"site"`
	DocType string `from:"query" json:"type"`
	Count   string `from:"query" json:"count"`
}

func (Request) Struct() interface{} {
	return &Request{}
}

func (s *Request) Validate() error {
	return nil
}
