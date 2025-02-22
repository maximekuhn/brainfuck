# Brainfuck
This project provides a handful set of tools for the [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) language.  
It currenctly includes:
- a lexer
- a parser
- an interpreter
- an interactive interpreter you can run from the terminal
- a web application running the interpreter (visit https://brainfuck.maximekuhn.dev for a demo)

## Usage as a library
To use the provided tools in Go code, you can simply add it as a dependecy in your `go.mod` file.
```
github.com/maximekuhn/brainfuck
```

## Usage as CLI tool(s)
To build and use the CLI tools, you will need:
- [Go](https://go.dev/) (project version is 1.23.x)
- [Task](https://taskfile.dev/)
- [Templ](https://templ.guide/)

Build the tools:
```shell
task build
```
This will create all executable(s) in the `./bin` directory.  
To use the different tools, simply run them without any arguments and help will be printed.

