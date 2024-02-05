generate-openapi:
	# generate go client
	rm -rf client
	openapi-generator-cli generate --input-spec openapi.yaml  \
        --generator-name go \
        --output client \
		--additional-properties=isGoSubmodule=true \
		--additional-properties=packageName=client \
		--additional-properties=structPrefix=true
	rm client/go.mod
	rm client/go.sum
	rm client/test/* # remove tests in client sdk

	# generate documentation
	openapi-generator-cli generate -i openapi.yaml  \
	--generator-name markdown
	
	# format go code
	go fmt ./...
	goimports -w .
