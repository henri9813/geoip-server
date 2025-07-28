package main

import (
	"encoding/json"
	"github.com/henri9813/geoip-server/pkg/geoip"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var router *mux.Router

func init() {
	err := geoip.Init(os.Getenv("DATABASE_PATH"))
	if err != nil {
		panic(err)
	}
	router = setupRouter()
}

func TestGetIP(t *testing.T) {
	req, err := http.NewRequest("GET", "/ip/137.74.157.33", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response geoip.IP
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Log(geoip.ASNs)
		t.Log(geoip.CIDRs)
		t.Fatal(err)
	}

	assert.Equal(t, "137.74.157.33", response.Ip)
	assert.Equal(t, "FR", response.CountryCode)
	assert.Equal(t, 16276, response.ASNID)
	assert.Equal(t, "137.74.0.0/16", response.CIDR)
}

func TestGetASN(t *testing.T) {
	req, err := http.NewRequest("GET", "/asn/16276", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response geoip.ASN
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "OVH SAS", response.Name)
	assert.Equal(t, 16276, response.Asn)
	assert.Equal(t, "FR", response.CountryCode)
}
