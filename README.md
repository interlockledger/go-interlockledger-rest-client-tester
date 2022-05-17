#  go-interlockledger-rest-client-tester

This is the test program for **go-interlockledger-rest-client**. It has been
designed to test calls to the API and debug the client code as a writing 
unit-tests for the client is very cumbersome.

It can also be used as an example about how the
**go-interlockledger-rest-client** API can be used. You will also need the
**InterlockLedger REST API** documentation as those examples explain how to
call the APIs but do not provide extensive information about how they work.

## How to run it

As a test tool, this program will compile if and only if it is parallel to the
source code of **go-interlockledger-rest-client**. For example:

```
./go-interlockledger-rest-client
./go-interlockledger-rest-client-teste
```

You can run the code from the directory `go-interlockledger-rest-client-teste`
by running the command:

```
$ go run cmd/main.go
```

This will print the documentation rquired to use the CLI commands that
implement each of the REST API calls.

## Configuration

This program requires a json configuration file with the following structure:

```json
{
    "basePath": "https://server:port",
    "keyFile": "key.pem",
    "certFile": "cert.pem"    
}
```

If this file is called `config.json`, it will be automatically loaded during the
startup of the program.

The files pointed by `keyFile` and `certFile` are the certificate and the key 
file of the certificate authorized to use the node.

## License

This library is released under a BSD-3-Clause License.
