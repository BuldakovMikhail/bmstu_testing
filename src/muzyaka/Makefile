test:
	rm -rf ./internal/domain/tests/unit/docs
	export ALLURE_OUTPUT_PATH="../docs" && go test ./internal/domain/tests/unit/... --race --parallel 11

allure:
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

.PHONY: test allure report