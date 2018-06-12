# Generating the Library


## About OpenAPI
The Grafeas libraries were created using the standard [Swagger (aka OpenAPI)](https://swagger.io/) specification.
This means that the API was defined once in a JSON file, which can then be used to auto-generate libraries for
all supported languages at once. This makes maintainence and upgrades much easier to implement

Along with the auto-generated portions of the library, we may have made some manual changes to the library files. 
These files are preserved between library generations through the [.swagger-codegen-ignore](../v1alpha1/.swagger-codegen-ignore) file. For this reason,
it is expected that new versions of the library are generated on top of previous ones, to preserve the manual changes.

## Regenerating the Library

To generate a new version of the library, simply run the following command:

```./generate_client_libs.sh```

This command will download the latest [OpenAPI spec file](https://github.com/grafeas/grafeas/blob/master/v1alpha1/proto/grafeas.swagger.json) from the main Grafeas repo, download the necessary version of
[swagger-codegen](https://github.com/swagger-api/swagger-codegen), and generate the library using the appropriate configuration options
