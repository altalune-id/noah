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

deploy-dev-id: zip
	cdk bootstrap
	cdk deploy "Dev-ID/*" --require-approval never --force
