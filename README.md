## Description

The TextSubstitutor package provides functionality for replacing placeholders in text with corresponding values. Only 31 lines of code

## Installation

To install the TextSubstitutor package, you can use the `go get` command:

```
go get github.com/singoesdeep/textsubstitutor
```

## Usage

Here is an example of how to use the TextSubstitutor package:

```go
package main

import (
	"fmt"
	"github.com/singoesdeep/textsubstitutor"
)

func main() {
	startsWith := "{{"
	endsWith := "}}"
	substitutor := textsubstitutor.NewTextSubstitutor(startsWith, endsWith)

	data := []textsubstitutor.Data{
		{Key: "name", Value: "singoesdeep"},
		{Key: "age", Value: "23"},
	}

	text := "My name is {{name}} and I am {{age}} years old."
	result := substitutor.ReplaceData(text, data)

	fmt.Println(result)
}
```

Output:
```
My name is singoesdeep and I am 23 years old.
```

## API

### type TextSubstitutor

```go
type TextSubstitutor struct {
	StartsWith string
	EndsWith   string
}
```

The TextSubstitutor struct represents a text substitutor with a specified starting and ending delimiter.

### func NewTextSubstitutor

```go
func NewTextSubstitutor(startsWith string, endsWith string) *TextSubstitutor
```

NewTextSubstitutor creates a new instance of the TextSubstitutor with the specified starting and ending delimiters.

### func (ur *TextSubstitutor) ReplaceData

```go
func (ur *TextSubstitutor) ReplaceData(text string, data []Data) string
```

ReplaceData replaces placeholders in the given text with the corresponding values from the provided data.

- `text`: The text containing placeholders.
- `data`: An array of Data objects representing the key-value pairs for substitution.

## Contributing

If you want to contribute to this package or have suggestions for improvements, please feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/your-username/textsubstitutor).

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

You can use this template to generate your README file for the TextSubstitutor package. Make sure to replace `"your-username"` with your actual GitHub username in the installation and contributing sections.