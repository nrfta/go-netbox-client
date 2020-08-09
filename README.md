# go-netbox-client

This code is generated using [swagger-codegen](https://github.com/swagger-api/swagger-codegen).
The build action is triggered off of the `ci` branch. That action will commit the generated code
to `master`.

## Local Builds

### Get Code Generator

The code generate requires Java 8 or above.

```shell script
wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.15/swagger-codegen-cli-2.4.15.jar -O swagger-codegen-cli.jar
```

### Validate the Swagger File

```shell script
java -jar ./ci/swagger-codegen-cli.jar validate -i netbox_swagger.json
```

### Generate Code

```shell script
java -jar swagger-codegen-cli.jar generate -l go \
            -c config.json \
            -i netbox_swagger.json \
            -o /path/to/output/dir
```
