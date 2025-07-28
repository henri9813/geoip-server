package geoip

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func loadASNs(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to open file", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	if _, err := r.Read(); err != nil {
		panic(err)
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, record := range records {
		if len(record) != 4 {
			return fmt.Errorf("invalid asn at line %d: %v", i, record)
		}

		asnId, err := strconv.Atoi(strings.ReplaceAll(record[0], "AS", ""))
		if err != nil {
			return fmt.Errorf("invalid asn id at line %d: %v", i, record)
		}

		ASNs[asnId] = ASN{
			Name:        record[1],
			Asn:         asnId,
			CountryCode: record[3],
		}
	}

	return nil
}

func loadCIDRs(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		if len(data) != 2 {
			return fmt.Errorf("invalid cidr: %s", scanner.Text())
		}

		asn_id, err := strconv.Atoi(data[1])
		if err != nil {
			return fmt.Errorf("invalid asn id for cidr %s: %v", data[0], err)
		}

		_, cidr, err := net.ParseCIDR(data[0])
		if err != nil {
			return fmt.Errorf("invalid cidr %s: %v", data[0], err)
		}

		asn, exist := ASNs[asn_id]
		if !exist {
			ASNs[asn_id] = ASN{
				Name:        "Private",
				Asn:         asn_id,
				CountryCode: "NA",
			}
		}

		CIDRs[data[0]] = CIDR{
			Cidr: cidr,
			Asn:  asn,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
