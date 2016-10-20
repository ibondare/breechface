// The service for the IP location info API
package main

import (
  "encoding/json"
  "github.com/oschwald/geoip2-golang"
  "github.com/ibondare/breechface/api/location/model"
  "log"
  "net"
  "net/http"
  "strings"
)

var context *geoip2.Reader

func main() {
  db, err := geoip2.Open("country.mmdb")

  if err != nil {
    log.Fatal(err)
  }

  context = db

  http.HandleFunc("/location/ip/", ipLocationHandler)
	http.ListenAndServe(":8080", nil)
}

// /location/ip/{ipAddr} URI handler
func ipLocationHandler(w http.ResponseWriter, r *http.Request) {
  values := strings.Split(r.URL.Path, "/")

  target := values[len(values)-1]

  if target == "ip" {
    log.Println("POST is unsupported")
  } else {
    // This is a GET with the IP address in the URI (/location/ip/ipAddr)
    ip := net.ParseIP(target)

    if ip != nil {
      countryData, err := model.LocateCountry(context, ip)
      if err != nil {
        log.Fatal(err)
      }

      b, err := json.Marshal(countryData)
    	if err != nil {
    		log.Println(err)
    	}

    	w.Write(b)
    }
  }
}
