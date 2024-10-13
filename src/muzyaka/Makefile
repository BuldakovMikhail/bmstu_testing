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

