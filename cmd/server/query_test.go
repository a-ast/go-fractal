package main

import (
	"net/url"
	"testing"
)

func TestGetIntParameters(t *testing.T) {

	v, _ := url.ParseQuery("h=100")
	query := NewQuery(v)

	got := query.GetInt("h", 0)
	want := 100

	if got != want {
		t.Fatalf(`Got %q, want %q`, got, want)
	}
}

func TestGetIntParametersReturrnError(t *testing.T) {

	v, _ := url.ParseQuery("h=a")
	query := NewQuery(v)

	query.GetInt("h", 100)

	if query.Error() == "" {
		t.Fatalf(`Got %q, want %q`, query.Error(), "Error message")
	}
}
