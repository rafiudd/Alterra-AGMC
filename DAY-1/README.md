Objectives:
1. Create Postman collection (JSON file)
2. Create Postman environment
3. Implement HTTP Method (GET, POST, UPDATE, DELETE) 

Tasks:
Do request to the following API target by using Postman environment, save the result using “Save Response” (Save as example), then export collection:

1. https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0
2. https://testnet.binance.vision
    - GET /api/v1/klines Get recent 1000 BTCUSDT klines data with 1 minute interval
    - GET /api/v1/klines Get BTCUSDT klines data with 1 day interval, start from 1 September 2022 to 7 September 2022 (UTC)
    - GET /api/v3/account Get information of your account

---
Note My Task:
- Now Postman can ``create variable in collection``, so import environment variable is ``not needed``
- [Postman Documenter](https://documenter.getpostman.com/view/4289441/2s7YYr74Me)
- [Postman Collection JSON](https://raw.githubusercontent.com/rafiudd/Golang-Mini-Course-Alterra/master/DAY-1/Advanced%20Golang%20Mini%20Course%20-%20Alterra.postman_collection.json)

![Screenshoot](https://raw.github.com/rafiudd/Golang-Mini-Course-Alterra/master/DAY-1/collection_variables.png)