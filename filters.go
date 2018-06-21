package godruid

// documentation for druid filtering: http://druid.io/docs/latest/querying/filters.html

type Filter struct {
	Type        string      `json:"type"`
	Dimension   string      `json:"dimension,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Pattern     string      `json:"pattern,omitempty"`
	Function    string      `json:"function,omitempty"`
	Field       *Filter     `json:"field,omitempty"`
	Fields      []*Filter   `json:"fields,omitempty"`
	Lower       string      `json:"lower,omitempty"`
	Upper       string      `json:"upper,omitempty"`
	LowerStrict bool        `json:"lowerStrict,omitempty"`
	UpperStrict bool        `json:"upperStrict,omitempty"`
	Ordering    string      `json:"ordering,omitempty"`
}

func FilterSelector(dimension string, value interface{}) *Filter {
	return &Filter{
		Type:      "selector",
		Dimension: dimension,
		Value:     value,
	}
}

func FilterRegex(dimension, pattern string) *Filter {
	return &Filter{
		Type:      "regex",
		Dimension: dimension,
		Pattern:   pattern,
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

const (
	filterOrderingLexicographic = "lexicographic"
	filterOrderingAlphanumeric  = "alphanumeric"
	filterOrderingNumeric       = "numeric"
	filterOrderingStrlen        = "strlen"
)

// FilterGreaterEqual: [dimension] >= [value]
func FilterGreaterEqual(dimension string, value string, ordering string) *Filter {
	return &Filter{
		Type:      "bound",
		Dimension: dimension,
		Lower:     value,
		Ordering:  ordering,
	}
}

// FilterGreater: [dimension] > [value]
func FilterGreater(dimension string, value string, ordering string) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Lower:       value,
		LowerStrict: true,
		Ordering:    ordering,
	}
}

// FilterLowerEqual: [dimension] <= [value]
func FilterLowerEqual(dimension string, value string, ordering string) *Filter {
	return &Filter{
		Type:      "bound",
		Dimension: dimension,
		Upper:     value,
		Ordering:  ordering,
	}
}

// FilterLower: [dimension] < [value]
func FilterLower(dimension string, value string, ordering string) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Upper:       value,
		UpperStrict: true,
		Ordering:    ordering,
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
