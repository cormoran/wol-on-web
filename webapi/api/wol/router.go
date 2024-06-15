package wol

import (
	"net/http"

	"github.com/ytsurui/wol-on-web/webapi/api"
)

func SendWolPacket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "HEAD":
	case "OPTIONS":
	case "POST":
		statusCode := sendWolPacket(w, r)
		api.LogPrint(r, statusCode)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	api.LogPrint(r, http.StatusMethodNotAllowed)
}
