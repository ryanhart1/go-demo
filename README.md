# Go Demo Project

## Description

This project consists of:
- A product service with REST endpoints for CRUD operations on products
- A product PostgreSQL DB
- An order service with a GraphQL schema and Mutation to create an order (only if the order's product is available in the product DB)
- An oder PostgreSQL DB

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation
- The project can be ran using the command `docker compose up --build`

## Usage

Product Service

- Create a Product.  
GET request to http://localhost:3000/createproduct with body of 
```
{
    "price": 199, 
    "code": "F22",
    "description" : "Computer"
}
```

- Fetch a product by id.
GET request to http://localhost:3000/products/{id} 

- Fetch all products.
GET request to http://localhost:3000/products

- Delete a product by id
DELETE request to http://localhost:3000/deleteproduct/{id}

Order Service

- Create an order.
```
mutation createOrder {
  createOrder(input: {id:1, customer_name: "Ryan Hart", orderItemId: "5" }) {
    customer_name
    orderItemId
  }
}
```

-Query all orders.
```
query orders {
  orders{
    id
    customer_name
    orderItemId
  }
}
```

-Query an order by ID.
```
query singleOrders {
  order(id: 1){
    id
    customer_name
    orderItemId
  }
}
```

## Demonstration

A demostration video can be viewed [here][link]

[link]:https://my.shell.com/:v:/r/personal/ryan_hart_shell_com/Documents/Recordings/Meeting-20241003_092201-Meeting%20Recording.mp4?csf=1&web=1&e=ey8WaA&nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJTdHJlYW1XZWJBcHAiLCJyZWZlcnJhbFZpZXciOiJTaGFyZURpYWxvZy1MaW5rIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXcifX0%3D
