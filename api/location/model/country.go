// Country lookup logic
package model

import (
  "github.com/oschwald/geoip2-golang"
  "net"
)

type CountryData struct {
  IPAddress net.IP `json:"ipAddress"`
  Name string `json:"countryName"`
  IsoCode string `json:"isoCode"`
}

func LocateCountry(db *geoip2.Reader, ipList []net.IP) (*[]CountryData, error) {
  result := make([]CountryData, len(ipList))

  for  i, ipAddr := range ipList {
    record, err := locateIpCountry(db, ipAddr)
    if (err != nil) {
      return nil, err
    }

    result[i] = record
  }

  return &result, nil
}

// Locate counrty data for a single IP address
func locateIpCountry(db *geoip2.Reader, ipAddr net.IP) (CountryData, error) {
  var result CountryData

  record, err := db.Country(ipAddr)

  if err != nil {
    return result, err
  }

  if record != nil {
    result.IPAddress = ipAddr
    result.Name = record.Country.Names["en"]
    result.IsoCode = record.Country.IsoCode
  }

  return result, nil
}
