gen:
	@protoc \
		--proto_path=proto "proto/account.proto" \
		--go_out=pkg/api/account \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/api/account \
		--go-grpc_opt=paths=source_relative
run-account:
	@air -c air/.account.toml
run-api:
	@air -c air/.api.toml
