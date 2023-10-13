package gosol

import "net/http"

// Client represents a JSON-RPC client with a base URL.
type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

type ClientOptions struct {
	HttpClient *http.Client
}

// RPCRequest represents a JSON-RPC request object.
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type RPCResponse[T any | Result[T] | int64] struct {
	ID      int64  `json:"id,omitempty"`
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Result  T      `json:"result,omitempty"`
}

type Result[T any] struct {
	Context *Context `json:"context,omitempty"`
	Value   *T       `json:"value,omitempty"`
}

type Context struct {
	APIVersion string `json:"apiVersion,omitempty"`
	Slot       int64  `json:"slot,omitempty"`
}

type AccountInfo struct {
	Data       []string `json:"data,omitempty"`
	Executable bool     `json:"executable,omitempty"`
	Lamports   int64    `json:"lamports,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	RentEpoch  int64    `json:"rentEpoch,omitempty"`
	Space      int64    `json:"space,omitempty"`
}

type GetBlockProps struct {
	Encoding                       string `json:"encoding"`
	MaxSupportedTransactionVersion int64  `json:"maxSupportedTransactionVersion"`
	TransactionDetails             string `json:"transactionDetails"`
	Rewards                        bool   `json:"rewards"`
}

type BlockResult struct {
	BlockHeight       int64                `json:"blockHeight"`
	BlockTime         interface{}          `json:"blockTime"`
	Blockhash         string               `json:"blockhash"`
	ParentSlot        int64                `json:"parentSlot"`
	PreviousBlockhash string               `json:"previousBlockhash"`
	Transactions      []TransactionElement `json:"transactions"`
}

type TransactionElement struct {
	Meta        Meta                   `json:"meta"`
	Transaction TransactionTransaction `json:"transaction"`
}

type Meta struct {
	Err               interface{}   `json:"err"`
	Fee               int64         `json:"fee"`
	InnerInstructions []interface{} `json:"innerInstructions"`
	LogMessages       []interface{} `json:"logMessages"`
	PostBalances      []int64       `json:"postBalances"`
	PostTokenBalances []interface{} `json:"postTokenBalances"`
	PreBalances       []int64       `json:"preBalances"`
	PreTokenBalances  []interface{} `json:"preTokenBalances"`
	Rewards           interface{}   `json:"rewards"`
	Status            Status        `json:"status"`
}

type Status struct {
	Ok interface{} `json:"Ok"`
}

type TransactionTransaction struct {
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}

type Message struct {
	AccountKeys     []string      `json:"accountKeys"`
	Header          Header        `json:"header"`
	Instructions    []Instruction `json:"instructions"`
	RecentBlockhash string        `json:"recentBlockhash"`
}

type Header struct {
	NumReadonlySignedAccounts   int64 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int64 `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int64 `json:"numRequiredSignatures"`
}

type Instruction struct {
	Accounts       []int64 `json:"accounts"`
	Data           string  `json:"data"`
	ProgramIDIndex int64   `json:"programIdIndex"`
}
