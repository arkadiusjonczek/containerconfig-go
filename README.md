# containerconfig-go

Create a go configuration struct from your container configuration environment variables.

Set rules to define if an environment variables is required or optional.

## Example

Use the library as shown below:

```go
import (
	"log"
	
	"github.com/arkadiusjonczek/containerconfig-go/configuration"
)

type CustomConfiguration struct {
    RequiredField string
    OptionalField string
}

func main() {
    builder := configuration.NewBuilder[CustomConfiguration]()
    builder.Env("RequiredField")
    builder.Env("OptionalField").SetOptional().WithDefault("optional")
    
    customConfiguration, err := builder.Build()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(customConfiguration.RequiredField)
    log.Println(customConfiguration.OptionalField)
}
```

You can find that example located inside the [main](main) directory.

## Usage

See the previous example in action:

```shell
$ go run main/main.go                    
2024/05/02 08:53:05 environment variable 'RequiredField' is required but was not found
exit status 1
```

```shell
$ RequiredField=Foo go run main/main.go                 
2024/05/02 08:53:20 Foo
2024/05/02 08:53:20 optional
```

```shell
$ RequiredField=Foo OptionalField=Bar go run main/main.go
2024/05/02 08:53:33 Foo
2024/05/02 08:53:33 Bar
```

## Custom Fields

If your configuration field name is different from your environment variable or file name then you can use custom field like:

```go
type CustomConfiguration struct {
	CustomField string
}
```

```go
builder := configuration.NewBuilder[CustomConfiguration]()
builder.Env("CUSTOM_FIELD").UseFieldName("CustomField")
```

The builder will then look for the environment variable `CUSTOM_FIELD` and save the value in the field `CustomField` inside of the `CustomConfiguration`.
