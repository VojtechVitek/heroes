package heroes

// We're using fork of WebRPC to generate server code and clients in TypeScript and Go.
// See https://hub.docker.com/repository/docker/golangcz/webrpc.

//go:generate echo "WebRPC generating ui/src/modules/api.gen.ts"
//go:generate sh -c "docker run --rm -v $PWD:/be golangcz/webrpc:v0.1.1 gen -schema=/be/rpc/api.go -target=ts -client > ./ui/src/modules/rpc.gen.ts"

//go:generate echo "WebRPC generating be/pkg/heroes/client.gen.go"
//go:generate sh -c "docker run --rm -v $PWD:/be golangcz/webrpc:v0.1.1 gen -schema=/be/rpc/api.go -target=go -client -pkg=heroes > ./pkg/heroes/client.gen.go"

//go:generate echo "WebRPC generating be/rpc/server.gen.go"
//go:generate sh -c "docker run --rm -v $PWD:/be golangcz/webrpc:v0.1.1 gen -schema=/be/rpc/api.go -target=go -server -pkg=rpc -extra=noTypes > ./rpc/server.gen.go"
