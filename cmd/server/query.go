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

func (p *Query) GetString(name string, defaultValue string) (result string) {
	if !p.url.Has(name) {
		return defaultValue
	}

	return p.url.Get(name)
}

func (p *Query) GetInt(name string, defaultValue int) (result int) {
	if !p.url.Has(name) {
		return defaultValue
	}

	text := p.url.Get(name)

	value, err := strconv.Atoi(text)

	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("parameter %v has invalid syntax for int value", name))
	}

	return value
}

func (p *Query) GetFloat(name string, defaultValue float32) (result float32) {
	if !p.url.Has(name) {
		return defaultValue
	}

	text := p.url.Get(name)

	value, err := strconv.ParseFloat(text, 32)

	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("parameter %v has invalid syntax for float value", name))
	}

	return float32(value)
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
