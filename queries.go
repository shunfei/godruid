package godruid

import (
	"bytes"
	"encoding/json"
)

// Check http://druid.io/docs/0.6.154/Querying.html#query-operators for detail description.

// The Query interface stands for any kinds of druid query.
type Query interface {
	setup()
	onResponse(content []byte) error
}

// ---------------------------------
// GroupBy Query
// ---------------------------------

type QueryGroupBy struct {
	QueryType        string                 `json:"queryType"`
	DataSource       string                 `json:"dataSource"`
	Dimensions       []DimSpec              `json:"dimensions"`
	Granularity      Granlarity             `json:"granularity"`
	LimitSpec        *Limit                 `json:"limitSpec,omitempty"`
	Having           *Having                `json:"having,omitempty"`
	Filter           *Filter                `json:"filter,omitempty"`
	Aggregations     []Aggregation          `json:"aggregations"`
	PostAggregations []PostAggregation      `json:"postAggregations,omitempty"`
	Intervals        []string               `json:"intervals"`
	Context          map[string]interface{} `json:"context,omitempty"`

	QueryResult []GroupbyItem `json:"-"`
}

type GroupbyItem struct {
	Version   string                 `json:"version"`
	Timestamp string                 `json:"timestamp"`
	Event     map[string]interface{} `json:"event"`
}

func (q *QueryGroupBy) setup() { q.QueryType = "groupBy" }
func (q *QueryGroupBy) onResponse(content []byte) error {
	res := new([]GroupbyItem)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// Search Query
// ---------------------------------

type QuerySearch struct {
	QueryType        string                 `json:"queryType"`
	DataSource       string                 `json:"dataSource"`
	Granularity      Granlarity             `json:"granularity"`
	Filter           *Filter                `json:"filter,omitempty"`
	Intervals        []string               `json:"intervals"`
	SearchDimensions []string               `json:"searchDimensions,omitempty"`
	Query            *SearchQuery           `json:"query"`
	Sort             *SearchSort            `json:"sort"`
	Context          map[string]interface{} `json:"context,omitempty"`

	QueryResult []SearchItem `json:"-"`
}

type SearchItem struct {
	Timestamp string     `json:"timestamp"`
	Result    []DimValue `json:"result"`
}

type DimValue struct {
	Dimension string `json:"dimension"`
	Value     string `json:"value"`
}

func (q *QuerySearch) setup() { q.QueryType = "search" }
func (q *QuerySearch) onResponse(content []byte) error {
	res := new([]SearchItem)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// SegmentMetadata Query
// ---------------------------------

type QuerySegmentMetadata struct {
	QueryType  string                 `json:"queryType"`
	DataSource string                 `json:"dataSource"`
	Intervals  []string               `json:"intervals"`
	ToInclude  *ToInclude             `json:"toInclude,omitempty"`
	Merge      interface{}            `json:"merge,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`

	QueryResult []SegmentMetaData `json:"-"`
}

type SegmentMetaData struct {
	Id        string                `json:"id"`
	Intervals []string              `json:"intervals"`
	Columns   map[string]ColumnItem `json:"columns"`
}

type ColumnItem struct {
	Type        string      `json:"type"`
	Size        int         `json:"size"`
	Cardinality interface{} `json:"cardinality"`
}

func (q *QuerySegmentMetadata) setup() { q.QueryType = "segmentMetadata" }
func (q *QuerySegmentMetadata) onResponse(content []byte) error {
	res := new([]SegmentMetaData)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// TimeBoundary Query
// ---------------------------------

type QueryTimeBoundary struct {
	QueryType  string                 `json:"queryType"`
	DataSource string                 `json:"dataSource"`
	Bound      string                 `json:"bound,omitempty"`
	Filter     *Filter                `json:"filter,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`

	QueryResult []TimeBoundaryItem `json:"-"`
}

type TimeBoundaryItem struct {
	Timestamp string       `json:"timestamp"`
	Result    TimeBoundary `json:"result"`
}

type TimeBoundary struct {
	MinTime string `json:"minTime"`
	MaxTime string `json:"maxTime"`
}

func (q *QueryTimeBoundary) setup() { q.QueryType = "timeBoundary" }
func (q *QueryTimeBoundary) onResponse(content []byte) error {
	res := new([]TimeBoundaryItem)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// Timeseries Query
// ---------------------------------

type QueryTimeseries struct {
	QueryType        string                 `json:"queryType"`
	DataSource       string                 `json:"dataSource"`
	Granularity      Granlarity             `json:"granularity"`
	Filter           *Filter                `json:"filter,omitempty"`
	Aggregations     []Aggregation          `json:"aggregations"`
	PostAggregations []PostAggregation      `json:"postAggregations,omitempty"`
	Intervals        []string               `json:"intervals"`
	Context          map[string]interface{} `json:"context,omitempty"`

	QueryResult []Timeseries `json:"-"`
}

type Timeseries struct {
	Timestamp string                 `json:"timestamp"`
	Result    map[string]interface{} `json:"result"`
}

func (q *QueryTimeseries) setup() { q.QueryType = "timeseries" }
func (q *QueryTimeseries) onResponse(content []byte) error {
	res := new([]Timeseries)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// TopN Query
// ---------------------------------

type QueryTopN struct {
	QueryType        string                 `json:"queryType"`
	DataSource       string                 `json:"dataSource"`
	Granularity      Granlarity             `json:"granularity"`
	Dimension        DimSpec                `json:"dimension"`
	Threshold        int                    `json:"threshold"`
	Metric           *TopNMetric            `json:"metric"`
	Filter           *Filter                `json:"filter,omitempty"`
	Aggregations     []Aggregation          `json:"aggregations"`
	PostAggregations []PostAggregation      `json:"postAggregations,omitempty"`
	Intervals        []string               `json:"intervals"`
	Context          map[string]interface{} `json:"context,omitempty"`

	QueryResult []TopNItem `json:"-"`
}

type TopNItem struct {
	Timestamp string                   `json:"timestamp"`
	Result    []map[string]interface{} `json:"result"`
}

func (q *QueryTopN) setup() { q.QueryType = "topN" }
func (q *QueryTopN) onResponse(content []byte) error {
	res := new([]TopNItem)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// Select Query
// ---------------------------------

type QuerySelect struct {
	QueryType   string                 `json:"queryType"`
	DataSource  string                 `json:"dataSource"`
	Intervals   []string               `json:"intervals"`
	Descending  bool                   `json:"descending"`
	Filter      *Filter                `json:"filter,omitempty"`
	Dimensions  []DimSpec              `json:"dimensions"`
	Metrics     []string               `json:"metrics"`
	Granularity Granlarity             `json:"granularity"`
	PagingSpec  map[string]interface{} `json:"pagingSpec,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"`

	QueryResult []SelectBlob `json:"-"`
}

// Select json blob from druid comes back as following:
// http://druid.io/docs/latest/querying/select-query.html
// the interesting results are in events blob which we
// call as 'SelectEvent'.
type SelectBlob struct {
	Timestamp string       `json:"timestamp"`
	Result    SelectResult `json:"result"`
}

type SelectResult struct {
	PagingIdentifiers map[string]int64 `json:"pagingIdentifiers"`
	Events            []SelectEvent    `json:"events"`
}

type SelectEvent struct {
	SegmentId string                 `json:"segmentId"`
	Offset    int64                  `json:"offset"`
	Event     map[string]interface{} `json:"event"`
}

func (q *QuerySelect) setup() { q.QueryType = "select" }
func (q *QuerySelect) onResponse(content []byte) error {
	res := new([]SelectBlob)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}

// ---------------------------------
// Scan Query
// ---------------------------------

// QueryScan is the model for scan query
type QueryScan struct {
	QueryType    string                 `json:"queryType"`
	DataSource   string                 `json:"dataSource"`
	Intervals    []string               `json:"intervals"`
	BatchSize    int64                  `json:"batchSize"`
	Limit        int64                  `json:"limit"`
	Order        string                 `json:"order"`
	Filter       *Filter                `json:"filter,omitempty"`
	Context      map[string]interface{} `json:"context,omitempty"`
	ResultFormat string                 `json:"resultFormat,omitempty"`

	QueryResult []ScanBlob `json:"-"`
}

// ScanBlob is the response to scan query
type ScanBlob struct {
	Columns   []string                 `json:"columns"`
	SegmentID string                   `json:"segmentId"`
	Events    []map[string]interface{} `json:"events"`
}

func (q *QueryScan) setup() { q.QueryType = "scan" }
func (q *QueryScan) onResponse(content []byte) error {
	res := new([]ScanBlob)
	d := json.NewDecoder(bytes.NewReader(content))
	d.UseNumber()
	if err := d.Decode(res); err != nil {
		return err
	}
	q.QueryResult = *res
	return nil
}
