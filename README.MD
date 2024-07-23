#rk7 xml client

Example:

```
package main

import (
	"log"
	"time"
	rk7client "xmlclient-test/rk7client"
)

const (
	CashServerIP   = "127.0.0.1"
	CashServerPort = 14450
	HttpUser       = "test"
	HttpPassword   = "test"
)

func main() {
	rkeeper := rk7client.NewClient(CashServerIP, CashServerPort, HttpUser, HttpPassword)


	input := []rk7client.RK7Command{
		{
			CMD:         rk7client.RK7CMD_GETREFDATA,
			RefName:     rk7client.RK7REF_EMPLOYEES,
			OnlyActrive: rk7client.ONLY_ACTIVE_TRUE,
			PropMask: "items.(Code,Name,Ident)",
		},
	}

	result, err := rkeeper.GetRefData(input)
	if err != nil {
		log.Println(err)
	}
	log.Printf("xml parse struct result: %+v", result)
}
```
