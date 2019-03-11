download:
	$(call download_database,GeoLite2-City)
	$(call download_database,GeoLite2-Country)
	$(call download_database,GeoLite2-ASN)

define download_database
	wget "https://geolite.maxmind.com/download/geoip/database/${1}.tar.gz" -P "/tmp/"
	tar -xzf "/tmp/${1}.tar.gz" -C data --wildcards "*.mmdb" --strip-components 1
	rm "/tmp/${1}.tar.gz"
endef
