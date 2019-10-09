# ava

Cryptocurrency Addresses Validator

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

At first, setup enabled address type:

```sh
export AVA_BITCOIN_ENABLED_ADDR_TYPES="P2PKH,P2SH"
```

And then, write some code:

```go

package main

import (
	"fmt"

	"github.com/LanfordCai/ava"
)

func main() {
	validator, err := ava.NewValidator("Bitcoin")
	if err != nil {
		panic(err.Error())
	}

	address := "19JeUHUvw23fwKeK1zZD4moKyxj1xn4Kxi"
	isTestnet := false
	isValid, msg :=validator.ValidateAddress(address, isTestnet)

	fmt.Printf("Address is valid?: %t\n", isValid)
	fmt.Printf("Extra message: %s\n", msg)
}

```

### Notice

For `EOS` addresses(account name), `ava` only valid it's format. A valid result doesn't mean the account has been registered on the blockchain.

### TODO

- [ ] Validate bech32 addresses for Bitcoin/BitcoinCash

### **License**
The **ava** is an open-source software licensed under the [MIT License](LICENSE).
