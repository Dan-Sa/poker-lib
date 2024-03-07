compile: clean
	protoc -I proto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./proto/poker.proto
	cd pb && go mod init "github.com/Dan-Sa/poker-lib/pb"
	cd pb && go get -v ./...
	cd pb && go mod tidy

clean:
	rm -rf pb/
	mkdir pb/
