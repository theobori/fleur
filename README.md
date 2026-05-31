# Gopher application framework

[![build badge](https://github.com/theobori/fleur/actions/workflows/build.yml/badge.svg)](https://github.com/theobori/fleur/actions/workflows/build.yml)

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

The name of the project is fleur, pronounced \\flœʁ\\, which means flower.

This GitHub repository is a KISS project that contains a framework and a CLI. The framework is used to build Gopher applications using the Go language. The goal is to enable users to easily create their own custom Gopher servers in accordance with RFC 1436. The CLI is an example of a Gopher application implemented using the framework.

## Getting started

Before continuing to read the documentation, make sure you have at least [Go](https://go.dev/dl/) version 1.26.3.

### Installation

If you want to use the fleur CLI, you can build and install the binary using the following command.

```bash
go install github.com/theobori/fleur/cmd/fleur
```

## Framework core components

There are a few core components to introduce in order to fully understand how the framework works.

### Route

A route is a pair consisting of a regular expression and a function used by a router.

### Router

This is an interface that manages routes. If one of the regexes matches the request path, the associated function is called. For example, a client requesting a path matching the regex `^/dice$` would receive a GopherMap page containing the result of a dice roll.

### Server

A server is essential since it is responsible for receiving Gopher requests from clients and responding to them. Before serving a file, it will ask its router if there is a routing rule for the path requested by the client. If there is one, the associated function will be called first to respond to the client.

Note that before serving a directory, the server will check if there is a `gophermap` file at its root; if so, it will be served instead of the directory.

Also, only files named `gophermap` and those with the `.gophermap` extension will be evaluated.

### Extension

An extension is a pair consisting of a regex and a function; it is similar to the routes in the router component, except that here, an extension is specific to the gophermap format. Creating an extension allows you, for example, to add, modify, or delete Gophermap item types. For instance, all standard RFC1436 item types are implemented using extensions in the Fleur CLI.

### Extension Manager

This is an interface for managing extensions; it is used by the evaluator to extend the Gophermap.

### Evaluator

An evaluator evaluates GopherMap files when the server is ready to serve them. This evaluation works line by line and allows the GopherMap format to be extended by adding extensions.

## CLI

The fleur CLI is a Gopher application that serves files. It optionally supports personal Gopherspaces on UNIX systems, with the virtual path `/~username/` being converted to `/home/username/public_gopher`.

It also implements an extension that lists files in the current directory. To use it, type `*` on a line. This feature is inspired by gophernicus.

### Help Message

Below is the CLI help message.

```text
Usage of ./fleur:
  -directory string
    	It specifies an input directory path that will be the root of the Gopher server (default "./fleur")
  -domain string
    	Gopher domain (default "localhost")
  -enable-auto-inline-text
    	Relax non compliant text error and convert to gophermap inline text
  -enable-personal-gopherspaces
    	Enable personal Gopherspaces, it allows each user of the system to serve its own files
  -port int
    	Gopher port (default 70)
  -verbose
    	Enable verbose logs.
```

## Examples

You can view the example of a custom Gopher server in the [examples](./examples) folder.

## Contribute

If you'd like to contribute to the project, please follow the instructions provided in the [CONTRIBUTING.md](./CONTRIBUTING.md) file.
