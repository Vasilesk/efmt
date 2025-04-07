# efmt

`efmt` is a Go library for enhanced error handling that allows you to attach structured metadata (key-value pairs) to errors. This helps to enrich error context while maintaining compatibility with Go's standard error interface.

## Installation

```bash
go get github.com/vasilesk/efmt
```

## Features

- Create errors with structured metadata
- Wrap existing errors with additional context
- Extract metadata from errors as key-value pairs
- Fully compatible with Go's standard error interface and unwrap mechanism

## Usage Examples

### Creating a new error with metadata

```go
import "github.com/vasilesk/efmt"

func validateUser(user User) error {
    if user.Age < 18 {
        return efmt.New("user is underage", 
            efmt.KV("user_id", user.ID),
            efmt.KV("age", user.Age),
            efmt.KV("required_age", 18),
        )
    }
    return nil
}
```

### Wrapping an existing error

```go
import "github.com/vasilesk/efmt"

func processRequest(req Request) error {
    user, err := fetchUser(req.UserID)
    if err != nil {
        return efmt.Wrap(err, "failed to fetch user", 
            efmt.KV("request_id", req.ID),
            efmt.KV("user_id", req.UserID),
        )
    }
    // Process user...
    return nil
}
```

### Extracting metadata from errors

```go
import (
    "fmt"
    "github.com/vasilesk/efmt"
)

func handleError(err error) {
    // Get all key-value pairs
    allValues := efmt.Values(err)
    fmt.Printf("Error context: %+v\n", allValues)
    
    // Get a specific value
    if requestID, ok := efmt.Value[string](err, "request_id"); ok {
        fmt.Printf("Request ID: %s\n", requestID)
    }
    
    // Get all key-value pairs as a slice
    pairs := efmt.ValuePairs(err)
    for _, kv := range pairs {
        fmt.Printf("%s: %v\n", kv.Key, kv.Value)
    }
}
```

### Working with error chains

```go
import "github.com/vasilesk/efmt"

func processData() error {
    data, err := fetchData()
    if err != nil {
        return efmt.Wrap(err, "data fetch failed", 
            efmt.KV("timestamp", time.Now()),
        )
    }
    
    err = validateData(data)
    if err != nil {
        return efmt.Wrap(err, "data validation failed", 
            efmt.KV("data_size", len(data)),
        )
    }
    
    return nil
}

// Later when handling the error:
func handleProcessingError(err error) {
    // This will extract all metadata from the entire error chain
    allValues := efmt.Values(err)
    fmt.Printf("Complete error context: %+v\n", allValues)
}
```

## Advanced Usage

You can combine `efmt` with other error handling packages as it maintains compatibility with Go's error unwrapping mechanism:

```go
import (
    "errors"
    
    "github.com/vasilesk/efmt"
)

// Check for specific error types while preserving metadata
if errors.Is(err, ErrNotFound) {
    // Handle not found case, but still access metadata
    values := efmt.Values(err)
    // ...
}
```

## License

See the [LICENSE](LICENSE) file for details.