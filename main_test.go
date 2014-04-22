package marketsapi

import (
    "time"
    "testing"
)

func TestStringToTime(t *testing.T) {
    dateString := "31日 15:28"
    now,err := time.Parse("2006-01-02 15:04","2014-04-01 09:10")
    if err != nil{
        t.Fatalf("error")
    }
    yyyymmddhhmm := StringToTime(dateString,now)
    assert_date := "2014-03-31 06:28:00 UTC"
    if yyyymmddhhmm != assert_date {
        t.Fatalf("error: date_yyyymmddhhmm = %+v, assert_date = %+v",yyyymmddhhmm,assert_date)
    }
}
