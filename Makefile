.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/account.proto
