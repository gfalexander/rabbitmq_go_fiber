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
    The expected result is:  ![RabbitMQ login page]()
    
    2. Default login to rabbit is ```guest``` in both inputs (user and password)

    3. Go to channels and check if MessageService is instanced and is running.

    4. send a request to sender
    
    ```
    curl \
    --request GET \
    --url 'http://localhost:3000/send?msg=test'
    ```

    5. Enjoy the service.