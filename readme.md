# Enum Generator

This Go tool generates enum definitions in various programming languages from a plain text input file. The input file should contain enum names and their corresponding integer values.

## Features

- Supports Kotlin, JavaScript, Swift, C#, and Go.
- Configurable output directory and package names.
- Option to export enums based on specified languages.

## Installation

Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

Clone this repository:

```sh
go install https://github.com/yourusername/enum-generator@latest
cd enum-generator
```

## Usage

Run the program with the following command:

```sh
enum-generator -input <path-to-input-file> -package <package-name> -output <output-directory> -langs <comma-separated-languages> -export
```

### Flags

- `-input` (default: `./enums.txt`): Path to the input file containing enum definitions.
- `-package` (default: `EnumPackage`): The package name for the generated files.
- `-output` (default: `lib`): Base directory for the output files.
- `-langs` (default: `kotlin,javascript,swift,csharp,golang`): Comma-separated list of languages to generate enums for.
- `-export` (default: `false`): Whether to export enums or not.

### Example

To generate enums for Kotlin and JavaScript and save them to the `lib` directory with a package name of `MyEnums`, use:

```sh
enum-generator -input ./enums.txt -package MyEnums -output lib -langs kotlin,javascript -export
```

## Input File Format

The input file should contain enum definitions in the following format:

```
EnumName = Value
```

For example:

```
StatusAccepted = 0
StatusRejected = 1
```

## Output

The output files will be generated in the specified output directory under language-specific subdirectories:

- `lib/kotlin/` for Kotlin enums.
- `lib/javascript/` for JavaScript enums.
- `lib/swift/` for Swift enums.
- `lib/csharp/` for C# enums.
- `lib/golang/` for Go enums.

Each file will contain the enums in the respective language's syntax.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to contribute.