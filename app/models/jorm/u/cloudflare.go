package u

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oknors/okno/app/models/jorm/c"
	"github.com/oknors/okno/app/models/jorm/cfg"
)

func CloudFlare() {

	coins := c.ReadAllCoins()
	for _, coin := range coins.C {
		slug := coin.Slug

		baseUrl := fmt.Sprintf(cfg.Credentials.CloudFlareAPI)
		client := new(http.Client)
		req, err := http.NewRequest("POST", baseUrl, nil)
		if err != nil {
		}
		//req.SetBasicAuth(c.User, c.Password)
		req.Header.Add("X-Auth-Email", cfg.Credentials.CloudFlareEmail)
		req.Header.Add("X-Auth-Key", cfg.Credentials.CloudFlareAPIkey)
		req.Header.Add("Content-Type", "application/json")
		args := make(map[string]interface{})
		//	args["data"] = "{\"type\":\"CNAME\",\"name\":" + slug + " ,\"content\":\"com-http.us\",\"ttl\":1,\"proxied\":true}"
		args["type"] = "CNAME"
		args["name"] = slug
		args["content"] = "com-http.us"
		args["ttl"] = 1
		args["proxied"] = true

		j, err := json.Marshal(args)
		if err != nil {
			fmt.Println(err)
		}
		//	fmt.Println("Blooblockblockblockblockooradb", args)
		req.Body = ioutil.NopCloser(strings.NewReader(string(j)))
		req.ContentLength = int64(len(string(j)))
		resp, err := client.Do(req)
		if err != nil {
			//return "", err
		}
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)

		var data map[string]interface{}
		json.Unmarshal(bytes, &data)
		if err, found := data["errors"]; found && err != nil {
			str, _ := json.Marshal(err)
			//return nil, errors.New(string(str))
			fmt.Println("Blooblockblockblockblockooradb", str)

		}

		if result, found := data["success"]; found {
			//return result, nil
			var dsata map[string]interface{}
			json.Unmarshal(bytes, &dsata)
			fmt.Println("Blooblockblockblockblockooradb", result)
			fmt.Println("BlooblockblockbdsatadsataSSSdsata", dsata)

		} else {
			//return nil, errors.New("no result")
		}
	}

}
