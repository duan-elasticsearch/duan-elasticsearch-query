package query

import (
	// "fmt"
	"reflect"
)

type QueryResponseHitsHitsDemoStruct struct {
	Index string `json:"_index,omitempty"`
	Type string `json:"_type,omitempty"`
	Id string `json:"_id,omitempty"`
	Score float64 `json:"_score,omitempty"`
	Source interface{} `json:"_source,omitempty"`
}

type QueryResponseHitsDemoStruct struct {
	Total struct {
		Value int64 `json:"value,omitempty"`
		Relation string `json:"relation,omitempty"`
	}
	MaxScore float64 `json:"max_score,omitempty"`
	Hits interface {} `json:"hits,omitempty"`
}

type QueryResponse struct {
	Took int64 `json:"took,omitempty"`
	TimedOut bool `json:"timed_out,omitempty"`
	Shards struct {
		Total int64 `json:"total,omitempty"`
		Successful int64 `json:"successful,omitempty"`
		Failed int64 `json:"failed,omitempty"`
	} `json:"_shards,omitempty"`
	Hits QueryResponseHitsDemoStruct `json:"hits,omitempty"`
}

func GetResponseObj (demoType reflect.Type) (res *QueryResponse) {
	tmpHitsDemoObj := QueryResponseHitsHitsDemoStruct {
		Source: reflect.New (demoType).Interface (),
	}
	tmpHitsDemoObjValue := reflect.ValueOf (tmpHitsDemoObj)

	tmpHitsDemoSliceType := reflect.SliceOf (tmpHitsDemoObjValue.Type ())

	tmpHitsSliceValue := reflect.MakeSlice (tmpHitsDemoSliceType, 0, 0)

	res = &QueryResponse {
		Hits: QueryResponseHitsDemoStruct {
			Hits: tmpHitsSliceValue.Interface (),
		},
	}

	return
}
