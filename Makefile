.PHONY: generate_api test_coverage

generate_api:
	bash scripts/genextapi.sh -p internal/api -s swagger.yaml

test_coverage:
	mkdir -p test
	make -r ./test
	./scripts/run-tests-with-coverage-reports.sh

run:
	docker compose up --build