package main

import (
	"fmt"
	"encoding/json"

	query "github.com/duan-elasticsearch/duan_elasticsearch_query/v5"
)

type Userinfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Content struct {
	Artical string `json:"artical,omitempty"`
}

func main () {
	req := query.Subnode {
		Bool: &query.Bool {
			Must: &[]query.Subnode {
				query.Subnode {
					Term: Userinfo {
						Username: "heheda",
						Password: "testpwd",
					},
				},
			},
			MustNot: &[]query.Subnode {
				query.Subnode {
					Match: Content {
						Artical: "天气不错",
					},
				},
			},
		},
	}

	res, err := json.Marshal (&req)
	fmt.Println (string (res))
	fmt.Println (err)
}
