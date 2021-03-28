## A3 Light Api (SOA3 Gate)

Частичная имплементация протокола SOA3GATE (LIGHT API) платежного сервиса А3 на Golang. 

В текущей версии реализован метод getPenalties (поиск бюджетных начислений в ГИС ГМП).


### Example

Для отправки запросов необходим авторизационный ключ и пара сертификатов.

В примере опущены проверки на ошибки.

```go
package main

import (
    a3gate "github.com/nikolayoskin/go-soa3gate"
)

func main() {
    authKey, isProd := "your_key", true
    certPath, pemPath := "./A3GateTrust.crt", "./A3GateProd.pem"

    config := a3gate.NewConfig(certPath, pemPath, authKey, isProd)
    a3, _ = a3gate.NewA3(config)
    
    // получаем начисление по УИН
    penalty, ok, _ := a3.GetPenaltyByBill("18810063140001779724", true)
    
    // проверка наличия результата в ответе
    if ok {
        println(penalty.AmountToPay)
    }

    // поиск начислений по документам одного типа
    stsNumbers := []string{"7300311726", "3123170887"}
    penalties, _ := a3.GetPenaltiesByType("sts", stsNumbers, true)

    for _, penalty := range penalties {
        println(penalty.BillNumber)
    }   
}
```