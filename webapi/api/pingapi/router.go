package pingapi

import (
	"net/http"

	"github.com/ytsurui/wol-on-web/webapi/api"
)

func SendPing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "GET":
		status := execPing(w, r)
		api.LogPrint(r, status)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}
