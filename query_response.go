package query

import (
	// "fmt"
	"reflect"
	"github.com/goinggo/mapstructure"
)

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
		Hits []struct {
			Index string `json:"_index,omitempty"`
			Type string `json:"_type,omitempty"`
			Id string `json:"_id,omitempty"`
			Score float64 `json:"_score,omitempty"`
			Source interface{} `json:"_source,omitempty"`
		} `json:"hits,omitempty"`
	} `json:"hits,omitempty"`
}

func (self *QueryResponse) CoverSource (demoType reflect.Type) {
	for numDemo, oneDemo := range self.Hits.Hits {
		tmpObj := reflect.New (demoType).Interface ()
		switch oneDemo.Source.(type) {

		case map[string]interface{}:
			mapstructure.Decode (oneDemo.Source, tmpObj)
			self.Hits.Hits[numDemo].Source = tmpObj
		}
	}
}
