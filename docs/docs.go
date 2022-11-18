// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/transaction/confirm": {
            "post": {
                "description": "Подтвердить перевод и разморозить средства",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Confirm transaction",
                "operationId": "confirm",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/transaction/reserve": {
            "post": {
                "description": "Заморозить средства пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Reserve balance",
                "operationId": "reserve",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/transaction/send": {
            "post": {
                "description": "Зачислить средства на счет пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Send balance",
                "operationId": "send",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/create-user": {
            "post": {
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create User",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/users": {
            "get": {
                "description": "Get All Users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get All Users",
                "operationId": "get-all-user",
                "responses": {}
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete User",
                "operationId": "Delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/{id}/balance": {
            "get": {
                "description": "Get User Balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_balance"
                ],
                "summary": "Get User Balance",
                "operationId": "get-user-balance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Transaction": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "id_order": {
                    "type": "integer"
                },
                "id_service": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Avito Tech Service",
	Description:      "Микросервис для работы с балансом пользователей",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}