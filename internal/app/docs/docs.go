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
        "contact": {
            "name": "GitHub",
            "url": "https://github.com/GlebSolncev/golang-modules"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/status": {
            "get": {
                "description": "All statuses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "status"
                ],
                "summary": "All statuses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/todo": {
            "get": {
                "description": "Get todos list with status items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "All todos items with statuses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Add new item to list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Add item to todos list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/todo/{ent.Todo.id}": {
            "get": {
                "description": "Show item of Todos list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Item from todos list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/todo/{id}": {
            "post": {
                "description": "Add new item to list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Update item in list TODOs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/todo/{id}/delete": {
            "get": {
                "description": "Remove item from todos list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Delete item from TODOs list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Response": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "string"
                },
                "payload": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Todo CRUD",
	Description:      "This is a simple CRUD for TODO list",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
