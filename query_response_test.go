package query

import (
	"fmt"
	"testing"
	"reflect"
	"encoding/json"
)

var test_res = `
{
  "took" : 7459,
  "timed_out" : false,
  "num_reduce_phases" : 5,
  "_shards" : {
    "total" : 1249,
    "successful" : 1248,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 10000,
      "relation" : "gte"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "000ab9f4d33e47cabc79c818f2d0b1b4760ac4ba9a0a8cc90b4f709fd31d1b1c3192e1290920629b96c18eb847432d05c9b512ed751f4a889f5582b37eb008ee_1",
        "_type" : "_doc",
        "_id" : "8ce3ea0ce80e8c4fa7a22ec48daaffdd6fbe943bffff340f8b2aff2e3085019d80e9cc0d98c115d1c55a107403e540ded08e37aaeaa71675703ac060fc83a91d",
        "_score" : 1.0,
        "_source" : {
          "file_path" : "/mnt/olddat",
          "content" : "aaaaaaaaa"
        }
      },
      {
        "_index" : "000ab9f4d33e47cabc79c818f2d0b1b4760ac4ba9a0a8cc90b4f709fd31d1b1c3192e1290920629b96c18eb847432d05c9b512ed751f4a889f5582b37eb008ee_1",
        "_type" : "_doc",
        "_id" : "b045d732e87e3e391f71340d0e0083625cb3d4b5303c20276ea1c16ab31b0ca4766a743ba82a9cf45e551517fb30baf20ca0a71cb6b53858d259734067c55033",
        "_score" : 1.0,
        "_source" : {
          "file_path" : "/mnt/olddat",
          "content" : "bbbbbbbbbb"
        }
      },
      {
        "_index" : "000ab9f4d33e47cabc79c818f2d0b1b4760ac4ba9a0a8cc90b4f709fd31d1b1c3192e1290920629b96c18eb847432d05c9b512ed751f4a889f5582b37eb008ee_1",
        "_type" : "_doc",
        "_id" : "2475f17054ca1b7ee2f3c40ce6f410cfbd571ae3b8efcc334fcdc782078ee4580df013c9990bbc61073fbb74e43c20f5dc034198a8d5b595408d9f285b7f1031",
        "_score" : 1.0,
        "_source" : {
          "file_path" : "/mnt/olddat",
          "content" : "ccccccccccc"
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

func TestBuildHitsNodeSlice (t *testing.T) {
	resObj := GetResponseObj (reflect.TypeOf (PasswdCrackType{}))

	if err := json.Unmarshal ([]byte (test_res), resObj); err != nil {
		panic (err)
	}

	fmt.Println (resObj)
}
