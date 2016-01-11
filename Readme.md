# go-validation [![Build Status](https://travis-ci.org/BakedSoftware/go-validation.svg?branch=master)](https://travis-ci.org/BakedSoftware/go-validation)

Provides validations for struct fields based on a validation tag

See godoc for more info: http://godoc.org/github.com/BakedSoftware/go-validation

# Usage Examples

```
type MyType struct {
    Name       string    `validation:"format=regexp:[A-Z][a-z]{3,12}"`
    Email      string    `validation:"format=email"`
    Category   string    `validation:"min_length=5 max_length=10"`
    Age        uint      `validation:"min=18"`
    Quantity   uint      `validation:"min=1 max=5"`
    Total      float32   `validation:"min=0"`
}
```
