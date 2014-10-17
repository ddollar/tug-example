.PHONY: build dev open psql release

build:
	docker build -t ddollar/tug-example .

dev:
	tug start -v

open:
	open http://localhost:5000

psql:
	tug run postgres psql -U postgres

release: build
	docker push ddollar/tug-example
