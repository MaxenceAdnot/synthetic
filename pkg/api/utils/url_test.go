package utils_test

import (
	"testing"

	"github.com/maxenceadnot/synthetic/pkg/api/utils"

	"github.com/stretchr/testify/assert"
)

var testURLs = []struct {
	rawURL     string
	isValidURL bool
}{
	{
		rawURL:     "http://example.com",
		isValidURL: true,
	},
	{
		rawURL:     "https://example.com",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/path",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/path?query",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/path?query#fragment",
		isValidURL: true,
	},
	{
		rawURL:     "http://user:pass@www.example.com/",
		isValidURL: true,
	},
	{
		rawURL:     "http://userpass:\\@www.example.com/",
		isValidURL: false,
	},
	{
		rawURL:     "",
		isValidURL: false,
	},
	{
		rawURL:     "example.com",
		isValidURL: false,
	},
	{
		rawURL:     "http://example.dev/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.中文网/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com:8080",
		isValidURL: true,
	},
	{
		rawURL:     "ftp://example.com",
		isValidURL: false,
	},
	{
		rawURL:     "ftp.example.com",
		isValidURL: false,
	},
	{
		rawURL:     "http://127.0.0.1/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/?query=%2F",
		isValidURL: true,
	},
	{
		rawURL:     "http://localhost:3000/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/?query",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com?query",
		isValidURL: true,
	},
	{
		rawURL:     "http://www.xn--froschgrn-x9a.net/",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.com/a-",
		isValidURL: true,
	},
	{
		rawURL:     "http://example.c_o_m/",
		isValidURL: false,
	},
	{
		rawURL:     "http://_example.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://example_example.com/",
		isValidURL: true,
	},
	{
		rawURL:     "xyz://example.com",
		isValidURL: false,
	},
	{
		rawURL:     ".com",
		isValidURL: false,
	},
	{
		rawURL:     "invalid.",
		isValidURL: false,
	},
	{
		rawURL:     "http://example.com/~user",
		isValidURL: true,
	},
	{
		rawURL:     "mailto:someone@example.com",
		isValidURL: false,
	},
	{
		rawURL:     "/abs/test/dir",
		isValidURL: false,
	},
	{
		rawURL:     "./rel/test/dir",
		isValidURL: false,
	},
	{
		rawURL:     "http://example-.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://-example.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://example_.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://.example.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://_example.com/",
		isValidURL: false,
	},
	{
		rawURL:     "http://example.com:80:80/",
		isValidURL: false,
	},
	{
		rawURL:     "http://example.com://8080",
		isValidURL: false,
	},
}

// TestIsURL tests the IsURL function with various URLs.
func TestIsURL(t *testing.T) {
	t.Run("Check if various urls are valid", func(t *testing.T) {
		for _, test := range testURLs {
			ret := utils.IsURL(test.rawURL)

			assert.Equal(t, test.isValidURL, ret, test.rawURL)
		}
	})
}

func BenchmarkIsUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IsURL("http://example.com")
	}
}

func TestIsHTTPMethod(t *testing.T) {
	t.Run("Check if various HTTP methods are valid", func(t *testing.T) {
		assert.True(t, utils.IsHTTPMethod("GET"))
		assert.True(t, utils.IsHTTPMethod("POST"))
		assert.True(t, utils.IsHTTPMethod("PUT"))
		assert.True(t, utils.IsHTTPMethod("DELETE"))
		assert.True(t, utils.IsHTTPMethod("PATCH"))
		assert.True(t, utils.IsHTTPMethod("HEAD"))
		assert.True(t, utils.IsHTTPMethod("OPTIONS"))
		assert.False(t, utils.IsHTTPMethod("TRACE"))
		assert.False(t, utils.IsHTTPMethod("CONNECT"))
		assert.False(t, utils.IsHTTPMethod("FOO"))
	})
}

func BenchmarkIsHTTPMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IsHTTPMethod("GET")
	}
}

func TestIsHTTPStatusCode(t *testing.T) {
	t.Run("Check if various HTTP status codes are valid", func(t *testing.T) {
		assert.True(t, utils.IsHTTPStatusCode(200), "200")
		assert.True(t, utils.IsHTTPStatusCode(201), "201")
		assert.True(t, utils.IsHTTPStatusCode(202), "202")
		assert.True(t, utils.IsHTTPStatusCode(300), "300")
		assert.True(t, utils.IsHTTPStatusCode(400), "400")
		assert.True(t, utils.IsHTTPStatusCode(500), "500")
		assert.True(t, utils.IsHTTPStatusCode(100), "100")
		assert.False(t, utils.IsHTTPStatusCode(600), "600")
		assert.False(t, utils.IsHTTPStatusCode(30), "30")
	})
}

func BenchmarkIsHTTPStatusCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.IsHTTPStatusCode(200)
	}
}
