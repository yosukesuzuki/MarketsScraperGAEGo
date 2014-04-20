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
        resp, err := client.Get("http://stocks.finance.yahoo.co.jp/stocks/list/indices?area=asia")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            //return
        }
        //c.Infof("response: %v",resp.Body)
        results := []Result{}
        doc, _ := goquery.NewDocumentFromResponse(resp)
        doc.Find("dl.lineFi").Each(func(_ int, s *goquery.Selection) {
            if s.Find("dt.title > a").Text() != "" {
                title := s.Find("dt.title > a").Text()
                c.Infof("title: %v",title)
                pricedate := s.Find("span.date").Text()
                price := s.Find("dd.fixWidth3 > strong").Text()
                diff := s.Find("dd.fixWidth2 > strong").Text()
                result := Result{title,pricedate,price,diff}
                results = append(results,result)
            }
        })
        r.JSON(200, map[string]interface{}{"results": results})
    })
    http.Handle("/", m)
}
