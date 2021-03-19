package u

import (
	"context"
	"fmt"
	"github.com/oknors/okno/app/jorm/c"
	"github.com/oknors/okno/pkg/utl"
	"log"


	"github.com/oknors/okno/app/cfg"

	cf "github.com/cloudflare/cloudflare-go"

)

func CloudFlare() {
	fmt.Println("CONFIGCONFIGCONFIGCONFIGCONFIGCONFIGCONFIG", cfg.CONFIG)
	ctx := context.Background()
	// Construct a new API object
	api, err := cf.New(cfg.CONFIG.CF.CloudFlareAPIkey, cfg.CONFIG.CF.CloudFlareEmail)
	utl.ErrorLog(err)

	//delAllCNameDNS(api, ctx, "com-http.us")
	createDNS(api, ctx, "com-http.us")


}

func createDNS(api *cf.API ,ctx context.Context, domain string){
	coins := c.LoadCoinsBase(true)
	// Fetch the zone ID
	id, err := api.ZoneIDByName(domain) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all records for a zone
	recs, err := api.DNSRecords(context.Background(), id, cf.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}
	var registrated []string
	for _, r := range recs {
		if r.Type ==  "CNAME"{
			registrated = append(registrated, r.Name )
		}
	}

	for _, coin := range coins.C {
		slug := coin.Slug
		fmt.Println("Eto for:", slug)
		//_, err := http.Get("https://" + slug + "." + domain)
		//if err != nil {

		for _, reg := range registrated {
			if slug+"."+domain == reg {
				fmt.Println("reg:", reg)
				fmt.Println("slug:", slug)
				fmt.Println("Ima:", slug+"."+domain)
			}else{
				id, err := api.ZoneIDByName(domain)
				utl.ErrorLog(err)
				t := true
				_, err = api.CreateDNSRecord(ctx, id, cf.DNSRecord{
					Type:    "CNAME",
					Name:    slug,
					Content: domain,
					TTL:     1,
					Proxied: &t,
				})
				utl.ErrorLog(err)
				fmt.Println("Created subdomain for:", slug)
			}
		}

		}
}

func delAllCNameDNS(api *cf.API ,ctx context.Context, domain string) {
	// Fetch the zone ID
	id, err := api.ZoneIDByName(domain) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all records for a zone
	recs, err := api.DNSRecords(context.Background(), id, cf.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range recs {
		if r.Type ==  "CNAME"{
			go delDNS(api, ctx, id, r.ID)
		}
	}
}

func delDNS(api *cf.API ,ctx context.Context, zoneId,id string) {
		err := api.DeleteDNSRecord(ctx, zoneId, id)
		utl.ErrorLog(err)
		fmt.Println("DeleteDNSRecord rrrrr:", id)
}
