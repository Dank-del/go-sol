package gosol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mitchellh/mapstructure"
)

// NewClient creates a new JSON-RPC client with the specified base URL.
func NewClient(baseURL string, opts *ClientOptions) *Client {
	if opts == nil {
		opts = DefaultClientOptions
	}
	return &Client{BaseURL: baseURL, HttpClient: opts.HttpClient}
}

// makeRequest is a private method to send a JSON-RPC request.
func (c *Client) makeRequest(method string, params interface{}) (map[string]interface{}, error) {
	requestPayload := RPCRequest{
		JSONRPC: JSONRPCVersion,
		ID:      1,
		Method:  method,
		Params:  params,
	}

	requestPayloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Post(c.BaseURL, "application/json", bytes.NewBuffer(requestPayloadBytes))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccountBalance retrieves the account balance for a given public key.
func (c *Client) GetAccountBalance(publicKey string) (*RPCResponse[Result[int64]], error) {
	params := []interface{}{
		publicKey,
		map[string]string{
			"encoding": "base58",
		},
	}

	response, err := c.makeRequest("getBalance", params)
	if err != nil {
		return nil, err
	}

	var rpcResponse RPCResponse[Result[int64]]
	err = mapstructure.Decode(response, &rpcResponse)
	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}

func (c *Client) GetAccountInfo(publicKey string) (*RPCResponse[Result[AccountInfo]], error) {
	params := []interface{}{
		publicKey,
		map[string]string{
			"encoding": "base58",
		},
	}

	response, err := c.makeRequest("getAccountInfo", params)
	if err != nil {
		return nil, err
	}

	var rpcResponse RPCResponse[Result[AccountInfo]]
	err = mapstructure.Decode(response, &rpcResponse)
	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}

func (c *Client) GetBlockHeight() (*RPCResponse[int64], error) {
	response, err := c.makeRequest("getBlockHeight", nil)
	if err != nil {
		return nil, err
	}

	var rpcResponse RPCResponse[int64]
	err = mapstructure.Decode(response, &rpcResponse)
	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}

func (c *Client) GetBlock(slotNumber int64, getBlockProps *GetBlockProps) (*RPCResponse[Result[BlockResult]], error) {
	if getBlockProps == nil {
		getBlockProps = &GetBlockProps{
			Encoding:                       "json",
			MaxSupportedTransactionVersion: 0,
			TransactionDetails:             "full",
			Rewards:                        false,
		}
	}
	params := []interface{}{
		slotNumber,
		getBlockProps,
	}

	response, err := c.makeRequest("getBlock", params)
	if err != nil {
		return nil, err
	}

	var rpcResponse RPCResponse[Result[BlockResult]]
	err = mapstructure.Decode(response, &rpcResponse)
	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}

// GetBlockProduction retrieves the block production information from the Solana RPC endpoint.
func (c *Client) GetBlockProduction() (*RPCResponse[Result[GetBlockProdResponse]], error) {
	response, err := c.makeRequest("getBlockProduction", nil)
	if err != nil {
		return nil, err
	}

	var rpcResponse RPCResponse[Result[GetBlockProdResponse]]
	err = mapstructure.Decode(response, &rpcResponse)
	if err != nil {
		return nil, err
	}

	return &rpcResponse, nil
}
