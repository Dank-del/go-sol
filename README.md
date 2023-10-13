# go-sol (WIP)

Simple, Opinionated, Lightweight, GoLang Solana Client

## Usage

```go
import (
    "github.com/Dank-del/go-sol"
    "fmt"
)

func main() {
    url := "https://api.devnet.solana.com"
    client := gosol.NewClient(url)
    publicKey := "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg"
    response, err := client.GetAccountBalance(publicKey)
    if err != nil {
        t.Error("Error:", err)
        return
    }
    fmt.Println(response)
}
```

## License

MIT License - see [LICENSE](LICENSE) for full text
