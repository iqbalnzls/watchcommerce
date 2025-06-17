# ⌚ WatchCommerce

WatchCommerce is a lightweight, containerized e-commerce API built with **Golang** and **PostgreSQL**, designed to manage watch brands, products, and orders. This project features a modular architecture, Swagger documentation, and Docker support for ease of setup and testing.

---

## 📦 Features


### ✅ Brand APIs
- **Create Brand**
    
    API to create a watch brands
    
    ```sh
    curl --location --request POST 'localhost:8000/api/v1/brand/save' \
    -header 'Content-Type: application/json' \
    --data-raw '{
                  "name" : "rolex"
    }'
    ```


### 📦 Product APIs
- **Create Product**

    API to create a watch product with spesific brand

    ```sh
    curl --location --request POST 'localhost:8000/api/v1/product/save' \
    --header 'Content-Type: application/json' \
    --data-raw '{
                  "brandID": 1,
                  "name": "daytona-XL",
                  "price": 29998,
                  "quantity": 100
    }'
    ```

- **Get Product By Its Brand**

    API to get a product by its brand `id`

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/product/brand/get?id=1'
    ```
  
- **Get Product By ID**
    
    API to get a product by its `id`

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/product/get?id=1'
    ```
  
    ### GraphQL APIs
- **Get Product By ID (GraphQL)**

    API to get a product by its `id` using GraphQL
    
    Endpoint: `http://localhost:8001/query`

    ```graphql
    query GetProduct($productId: Int!) {
      product(id: $productId) {
        id
        name
        price
      }
    }
    ```

### 🛒 Order APIs
- **Create Order**

    API to make an order

    ```sh
    curl --location --request POST 'localhost:8000/api/v1/order/save' \
    --header 'Content-Type: application/json' \
    --data-raw '{
          "orderDetails": [
              {
                  "productID": 1,
                  "quantity": 1
              },
              {
                  "productID": 2,
                  "quantity": 3
              }
          ]
    }'
    ```
  
- **Get Order**

    API to get an order that has been done

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/order/get?id=1'
    ```


## 🚀 Getting Started


1. Ensure you have run docker on your local computer, then type the following command in the terminal
    ```sh
    make run
    ```
    This will spin up both the app and PostgreSQL using Docker.  


2. Run Migrations 
   
   Once the containers are up, you need to run the migrations to set up the database schema. You can do this by executing:

   ```sh
   docker exec -it watchcommerce_db psql -U commerce -W watchcommerce
   ```

3. Stop the containers

    After you are done, you can stop the containers by running:

    ```sh
    make stop
    ```


### 🧪 Testing

- To run the tests, you can use the following command:

    ```sh
    make test
    ```
  
### 📖 Documentation
  
- The API documentation is available via Swagger. You can access it by navigating to:
    ```sh
    http://localhost:8000/swagger/index.html#/
    ```