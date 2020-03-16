package query

import (
	// "fmt"
	"reflect"
	"encoding/json"
)

type QueryResponseHitsOne struct {
	Index string `json:"_index,omitempty" validate:"required"`
	Type string `json:"_type,omitempty" validate:"required"`
	Id string `json:"_id,omitempty" validate:"required"`
	Score float64 `json:"_score,omitempty"`
	Source json.RawMessage `json:"_source,omitempty"`
	SourceObj interface{} `json:"_source_obj,omitempty"`
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
		Total struct {
			Value int64 `json:"value,omitempty"`
			Relation string `json:"relation,omitempty"`
		} `json:"total,omitempty"`
		MaxScore float64 `json:"max_score,omitempty"`
		Hits []QueryResponseHitsOne `json:"hits,omitempty"`
	} `json:"hits,omitempty"`
}

func (self *QueryResponse) CoverSource (demoType reflect.Type) (err error) {
	for oneDataNum, oneData := range self.Hits.Hits {
		self.Hits.Hits[oneDataNum].SourceObj = reflect.New (demoType).Interface ()
		if err = json.Unmarshal (oneData.Source, self.Hits.Hits[oneDataNum].SourceObj); err != nil {
			return
		}
		self.Hits.Hits[oneDataNum].Source = []byte{};
	}

	return
}

func (self *QueryResponseHitsOne) CoverHitsOne (demoType reflect.Type) (err error) {
	self.SourceObj = reflect.New (demoType).Interface ()
	if err = json.Unmarshal (self.Source, self.SourceObj); err != nil {
		return
	}
	self.Source = []byte{};

	return
}
