package google

type SearchRequest struct {
	WebSite string `from:"query" json:"site"`
	DocType string `from:"query" json:"type"`
	Count   string `from:"query" json:"count"`
}

func (SearchRequest) Struct() interface{} {
	return &SearchRequest{}
}

func (s *SearchRequest) Validate() error {
	return nil
}
