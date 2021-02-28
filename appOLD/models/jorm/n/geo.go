package n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/appOLD/models/jorm/jdb"
)

type GeoResponse struct {
	Status      string
	Description string
	Data        struct {
		Geo struct {
			Host          string  `json:"host"`
			IP            string  `json:"ip"`
			RDSN          string  `json:"rdns"`
			ASN           float64 `json:"asn"`
			ISP           string  `json:"isp"`
			CountryName   string  `json:"country_name"`
			CountryCode   string  `json:"country_code"`
			RegionName    string  `json:"region_name"`
			RegionCode    string  `json:"region_code"`
			City          string  `json:"city"`
			PostalCode    string  `json:"postal_code"`
			ContinentName string  `json:"continent_name"`
			ContinentCode string  `json:"continent_code"`
			Latitude      float64 `json:"latitude"`
			Longitude     float64 `json:"longitude"`
			MetroCode     string  `json:"metro_code"`
			Timezone      string  `json:"timezone"`
			Datetime      string  `json:"datetime"`
		}
	}
}

func GetGeoIP(ip string) (n NodeInfo) {

	if jdb.DB.Read(cfg.Web+"/geo", ip, &n) != nil {

		if ip[:4] == "10.0" {
			ip = "212.62.35.158"
		}
		fmt.Println("laaaaaa", ip)
		resp, err := http.Get("https://tools.keycdn.com/geo.json?host=" + ip)
		if err != nil {
		}
		defer resp.Body.Close()
		mapBody, err := ioutil.ReadAll(resp.Body)
		var g GeoResponse
		err = json.Unmarshal(mapBody, &g)
		if err != nil {
		}
		geo := g.Data.Geo
		n.IP = ip
		n.Rdns = geo.RDSN
		n.ISP = geo.ISP
		n.CountryName = geo.CountryName
		n.CountryCode = geo.CountryCode
		n.RegionName = geo.RegionName
		n.RegionCode = geo.RegionCode
		n.City = geo.City
		n.Zipcode = geo.PostalCode
		n.ContinentName = geo.ContinentName
		n.ContinentCode = geo.ContinentCode
		n.Latitude = fmt.Sprintf("%.4f", geo.Latitude)
		n.Longitude = fmt.Sprintf("%.4f", geo.Longitude)
		n.Postcode = geo.PostalCode
		n.Timezone = geo.Timezone
		jdb.DB.Write(cfg.Web+"/geo", ip, n)
	}
	return n
}
