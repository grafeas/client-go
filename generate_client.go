// Downloads v1beta1 grafeas.swagger.json, swagger-codegen CLI tool,
// and generates Go client library from the downloaded Swagger spec.

//go:generate wget https://github.com/grafeas/grafeas/raw/master/proto/v1beta1/swagger/grafeas.swagger.json
//go:generate wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.5/swagger-codegen-cli-2.4.5.jar -O swagger-codegen-cli.jar
//go:generate java -jar ./swagger-codegen-cli.jar generate -i ./grafeas.swagger.json -l go -o 0.1.0 -c ./config.go.json

package generate_client
