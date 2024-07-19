# frete-rapido

A toy app made as a backend challenge at Frete Rapido.


## Requisites
* Docker
* Docker Compose

## Usage
* Clone this repository
```
git clone https://github.com/rdrgmt/frete-rapido.git
```

* Navigate to project directory
```
cd frete-rapido
```

* Run Docker Compose
```
docker compose up --build
```

## Endpoints

### [POST] .../quote

* Request
```bash
curl --location 'http://localhost:8080/quote' \
--header 'Content-Type: application/json' \
--data '{
    "recipient": {
        "address": {
            "zipcode": "29161376"
        }
    },
    "volumes": [
        {
            "category": "7",
            "amount": 1,
            "unitary_weight": 5,
            "price": 349,
            "sku": "abc-teste-123",
            "height": 0.2,
            "width": 0.2,
            "length": 0.2
        },
        {
            "category": "7",
            "amount": 2,
            "unitary_weight": 4,
            "price": 556,
            "sku": "abc-teste-527",
            "height": 0.4,
            "width": 0.6,
            "length": 0.15
        }
    ]
}'
```
* Response

```json
{
    "carrier": [
        {
            "name": "JADLOG",
            "service": "Rodoviário",
            "deadline": 5,
            "price": 39.75
        },
        {
            "name": "PRESSA FR (TESTE)",
            "service": "Rodoviário",
            "deadline": 0,
            "price": 58.95
        },
        {
            "name": "BTU BRASPRESS",
            "service": "Rodoviário",
            "deadline": 4,
            "price": 78.63
        },
        {
            "name": "RAPIDÃO FR (TESTE)",
            "service": "Rodoviário",
            "deadline": 5,
            "price": 176.58
        }
    ]
}
```

### [GET] .../metrics?last_quotes=

* Request
```bash
curl --location 'http://localhost:8080/metrics?last_quotes=6'
```

* Response
```json
{
    "metrics": [
        {
            "results_per_carrier": {
                "BTU BRASPRESS": 6,
                "JADLOG": 6,
                "PRESSA FR (TESTE)": 6,
                "RAPIDÃO FR (TESTE)": 6
            },
            "total_price_per_carrier": {
                "BTU BRASPRESS": 471.78,
                "JADLOG": 238.5,
                "PRESSA FR (TESTE)": 353.7,
                "RAPIDÃO FR (TESTE)": 1059.48
            },
            "avg_price_per_carrier": {
                "BTU BRASPRESS": 78.63,
                "JADLOG": 39.75,
                "PRESSA FR (TESTE)": 58.95,
                "RAPIDÃO FR (TESTE)": 176.58
            },
            "cheapest_freight": {
                "BTU BRASPRESS": 78.63,
                "JADLOG": 39.75,
                "PRESSA FR (TESTE)": 58.95,
                "RAPIDÃO FR (TESTE)": 176.58
            },
            "priciest_freight": {
                "BTU BRASPRESS": 78.63,
                "JADLOG": 39.75,
                "PRESSA FR (TESTE)": 58.95,
                "RAPIDÃO FR (TESTE)": 176.58
            }
        }
    ]
}
```
