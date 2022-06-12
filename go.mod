module github.com/liuxiong332/go-futu-api

go 1.12

require github.com/golang/protobuf v1.5.0

require (
	github.com/futuopen/ftapi4go/pb v0.0.1
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/futuopen/ftapi4go/pb v0.0.1 => ./pb/github.com/futuopen/ftapi4go/pb
