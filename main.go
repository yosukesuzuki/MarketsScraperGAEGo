package marketsapi

import (
    "appengine"
    "appengine/urlfetch"
    "net/http"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/PuerkitoBio/goquery"
)

type Result struct {
    Title string
    PriceTime    string
    Price      string
    Diff      string
}


func init() {
    m := martini.Classic()
    m.Use(render.Renderer())
    m.Get("/", func() string {
        return "Hello world!"
    })
    m.Get("/api/Markets", func(w http.ResponseWriter,r render.Render,req *http.Request) {
        c := appengine.NewContext(req)
        client := urlfetch.Client(c)
        resp, err := client.Get("http://www.nikkei.com/markets/kaigai/worldidx.aspx")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            //return
        }
        //c.Infof("response: %v",resp.Body)
        results := []Result{}
        doc, _ := goquery.NewDocumentFromResponse(resp)
        doc.Find("div.mk-world_market div table tr").Each(func(_ int, s *goquery.Selection) {
                title := s.Find("th").Text()
                price := s.Find("th").Next().Text()
                diff := s.Find("td:nth-child(3)").Text()
                pricedate := s.Find("td:nth-child(4)").Text()
                result := Result{title,pricedate,price,diff}
                results = append(results,result)
        })
        r.JSON(200, map[string]interface{}{"results": results})
    })
    http.Handle("/", m)
}
