# UUID Package

This package provides utilities for UUID (Universally Unique Identifier) generation and manipulation using only Go's standard library.

## Features
- Generate new UUIDs (v4)
- Parse UUID from string
- Convert UUID to string
- Validate a UUID string
- Compare two UUIDs
- Marshal/Unmarshal UUID to/from JSON
- Generate nil (zero value) UUID

## Usage Example
```go
import "yourmodule/uuid"

// Generate a new UUID (v4)
id := uuid.NewUUID()

// Convert to string
s := id.String()

// Validate a UUID string
valid := uuid.Validate(s)

// Parse from string
parsed, err := uuid.Parse(s)

// Compare UUIDs
isEqual := id.Equal(parsed)

// Nil UUID
zero := uuid.Nil()

// Marshal/Unmarshal JSON
b, _ := json.Marshal(id)
var u2 uuid.UUID
_ = json.Unmarshal(b, &u2)
```

## Testing
Run unit tests:
```sh
go test ./uuid
```
