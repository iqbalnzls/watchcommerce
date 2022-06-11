# watchcommerce

watchcommerce is a simple crud functionality for "mini" e-commerce that focus on watch like its name

## Features
There are some functionality that can be done, like 

- **Create Brand** `/api/v1/brand/save`
    
    API to create a watch brands
    
    ```sh
    curl --location --request POST 'localhost:8000/api/v1/brand/save' \
    -header 'Content-Type: application/json' \
    --data-raw '{
                  "name" : "rolex"
    }'
    ```



- **Create Product** `/api/v1/product/save`

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

- **Get Product By Its Brand** `/api/v1/product/brand/get`

    API to get a product by its brand `id`

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/product/brand/get?id=1'
    ```
  
- **Get Product By ID** `/api/v1/product/get`
    
    API to get a product by its `id`

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/product/get?id=1'
    ```
  
- **Create Order** `/api/v1/order/save`

    API to make a transaction/order

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
  
- **Get Order** (`/api/v1/order/get`)

    API to get a transaction/order that has been done

    ```sh
    curl --location --request GET 'localhost:8000/api/v1/order/get?id=1'
    ```
  

## How to run ?


- Make sure u have in the project directory, then just type the following command in ur terminal

    ```sh
    make run
    ```
  
    Then, u have to execute the `watchcommerce.sql` in the migrations directory in the postgres container. To do this, first, type the following command

    ```sh
    docker exec -it watchcommerce_db psql -U commerce -W watchcommerce
    ```
  
    After that, u can execute the sql




- U can stop the container by

    ```sh
    make stop
    ```



- To run unit test by 

    ```sh
    make test
    ```