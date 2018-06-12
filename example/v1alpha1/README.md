# Local Grafeas Testing Sample

To help support testing, we have developed a Grafeas test server you can run locally. 
This directory holds a sample go script that can be against the test server. 
To see it in action, run the following commands:

```
# set where to download the go projects
GOPATH=$(pwd)

# download projects
go get github.com/Grafeas/Grafeas/samples/server/go-server/api/server/main
go get github.com/Grafeas/client-go/example/v1alpha1

# run the Grafeas test server
go run src/github.com/Grafeas/Grafeas/samples/server/go-server/api/server/main/main.go

# run the client example (in another terminal window)
go run src/github.com/grafeas/client-go/example/v1alpha1/main.go
