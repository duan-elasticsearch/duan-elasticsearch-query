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

func main () {
	req := query.Subnode {
		Term: &Userinfo {
			Username: "heheda",
			Password: "testpsd",
		},
	}

	res, err := json.Marshal (&req)
	fmt.Println (string (res))
	fmt.Println (err)
}
