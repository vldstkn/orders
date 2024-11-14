gen:
	@protoc \
		--proto_path=proto "proto/account.proto" \
		--go_out=proto/gen/account \
		--go_opt=paths=source_relative \
		--go-grpc_out=proto/gen/account \
		--go-grpc_opt=paths=source_relative
run-account:
	@air -c air/.account.toml
