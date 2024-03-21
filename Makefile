.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/account.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/wechat.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/behaviour.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/message.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/user/score.proto
