package main

import (
  "regexp"
  "strconv"
  "strings"
)

type Rss struct {
  Channel []Channel `xml:"channel"`
}

type Channel struct {
  Item []Item `xml:"item"`
}

type Item struct {
  Title       string `xml:"title"`
  Link        string `xml:"link"`
  PubDate     string `xml:"pubDate"`
  Description string `xml:"description"`
}

func (item Item) Price() int {
  priceString := captureRegex(`Cena: <b><b>([0-9,]+)`, item.Description)
  priceString = strings.Replace(priceString, ",", "", -1)
  price, err := strconv.Atoi(priceString)
  if err != nil { }
  return price
}

func (item Item) Manufacturer() string {
  return captureRegex(`Marka: <b>(.*?)<`, item.Description)
}

func (item Item) Model() string {
  return captureRegex(`Modelis: <b>(.*?)<`, item.Description)
}

func (item Item) Year() int {
  yearString := captureRegex(`Gads: <b>(.*?)<`, item.Description)
  year, err := strconv.Atoi(yearString)
  if err != nil { }
  return year
}

func (item Item) EngineCapacity() string {
  return captureRegex(`Tilp.: <b>(.*?)<`, item.Description)
}

func (item Item) Mileage() string {
  return captureRegex(`Nobrauk.: <b>(.*?)<`, item.Description)
}

func captureRegex (regex, text string) string {
  var regexObj = regexp.MustCompile(regex)
  if len(regexObj.FindStringSubmatch(text)) == 0 {
    return ""
  } else {
    return regexObj.FindStringSubmatch(text)[1]
  }
}

// Sample ss.lv RSS
// <rss>
// <channel>
// <item>
//    <title><![CDATA[Автозапуск - 200-300M.е прогретый автомобиль ...]]></title>
//    <link>http://www.ss.lv/msg/lv/transport/cars/toyota/corolla/bimno.html</link>
//    <pubDate>Fri, 27 Sep 2013 16:14:19 +0300</pubDate>
//    <description><![CDATA[<a href="http://www.ss.lv/msg/lv/transport/cars/toyota/corolla/bimno.html"><img align=right border=0 src="http://i.ss.lv/images/2013-09-27/316642/VHwPGkxjQ1w=/1.t.jpg" width="160" height="120" alt=""></a>
//       Marka: <b><b>Toyota<br>Corolla</b></b><br/>Modelis: <b><b>Corolla</b></b><br/>Gads: <b><b>2003</b></b><br/>Tilp.: <b><b>1.6</b></b><br/>Nobrauk.: <b><b>107</b> tūkst.</b><br/>Cena: <b><b>3,163</b>  Ls<br><div class=cc2>4,500 €</div></b><br/><br/><b><a href="http://www.ss.lv/msg/lv/transport/cars/toyota/corolla/bimno.html">Apskatīt sludinājumu</a></b><br/><br/>
//       ]]>
//     </description>

// </item>
// </channel>
// </rss>
