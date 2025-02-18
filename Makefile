format:
	gofmt -s -w .

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/restapi/bootstrap restapi/*.go
	env GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ./bin/websocket/bootstrap websocket/*.go

zip: build
	@zip -j -9 ./bin/restapi.zip ./bin/restapi/bootstrap ./config.yaml
	@zip -j -9 ./bin/websocket.zip ./bin/websocket/bootstrap ./config.yaml

deploy-local: zip
	cdklocal bootstrap
	cdklocal deploy "Local/*" --require-approval never --force

deploy: zip
	cdk deploy "${STAGE}/*" --require-approval never
