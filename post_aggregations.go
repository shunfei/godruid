package godruid

type PostAggregatable interface{}

type PostAggregation struct {
    Type       string             `json:"type"`
    Name       string             `json:"name,omitempty"`
    Fn         string             `json:"fn,omitempty"`
    Fields     []PostAggregatable `json:"fields,omitempty"`
    FieldNames []string           `json:"fieldNames,omitempty"`
    Function   string             `json:"function,omitempty"`
}

type PostAggregator struct {
    Type      string      `json:"type"`
    FieldName string      `json:"fieldName,omitempty"`
    Name      string      `json:"name,omitempty"`
    Value     interface{} `json:"value,omitempty"`
}

func PostAggArithmetic(name, fn string, fields []PostAggregatable) PostAggregatable {
    return &PostAggregation{
        Type:   "arithmetic",
        Name:   name,
        Fn:     fn,
        Fields: fields,
    }
}

func PostAggFieldAccessor(fieldName string) PostAggregatable {
    return &PostAggregator{
        Type:      "fieldAccess",
        FieldName: fieldName,
    }
}

func PostAggConstant(name string, value interface{}) PostAggregatable {
    return &PostAggregator{
        Type:  "constant",
        Name:  name,
        Value: value,
    }
}

func PostAggJavaScript(name, function string, fieldNames []string) PostAggregatable {
    return &PostAggregation{
        Type:       "javascript",
        Name:       name,
        FieldNames: fieldNames,
        Function:   function,
    }
}

func PostAggFieldHyperUnique(fieldName string) PostAggregatable {
    return &PostAggregator{
        Type:      "hyperUniqueCardinality",
        FieldName: fieldName,
    }
}
