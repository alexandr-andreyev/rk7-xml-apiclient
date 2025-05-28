#rk7 xml client

Example:

```
package main

import (
	"log"
	"time"
	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

const (
	CashServerIP   = "127.0.0.1"
	CashServerPort = 14450
	HttpUser       = "test"
	HttpPassword   = "test"
)

func main() {
	rkeeper := rk7client.NewClient(CashServerIP, CashServerPort, HttpUser, HttpPassword, 10 * time.Second)


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

	for _, result := range response.CommandResult[0].Data {
		for _, item := range result.(rk7client.RK7Reference).Items {
			fmt.Println(item.Ident)
			fmt.Printf("%+v", item)
		}
	}
}
```
```
	//example get reflist versions
	result, err := rkeeper.GetRefList()
	if err != nil {
		log.Printf("error: %s", err)
	}
	// Ищем RK7RefList в данных
	for _, data := range result.CommandResult[0].Data {
		if refList, ok := data.(rk7client.RK7RefList); ok {
			// Теперь refList содержит все RK7Reference
			log.Printf("Total references: %s", refList.Count)
			for _, ref := range refList.RK7Reference {
				log.Printf("Ref: %s, Count: %s, Version: %s",
					ref.RefName, ref.Count, ref.DataVersion)
			}
		}
	}
```

TODO:
1. ADD METHODS  
CreateOrder  
SaveOrder  
SetRefData  