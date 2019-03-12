package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

var cityDatabase *geoip2.Reader
var countryDatabase *geoip2.Reader
var asnDatabase *geoip2.Reader

func main()  {
	fmt.Println("Starting server")

	cityDatabase, _ = loadDatabase("City")
	countryDatabase, _ = loadDatabase("Country")
	asnDatabase, _ = loadDatabase("ASN")

	r := mux.NewRouter()
	r.HandleFunc("/{database}/{address}", geoIPHandler)

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func geoIPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := net.ParseIP(vars["address"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ip interface{}
	var err error

	switch database := vars["database"]; database {
	case "asn":
		ip, err = asnDatabase.ASN(address)
		if err != nil {
			_,_ = fmt.Fprintln(w, "IP not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	case "city":
		ip, err = cityDatabase.City(address)
		if err != nil {
			_,_ = fmt.Fprintln(w, "IP not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	case "country":
		ip, err = countryDatabase.ASN(address)
		if err != nil {
			_,_ = fmt.Fprintln(w, "IP not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	default:
		_, _ = fmt.Fprintln(w, "Unknown database")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ipJSON, err := json.Marshal(ip)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = fmt.Fprint(w, string(ipJSON))
}

func loadDatabase(name string) (database *geoip2.Reader, err error){
	db, err := geoip2.Open("data/GeoLite2-" + name + ".mmdb")
	if err != nil {
		return nil, errors.New("can't work with it")
	}

	return db, nil
}
