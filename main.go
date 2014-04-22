package marketsapi

import (
    "fmt"
    "time"
    "strings"
    "strconv"
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


func StringToTime(dateString string,t time.Time) string{
    t_jst := t.Add(9*time.Hour)
    datetext := strings.Replace(dateString,"日","",-1)
    datetextarr := strings.Split(datetext," ")
    dayNum, _ := strconv.Atoi(datetextarr[0])
    if dayNum > t_jst.Day() {
        t_jst = t_jst.AddDate(0,-1,0)
    }
    yyyymm := t_jst.Format("2006-01")
    yyyymmddhhmm := yyyymm +"-"+ datetext
    date_yyyymmddhhmm,err := time.Parse("2006-01-02 15:04",yyyymmddhhmm)
    if err != nil{
        fmt.Println("error")
    }
    date_yyyymmddhhmm = date_yyyymmddhhmm.Add(-9*time.Hour)
    return date_yyyymmddhhmm.Format("2006-01-02 15:04:05 MST")
}

func Indexes () map[string]string{
    indexes := map[string]string{
        "日経平均（円）":"Nikkei225",
        "ドル・円":"USD/JPY",
        "ユーロ・円":"EURO/JPY",
        "ユーロ・ドル":"EURO/USD",
        "ドル・中国人民元":"USD/CNY",
        "NYダウ工業株30種（ドル）":"DJIA",
        "ナスダック":"Nasdaq",
        "英FTSE100":"FTSE100",
    }
    return indexes
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
        indexes := Indexes()
        results := []Result{}
        doc, _ := goquery.NewDocumentFromResponse(resp)
        doc.Find("div.mk-world_market div table tr").Each(func(_ int, s *goquery.Selection) {
                title := s.Find("th").Text()
                title = strings.Trim(strings.Replace(title,"※","",-1)," ")
                if val,ok := indexes[title]; ok {
                    price := s.Find("th").Next().Text()
                    diff := s.Find("td:nth-child(3)").Text()
                    pricetime := s.Find("td:nth-child(4)").Text()
                    t := time.Now()
                    pricetime = StringToTime(pricetime,t)
                    result := Result{val,pricetime,price,diff}
                    results = append(results,result)
                }
        })
        r.JSON(200, map[string]interface{}{"results": results})
    })
    http.Handle("/", m)
}
