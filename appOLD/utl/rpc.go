package utl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Endpoint is the contact details for a JSONRPC endpoint
type Endpoint struct {
	User     string
	Password string
	Host     string
	Port     int64
}

// MakeRequest queries a JSONRPC endpoint and returns a map of the response
func (c *Endpoint) MakeRequest(method string, params interface{}) (interface{}, error) {
	baseURL := fmt.Sprintf("http://%s:%d", c.Host, c.Port)
	client := new(http.Client)
	req, err := http.NewRequest("POST", baseURL, nil)
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
		return nil, errors.New(string(str))
	}

	if result, found := data["result"]; found {
		return result, nil
	}
	return nil, errors.New("no result")
}

// NewClient generates a new jsonRPC client
func NewClient(user string, password string, host string, port int64) *Endpoint {
	c := Endpoint{user, password, host, port}
	return &c
}

// func (c *Endpoint) MakeRequest(method string, params interface{}) (interface{}, error) {
// 	baseURL := fmt.Sprintf("https://%s:%d", c.Host, c.Port)
// 	cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
// 	if err != nil {
// 		log.Printf("Error: %s when load client keys", err)
// 	}

// 	if len(cert.Certificate) != 2 {
// 		log.Printf("client1.crt should have 2 concatenated certificates: client + CA")
// 	}

// 	ca, err := x509.ParseCertificate(cert.Certificate[1])
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	certPool := x509.NewCertPool()
// 	certPool.AddCert(ca)

// 	config := tls.Config{
// 		Certificates: []tls.Certificate{cert},
// 		//InsecureSkipVerify: true,
// 		RootCAs: certPool,
// 	}

// 	conn, err := tls.Dial("tcp", baseURL, &config)
// 	if err != nil {
// 		log.Printf("Error: %s when dialing", err)
// 	}
// 	defer conn.Close()
// 	log.Println("Client connected to :", conn.RemoteAddr())
// 	rpcClient := rpc.NewClient(conn)
// 	var reply int
// 	if err := rpcClient.Call("MyServer.Sum", &RpcCall{
// 		Jsonrpc: "1.0",
// 		Id:      "BitNodes",
// 		Method:  method,
// 		params:  params,
// 	}, &reply); err != nil {
// 		log.Printf("Failed to call RPC", err)
// 	}
// 	log.Printf("Returned result is %d", reply)
// 	return rpcClient, err
// }
