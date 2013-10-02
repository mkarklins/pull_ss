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
    fmt.Println(rss.Channel[0].Item[i].Price())
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
