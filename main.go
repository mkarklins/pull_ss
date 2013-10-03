package main

import (
  "time"
  "encoding/base64"
  // "fmt"
  "io"
  "crypto/md5"
)

func main() {
  var feedMemmory []string
  for {
    feedItems := fetch()
    for i := 0; i < len(feedItems); i ++ {
      feedChecksum := checksum(feedItems[i].Description)
      if isInteresting(feedItems[i]) && !contains(feedMemmory, feedChecksum) {
        feedMemmory = append(feedMemmory, checksum(feedItems[i].Description))
        notify(feedItems[i].Link)
      }
    }
    time.Sleep(time.Second * 3)
  }
}

func contains(list []string, elem string) bool {
  for _, t := range list { if t == elem { return true } }
  return false
}

func checksum(text string) string {
  h := md5.New()
  io.WriteString(h, text)
  // A very strange way to convert bytes to string.
  return base64.URLEncoding.EncodeToString((h.Sum([]byte{})))
}

func isInteresting(item Item) bool {
  return ( item.ForSale() &&
           (item.Manufacturer() == "Bmw" || item.Manufacturer() == "Audi") &&
           (item.Model() == "320" || item.Model() == "318" || item.Model() == "325" || item.Model() == "330" || item.Model() == "A4") &&
           item.Price() > 7000 &&
           item.Price() < 9000)
}