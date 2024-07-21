# Logic Compiler (lgc)

Logic Compiler (lgc) is a tool designed for compiling logic expressions into a form that can be executed or analyzed further. It is built in Go and leverages the power of modern compiler design to process `.lgc` files. Specifically, it generates truth tables from these `.lgc` files, allowing for a comprehensive analysis of logic expressions.

## Getting Started

To use lgc, you need to have Go installed on your machine. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/matwate/golgc.git
cd golgc
```

### Build the project

```sh
go build 
```

### Usage

To compile a .lgc file, simply run:
```sh
./lgc <path_to_file>
```
It will print out its truth table and also write to an `.lgout` file

Make sure the file has the .lgc extension.

### Writing `.lgc` files

`.lgc` files are written using a simple syntax to represent logic expressions.

Here's an example of how to write a `.lgc` file


```lgc
a + b * !c  => (d <=> f )
```
### `.lgc` Cheatsheet

`+`: Or \
`*`: And \
`!`: Not \
`=>`: Implies\
`<=>`: If and only If \


### Next Steps
The next feature in development for lgc is to simplify logic expressions using De Morgan's laws and equivalence checking.

This will allow users to optimize their logic expressions for both readability and computational efficiency, making it easier to integrate these expressions into larger systems or analyses.


### Contributing

Contributions are welcome! Please feel free to submit a pull request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
