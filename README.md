# degg - Dumb Enum Generator for Go

**Table of Contents**
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Features](#features)
- [Usage](#usage)
- [Declaration Files](#declaration-files)
- [Generated Code](#generated-code)
- [Examples](#examples)
- [Development](#development)
- [License](#license)

## Installation

You can build the binary from source using the `just` utility:

```bash
just build degg
```

The binary will be created in the `./bin/` directory.

## Quick Start

1. Create a declaration file (e.g., `colors.yml`):

```yaml
name: Color
package: color
type: string
values:
  - RED
  - BLUE
  - GREEN
```

2. Generate the enum:

```bash
./bin/degg -i colors.yml -o color_enum.go
```

3. Use the enum in your code:

```go
package main

import (
    "fmt"
    "./color"
)

func main() {
    c := color.RED
    fmt.Println(c)                 // "RED"
    fmt.Println(c.Code())          // "Color.RED"
    fmt.Println(c.Repr())          // "Color("RED")"

    // Validation
    val, err := color.FromName("green")  // Case-insensitive
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(val)               // "GREEN"
}
```

## Features

- **Type Safety**: Generate type-safe enums based on `string` or `int` types
- **Multiple Input Formats**: Support for YAML or JSON declaration files
- **Rich Helper Methods**:
  - Conversion between underlying types and enum values
  - Case-insensitive string matching
  - Validation of enum values
  - String representation methods
- **Automatic Code Formatting**: Generated code is automatically formatted
- **Simple API**: Intuitive declaration format and command-line interface

## Usage

```bash
./bin/degg -i <declaration_file> -o <output_go_file>
```

### Command Line Arguments

| Argument | Description |
|----------|-------------|
| `-i, --input` | Path to the input declaration file (.yaml, .yml, or .json) |
| `-o, --output` | Path to the output Go file (.go) |

## Declaration Files

Declaration files define the properties of the enum to be generated, using YAML or JSON format.

### Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Name of the enum type (e.g., `Color`). Must be a valid Go identifier starting with an uppercase letter. |
| `package` | string | Yes | Go package name for the generated file (e.g., `color`). Must be a valid Go package name. |
| `type` | string | Yes | Underlying type for the enum. Must be either `string` or `int`. |
| `values` | []string | * | List of enum keys. If `type` is `string`, the key itself is used as the value. If `type` is `int`, values are assigned sequentially starting from 0. |
| `named-values` | []map[string]string | * | List of key-value pairs for explicit value assignment. For `int` type, string values must be valid integers. |

> **Note**: You must use either `values` or `named-values`, not both.

### Example Declarations

#### String Enum (YAML)

```yaml
name: Color
package: color
type: string
values:
  - RED
  - BLUE
  - GREEN
  - WHITE
  - BLACK
```

#### Integer Enum (YAML)

```yaml
name: Month
package: month
type: int
values:
  - JANUARY
  - FEBRUARY
  - MARCH
  # ... other months
```

#### Named Values Enum (JSON)

```json
{
  "name": "Capital",
  "package": "capital",
  "type": "string",
  "named-values": [
    {"AUSTRALIA": "Canberra"},
    {"BRAZIL": "Brasilia"},
    {"CANADA": "Ottawa"},
    {"CHINA": "Beijing"}
  ]
}
```

#### Integer With Explicit Values (JSON)

```json
{
  "name": "Status",
  "package": "process",
  "type": "int",
  "named-values": [
    { "PENDING": "10" },
    { "RUNNING": "20" },
    { "FAILED": "-1" }
  ]
}
```

## Generated Code

The generated Go file includes:

### Type Definition

```go
type (
    EnumName underlyingType
)
```

### Constants

```go
const (
    _enumName = "EnumName"

    VALUE1 EnumName = value1
    VALUE2 EnumName = value2
    // ...

    _unknown EnumName = defaultValue
)

var (
    ErrInvalidEnumName = errors.New("invalid value for EnumName, must be one of [VALUE1, VALUE2, ...]")
)
```

### Helper Functions

```go
// Returns a slice of all valid enum values
func Values() []EnumName

// Returns a slice of all valid enum keys as strings
func StringValues() []string

// Converts the underlying type value to the enum type
func FromValue(v underlyingType) (EnumName, error)

// Converts a string (case-insensitive) to the enum type
func FromName(s string) (EnumName, error)
```

### Methods

```go
// Returns the string representation
func (e EnumName) String() string

// Returns the integer representation
func (e EnumName) Int() int

// Returns a string in the format "EnumName.KEY"
func (e EnumName) Code() string

// Returns a string in the format "EnumName(Value)"
func (e EnumName) Repr() string
```

## Examples

The repository includes several examples demonstrating different enum types:

- **String Enum**: [examples/color](examples/color)
- **Integer Enum**: [examples/month](examples/month)
- **Named Values**: [examples/capital](examples/capital)

## Development

```bash
# Build the binary
just build degg

# Run with arguments
just run degg -- -i <input> -o <output>

# Run tests
just test
```

## License

This project is licensed under the [CC0-1.0 License](LICENSE).
