package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/henri9813/geoip-server/pkg/geoip"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ip/{ipAddress}", getIP).Methods("GET")
	r.HandleFunc("/asn/{asn}", getASN).Methods("GET")

	return r
}

func main() {
	if err := geoip.Init(os.Getenv("DATABASE_PATH")); err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":80", setupRouter()))
}

func getIP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipAddress := vars["ipAddress"]

	ip, ok := geoip.IPLookup(net.ParseIP(ipAddress))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "IP not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ip)
}

func getASN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	asn, err := strconv.Atoi(vars["asn"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Invalid ASN")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(geoip.ASNs[asn])
}
