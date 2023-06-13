## Rundoo


## Pre-requisites

- Install docker

## How to run

1. Clone the repository
2. Run `docker-compose up --build` in the root directory

The above will start the http server on port 8080. You can access the server at http://localhost:8080

## APIs

### POST /api/v1/products

### Request Body

```json
{
    "category": "category",
    "name": "name",
    "sku": "sku123"
}
```

### Response

```json
{
  "id": "4aa2017b-a2fa-4cc5-8e8e-c55930b4169f",
  "category": "category",
  "name": "name",
  "sku": "sku123"
}
```

### GET /api/v1/products/search

### Request Query Params

- `query`: string: String to search for
- `limit`: int: Number of results to return

### Response 

```json
[
  {
    "id": "4aa2017b-a2fa-4cc5-8e8e-c55930b4169f",
    "category": "category",
    "name": "name",
    "sku": "sku123"
  }
]
```

