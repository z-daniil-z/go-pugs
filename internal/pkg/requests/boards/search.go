package boards

type AutoRequest struct {
	Mark  string `from:"query" json:"mark"`
	Model string `from:"query" json:"model"`
	Page  string `from:"query" json:"page"`
}

func (AutoRequest) Struct() interface{} {
	return &AutoRequest{}
}

func (s *AutoRequest) Validate() error {
	return nil
}
