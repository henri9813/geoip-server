all: download test

download:
	wget https://bgp.tools/table.txt -O data/table.txt
	wget https://bgp.tools/asns.csv -O data/asns.csv

test:
	docker compose run --rm test
