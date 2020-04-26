package query

import (
	"fmt"
	"testing"
	"reflect"
	"encoding/json"
)

var test_res = `{
  "took" : 2,
  "timed_out" : false,
  "_shards" : {
    "total" : 5,
    "successful" : 5,
    "failed" : 0
  },
  "hits" : {
    "total" : 77,
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1149938446",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548225148"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1137152439",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548215545"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1061707874",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548226144"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1236310072",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548219129"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1326588821",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548229190"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1146738734",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548216244"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1359190838",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548176016"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1149066472",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1545716748"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1220111865",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548223187"
        }
      },
      {
        "_index" : "group_data",
        "_type" : "telegram",
        "_id" : "1342029169",
        "_score" : 1.0,
        "_source" : {
          "create_time" : "1548174694"
        }
      }
    ]
  }
}
`

type PasswdCrackType struct {
	CreateTime string `json:"create_time,omitempty"`
}

func TestBuildHitsNodeSlice (t *testing.T) {
	resObj := QueryResponse {}

	if err := json.Unmarshal ([]byte (test_res), &resObj); err != nil {
		panic (err)
	}

	fmt.Println (resObj)
	resObj.CoverSource (reflect.TypeOf (PasswdCrackType{}))
	fmt.Println (resObj)
}
