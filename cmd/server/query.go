package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	url    url.Values
	errors []error
}

func NewQuery(url url.Values) *Query {
	return &Query{
		url:    url,
		errors: make([]error, 0),
	}
}

// func (p Query) GetString(name string) string {

// }

func (p *Query) GetInt(name string, defaultValue int) (result int) {
	if !p.url.Has(name) {
		return defaultValue
	}

	text := p.url.Get(name)

	value, err := strconv.Atoi(text)

	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("Parameter %v has invalid syntax for int value", name))
	}

	return value
}

func (p *Query) Error() string {
	if len(p.errors) == 0 {
		return ""
	}

	result := ""
	for _, err := range p.errors {
		result = result + " " + err.Error()
	}

	return strings.TrimSpace(result)
}
