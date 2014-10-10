package godruid

type Having struct {
    Type        string      `json:"type"`
    Aggregation string      `json:"aggregation,omitempty"`
    Value       interface{} `json:"nalue,omitempty"`
    HavingSpec  *Having     `json:"havingSpec,omitempty"`
    HavingSpecs []*Having   `json:"havingSpecs,omitempty"`
}

func HavingEqualTo(agg string, value interface{}) *Having {
    return &Having{
        Type:        "equalto",
        Aggregation: agg,
        Value:       value,
    }
}

func HavingGreaterThan(agg string, value interface{}) *Having {
    return &Having{
        Type:        "greaterThan",
        Aggregation: agg,
        Value:       value,
    }
}

func HavingLessThan(agg string, value interface{}) *Having {
    return &Having{
        Type:        "lessThan",
        Aggregation: agg,
        Value:       value,
    }
}

func HavingAnd(havings ...*Having) *Having {
    return joinHavings(havings, "and")
}

func HavingOr(havings ...*Having) *Having {
    return joinHavings(havings, "or")
}

func HavingNot(having *Having) *Having {
    return &Having{
        Type:       "not",
        HavingSpec: having,
    }
}

func joinHavings(havings []*Having, connector string) *Having {
    fLen := len(havings)
    if fLen == 0 {
        return nil
    }
    if fLen == 1 {
        return havings[0]
    }

    return &Having{
        Type:        connector,
        HavingSpecs: havings,
    }
}
