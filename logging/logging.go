package logging

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func LogHttp(r *http.Request) {
	log.Printf("%s %s %s %s", r.Method, r.URL, r.RemoteAddr, r.UserAgent())
}

func LogHttpErr(w http.ResponseWriter, r *http.Request, err error, status int) {
	log.Errorf("Returned error \"%s\" (HTTP %s) to %s", err.Error(), http.StatusText(status), r.RemoteAddr)
	serialErr, _ := json.Marshal(err.Error())
	http.Error(w, string(serialErr), status)
}
