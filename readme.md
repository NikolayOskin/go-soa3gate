## A3 Light Api (SOA3 Gate)

Golang имплементация протокола SOA3GATE (LIGHT API) платежного сервиса А3.


### How to use

Для получения данных необходим авторизационный ключ и пара сертификатов.

В данном примере опущены проверки на ошибки.

```go
package main

import (
    a3 "github.com/nikolayoskin/go-soa3gate"
)

func main() {
    authKey, isProd := "your_key", true
    certPath, pemPath := "./A3GateTrust.crt", "./A3Prod.pem"

    config := a3.NewConfig(certPath, pemPath, authKey, isProd)
    a3api, _ = a3.NewA3(config)
    
    // получаем штраф по УИН
    penalty, ok, _ := a3api.FetchByBill("18810063140001779724", true)
    
    // проверяем есть ли результат в ответе
    if ok {
       print(penalty.AmountToPay)
       print(penalty.BillDate)
    }
}
```