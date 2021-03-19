package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type JsonRpc struct {
	User     string
	Password string
	Host     string
	Port     int64
}

func (c *JsonRpc) MakeRequest(method string, params interface{}) (interface{}, error) {
	baseUrl := fmt.Sprintf("http://%s:%d", c.Host, c.Port)
	client := new(http.Client)
	req, err := http.NewRequest("POST", baseUrl, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(c.User, c.Password)
	req.Header.Add("Content-Type", "text/plain")

	args := make(map[string]interface{})
	args["jsonrpc"] = "1.0"
	args["id"] = "BitNodes"
	args["method"] = method
	args["params"] = params

	j, err := json.Marshal(args)
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println("Blooblockblockblockblockooradb", args)

	req.Body = ioutil.NopCloser(strings.NewReader(string(j)))
	req.ContentLength = int64(len(string(j)))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(bytes, &data)
	if err, found := data["error"]; found && err != nil {
		str, _ := json.Marshal(err)
		return "", errors.New(string(str))
	}

	if result, found := data["result"]; found {
		return result.(interface{}), nil
	} else {
		return "", errors.New("no result")
	}
}
func NewClient(user string, password string, host string, port int64) *JsonRpc {
	c := JsonRpc{user, password, host, port}
	return &c
}
