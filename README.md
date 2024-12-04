# 🆗 Remote OK Jobs API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/remote-ok-jobs.svg)](https://pkg.go.dev/github.com/go-api-libs/remote-ok-jobs/pkg/remoteokjobs)
[![Official Documentation](https://img.shields.io/badge/docs-API-blue)](https://www.remoteok.com)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/remote-ok-jobs)](https://goreportcard.com/report/github.com/go-api-libs/remote-ok-jobs)
![Code Coverage](https://img.shields.io/badge/coverage-94%25-brightgreen)
![API Health](https://img.shields.io/badge/API_health-70%25-yellowgreen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Remote OK is the #1 remote-only jobs board on the web. It has over 30,000+ remote job listings, adding new ones every day. With this API, now you can build apps with our data feed which covers 80% of remote jobs on the web. ([Source](https://www.remoteok.com))

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/remote-ok-jobs/pkg/remoteokjobs
```

## Usage

### Example 1: 

```go
package main

import (
	"context"

	"github.com/go-api-libs/remote-ok-jobs/pkg/remoteokjobs"
)

func main() {
	c, err := remoteokjobs.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	getAPIOkJSONResponse, err := c.GetAPI(ctx)
	if err != nil {
		panic(err)
	}

	// Use getAPIOkJSONResponse array
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/remote-ok-jobs/pkg/remoteokjobs): The Go reference documentation for the client package.
- [**Official Documentation**](https://www.remoteok.com): The official API documentation.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/remote-ok-jobs): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/remote-ok-jobs).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
