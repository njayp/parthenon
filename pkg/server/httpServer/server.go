package httpServer

import (
	"fmt"
	"net/http"

	"k8s.io/klog/v2"
)

func Start(port int) error {
	http.HandleFunc("/livez/", livez)

	klog.Infof("http listening on port %v", port)
	address := fmt.Sprintf(":%v", port)
	return http.ListenAndServe(address, nil)
}
