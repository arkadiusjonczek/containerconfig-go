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
    ruleset := configuration.NewRuleSet()
    ruleset.Required("RequiredField")
    ruleset.Optional("OptionalField", "optional")
    
    customConfiguration, err := configuration.Create[CustomConfiguration](ruleset)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(customConfiguration.RequiredField)
    log.Println(customConfiguration.OptionalField)
}
```

You can find that example located inside the main directory.

## Usage

See the previous example in action:

```shell
$ go run main/main.go                    
2024/05/01 20:55:19 environment variable 'RequiredField' is required but was not found
exit status 1
```

```shell
$ RequiredField=Foo go run main/main.go                 
2024/05/01 20:56:14 Foo
2024/05/01 20:56:14 optional
```

```shell
$ RequiredField=Foo OptionalField=Bar go run main/main.go
2024/05/01 20:56:41 Foo
2024/05/01 20:56:41 Bar
```
