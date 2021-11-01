package main

import (
	"net/http"
	logs "snakealive/m/internal/logger"
	cnst "snakealive/m/pkg/constants"
)

func main() {
	logs.BuildLogger()
	logger := logs.GetLogger()

	fs := http.FileServer(http.Dir(cnst.StaticPath))
	http.Handle("/", fs)

	logger.Info("starting server at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Fatal("failed to start file server")
		return
	}
}
