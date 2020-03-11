package main

import (
	"fmt"
	"reflect"
	"encoding/json"

	"github.com/duan-elasticsearch/duan_elasticsearch_query/v7"
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
    "total" : 4,
    "max_score" : 0.2876821,
    "hits" : [
      {
        "_index" : "password",
        "_type" : "content",
        "_id" : "1",
        "_score" : 0.2876821,
        "_source" : {
          "h_type" : "tar",
          "key" : "par",
          "value" : "123456789"
        }
      },
      {
        "_index" : "password",
        "_type" : "content",
        "_id" : "3",
        "_score" : 0.2876821,
        "_source" : {
          "h_type" : "tar",
          "key" : "par",
          "value" : "123456789"
        }
      },
      {
        "_index" : "password",
        "_type" : "content",
        "_id" : "2",
        "_score" : 0.18232156,
        "_source" : {
          "h_type" : "tar",
          "key" : "par",
          "value" : "123456789"
        }
      },
      {
        "_index" : "password",
        "_type" : "content",
        "_id" : "4",
        "_score" : 0.18232156,
        "_source" : {
          "h_type" : "arp",
          "key" : "dob",
          "value" : "123456789"
        }
      }
    ]
  }
}
`

type PasswdCrackType struct {
	HType string `json:"h_type,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func main () {
	resObj := query.GetResponseObj (reflect.TypeOf (PasswdCrackType{}))

	if err := json.Unmarshal ([]byte (test_res), resObj); err != nil {
		panic (err)
	}

	fmt.Println (resObj)
}
