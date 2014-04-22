package marketsapi

import (
    "testing"
)

func TestStringToTime(t *testing.T) {
    dateString := "31日 15:28"
    yyyymmddhhmm := StringToTime(dateString)
    assert_date := "2014-03-31 06:28:00 UTC"
    if yyyymmddhhmm != assert_date {
        t.Fatalf("error: date_yyyymmddhhmm = %+v, assert_date = %+v",yyyymmddhhmm,assert_date)
    }
}
