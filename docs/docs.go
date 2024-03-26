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
        "/dentists": {
            "get": {
                "description": "Get all dentists from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Get all dentists",
                "responses": {
                    "200": {
                        "description": "List of dentists",
                        "schema": {
                            "type": "slice"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing dentist in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Update existing dentist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Updated data of the dentist",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dentist updated successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    },
                    "400": {
                        "description": "Invalid dentist data or missing dentist ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Token not found or invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new dentist in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Create a new dentist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Data of the dentist to create",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Dentist created successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    },
                    "400": {
                        "description": "Invalid dentist data or missing required fields",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Token not found or invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dentists/{id}": {
            "get": {
                "description": "Get a dentist from the system by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Get dentist by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dentist found successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dentist not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a dentist from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Delete dentist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Dentist deleted successfully"
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Token not found or invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update the license of a dentist in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dentists"
                ],
                "summary": "Update dentist license",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dentist updated successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentist"
                        }
                    },
                    "400": {
                        "description": "Invalid ID, invalid JSON or missing license data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Token not found or invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dentist not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Dentist": {
            "type": "object",
            "required": [
                "FirstName",
                "LastName",
                "License"
            ],
            "properties": {
                "FirstName": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "LastName": {
                    "type": "string"
                },
                "License": {
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