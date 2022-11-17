package heroes

//go:generate go2webrpc --version
//go:generate webrpc-gen --version

//go:generate echo "Generating webrpc.json schema from API interface"
//go:generate go2webrpc -schema=./rpc -interface=HeroesServer -out ./rpc/webrpc.json

//go:generate echo "Generating rpc/server.gen.go"
//go:generate webrpc-gen -schema=./rpc/webrpc.json -target=../../webrpc/gen-golang -pkg=rpc -server -types=false -out=./rpc/server.gen.go

//go:generate echo "Generating pkg/hubs/rpc_client.go"
//go:generate webrpc-gen -schema=./rpc/webrpc.json -target=../../webrpc/gen-golang -pkg=heroes -client -out=./pkg/heroes/client.gen.go

//go:generate echo "Generating ../../hubs-frontend/applications/next/src/restApi/api.ts"
//go:generate webrpc-gen -schema=./rpc/webrpc.json -target=typescript -client -out=./ui/src/modules/rpc.gen.ts
