# ImmutableStringMap

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
| indexkey get | 80.4 ns/op | 28.5 ns/op |
| noindexkey get | 710 ns/op | 28.5 ns/op |
| range | 445 ns/op | 167 ns/op |



