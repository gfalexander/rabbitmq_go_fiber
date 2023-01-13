# **Go Fiber RabbitMQ**
   This project was created purposed by being a RabbitMQ messenger example.
## **About**
    Project for RabbitMQ studying using Golang and Fiber

## Quick Start

- Start
    1. Install Docker, Docker compose
    2. Run Makefile
    ```Shell 
    make run
    ```

- Make sure is running
    1. Access the RabbitMQ interface in browser
    ```
    http://localhost:15672/
    ```
    The expected result is:  ![RabbitMQ login page](https://user-images.githubusercontent.com/61751336/212337413-01e2f8e3-ae2d-459f-a68d-f35f04c106d5.PNG)
    
    2. Default login to rabbit is ```guest``` in both inputs (user and password)
    
    3. Go to Channels and check if the 2 services (sender and receiver) are connected
      ![RabbitMQ channel page](https://user-images.githubusercontent.com/61751336/212338077-2ed0218b-911c-4b1a-8bd8-776454b7d1c6.PNG)
    
    4. Go to queues and check if MessageService is instanced and is running.

    5. send a request to sender
    
    ```
    curl \
    --request GET \
    --url 'http://localhost:3000/send?msg=test'
    ```

    5. Enjoy the service.
