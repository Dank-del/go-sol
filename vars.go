package gosol

import "net/http"

var (
	// DefaultHTTPClient is the default HTTP client for the Solana API.
	DefaultHTTPClient = &http.Client{}
	// DefaultClientOptions is the default options for the Solana API client.
	DefaultClientOptions = &ClientOptions{
		HttpClient: DefaultHTTPClient,
	}
)
