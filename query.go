package query

type Query struct {
	Bool        *Bool       `json:"bool,omitempty"`
	Term        interface{} `json:"term,omitempty"`
	Terms       interface{} `json:"terms,omitempty"`
	Match       interface{} `json:"match,omitempty"`
	MatchPhrase interface{} `json:"match_phrase,omitempty"`
	MatchAll    interface{} `json:"match_all,omitempty"`
	Wildcard    interface{} `json:"wildcard,omitempty"`
	Exists      interface{} `json:"exists,omitempty"`
	Regexp      interface{} `json:"regexp,omitempty"`
	Nested      interface{} `json:"nested,omitempty"`

	Range interface{} `json:"range,omitempty"`

	SimpleQueryString *QueryString `json:"simple_query_string,omitempty"`
	QueryString       *QueryString `json:"query_string,omitempty"`
}
