// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account": {
            "get": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "get account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "get account",
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.AccountEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/account/change_currency": {
            "patch": {
                "description": "change currency of account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "change currency",
                "parameters": [
                    {
                        "description": "currency that is for change",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.ChangeCurrencyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.AccountEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/account/change_name": {
            "patch": {
                "description": "change name of account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "change name",
                "parameters": [
                    {
                        "description": "name that is for change",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.NameInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.AccountEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/account/signin": {
            "post": {
                "description": "signing in to account that already signed up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "signing in to account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.LoginInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.TokenBundleOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/account/signup": {
            "post": {
                "description": "create new account with specified mobile and expected info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "signing up a new account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.LoginInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.TokenBundleOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/bank/all": {
            "get": {
                "security": [
                    {
                        "Authenticate bearer": []
                    }
                ],
                "description": "get all bank information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "get all bank",
                "responses": {
                    "200": {
                        "description": "Always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/services.BankEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/bank/create": {
            "post": {
                "security": [
                    {
                        "Authenticate bearer": []
                    }
                ],
                "description": "create bank account with specified input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "create bank account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.BankAccountInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.BankAccountEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/bank/{id}": {
            "get": {
                "security": [
                    {
                        "Authenticate bearer": []
                    }
                ],
                "description": "get bank by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bank"
                ],
                "summary": "get bank",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bank id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.BankAccountEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/expenses/create": {
            "post": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "create expenses with expenses input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "create expenses",
                "parameters": [
                    {
                        "description": "expenses input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.ExpensesInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.ExpensesEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/expenses/get_all": {
            "get": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "get all expenses with specified id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "get all expenses",
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/services.ExpensesEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/expenses/get_in_month/{year}/{month}": {
            "get": {
                "description": "get expenses in the specified month",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "get expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/services.ExpensesEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/expenses/get_period_time": {
            "get": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "get expenses in period time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "get expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "start year",
                        "name": "fromYear",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "start month",
                        "name": "fromMonth",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "end year",
                        "name": "toYear",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "end month",
                        "name": "toMonth",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/services.ExpensesEntity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/expenses/{id}": {
            "get": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "get expenses by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "get expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "expenses id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.ExpensesEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "update expenses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "update expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "expenses id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "expenses input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.ExpensesInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/services.ExpensesEntity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authenticate Bearer": []
                    }
                ],
                "description": "delete expenses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "delete expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "expenses id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "always returns status 200 but body contains errors",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/accounting.BaseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "accounting.BaseResult": {
            "type": "object",
            "properties": {
                "errors": {
                    "description": "Errors provides list off error that occurred in processing request",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "description": "Result single/array of any type (object/number/string/boolean) that returns as response"
                },
                "result_count": {
                    "description": "ResultCount specified number of records that returned in result_count field expected result been array.",
                    "type": "integer"
                }
            }
        },
        "services.AccountEntity": {
            "type": "object",
            "properties": {
                "bank_account": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.BankAccountEntity"
                    }
                },
                "currency_type": {
                    "description": "CurrencyType is currency type between Rial, Dollar, Dinar, ...",
                    "allOf": [
                        {
                            "$ref": "#/definitions/services.CurrencyType"
                        }
                    ]
                },
                "email": {
                    "description": "Email",
                    "type": "string"
                },
                "full_name": {
                    "description": "FullName",
                    "type": "string"
                },
                "mobile": {
                    "description": "Mobile",
                    "type": "string"
                },
                "password": {
                    "description": "Password",
                    "type": "string"
                },
                "user_name": {
                    "description": "UserName",
                    "type": "string"
                }
            }
        },
        "services.BankAccountEntity": {
            "type": "object",
            "properties": {
                "account_id": {
                    "description": "AccountID",
                    "type": "integer"
                },
                "bank_number": {
                    "description": "BankNumber",
                    "type": "integer"
                },
                "bank_slug": {
                    "description": "BankSlug",
                    "type": "string"
                },
                "name": {
                    "description": "Name",
                    "type": "string"
                }
            }
        },
        "services.BankAccountInput": {
            "type": "object",
            "properties": {
                "bank_number": {
                    "description": "BankNumber",
                    "type": "integer"
                },
                "bank_slug": {
                    "description": "BankSlug",
                    "type": "string"
                },
                "name": {
                    "description": "Name",
                    "type": "string"
                }
            }
        },
        "services.BankEntity": {
            "type": "object",
            "properties": {
                "bank_slug": {
                    "description": "BankSlug",
                    "type": "string"
                },
                "name": {
                    "description": "Name",
                    "type": "string"
                }
            }
        },
        "services.ChangeCurrencyInput": {
            "type": "object",
            "properties": {
                "currency_type": {
                    "description": "CurrencyType",
                    "allOf": [
                        {
                            "$ref": "#/definitions/services.CurrencyType"
                        }
                    ]
                }
            }
        },
        "services.CurrencyType": {
            "type": "string",
            "enum": [
                "IRT",
                "USD",
                "EUR",
                "AED",
                "TRY"
            ],
            "x-enum-varnames": [
                "Rial",
                "Dollar",
                "Euro",
                "Dirham",
                "Lier"
            ]
        },
        "services.DateAndTime": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "hour": {
                    "type": "integer"
                },
                "minute": {
                    "type": "integer"
                },
                "month": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "services.ExpensesEntity": {
            "type": "object",
            "properties": {
                "account_id": {
                    "description": "AccountID",
                    "type": "integer"
                },
                "amount": {
                    "description": "Amount",
                    "type": "number"
                },
                "bank_id": {
                    "description": "BankID",
                    "type": "integer"
                },
                "bank_name": {
                    "description": "BankName",
                    "type": "string"
                },
                "bank_number": {
                    "description": "BankNumber",
                    "type": "integer"
                },
                "category": {
                    "description": "Category",
                    "type": "string"
                },
                "day": {
                    "description": "Day is day that this expenses is done",
                    "type": "integer"
                },
                "hour": {
                    "description": "Hour is hour that this expenses is done",
                    "type": "integer"
                },
                "minute": {
                    "description": "Minute is minute that this expenses is done",
                    "type": "integer"
                },
                "month": {
                    "description": "Month is month that this expenses is done",
                    "type": "integer"
                },
                "notes": {
                    "description": "Notes",
                    "type": "string"
                },
                "year": {
                    "description": "Year is year that this expenses is done",
                    "type": "integer"
                }
            }
        },
        "services.ExpensesInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "bank_id": {
                    "type": "integer"
                },
                "bank_name": {
                    "type": "string"
                },
                "bank_number": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "date": {
                    "$ref": "#/definitions/services.DateAndTime"
                },
                "note": {
                    "type": "string"
                }
            }
        },
        "services.LoginInputDTO": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "Password is password for sign up",
                    "type": "string"
                },
                "user_name": {
                    "description": "Username is user name for sign up",
                    "type": "string"
                }
            }
        },
        "services.NameInput": {
            "type": "object",
            "properties": {
                "full_name": {
                    "description": "FullName",
                    "type": "string"
                }
            }
        },
        "services.TokenBundleOutput": {
            "type": "object",
            "properties": {
                "expire": {
                    "description": "Expire time of Token and CentrifugeToken",
                    "type": "string"
                },
                "refresh": {
                    "description": "Refresh token string used for refreshing authentication and give fresh token",
                    "type": "string"
                },
                "token": {
                    "description": "Token is JWT/PASETO token staring for storing in client side as access token",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
