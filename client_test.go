package godruid

import (
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestGroupby(t *testing.T) {
    Convey("TestGroupby", t, func() {
        query := &QueryGroupBy{
            DataSource:       "campaign",
            Intervals:        []string{"2014-09-01T00:00/2020-01-01T00"},
            Granularity:      GranAll,
            Filter:           FilterJavaScript("hour", "function(x) { return(x >= 1) }"),
            LimitSpec:        LimitSpecDefault(5),
            Dimensions:       []string{"campaign_id"},
            Aggregations:     []Aggregation{AggCount("count"), AggLongSum("impressions", "impressions")},
            PostAggregations: []PostAggregatable{PostAggArithmetic("imp/count", "/", []PostAggregatable{PostAggFieldAccessor("impressions"), PostAggFieldAccessor("count")})},
            // PostAggregations: []PostAggregatable{PostAggFieldAccessor("impressions")},
        }
        client := Client{
            Url:   "http://192.168.10.60:8009",
            Debug: true,
        }

        err := client.DoQuery(query)
        So(err, ShouldEqual, nil)

        fmt.Println("requst", client.LastRequest)
        fmt.Println("response", client.LastResponse)

        fmt.Printf("query.QueryResult:\n%v", query.QueryResult)

    })
}

func TestSearch(t *testing.T) {
    // return
    Convey("TestSearch", t, func() {
        query := &QuerySearch{
            DataSource:       "campaign",
            Intervals:        []string{"2014-09-01T00:00/2020-01-01T00"},
            Granularity:      GranAll,
            SearchDimensions: []string{"campaign_id", "hour"},
            Query:            SearchQueryInsensitiveContains(1313),
            Sort:             SearchSortLexicographic,
        }
        client := Client{
            Url:   "http://192.168.10.60:8009",
            Debug: true,
        }

        err := client.DoQuery(query)
        So(err, ShouldEqual, nil)

        fmt.Println("requst", client.LastRequest)
        fmt.Println("response", client.LastResponse)

        fmt.Printf("query.QueryResult:\n%v", query.QueryResult)

    })
}
