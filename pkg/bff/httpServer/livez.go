package httpServer

import (
	"fmt"
	"net/http"
)

func livez(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "http ok")
}
