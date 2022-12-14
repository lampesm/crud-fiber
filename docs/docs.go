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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/account/create": {
            "post": {
                "description": "create a account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "create a account",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "User",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializers.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/account/delete/{id}": {
            "delete": {
                "description": "delete a account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "delete a account",
                "operationId": "delete-account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/account/read/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Show a account",
                "operationId": "Show-account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/account/update/{id}": {
            "put": {
                "description": "update a account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "update a account",
                "operationId": "update-account",
                "parameters": [
                    {
                        "description": "User",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializers.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "create a JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authenticate"
                ],
                "summary": "create a JWT",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializers.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "serializers.Login": {
            "type": "object",
            "required": [
                "pass",
                "user"
            ],
            "properties": {
                "pass": {
                    "type": "string",
                    "minLength": 3
                },
                "user": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                }
            }
        },
        "serializers.User": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 4
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3005",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRUD Fiber",
	Description:      "This is a  swagger for CURD Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
