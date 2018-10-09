package godruid

type Filter struct {
	Type        string           `json:"type"`
	Dimension   string           `json:"dimension,omitempty"`
	Value       interface{}      `json:"value,omitempty"`
	Values      []interface{}    `json:"values,omitempty"`
	Pattern     string           `json:"pattern,omitempty"`
	Function    string           `json:"function,omitempty"`
	Lower       string           `json:"lower,omitempty"`
	Upper       string           `json:"upper,omitempty"`
	LowerStrict *bool            `json:"lowerStrict,omitempty"`
	UpperStrict *bool            `json:"upperStrict,omitempty"`
	Ordering    string           `json:"ordering,omitempty"`
	Field       *Filter          `json:"field,omitempty"`
	Fields      []*Filter        `json:"fields,omitempty"`
	SearchSpec  *SearchQuerySpec `json:"query,omitempty"`
}

type SearchQuerySpec struct {
	Type          string   `json:"type"`
	Value         string   `json:"value,omitempty"`
	Values        []string `json:"values,omitempty"`
	CaseSensitive bool     `json:"caseSensitive,omitempty"`
}

func FilterSelector(dimension string, value interface{}) *Filter {
	return &Filter{
		Type:      "selector",
		Dimension: dimension,
		Value:     value,
	}
}

// Filter for <= operator
func FilterLte(dimension string, value string, ordering string) *Filter {
	upperStrict := false
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Upper:       value,
		UpperStrict: &upperStrict,
		Ordering:    ordering,
	}
}

// Filter for < operator
func FilterLt(dimension string, value string, ordering string) *Filter {
	upperStrict := true
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Upper:       value,
		UpperStrict: &upperStrict,
		Ordering:    ordering,
	}
}

// Filter for >= operator
func FilterGte(dimension string, value string, ordering string) *Filter {
	lowerStrict := false
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Lower:       value,
		LowerStrict: &lowerStrict,
		Ordering:    ordering,
	}
}

// Filter for > operator
func FilterGt(dimension string, value string, ordering string) *Filter {
	lowerStrict := true
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Lower:       value,
		LowerStrict: &lowerStrict,
		Ordering:    ordering,
	}
}

// Filter for range filter, i.e. lower <= value <= upper
func FilterRangeIncl(dimension string, lower string, upper string,
	ordering string) *Filter {
	upperStrict := false
	lowerStrict := false
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Lower:       lower,
		Upper:       upper,
		LowerStrict: &lowerStrict,
		UpperStrict: &upperStrict,
		Ordering:    ordering,
	}
}

// Filter for range operation. Takes input lowerStrict and upperStrict
// options. If lowerStrict = true, lower < value else lower <= value.
// If upperStrict = true, value < upper else value <= upper.
func FilterRange(dimension string, lower string, upper string,
	lowerStrict bool, upperStrict bool,
	ordering string) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Lower:       lower,
		Upper:       upper,
		LowerStrict: &lowerStrict,
		UpperStrict: &upperStrict,
		Ordering:    ordering,
	}
}

func FilterCaseSensitiveContains(dimension string, value string) *Filter {
	return &Filter{
		Type:      "search",
		Dimension: dimension,
		SearchSpec: &SearchQuerySpec{
			Type:          "contains",
			Value:         value,
			CaseSensitive: true,
		},
	}
}

func FilterCaseInsensitiveContains(dimension string, value string) *Filter {
	return &Filter{
		Type:      "search",
		Dimension: dimension,
		SearchSpec: &SearchQuerySpec{
			Type:  "insensitive_contains",
			Value: value,
		},
	}
}

func FilterRegex(dimension, pattern string) *Filter {
	return &Filter{
		Type:      "regex",
		Dimension: dimension,
		Pattern:   pattern,
	}
}

func FilterInClause(dimension string, values []interface{}) *Filter {
	return &Filter{
		Type:      "in",
		Dimension: dimension,
		Values:    values,
	}
}

func FilterJavaScript(dimension, function string) *Filter {
	return &Filter{
		Type:      "javascript",
		Dimension: dimension,
		Function:  function,
	}
}

func FilterAnd(filters ...*Filter) *Filter {
	return joinFilters(filters, "and")
}

func FilterOr(filters ...*Filter) *Filter {
	return joinFilters(filters, "or")
}

func FilterNot(filter *Filter) *Filter {
	return &Filter{
		Type:  "not",
		Field: filter,
	}
}

func joinFilters(filters []*Filter, connector string) *Filter {
	// Remove null filters.
	p := 0
	for _, f := range filters {
		if f != nil {
			filters[p] = f
			p++
		}
	}
	filters = filters[0:p]

	fLen := len(filters)
	if fLen == 0 {
		return nil
	}
	if fLen == 1 {
		return filters[0]
	}

	return &Filter{
		Type:   connector,
		Fields: filters,
	}
}
