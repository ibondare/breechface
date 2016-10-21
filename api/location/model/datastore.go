// Datastore intialization/handling
package model

import (
  "github.com/oschwald/geoip2-golang"
)

const DefaultPath = "./country.mmdb"

var dataStore *geoip2.Reader

// Open the datastore identified by the file path
func Open(path string) error {
  db, err := geoip2.Open(path)

  if err != nil {
    return err
  }

  dataStore = db

  return nil
}

// Close the dataStore
func Close() {
  if dataStore != nil {
    dataStore.Close()
  }
}
