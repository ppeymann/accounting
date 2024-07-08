# API Documentation

### <span style="color: gray;">This document is for an accounting application in which you can enter your expenses and have a statistic of your expenses.</span>

### 1 - First, take a clone from the project whose address is in README

### 2 - install Docker

### 3 - run docker compose command (compose in README)

### 4 - run build command (build in README)

### You can see all the endpoints in Swagger and I explain api below

(base Url is: http://localhost:8080/api/v1)

## account service

## sign up:

    method: POST
    `/account/signup`

    this api is for register in application
    get: token and refresh token and expire time that should save it in localstorage

## sign in:

    method: POST
    `/account/signin`

    this api is for login in application
    get: Like sign up

## account:

    method: GET
    `/account`

    this api is for get account information
    get: account information like: bank information, expenses information, ...

## change currency:

    method: PATCH
    `/account/change_currency`

    this api is for change currency type in account information
    *(U should send "IRT", "USD")
    get: account information

## change name:

    method: PATCH
    `/account/change_name`

    this api is for change name in account information
    get: account information

## bank service

## create bank account:

    method: POST
    `/bank/create`

    this api is for create bank account
    *(for fill bank_slug U should get all bank slugs from bank slug api and put a one of them here)
    get: bank account information

## get all banks:

    method: GET
    `/bank/all`

    with this api U can get all bank slugs
    *(bank means a name of a bank and unique slug)
    get: array of banks (slug and name)

## get bank account:

    method: GET
    `/bank/{id}`

    with this api U can get a bank account that user have
    get: a bank account

## update bank account:

    method: PUT
    `/bank/{id}`

    can update bank account
    get: bank account information

## delete bank account

    method: DELETE
    `/bank/{id}`

    can delete bank account
    get: bank account ID

## expenses service

## create expenses:

    method: POST
    `/expenses/create`

    can create a expenses information
    get: expenses information that made

## get all expense:

    method: GET
    `/expenses/{id}`

    can get all expenses for user
    get: array of expenses

## get expenses in specified month and year:

    method: GET
    `/expenses/get_in_month/{year}/{month}`

    U can send month and year and get all expenses that is for this month and year
    get: array of expenses

## get expenses in period time:

    method: GET
    `/expenses/get_period_time`

    You can send a period of time and receive the corresponding expenses
    *(should send parameters in query route like this
        `expenses/get_period_time?fromYear=1401&fromMonth=3&toYear=1408&toMonth=12`
    )
    get: array of expenses

## get expense by ID:

    method: GET
    `/expenses/{id}`

    can get expenses with specified ID
    get: expenses information

## update expenses:

    method: PUT
    `/expenses/{id}`

    can send expenses information and update it
    get: expenses information

## delete expenses

    method: DELETE
    `/expenses/get_all`

    can delete expenses
    get: expenses ID
