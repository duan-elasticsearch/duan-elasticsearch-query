package main

import (
	"fmt"
	"encoding/json"

	"github.com/duan-elasticsearch/duan_elasticsearch/v5/query"
)

type Userinfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func main () {
	req := query.Query {
		Term: &Userinfo {
			Username: "heheda",
			Password: "testpsd",
		},
	}

	res, err := json.Marshal (&req)
	fmt.Println (string (res))
	fmt.Println (err)
}
