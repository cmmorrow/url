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
	{cmd.URI{Domain: "myhost.com"}, "myhost.com"},
	{cmd.URI{Domain: "myhost.com", Port: ""}, "myhost.com"},
	{cmd.URI{Domain: "myhost.com", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Domain: "myhost.com:5000"}, "myhost.com:5000"},
	{cmd.URI{Domain: "myhost.com:8000", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Port: "5000"}, ""},
	{cmd.URI{Domain: "myhost.com:5000:8000"}, "myhost.com:5000:8000"},
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

var urlTests = []uriTest{
	{cmd.URI{Scheme: "http", Domain: "test.com"}, "http://test.com"},
	{cmd.URI{Scheme: "https", Domain: "test.com:8888"}, "https://test.com:8888"},
	{cmd.URI{Scheme: "https", Domain: "test.com", Port: "8888"}, "https://test.com:8888"},
	{cmd.URI{Scheme: "https", Domain: "test.com", Query: "foo=bar&bar=baz"}, "https://test.com?foo=bar&bar=baz"},
	{cmd.URI{Scheme: "https", Domain: "test.com", RawParams: []string{"foo=bar", "bar=baz"}}, "https://test.com?bar=baz&foo=bar"},
	{cmd.URI{Scheme: "https", Domain: "test.com", Params: map[string]interface{}{"foo": "bar", "bar": "baz"}}, "https://test.com?bar=baz&foo=bar"},
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
