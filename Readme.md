# go-validation ![build status](https://ci.gitlab.com/projects/6771/status.png?ref=master)

Provides validations for struct fields based on a validation tag

See godoc for more info: http://godoc.org/gitlab.com/bkdsw/go-validation

# Usage Examples

```
type MyType struct {
    Name       string    `validation:"format=regexp:[A-Z][a-z]{3,12}"`
    Email      string    `validation:"format=email"`
    Age        uint      `validation:"min=18"`
    Quantity   uint      `validation:"min=1 max=5"`
    Total      float32   `validation:"min=0"`
}
```
