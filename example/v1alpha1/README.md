# Sample go server that uses the Grafeas go client library

To run this server against the Grafeas reference implementation run the following:

```
go get github.com/Grafeas/Greafeas/samples/server/go-server/api/
go get github.com/Grafeas/client-go
go run grafeas/samples/server/go-server/api/server/main/main.go

# v1alpha1 client-go example relies on the existence of project "best-vuln-scanner"
# Use the following curl command to create the project "best-vuln-scanner"
curl -H "Accept: application/json" -H "Content-type: application/json" -d "{\"name\":\"projects/best-vuln-scanner\"}" -X POST  http://localhost:8080/v1alpha1/projects

go run client-go/example/v1alpha1/main.go
