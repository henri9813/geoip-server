package geoip

import "net"

type ASN struct {
	Name        string `json:"name"`
	Asn         int    `json:"asn"`
	CountryCode string `json:"country_code"`
}

type CIDR struct {
	Cidr *net.IPNet `json:"cidr"`
	Asn  ASN        `json:"asn"`
}

type IP struct {
	Ip          string `json:"ip"`
	CountryCode string `json:"country_code"`
	ASNName     string `json:"asn_name"`
	ASNID       int    `json:"asn_id"`
	CIDR        string `json:"cidr"`
}
