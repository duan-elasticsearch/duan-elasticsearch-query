package query

type Query struct {
	Bool *Bool `json:"bool,omitempty"`
	Term interface{} `json:"term,omitempty"`
	Terms interface{} `json:"terms,omitempty"`
	Match interface{} `json:"match,omitempty"`
	MatchPhrase interface{} `json:"match_phrase,omitempty"`
	MatchAll interface{} `json:"match_all,omitempty"`
	Wildcard interface{} `json:"wildcard,omitempty"`
	Regexp interface{} `json:"regexp,omitempty"`

	Range interface{} `json:"range,omitempty"`

	SimpleQueryString struct {
		Query string `json:"query,omitempty"`
		Fields []string `json:"fields,omitempty"`
		DefaultOperator string `json:"default_operatior,omitempty"`
	} `json:"simple_query_string"`
	QueryString struct {
		Query string `json:"query,omitempty"`
		Fields []string `json:"fields,omitempty"`
		DefaultOperator string `json:"default_operatior,omitempty"`
	} `json:"query_string"`
}
