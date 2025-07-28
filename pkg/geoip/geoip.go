package geoip

import (
	"fmt"
	"net"
)

var ASNs = make(map[int]ASN)
var CIDRs = make(map[string]CIDR)

func Init(databasePath string) error {
	if err := loadASNs(fmt.Sprintf("%s/asns.csv", databasePath)); err != nil {
		return fmt.Errorf("unable to open ASN database: %v", err)
	}
	if err := loadCIDRs(fmt.Sprintf("%s/table.txt", databasePath)); err != nil {
		return fmt.Errorf("unable to open CIDR database: %v", err)
	}

	return nil
}

func IPLookup(ip net.IP) (*IP, bool) {
	for _, cidr := range CIDRs {
		if cidr.Cidr.Contains(ip) {
			ipStruct := &IP{
				Ip:          ip.String(),
				CountryCode: cidr.Asn.CountryCode,
				ASNName:     cidr.Asn.Name,
				ASNID:       cidr.Asn.Asn,
				CIDR:        cidr.Cidr.String(),
			}

			return ipStruct, true
		}
	}

	return nil, false
}
