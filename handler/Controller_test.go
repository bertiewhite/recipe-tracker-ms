package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Test hello world returns hello world", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://dummy.org", nil)
		resp := httptest.NewRecorder()

		HelloWorld(resp, req)

		final := resp.Result()

		assert.Equal(t, 200, final.StatusCode)
		assert.Equal(t, HelloWorldMessage, resp.Body.String())
	})

	t.Run("Add Route adds a handler to server", func(t *testing.T) {
		path := "/path"

		count := 0
		increment := func(w http.ResponseWriter, r *http.Request) {
			count++
		}

		h := NewHandler()
		h.AddRoute(path, increment)

		go func() {
			err := h.Start()
			assert.Nil(t, err)
		}()

		_, err := http.Get("http://" + Host + ":" + Port + path)
		assert.Nil(t, err)
		assert.Equal(t, 1, count)
	})
}
