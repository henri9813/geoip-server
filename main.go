package main

import (
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func main()  {
	fmt.Println("Starting server")

	cityDatabase, err := loadDatabase("City")
	if err != nil {
		log.Fatal(err.Error())
	}

	countryDatabase, err := loadDatabase("Country")
	if err != nil {
		log.Fatal(err.Error())
	}

	asnDatabase, err := loadDatabase("ASN")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cityDatabase.City(net.ParseIP("8.8.8.8")))
	fmt.Println(countryDatabase.Country(net.ParseIP("8.8.8.8")))
	fmt.Println(asnDatabase.ASN(net.ParseIP("8.8.8.8")))
}

func loadDatabase(name string) (database *geoip2.Reader, err error){
	db, err := geoip2.Open("data/GeoLite2-" + name + ".mmdb")
	if err != nil {
		return nil, errors.New("can't work with it")
	}

	return db, nil
}
