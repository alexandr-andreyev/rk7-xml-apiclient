#rk7 xml client

Example:
package main

import (
"encoding/json"
"log"
rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

const (
CashServerIP = "127.0.0.1"
CashServerPort = 14450
HttpUser = "test"
HttpPassword = "test"
)

func main() {
rkeeper := rk7client.NewClient(CashServerIP, CashServerPort, HttpUser, HttpPassword)
result, err := rkeeper.GetRefData("Employees", "true", "0", "1", "items.(IDENT, CODE, NAME, gen\*)")
if err != nil {
log.Println(err)
}
log.Printf("%+v", result)
log.Println()
b, err := json.Marshal(result)
if err != nil {
log.Println(err)
return
}
log.Println(string(b))
}
