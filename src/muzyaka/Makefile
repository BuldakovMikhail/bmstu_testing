#test:
#	export ALLURE_OUTPUT_PATH="../docs" && go test ./internal/domain/tests/unit/... --race --parallel 11
#
#allure:
#	allure generate ./internal/domain/tests/unit/docs/allure-results  --clean -o allure-reports
#	mkdir -p ./internal/domain/tests/unit/docs/allure-results/history && cp -r ./allure-report/history/* ./internal/domain/tests/unit/docs/allure-results/history/ || true
#	allure serve ./internal/domain/tests/unit/docs/allure-results -p 4000
#
#report: test allure
#
#.PHONY: test allure report

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

