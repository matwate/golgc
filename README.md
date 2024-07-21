# Logic Compiler (lgc)

Logic Compiler (lgc) is a tool designed for compiling logic expressions into a form that can be executed or analyzed further. It is built in Go and leverages the power of modern compiler design to process `.lgc` files.

## Getting Started

To use lgc, you need to have Go installed on your machine. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/matwate/lgc.git
cd lgc
```
Build the project:

Usage
To compile a .lgc file, simply run:

Make sure the file has the .lgc extension.

Development
The project is structured as follows:

main.go: Entry point of the application.
root.go: Defines the CLI interface and the main logic for file processing.
parser.go: Contains the logic for parsing .lgc files.
semantic.go: Handles semantic analysis of the parsed logic expressions.
simplify.go, Truthtable.go: Additional modules for logic simplification and truth table generation.
Dependencies
cobra: A library for creating powerful modern CLI applications.
mousetrap: A library used by cobra for Windows support.
Dependencies are managed using Go modules.

Contributing
Contributions are welcome! Please feel free to submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.