package godruid

import (
	"encoding/json"
)

type Aggregator struct {
	Type        string   `json:"type"`
	Name        string   `json:"name,omitempty"`
	FieldName   string   `json:"fieldName,omitempty"`
	FieldNames  []string `json:"fieldNames,omitempty"`
	FnAggregate string   `json:"fnAggregate,omitempty"`
	FnCombine   string   `json:"fnCombine,omitempty"`
	FnReset     string   `json:"fnReset,omitempty"`
	ByRow       bool     `json:"byRow,omitempty"`
}

type FilteredAggregator struct {
	Filter      *Filter       `json:"filter,omitempty"`
	Aggregator  Aggregator    `json:"aggregator,omitempty"`
}

type Aggregation struct {
	Aggregator
	FilteredAggregator
}

func AggRawJson(rawJson string) Aggregation {
	agg := &Aggregator{}
	json.Unmarshal([]byte(rawJson), agg)
	return Aggregation{*agg, FilteredAggregator{}}
}

func AggCount(name string) Aggregation {
	agg := Aggregator{
		Type: "count",
		Name: name,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggLongSum(name, fieldName string) Aggregation {
	agg := Aggregator{
		Type:      "longSum",
		Name:      name,
		FieldName: fieldName,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggDoubleSum(name, fieldName string) Aggregation {
	agg := Aggregator{
		Type:      "doubleSum",
		Name:      name,
		FieldName: fieldName,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggMin(name, fieldName string) Aggregation {
	agg := Aggregator{
		Type:      "min",
		Name:      name,
		FieldName: fieldName,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggMax(name, fieldName string) Aggregation {
	agg := Aggregator{
		Type:      "max",
		Name:      name,
		FieldName: fieldName,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggJavaScript(name, fnAggregate, fnCombine, fnReset string, fieldNames []string) Aggregation {
	agg := Aggregator{
		Type:        "javascript",
		Name:        name,
		FieldNames:  fieldNames,
		FnAggregate: fnAggregate,
		FnCombine:   fnCombine,
		FnReset:     fnReset,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func AggCardinality(name string, fieldNames []string, byRow ...bool) Aggregation {
	isByRow := false
	if len(byRow) != 0 {
		isByRow = byRow[0]
	}
	agg := Aggregator{
		Type:       "cardinality",
		Name:       name,
		FieldNames: fieldNames,
		ByRow:      isByRow,
	}
	return Aggregation{agg, FilteredAggregator{}}
}

func FilteredAgg(filter *Filter, aggregation Aggregation) Aggregation {
	agg := Aggregator{
		Type: "filtered",
	}
	filterAgg := FilteredAggregator{
		Filter: filter,
		// aggregation.Aggregator should be part of FilteredAggregator
		// since this is a different json property field "aggregator"
		Aggregator: aggregation.Aggregator,
	}
	return Aggregation{agg, filterAgg}
}
