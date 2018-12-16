# ava

Crypto-Addresses Validator

### Installation

Install the package using
```go
$ go get github.com/LanfordCai/ava
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/LanfordCai/ava"
```

### Example

***Validate a `Bitcoin` address***

```go

package main

import (
	"fmt"

	"github.com/LanfordCai/ava"
)

func main() {
	a := ava.Address{
		ChainName: "Bitcoin",
		Address:   "3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9",
		IsTestnet: false,
	}

	if a.IsValid() {
		fmt.Printf("%s is a valid %s address!", a.Address, a.ChainName)
	} else {
		fmt.Printf("%s is an invalid %s address!", a.Address, a.ChainName)
	}
}

```

### Notice

For `EOS` addresses(account name), `ava` only valid it's format. A valid result doesn't mean the account has been registered on the blockchain.

### TODO

- [ ] Validate bech32 addresses for Bitcoin/BitcoinCash

### **License**
The **ava** is an open-source software licensed under the [MIT License](LICENSE).
