# Download v1alpha1 grafeas.swagger.json.
wget https://github.com/grafeas/grafeas/raw/master/v1alpha1/proto/grafeas.swagger.json

# Download swagger-codegen CLI tool.
wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar -O swagger-codegen-cli.jar

# Generate libraries from grafeas.swagger.json.
java -jar ./swagger-codegen-cli.jar generate \
    -i ./grafeas.swagger.json \
    -l go \
    -o ../v1alpha1 \
    -c ./config.go.json \
