# goten

Go SDK for kintone REST API

[![Go Reference](https://pkg.go.dev/badge/github.com/goqoo-on-kintone/goten.svg)](https://pkg.go.dev/github.com/goqoo-on-kintone/goten)
[![Go Report Card](https://goreportcard.com/badge/github.com/goqoo-on-kintone/goten)](https://goreportcard.com/report/github.com/goqoo-on-kintone/goten)

English | [日本語](/README.ja.md)

## Features

- **Type-safe**: Type-safe record operations using Go 1.18+ generics
- **Facade Pattern**: Intuitive API design inspired by the official JS SDK
- **JS SDK Compatible**: Full support for all official JavaScript SDK features
- **Multiple Auth Methods**: API token, password, and Basic authentication
- **context.Context Support**: Timeout and cancellation handling

## Installation

```bash
go get github.com/goqoo-on-kintone/goten
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/goqoo-on-kintone/goten"
    "github.com/goqoo-on-kintone/goten/auth"
    "github.com/goqoo-on-kintone/goten/record"
)

// Define record type (auto-generation with gotenks is recommended)
// https://github.com/goqoo-on-kintone/gotenks
type MyRecord struct {
    ID struct {
        Value string `json:"value"`
    } `json:"$id"`
    Title struct {
        Value string `json:"value"`
    } `json:"title"`
}

func main() {
    // Create client
    client := goten.NewClient(goten.ClientConfig{
        BaseURL: "https://your-domain.cybozu.com",
        Auth:    auth.APITokenAuth{Token: os.Getenv("KINTONE_API_TOKEN")},
    })

    ctx := context.Background()

    // Get records (type-safe with generics)
    result, err := record.GetRecords[MyRecord](ctx, client.Record, record.GetRecordsParams{
        App:   "1",
        Query: "CreatedTime > TODAY()",
    })
    if err != nil {
        panic(err)
    }

    for _, rec := range result.Records {
        fmt.Printf("ID: %s, Title: %s\n", rec.ID.Value, rec.Title.Value)
    }
}
```

## API Reference

### RecordClient

| Method | Description |
|--------|-------------|
| `GetRecord[T]` | Get a single record |
| `GetRecords[T]` | Get multiple records |
| `GetAllRecords[T]` | Get all records (auto-paging) |
| `AddRecord` | Add a record |
| `AddRecords` | Add multiple records |
| `UpdateRecord` | Update a record |
| `UpdateRecords` | Update multiple records |
| `DeleteRecords` | Delete records |
| `UpsertRecord` | Upsert (update if exists, add if not) |
| `CreateCursor` | Create a cursor |
| `GetRecordsByCursor[T]` | Get records by cursor |
| `DeleteCursor` | Delete a cursor |
| `GetRecordComments` | Get comments |
| `AddRecordComment` | Add a comment |
| `DeleteRecordComment` | Delete a comment |
| `UpdateRecordStatus` | Update record status |
| `UpdateRecordsStatus` | Update multiple record statuses |

### AppClient

| Method | Description |
|--------|-------------|
| `GetApp` | Get app info |
| `GetApps` | Get multiple apps info |
| `AddPreviewApp` | Create app (preview) |
| `CopyApp` | Copy app |
| `DeployApp` | Deploy app |
| `GetDeployStatus` | Get deploy status |
| `GetFormFields` | Get form fields |
| `AddFormFields` | Add fields |
| `UpdateFormFields` | Update fields |
| `DeleteFormFields` | Delete fields |
| `GetFormLayout` | Get form layout |
| `UpdateFormLayout` | Update form layout |
| `GetViews` | Get views |
| `UpdateViews` | Update views |
| `GetAppSettings` | Get app settings |
| `UpdateAppSettings` | Update app settings |
| `GetAppCustomize` | Get customization settings |
| `UpdateAppCustomize` | Update customization settings |
| `GetProcessManagement` | Get process management settings |
| `UpdateProcessManagement` | Update process management settings |
| `GetAppAcl` | Get app permissions |
| `UpdateAppAcl` | Update app permissions |
| `GetFieldAcl` | Get field permissions |
| `UpdateFieldAcl` | Update field permissions |
| `GetRecordAcl` | Get record permissions |
| `UpdateRecordAcl` | Update record permissions |

### SpaceClient

| Method | Description |
|--------|-------------|
| `GetSpace` | Get space info |
| `UpdateSpace` | Update space |
| `DeleteSpace` | Delete space |
| `GetSpaceMembers` | Get space members |
| `UpdateSpaceMembers` | Update members |
| `AddThread` | Add thread |
| `UpdateThread` | Update thread |
| `AddThreadComment` | Add thread comment |
| `AddGuests` | Add guest users |
| `AddGuestsToSpace` | Add guests to guest space |
| `UpdateSpaceGuests` | Update guest members |
| `DeleteGuests` | Delete guest users |

### FileClient

| Method | Description |
|--------|-------------|
| `Upload` | Upload file |
| `Download` | Download file |

### BulkRequestClient

| Method | Description |
|--------|-------------|
| `Send` | Execute bulk request (max 20 requests) |

## Authentication

```go
// API Token Authentication
auth.APITokenAuth{Token: "your-api-token"}

// Password Authentication (kintone only)
auth.PasswordAuth{
    Username: "EXAMPLE_USER",
    Password: "CHANGEME",
}

// Basic Authentication (for proxy, etc.)
auth.BasicAuth{
    Username: "EXAMPLE",
    Password: "CHANGEME",
}
```

## Guest Space Support

```go
client := goten.NewClient(goten.ClientConfig{
    BaseURL:      "https://your-domain.cybozu.com",
    Auth:         auth.APITokenAuth{Token: "token"},
    GuestSpaceID: intPtr(123),  // Guest space ID
})
```

## Bulk Request

```go
// Convenient building with Builder
builder := bulk.NewBuilder()
builder.
    AddRecord("1", record1).
    AddRecord("1", record2).
    UpdateRecord("1", "100", updates, "")

ctx := context.Background()
result, err := client.Bulk.Send(ctx, bulk.SendParams{
    Requests: builder.Build(),
})
```

## Development

```bash
# Build
go build ./...

# Test
go test ./...

# Format
go fmt ./...
```

## Documentation

- [Design Document](docs/DESIGN.md) - Architecture and design philosophy
- [API Specification](docs/API.md) - Public interface definitions
- [TODO](TODO.md) - Implementation status and future plans

## License

MIT License

## Related Links

- [kintone REST API Documentation](https://kintone.dev/en/docs/kintone/rest-api/)
- [Official JavaScript SDK](https://github.com/kintone/js-sdk)
- [gotenks](https://github.com/goqoo-on-kintone/gotenks) - CLI tool to auto-generate Go type definitions from kintone apps
