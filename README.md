# Checkout-backend

**Installation**
```
make start
```

**Stop App**
```
make stop
```

**API**
```
POST localhost:9000/checkout

Example body request:
{
    "cashier_name": "Dirga",
    "products": [
        {
            "sku": "120P90",
            "quantity": 1000
        },
        {
            "sku": "43N23P",
            "quantity": 1000
        },
        {
            "sku": "234234",
            "quantity": 1000
        }
    ]   
}
```
