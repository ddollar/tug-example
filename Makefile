.PHONY: dev open psql

dev:
	tug start -v

open:
	open http://localhost:5000

psql:
	psql -U postgres -h localhost -p 5100
