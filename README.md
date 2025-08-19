# Roman Numeral Converter CLI

## Overview

This is a simple command-line tool written in Go that converts Arabic numerals (integers) into their Roman numeral equivalents. The tool accepts a single integer as a command-line argument or runs in an interactive mode where users can input numbers and receive their Roman numeral representations. It handles numbers from 0 to 3999, with special cases for negative numbers (returns an empty string) and zero (returns "N").

## Features

- Converts positive integers to Roman numerals
- Supports both command-line argument input and interactive mode
- Handles edge cases: negative numbers return an empty string, and zero returns "N"
- Includes comprehensive unit tests to ensure accurate conversions
- Simple and lightweight with no external dependencies beyond the Go standard library

## Prerequisites

- Go 1.21 or higher installed

## Installation

1. Clone the repository:

```bash
git clone https://github.com/thxrsxm/roman.git
```

2. Navigate to the project directory:

```bash
cd roman
```

3. Build the binary:

```bash
go build -o roman
```

## Usage

### Command-line argument mode

Run the tool with a number as an argument to convert it to a Roman numeral:

```bash
./roman 2025
```

**Output:** `MMXXV`

### Interactive mode

Run the tool without arguments to enter interactive mode:

```bash
./roman
```

Input a number when prompted, or type `exit` to quit, or `info` to see the author's name:

```
roman: 2025
MMXXV

roman: info
Created by Erik Andrè Thürsam.

roman: exit
```

## Limitations

- Negative numbers return an empty string
- Non-integer inputs in the CLI will result in an error message

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss improvements or bugs.

## License

This project is licensed under the MIT License.

## Author

Created by thxrsxm.
