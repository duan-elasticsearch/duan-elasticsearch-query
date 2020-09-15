package query

type QueryString struct {
	Query string `json:"query,omitempty"`
	Fields []string `json:"fields,omitempty"`
	DefaultOperator string `json:"default_operatior,omitempty"`
}
