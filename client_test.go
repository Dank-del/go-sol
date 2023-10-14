package gosol_test

import (
	"encoding/json"
	"fmt"
	"testing"

	gosol "github.com/Dank-del/go-sol"
)

func TestGetAccountBalance(t *testing.T) {
	client := gosol.NewClient(gosol.DevnetURL, nil)
	publicKey := "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg"
	response, err := client.GetAccountBalance(publicKey)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	DisplayJson(response)
}

func TestGetAccountInfo(t *testing.T) {
	client := gosol.NewClient(gosol.DevnetURL, nil)
	publicKey := "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg"
	accountInfo, err := client.GetAccountInfo(publicKey)
	if err != nil {
		t.Error("Error:", err)
		return
	}
	DisplayJson(accountInfo)
}

func TestGetBlockHeight(t *testing.T) {
	client := gosol.NewClient(gosol.DevnetURL, nil)
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
func TestGetBlock(t *testing.T) {
	client := gosol.NewClient(gosol.DevnetURL, nil)
	slotNumber := int64(430)
	response, err := client.GetBlock(slotNumber, nil)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	DisplayJson(response)
}
func TestGetBlockProduction(t *testing.T) {
	client := gosol.NewClient(gosol.DevnetURL, nil)
	response, err := client.GetBlockProduction()
	if err != nil {
		t.Error("Error:", err)
		return
	}

	DisplayJson(response)
}
