# GobrePro
A headless libreoffice instance wrapped in GRPC.

<img src="gopher.png" alt="A cool looking gopher mascot" width="200"/>

## Why
This would be used when you need to quickly convert document types. For example a Microsoft Word Document to a PDF. 

## References
Libreoffice Headless: https://help.libreoffice.org/latest/he/text/shared/guide/start_parameters.html

GRPC: https://pkg.go.dev/google.golang.org/grpc

## Running
This can largely change based on your configuration, so keep that in mind. 

### Using protoc
In order to make modifications or expand the API you must use the protoc tool 
to generate the protobuf files without having to write them yourself. 
Run the command below to re-generate these files: 

`protoc proto/service.proto --go_out=proto/ --go_opt=paths=source_relative 
--proto_path=proto/ --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative`

Then copy those generated files to proto/gobre and your client. 

### Running with just Golang
`go run .`

### Development 
`docker compose watch`

### Production 
Build and deploy to your cluster based on your configuration. 

## Usage 
1) Once the docker is running find the ip address
2) Make a GPRC request like `<ip address>:8081` as the server and `Gobre/HandleRequest` as the handler.
3) Have a message formatted like
```
{
    "FileData": "raw file data",
    "NewFileType": "pdf",
    "OriginalFileType": "docx"
}
```
4) Make GRPC call.

*Note you can use reflection on your client to get the avaiable service definitions and data structures. 

## Running Tests
```
go test -v -cover ./... -coverprofile test_data/cover.out
go tool cover -html test_data/cover.out -o test_data/cover.html
```

## Contributing 
Please make an issue with a detailed description of your feature or problem. 
