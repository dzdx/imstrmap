# ImmutableStringMap


Compared with the string map, ImmutableStringMap memory usage is small, you can specify the keys for index accelerate Get method.

More suitable for the map that resident memory, and will not change after initialization 


## example

```
var (
    indexerFactory = imstrmap.NewIndexerFactory([]string{"indexkey"})
)

func run(){
    m := imstrmap.FromMap(map[string]string{...}, indexerFactory)
    value, ok := m.Get("a")
    m.Range(func(key string, value string){
        fmt.Println(key, value)
    }
}
```

## benchmark

goos: darwin

goarch: amd64


### testdata:
```
{
    "locality": "vlocality",
    "a": "a", 
    "b": "b", 
    "c": "c", 
    "das": "das", 
    "huhqw": "huhqw", 
    "uyoqw": "uyoqw", 
    "y9qw": "y9qw", 
    "juioq": "juioq", 
    "qqeq": "qqeq", 
    "vqrqasas": "vqrqasas", 
    "hqw": "hqw", 
    "asdqw": "asdqw", 
    "asqwqwe": "asqwqwe"
}
```

### memory

| count | ImmutableStringMap | map[string]string |
|---|---|---|
| 10000 |  3735552 | 12820480 |

### time

|action | ImmutableStringMap | map[string]string |
|---|---|---|
| indexkey get | 77.2 ns/op | 28.5 ns/op |
| noindexkey get | 324 ns/op | 28.5 ns/op |
| range | 445 ns/op | 167 ns/op |
| toMap | 1096 ns/op |  |


### API

|name | intro |
|---|---|
| NewIndexerFactory | specify keys to create an index factory that can accelerate Get method|
| FromMap | create an ImmutableStringMap from a map[string]string and an indexerfactory|
| .Get | get value by key store in this map |
| .Range | iterate pass k, v to the func in args |
| .Map | convert current ImmutableStringMap to a map[string]string |
