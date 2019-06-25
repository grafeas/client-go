# Grafeas - Go Client Library

[Grafeas](https://github.com/grafeas/grafeas) is an artifact metadata API.
This is a client library in Go that utilizes the Grafeas API.

## Generating the Library

The Grafeas client libraries are generated using the standard [Swagger (aka OpenAPI)](https://swagger.io/) specification.
The API is defined once in a JSON file, which is then used to auto-generate libraries for all supported languages at once.
This simplifies maintainance and upgrades.

Along with the auto-generated portions of the library, there may be manual changes to the library files.
These files are preserved between library generations through the [.swagger-codegen-ignore](../0.1.0/.swagger-codegen-ignore) file.
Therefore, it is expected that new versions of the library are generated on top of previous ones, to preserve the manual changes.

## Regenerating the Library

To generate a new version of the library, run the following command:

```bash
go generate
```

This command will download the latest [OpenAPI spec file](https://github.com/grafeas/grafeas/blob/master/proto/v1beta1/swagger/grafeas.swagger.json) from the main Grafeas repo, download the necessary version of
[swagger-codegen](https://github.com/swagger-api/swagger-codegen), and generate the library using the appropriate configuration options.

### Upgrading client library version

The [config.go.json](config.go.json) sets auto-generation options.
When upgrading the client version, make sure to update `packageVersion` field
before running the `go generate` command. All available options can be found by
running:

```bash
java -jar swagger-codegen-cli.jar config.go.json -l go
```
