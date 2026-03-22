# gob

A Go utility library providing common object, array, map, string, and cache operations with generics support.

Co-created with [@johnjerrico](https://github.com/johnjerrico).

## Installation

```
go get github.com/louvri/gob
```

Requires Go 1.25.0+

## Packages

### `arr` — Slice utilities

```go
import "github.com/louvri/gob/arr"

// Search (generic)
arr.Search([]string{"a", "b", "c"}, "b") // 1
arr.Search([]int{10, 20, 30}, 20)         // 1

// Insert at index
arr.Insert([]string{"a", "c"}, "b", 1) // ["a", "b", "c"]

// Copy with exclusion
arr.Copy([]string{"a", "b", "c"}, []string{"b"}) // ["a", "c"]

// Symmetric difference
arr.Unique([]string{"a", "b", "c"}, []string{"b", "c", "d"}) // ["a", "d"]

// Trim empty strings
arr.Trim([]string{"a", "", " ", "b"}) // ["a", "b"]

// Boolean lookup index
idx := arr.Index([]string{"x", "y"})
idx["x"] // true

// Functional helpers
arr.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })             // [2, 4, 6]
arr.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })   // [2, 4]
arr.Reduce([]int{1, 2, 3}, 0, func(acc, n int) int { return acc + n }) // 6
```

### `mp` — Map utilities

```go
import "github.com/louvri/gob/mp"

// Copy excluding keys
mp.Copy(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"b"})
// map[a:1 c:3]

// Copy only specified keys
mp.CopyOnly(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"})
// map[a:1 c:3]

// Find first existing key
key, found := mp.Search(map[string]int{"a": 1, "b": 2}, []string{"z", "b"})
// "b", true
```

### `str` — String utilities

```go
import "github.com/louvri/gob/str"

str.ExtractNumberFromText("abc123def")        // "123"
str.ExtractAlfaNumericFromText("Hello, World! 123") // "Hello World 123"
str.SplitOnNotEmpty("a,,b,,c", ",")           // ["a", "b", "c"]
str.RemoveUncommonCharacters("hello\u00A1world") // "helloworld"
```

### `object` — Reflection-based struct operations

```go
import "github.com/louvri/gob/object"

type User struct {
    Name  string
    Age   int64
    Score float64
}

// Get/set fields via reflection
ref := object.Ref(&user)
prop := object.Prop(ref, "Name")
val := object.Get(prop)

// Check if zero value
object.IsEmpty(prop) // true if ""

// Assign with type coercion
object.Assign(ref.FieldByName("Age"), "Age", int64(25))

// Patch non-empty fields from one struct to another
object.Patch(&target, &source)

// Compare on non-empty filter fields
object.EqualOnNonEmpty(&data, &filter)

// Flatten struct to string slice
object.Flatten(&user, []string{"Name", "Score"}) // ["alice", "3.14"]

// Iterate fields
object.Iterate(&user, func(key string, value any, isempty bool) {
    // ...
})

// Random text generation (crypto/rand)
object.GenerateRandomText(16) // "aBcDeFgHiJkLmNoP"

// Unique ID generation (UUID + FNV hash)
object.GenerateRunningNumbers() // 3129573891
```

### `cache` — Thread-safe memoization

```go
import "github.com/louvri/gob/cache"

c := cache.New(5 * time.Minute)
c.Memoize("expensive result")
if c.Exists() {
    val := c.Get()
}
```

### `node` — Machine identification

```go
import "github.com/louvri/gob/node"

id, err := node.GetID() // combines /etc/machine-id with timestamp
```

## License

See [LICENSE](LICENSE).
