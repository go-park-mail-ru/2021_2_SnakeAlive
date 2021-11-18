package main

import (
	"net/http"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"
)

func main() {
	l := logs.BuildLogger()

	fs := http.FileServer(http.Dir(cnst.StaticPath))
	http.Handle("/", fs)

	l.Logger.Info("starting server at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		l.Logger.Fatal("failed to start file server")
		return
	}
}
