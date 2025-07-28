package geoip

import "testing"

func TestLoadASNsInvalidRowLength(t *testing.T) {
	err := loadASNs("/app/mocks/asns-invalid-row-length.csv")
	if err == nil {
		t.Error("should not be able to load ASNs with invalid row length")
	}
}

func TestLoadASNsInvalidID(t *testing.T) {
	err := loadASNs("/app/mocks/asns-bad-id.csv")
	if err == nil {
		t.Error("should not be able to load ASNs with invalid id")
	}
}

func TestLoadCIDRsInvalidRowLength(t *testing.T) {
	err := loadCIDRs("/app/mocks/table-invalid-row-length.txt")
	if err == nil {
		t.Error("should not be able to load CIDRs with invalid row length")
	}
}

func TestLoadCIDRsInvalidCIDR(t *testing.T) {
	err := loadCIDRs("/app/mocks/table-invalid-cidr.txt")
	if err == nil {
		t.Error("should not be able to load CIDRs with invalid row length")
	}
}

func TestLoadCIDRsWrongASN(t *testing.T) {
	err := loadCIDRs("/app/mocks/table-invalid-asn.txt")
	if err == nil {
		t.Error("all CIDRs should have a existing ASN")
		t.Error(ASNs)
		t.Error(CIDRs)
	}
}
