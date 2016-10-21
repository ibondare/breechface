// Country lookup logic
package model

import (
  "net"
)

type CountryData struct {
  IPAddress net.IP `json:"ipAddress"`
  Name string `json:"countryName"`
  IsoCode string `json:"isoCode"`
}

func LocateCountry(ipList []net.IP) (*[]CountryData, error) {
  result := make([]CountryData, len(ipList))

  for  i, ipAddr := range ipList {
    record, err := locateIpCountry(ipAddr)
    if (err != nil) {
      return nil, err
    }

    result[i] = record
  }

  return &result, nil
}

// Locate counrty data for a single IP address
func locateIpCountry(ipAddr net.IP) (CountryData, error) {
  var result CountryData

  record, err := dataStore.Country(ipAddr)

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
