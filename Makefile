.PHONY: test allure report

test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/mnt/c/sem7/bmstu_testing/src/muzyaka" && go test ./internal/domain/tests/... --race --parallel 11

allure:
	[ -d allure-reports ] && cp -rf allure-reports/history allure-results || :
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
 	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	go test ./internal/domain/tests/unit/... --race

local-unit:
	export ALLURE_OUTPUT_PATH="/mnt/c/sem7/bmstu_testing/src/muzyaka" && \
 	go test ./internal/domain/tests/unit/... --race

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
	go test ./internal/domain/tests/integration/... --race

local-integration:
	export ALLURE_OUTPUT_PATH="/mnt/c/sem7/bmstu_testing/src/muzyaka" && \
	go test ./internal/domain/tests/integration/... --race

ci-e2e:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
    export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test ./e2e/... --race

local-e2e:
	export ALLURE_OUTPUT_PATH="/mnt/c/sem7/bmstu_testing/src/muzyaka" && \
	go test ./e2e/... --race


ci-concat-reports:
	mkdir allure-results
	cp unit-allure/* allure-results/
	cp integration-allure/* allure-results/
	cp e2e-allure/* allure-results/
	cp environment.properties allure-results

.PHONY: test allure report ci-unit local-unit ci-integration local-integration ci-e2e local-e2e ci-concat-reports