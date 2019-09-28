package query

import (
	"reflect"
)

type HitsNode struct {
	Index string `json:"_index,omitempty"`
	Type string `json:"_type,omitempty"`
	Id string `json:"_id,omitempty"`
	Score float64 `json:"_score,omitempty"`
	Source interface{} `json:"_source,omitempty"`
}

type QueryResponse struct {
	Took int64 `json:"took,omitempty"`
	TimedOut bool `json:"timed_out,omitempty"`
	Shards struct {
		Total int64 `json:"total,omitempty"`
		Successful int64 `json:"successful,omitempty"`
		Failed int64 `json:"failed,omitempty"`
	} `json:"_shards,omitempty"`
	Hits struct {
		Total int64 `json:"total,omitempty"`
		MaxScore float64 `json:"max_score,omitempty"`
		Hits interface{} `json:"hits,omitempty"`
	} `json:"hits,omitempty"`
}

func BuildHitsNodeSlice (demo interface{}) (res []HitsNode) {
	res = reflect.MakeSlice (reflect.TypeOf (demo), 0, 0).New ().Interface ()
}
