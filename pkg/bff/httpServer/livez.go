package httpServer

import (
	"fmt"
	"net/http"

	"k8s.io/klog/v2"
)

func livez(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(w, "http ok")
	klog.Error(err)
}
