package main

import (
  "log"
  "regexp"
  "strconv"
  "strings"
  "encoding/xml"
  "io/ioutil"
  "net/http"
)

type Rss struct {
  Channel []Channel `xml:"channel"`
}

type Channel struct {
  Items []Item `xml:"item"`
}

type Item struct {
  Title       string `xml:"title"`
  Link        string `xml:"link"`
  PubDate     string `xml:"pubDate"`
  Description string `xml:"description"`
}

func (item Item) Price() int {
  priceString := item.Capture(`Cena: <b><b>([0-9,]+)`)
  priceString = strings.Replace(priceString, ",", "", -1)
  price, err := strconv.Atoi(priceString)
  if err != nil { }
  return price
}

func (item Item) Manufacturer() string {
  return item.Capture(`Marka: <b><b>(.*?)<`)
}

func (item Item) Model() string {
  return item.Capture(`Modelis: <b><b>(.*?)<`)
}

func (item Item) Year() int {
  yearString := item.Capture(`Gads: <b><b>(.*?)<`)
  year, err := strconv.Atoi(yearString)
  if err != nil { }
  return year
}

func (item Item) EngineCapacity() string {
  return item.Capture(`Tilp.: <b>(.*?)<`)
}

func (item Item) Mileage() string {
  return item.Capture(`Nobrauk.: <b>(.*?)<`)
}

func (item Item) ForSale() bool {
  return item.Price() > 0
}

// Returns a list of car feed items.
func fetch() []Item {
  rss := &Rss{}
  resp, err := http.Get("http://www.ss.lv/lv/transport/cars/rss/")
  if err != nil { log.Fatal(err) }

  // Always closes the body, even if an exception occures.
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  body = body

  err = xml.Unmarshal([]byte(body), rss)
  // err = xml.Unmarshal([]byte(fixture), rss)
  if err != nil { log.Fatal(err) }
  return rss.Channel[0].Items
}

func (item Item) Capture(regex string) string {
  var regexObj = regexp.MustCompile(regex)
  if len(regexObj.FindStringSubmatch(item.Description)) == 0 {
    return ""
  } else {
    return regexObj.FindStringSubmatch(item.Description)[1]
  }
}


// Sample ss.lv RSS
var fixture = `<rss>
<channel>
  <item>
  <title>
  <![CDATA[
  Tikko no Vacijas. Volvo V40 1.9 dizels. Pirma registracija - Junis 2004. Latvijā nav eksploatēts. Elektro logi. Elektro regul ...
  ]]>
  </title>
  <link>
  http://www.ss.lv/msg/lv/transport/cars/volvo/v40/cbcxh.html
  </link>
  <pubDate>Thu, 03 Oct 2013 21:09:32 +0300</pubDate>
  <description>
  <![CDATA[
  <a href="http://www.ss.lv/msg/lv/transport/cars/volvo/v40/cbcxh.html"><img align=right border=0 src="http://i.ss.lv/images/2013-08-15/311253/VHwIHk1gRFw=/1.t.jpg" width="160" height="120" alt=""></a>
   Marka: <b><b>Bmw<br>V40</b></b><br/>Modelis: <b><b>320</b></b><br/>Gads: <b><b>2008</b></b><br/>Tilp.: <b><b>1.9D</b></b><br/>Nobrauk.: <b><b>215</b> tūkst.</b><br/>Cena: <b><b>8,600</b> Ls<br><div class=cc2>3,700 €</div></b><br/><br/><b><a href="http://www.ss.lv/msg/lv/transport/cars/volvo/v40/cbcxh.html">Apskatīt sludinājumu</a></b><br/><br/>
  ]]>
  </description>
  </item>
</channel>
</rss>
`