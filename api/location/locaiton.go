// The service for the IP location info API
package main

import (
  "encoding/json"
  "github.com/ibondare/breechface/api/location/model"
  "github.com/oschwald/geoip2-golang"
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
// Batching is supported with a comma separated list, e.g.:
//   /location/ip/{ipAddr1},{ipAddr2},{ipAddr3}
func ipLocationHandler(w http.ResponseWriter, r *http.Request) {
  uriElements := strings.Split(r.URL.Path, "/")
  target := uriElements[len(uriElements)-1]

  if len(target) > 0 {
    valueList := strings.Split(target, ",")

    ipList := make([]net.IP, len(valueList))

    var ip net.IP

    for i, rawValue := range valueList {
      ip = net.ParseIP(rawValue)

      if (ip != nil) {
        ipList[i] = ip
      }
    }

    countryData, err := model.LocateCountry(context, ipList)
    if err != nil {
      log.Println(err)
    }

    b, err := json.Marshal(countryData)
  	if err != nil {
  		log.Println(err)
  	}

  	w.Write(b)
  }
}
