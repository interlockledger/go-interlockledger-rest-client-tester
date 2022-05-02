#  go-interlockledger-rest-client-tester

This is the test program for go-interlockledger-rest-client. It has been
designed to test calls to the API and debug the client code as a writing 
unit-tests for the client is very cumbersome.

It can also be used as an example about how the API can be used.

## How to run it

Just checkout this repository in parallel to the code within this repository.
For example:

```
./go-interlockledger-rest-client
./go-interlockledger-rest-client-teste
```

You can also create a workspace if it is a better way to manage the projects.

## Configuration

This directory contains a very simple program that illustrates how the client
can be used.

This program will require:

- The client certificate with the key in PEM format;
- A proper `config.json` with the URL of the node;

Just run the program to see how the API works.

## License

This library is released under a BSD-3-Clause License.
