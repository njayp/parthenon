package httpServer

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port int) error {
	http.HandleFunc("/livez/", livez)

	log.Printf("http listening on port %v", port)
	address := fmt.Sprintf(":%v", port)
	return http.ListenAndServe(address, nil)
}
