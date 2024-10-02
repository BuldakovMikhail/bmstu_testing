test:
	rm -rf ./internal/domain/tests/unit/docs
	export ALLURE_OUTPUT_PATH="../docs" && go test ./internal/domain/tests/unit/... --race --parallel 11

allure:
	allure generate ./internal/domain/tests/unit/docs/allure-results -o allure-reports
	allure serve ./internal/domain/tests/unit/docs/allure-results -p 4000

report: test allure

.PHONY: test allure report