# Download v1beta1 grafeas.swagger.json.
wget https://github.com/grafeas/grafeas/raw/master/proto/v1beta1/swagger/grafeas.swagger.json

# Download swagger-codegen CLI tool.
wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.5/swagger-codegen-cli-2.4.5.jar -O swagger-codegen-cli.jar

# Generate libraries from grafeas.swagger.json.
java -jar ./swagger-codegen-cli.jar generate \
    -i ./grafeas.swagger.json \
    -l go \
    -o 0.1.0 \
    -c ./config.go.json \
