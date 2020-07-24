package main

import (
	"log"
	"net/http"
	"strconv"

	"./api/health"
	"./api/machines"
	"./api/pingapi"
	"./api/wol"

	"github.com/gorilla/mux"
)

func httpHandler() {
	httpRouter := mux.NewRouter().StrictSlash(false)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	httpRouter.PathPrefix("/static/").Handler(s)

	httpRouter.HandleFunc("/health", health.HealthResponse)
	httpRouter.HandleFunc("/health/", health.HealthResponse)

	httpRouter.HandleFunc("/api/machines", machines.GetMachineList)
	httpRouter.HandleFunc("/api/machines/", machines.GetMachineList)
	httpRouter.HandleFunc("/api/machines/item", machines.GetMachineItem)
	httpRouter.HandleFunc("/api/machines/item/", machines.GetMachineItem)

	httpRouter.HandleFunc("/api/wol", wol.SendWolPacket)
	httpRouter.HandleFunc("/api/wol/", wol.SendWolPacket)

	httpRouter.HandleFunc("/api/ping", pingapi.SendPing)
	httpRouter.HandleFunc("/api/ping/", pingapi.SendPing)

	httpRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("location", "/static/index.html")
		w.WriteHeader(http.StatusMovedPermanently)
	})

	port := strconv.Itoa(machines.GetHttpPortNum())
	log.Fatal(http.ListenAndServe(":"+port, httpRouter))
}

func main() {
	machines.InitConfig()
	httpHandler()
}
