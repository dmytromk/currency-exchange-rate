# currency-exchange-rate

## General Information

This is a REST API API application for currency exchange rate (USD to UAH). It allows:
1. 3 API requests:
    1. GET /rate 
    2. POST /subscribe (expects "email: {body}"  payload)
    3. POST /test_email (test endpoint for sending all emails)
2. Sending email with current USD to UAH exchange rate to all subscribers each day (midnight UTC)

## Local Deployment


### Docker startup

Inside of the repository directory run the following commands:
```shell
docker-compose up --build
```

### Docker shutdown

```shell
docker-compose down
```

### Environment variable

You can find all relevant env variables in .env file inside the repository.

## Checking email

In order to check all the relevant functions of the application do the following:
1. Click on the following [link](https://www.wpoven.com/tools/free-smtp-server-for-testing).
2. Add notify@mysite.com <br>
    <img width="400" src="https://github.com/dmytromk/currency-exchange-rate/assets/96624185/6f3a6b3d-1c56-4cf6-a695-e46051e37132">
3. Open Postman and import .json collection from ./postman/postman_collection.json inside the repository
4. Execute POST "subscribe" request. <br>
    <img width="700" src="https://github.com/dmytromk/currency-exchange-rate/assets/96624185/eb0af63e-3d3a-4d05-abac-013300028d8e">
5. Execute POST "test_email" request. <br>
    <img width="700" src="https://github.com/dmytromk/currency-exchange-rate/assets/96624185/f6ae45a2-4c86-4d4f-9ce6-404f08cd89d3">
6. Check the first page once again. <br>
    <img width="600" src="https://github.com/dmytromk/currency-exchange-rate/assets/96624185/04e57082-ea1b-488c-b91a-f83a6980c356">
