package duan_elasticsearch_query

type Bool struct {
	Must *[]Subnode `json:"must,omitempty"`
	MustNot *[]Subnode `json:"must_not,omitempty"`
	Should *[]Subnode `json:"should,omitempty"`
}
