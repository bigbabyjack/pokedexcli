package urlutils

import (
	"net/url"
	"testing"
)

func TestParseURL(t *testing.T) {
	rawURL := "https://user:pass@example.com:8080/path?query1=123&query2=456#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	if parsedURL.Scheme != "https" {
		t.Errorf("Expected scheme 'https', got '%s'", parsedURL.Scheme)
	}

	if parsedURL.User.Username() != "user" {
		t.Errorf("Expected username 'user', got '%s'", parsedURL.User.Username())
	}

	password, _ := parsedURL.User.Password()
	if password != "pass" {
		t.Errorf("Expected password 'pass', got '%s'", password)
	}

	if parsedURL.Host != "example.com:8080" {
		t.Errorf("Expected host 'example.com:8080', got '%s'", parsedURL.Host)
	}

	if parsedURL.Path != "/path" {
		t.Errorf("Expected path '/path', got '%s'", parsedURL.Path)
	}

	if parsedURL.RawQuery != "query1=123&query2=456" {
		t.Errorf("Expected raw query 'query1=123&query2=456', got '%s'", parsedURL.RawQuery)
	}

	if parsedURL.Fragment != "fragment" {
		t.Errorf("Expected fragment 'fragment', got '%s'", parsedURL.Fragment)
	}
}

func TestModifyURL(t *testing.T) {
	rawURL := "https://example.com/path?query=123"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	parsedURL.Scheme = "http"
	parsedURL.Host = "newexample.com"
	parsedURL.Path = "/newpath"
	parsedURL.RawQuery = "newquery=456"
	parsedURL.Fragment = "newfragment"

	expectedURL := "http://newexample.com/newpath?newquery=456#newfragment"
	if parsedURL.String() != expectedURL {
		t.Errorf("Expected modified URL '%s', got '%s'", expectedURL, parsedURL.String())
	}
}

func TestQueryParams(t *testing.T) {
	rawURL := "https://example.com/path?query1=123&query2=456"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	queryParams := parsedURL.Query()
	if queryParams.Get("query1") != "123" {
		t.Errorf("Expected query1 '123', got '%s'", queryParams.Get("query1"))
	}

	if queryParams.Get("query2") != "456" {
		t.Errorf("Expected query2 '456', got '%s'", queryParams.Get("query2"))
	}
}
