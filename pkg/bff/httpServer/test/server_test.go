package test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/njayp/parthenon/pkg/bff/httpServer"
)

func TestHTTP(t *testing.T) {
	port := 8080
	go httpServer.Start(port)
	t.Run("livez get request", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%v/livez/", port))
		if err != nil {
			t.Fatal(err.Error())
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err.Error())
		}
		text := string(body)
		expected := "ok"
		if !strings.Contains(text, expected) {
			t.Errorf("response text '%s' does not contian '%s'", text, expected)
		}
	})
}
