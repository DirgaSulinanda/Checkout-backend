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
```

Example body request:
```
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

Output
```
{
    "status": "success",
    "cashier_name": "Dirga",
    "products": [
        {
            "sku": "120P90",
            "name": "Google Home",
            "quantity": 10,
            "price": 49.99,
            "total_price": 449.91,
            "original_price": 499.90000000000003,
            "promos": [
                "Buy 2 get 1 free Google Home"
            ],
            "is_out_of_stock": true
        },
        {
            "sku": "43N23P",
            "name": "MacBook Pro",
            "quantity": 5,
            "price": 5399.99,
            "total_price": 26999.949999999997,
            "original_price": 26999.949999999997,
            "promos": null,
            "is_out_of_stock": true
        },
        {
            "sku": "234234",
            "name": "Raspberry Pi B",
            "quantity": 2,
            "price": 30,
            "total_price": 30,
            "original_price": 60,
            "promos": [
                "Free Rasp Pi"
            ],
            "is_out_of_stock": true
        }
    ],
    "sub_total": 27479.859999999997,
    "original_price": 27559.85
}
```

Gql query
```
query{
  Checkout(cashierName: "Dirga", products: [{sku: "120P90",quantity: 1000},{sku: "43N23P",quantity: 1000},{sku: "234234",quantity: 1000}]) {
    data {
        status
        cashierName
        products{
            sku
            name
            quantity
            price
            totalPrice
            originalPrice
            promos
            isOutOfStock
        }
        subTotal
        originalPrice
    }
  }
}
```
