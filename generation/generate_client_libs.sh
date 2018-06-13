# Download grafeas.swagger.json
wget https://github.com/grafeas/grafeas/raw/master/v1alpha1/proto/grafeas.swagger.json grafeas.swagger.json

# Download swagger-codegen cli tool
wget https://oss.sonatype.org/content/repositories/snapshots/io/swagger/swagger-codegen-cli/2.4.0-SNAPSHOT/swagger-codegen-cli-2.4.0-20180611.162651-269.jar -O swagger-codegen-cli.jar

# Generate libraries from grafeas.swagger.json
java -jar ./swagger-codegen-cli.jar generate \
    -i ./grafeas.swagger.json \
    -l go \
    -o ../v1alpha1 \
    -c ./config.go.json \

