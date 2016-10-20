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

func LocateCountry(db *geoip2.Reader, ipAddr net.IP) (*CountryData, error) {

  record, err := db.Country(ipAddr)

  if err != nil {
    return nil, err
  }

  var countryData *CountryData;

  if record != nil {
    countryData = new(CountryData)

    countryData.IPAddress = ipAddr
    countryData.Name = record.Country.Names["en"]
    countryData.IsoCode = record.Country.IsoCode
  }

  return countryData, nil
}
