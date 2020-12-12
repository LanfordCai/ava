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

```go
package main

import (
	"fmt"

	"github.com/LanfordCai/ava"
)

func main() {
	validator := &ava.Bitcoin{}
	addr := "19JeUHUvw23fwKeK1zZD4moKyxj1xn4Kxi"
	result := validator.ValidateAddress(addr, ava.Mainnet)

	fmt.Printf("Address is valid?: %t\n", result.IsValid)
	fmt.Printf("Address type: %s\n", result.Type)

	// Address is valid?: true
	// Address type: P2PKH
}
```

### **License**
The **ava** is an open-source software licensed under the [MIT License](LICENSE).
