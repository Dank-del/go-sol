package gosol_test

import (
	"encoding/json"
	"fmt"
	"testing"

	gosol "github.com/Dank-del/go-sol"
)

func TestClient(t *testing.T) {
	url := "https://api.devnet.solana.com"
	client := gosol.NewClient(url)
	publicKey := "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg"
	response, err := client.GetAccountBalance(publicKey)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	DisplayJson(response)

	accountInfo, err := client.GetAccountInfo(publicKey)
	if err != nil {
		t.Error("Error:", err)
		return
	}
	DisplayJson(accountInfo)

	bh, err := client.GetBlockHeight()
	if err != nil {
		t.Error("Error:", err)
		return
	}
	DisplayJson(bh)
}

func DisplayJson(r any) {
	ok, err := json.MarshalIndent(r, "  ", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
		return
	}

	fmt.Println("Response:", string(ok))
}
