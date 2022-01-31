package cmd_test

import (
	"testing"

	"github.com/cmmorrow/url/cmd"
)

type uriTest struct {
	uri      cmd.URI
	expected string
}

var buildHostnameTests = []uriTest{
	{cmd.URI{Host: "myhost.com"}, "myhost.com"},
	{cmd.URI{Host: "myhost.com", Port: ""}, "myhost.com"},
	{cmd.URI{Host: "myhost.com", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Host: "myhost.com:5000"}, "myhost.com:5000"},
	{cmd.URI{Host: "myhost.com:8000", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Port: "5000"}, ""},
	{cmd.URI{Host: "myhost.com:5000:8000"}, "myhost.com:5000:8000"},
}

func TestBuildHostname(t *testing.T) {
	for _, test := range buildHostnameTests {
		out := test.uri.BuildHostname()
		if out != test.expected {
			t.Fatalf("Expected '%s', got %s", test.expected, out)
		}
	}
}

func TestBuildHostnameEmpty(t *testing.T) {
	uri := cmd.URI{}
	out := uri.BuildHostname()
	if out != "" {
		t.Fail()
	}
}

var valueTests = []uriTest{
	{cmd.URI{RawParams: []string{"foo=bar", "bar=baz"}}, "bar=baz&foo=bar"},
	{cmd.URI{RawParams: []string{"foo=bar", "bar=baz", "foo=baz"}}, "bar=baz&foo=bar&foo=baz"},
	{cmd.URI{RawParams: []string{"foo=bar", "bar=baz", "=foo"}}, "=foo&bar=baz&foo=bar"},
	{cmd.URI{RawParams: []string{"foo=bar", "bar=baz", "foo="}}, "bar=baz&foo=bar&foo="},
	{cmd.URI{RawParams: []string{"foo=bar baz"}}, "foo=bar+baz"},
	{cmd.URI{RawParams: []string{"foo=%%#$"}}, "foo=%25%25%23%24"},
	{cmd.URI{RawParams: []string{"foo"}}, "foo="},
	{cmd.URI{Params: map[string]interface{}{"foo": "bar", "bar": "baz"}}, "bar=baz&foo=bar"},
	{cmd.URI{Params: map[string]interface{}{"foo": []interface{}{"bar", "baz"}, "bar": "baz"}}, "bar=baz&foo=bar&foo=baz"},
	{cmd.URI{Params: map[string]interface{}{"foo": "bar", "bar": "baz", "": "foo"}}, "=foo&bar=baz&foo=bar"},
	{cmd.URI{Params: map[string]interface{}{"foo": []interface{}{"bar", ""}, "bar": "baz"}}, "bar=baz&foo=bar&foo="},
	{cmd.URI{Params: map[string]interface{}{"foo": "bar baz"}}, "foo=bar+baz"},
	{cmd.URI{Params: map[string]interface{}{"foo": "%%#$"}}, "foo=%25%25%23%24"},
	{cmd.URI{Params: map[string]interface{}{"foo": ""}}, "foo="},
}

func TestBuildValues(t *testing.T) {
	for _, test := range valueTests {
		values := test.uri.BuildValues()
		out := values.Encode()
		if out != test.expected {
			t.Fatalf("Expected '%s', got %s", test.expected, out)
		}

	}
}

func TestBuildValuesEmpty(t *testing.T) {
	uri := cmd.URI{}
	out := uri.BuildValues()
	if len(out) != 0 {
		t.Fatalf("Expected []map, got %v", out)
	}
}

var setUserTests = []uriTest{
	{cmd.URI{User: "wanda"}, "wanda"},
	{cmd.URI{User: "wanda:1234"}, "wanda:1234"},
	{cmd.URI{User: ""}, ""},
	{cmd.URI{User: "wanda:1234:blah"}, "wanda:1234"},
}

func TestSetUser(t *testing.T) {
	for _, test := range setUserTests {
		out := test.uri.SetUser().String()
		if out != test.expected {
			t.Fatalf("Expected '%s', got %s", test.expected, out)
		}

	}
}

var urlTests = []uriTest{
	{cmd.URI{Scheme: "http", Host: "test.com"}, "http://test.com"},
	{cmd.URI{Scheme: "https", Host: "test.com:8888"}, "https://test.com:8888"},
	{cmd.URI{Scheme: "https", Host: "test.com", Port: "8888"}, "https://test.com:8888"},
	{cmd.URI{Scheme: "https", Host: "test.com", Query: "foo=bar&bar=baz"}, "https://test.com?foo=bar&bar=baz"},
	{cmd.URI{Scheme: "https", Host: "test.com", RawParams: []string{"foo=bar", "bar=baz"}}, "https://test.com?bar=baz&foo=bar"},
	{cmd.URI{Scheme: "https", Host: "test.com", Params: map[string]interface{}{"foo": "bar", "bar": "baz"}}, "https://test.com?bar=baz&foo=bar"},
	{cmd.URI{Scheme: "mailto", UriPath: "nobody@email.com"}, "mailto:nobody@email.com"},
	{cmd.URI{Scheme: "http", User: "wanda", Host: "test.com"}, "http://wanda@test.com"},
	{cmd.URI{Scheme: "http", User: "wanda:1234", Host: "test.com"}, "http://wanda:1234@test.com"},
}

func TestAsURL(t *testing.T) {
	for _, test := range urlTests {
		u := test.uri.AsURL()
		out := u.String()
		if out != test.expected {
			t.Fatalf("Expected '%s', got %s", test.expected, out)
		}
	}
}
