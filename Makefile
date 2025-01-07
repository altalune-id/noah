format:
	gofmt -s -w .

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/restapi/bootstrap restapi/main.go

zip: build
	@zip -j -9 ./bin/restapi.zip ./bin/restapi/bootstrap

deploy-local: zip
	cdklocal bootstrap
	cdklocal deploy "Local/*" --require-approval never --force

deploy: zip
	echo "stage: ${STAGE}"
	echo "account id: ${AWS_ACCOUNT_ID}"
	echo "region: ${AWS_REGION}"
	echo "role: ${AWS_ROLE_OIDC}"
