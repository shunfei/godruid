package godruid

import (
    "encoding/json"
)

type PostAggregation struct {
    Type       string            `json:"type"`
    Name       string            `json:"name,omitempty"`
    Value      interface{}       `json:"value,omitempty"`
    Fn         string            `json:"fn,omitempty"`
    Fields     []PostAggregation `json:"fields,omitempty"`
    FieldName  string            `json:"fieldName,omitempty"`
    FieldNames []string          `json:"fieldNames,omitempty"`
    Function   string            `json:"function,omitempty"`
}

// Return the aggregations which this post-aggregation used.
// It could be helpful while automatically filling the aggregations base on post-aggregations.
func (pa PostAggregation) GetRelativeAggs() (aggs []string) {
    if pa.FieldName != "" {
        aggs = append(aggs, pa.FieldName)
    }
    aggs = append(aggs, pa.FieldNames...)
    for _, spa := range pa.Fields {
        aggs = append(aggs, spa.GetRelativeAggs()...)
    }
    return
}

func PostAggRawJson(rawJson string) PostAggregation {
    pa := &PostAggregation{}
    json.Unmarshal([]byte(rawJson), pa)
    return *pa
}

func PostAggArithmetic(name, fn string, fields []PostAggregation) PostAggregation {
    return PostAggregation{
        Type:   "arithmetic",
        Name:   name,
        Fn:     fn,
        Fields: fields,
    }
}

func PostAggFieldAccessor(fieldName string) PostAggregation {
    return PostAggregation{
        Type:      "fieldAccess",
        FieldName: fieldName,
    }
}

func PostAggConstant(name string, value interface{}) PostAggregation {
    return PostAggregation{
        Type:  "constant",
        Name:  name,
        Value: value,
    }
}

func PostAggJavaScript(name, function string, fieldNames []string) PostAggregation {
    return PostAggregation{
        Type:       "javascript",
        Name:       name,
        FieldNames: fieldNames,
        Function:   function,
    }
}

func PostAggFieldHyperUnique(fieldName string) PostAggregation {
    return PostAggregation{
        Type:      "hyperUniqueCardinality",
        FieldName: fieldName,
    }
}
