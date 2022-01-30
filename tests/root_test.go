package cmd_test

import (
	"testing"

	"github.com/cmmorrow/url/cmd"
)

type uriTest struct {
	uri      cmd.URI
	expected string
}

var hostnameTests = []uriTest{
	{cmd.URI{Domain: "myhost.com"}, "myhost.com"},
	{cmd.URI{Domain: "myhost.com", Port: ""}, "myhost.com"},
	{cmd.URI{Domain: "myhost.com", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Domain: "myhost.com:5000"}, "myhost.com:5000"},
	{cmd.URI{Domain: "myhost.com:8000", Port: "5000"}, "myhost.com:5000"},
	{cmd.URI{Port: "5000"}, ""},
	{cmd.URI{Domain: "myhost.com:5000:8000"}, "myhost.com:5000:8000"},
}

func TestBuildHostname(t *testing.T) {
	for _, test := range hostnameTests {
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
