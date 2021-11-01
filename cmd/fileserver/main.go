package main

import (
	"fmt"
	"net/http"
	cnst "snakealive/m/pkg/constants"
)

func main() {
	fs := http.FileServer(http.Dir(cnst.StaticPath))
	http.Handle("/", fs)

	fmt.Println("starting file server at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("failed to start file server:", err)
		return
	}
}
