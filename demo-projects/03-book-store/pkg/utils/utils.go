package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// []byte(body) 将body转换为[]byte类型
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
