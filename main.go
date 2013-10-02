package main

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  rss := &Rss{}

  resp, err := http.Get("http://www.ss.lv/lv/transport/cars/rss/")
  if err != nil { fmt.Println("Could not get rss") }

  // Always closes the body, even if an exception occures.
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  err = xml.Unmarshal([]byte(body), rss)
  if err != nil { fmt.Println("Marshalling failed") }

  for i := 0; i < len(rss.Channel[0].Item); i++ {
    // For testing purposes print out price.
    fmt.Println(rss.Channel[0].Item[i].Price())
  }
}
