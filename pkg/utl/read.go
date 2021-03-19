package utl

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Images is the list of differently scaled logo images for each coin
type Images struct {
	Img16  string `json:"img16"`
	Img32  string `json:"img32"`
	Img64  string `json:"img64"`
	Img128 string `json:"img128"`
	Img256 string `json:"img256"`
}

// GetJSON reads a JSON file and returns an map containing
// the parsed data
func GetJSON(completeURL string) (interface{}, error) {
	resp, err := http.Get(completeURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)

	var mapBodyInterface interface{}
	json.Unmarshal(mapBody, &mapBodyInterface)
	return mapBodyInterface, nil
}

// GetIMG loads a logo from the database and generates the various sized
// icons from it
func GetIMG(url, path, coin string) Images{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Problem Insert", err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	img16, _ := imageResize(content, options{Width: 16, Height: 16})
	img32, _ := imageResize(content, options{Width: 32, Height: 32})
	img64, _ := imageResize(content, options{Width: 64, Height: 64})
	img128, _ := imageResize(content, options{Width: 128, Height: 128})
	img256, _ := imageResize(content, options{Width: 256, Height: 256})
	imgs := Images{
		Img16:  base64.StdEncoding.EncodeToString(img16),
		Img32:  base64.StdEncoding.EncodeToString(img32),
		Img64:  base64.StdEncoding.EncodeToString(img64),
		Img128: base64.StdEncoding.EncodeToString(img128),
		Img256: base64.StdEncoding.EncodeToString(img256),
	}
	//Create a empty file
	ioutil.WriteFile(path+"/"+ coin +"16.png", img16, 777)
	ioutil.WriteFile(path+"/"+ coin +"32.png", img32, 777)
	ioutil.WriteFile(path+"/"+ coin +"64.png", img64, 777)
	ioutil.WriteFile(path+"/"+ coin +"128.png", img128, 777)
	ioutil.WriteFile(path+"/"+ coin +"256.png", img256, 777)
	return imgs
}
