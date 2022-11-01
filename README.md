# test-task-evo

How to install:

1. Clone this repository: `git clone https://github.com/arimatakao/test-task-evo.git`
2. Download dependency: `go mod download`
3. First run: `make parse`
4. Others run: `make`
5. For build project: `make build`

## Documentation

| route                      | method |
| -------------------------- | ------ |
| /                          | POST   |
| /transactionid/{id:[0-9]+} | GET    |
| /terminalid                | GET    |
| /status                    | GET    |
| /paymenttype               | GET    |
| /date                      | GET    |
| /paymentnarrative          | GET    |

------

**/ POST**
request: none
respone: 

```
Test task for EVO Finctech
```

------

**/transactionid/{id:[0-9]+} GET**

request: none
response:

```json
{
	"TransactionId": "123",
    "RequestId": "123",
    "TerminalId": "123",
    "PartnerObjectId": "123", 
    "AmountTotal": "123",
    "AmountOriginal": "123",
    "CommissionPS": 1.23,
    "CommissionClient": 1.23, 
    "CommissionProvider": 1.23,
    "DateInput": "2022-08-23 11:49:02",
    "DatePost": "2022-08-23 11:49:02",
    "Status": "accepted",
    "PaymentType": "cash", 
    "PaymentNumber": "PS12345678",
    "ServiceId": "12345",
    "Service": "abc",
    "PayeeId": 12345678,      
    "PayeeName": "privat",
    "PayeeBankMfo": 123456,
    "PayeeBankAccount": "US123456",
    "PaymentNarrative": "abc"
}
```

------

**/terminalid GET**

request:

```json
{
	"TerminalId": "1111"
}
```

response:

```json
{
	
    "obj_TerminalId":
    	[
        	{
                ...
    			"TerminalId": "1111",
    			...
            },
            ...
    	]
}
```

------

**/status GET**

request:

```json
{
	"Status": "accepted"
}
```

response:

```json
{
    "obj_Status":
    	[
        	{
                ...
    			"Status": "accepted",
    			...
            },
            ...
    	]
}
```

------

**/paymenttype GET**

request:

```json
{
	"PaymentType": "cash"
}
```

response:

```json
{
    "obj_PaymentType":
    	[
        	{
                ...
				"PaymentType": "cash",
				...
            },
            ...
    	]
}
```

------

**/date GET**

request:

```json
{
	"from_DatePost": "2022-08-23 11:00:00",
	"to_DatePost": "2022-08-23 12:00:00"
}
```

response:

```json
{
	"obj_from_DatePost_to_DatePost":
    	[
        	{
                ...
                "DatePost": "2022-08-23 11:01:23",
                ...
            },
            ...
    	]
}
```

------

**/paymentnarrative GET**

request:

```json
{
	"PaymentNarrative": "271"
}
```

response

```
{

	"obj_PaymentNarrative":
    	[
        	{
                ...
                "PaymentNarrative": "A22/27144"
            },
            ...
    	]
}
```

