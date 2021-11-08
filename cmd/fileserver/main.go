package main

import (
	"net/http"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"
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
