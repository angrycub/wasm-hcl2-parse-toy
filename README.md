This is an implementation heavily inspired by:

<https://golangbot.com/webassembly-using-go/>

This will make a little page that takes an HCL2 job and emits the full json
jobspec based on it.


## Build and deploy the json.wasm file to the server's asset folder

```
cd cmd/wasm/  
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm
```

# Run the server process

```
cd cmd/server/
go run main.go
```

