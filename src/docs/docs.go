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
        "/process/logs/{name}/{endOffset}/{limit}": {
            "get": {
                "description": "Retrieves the process logs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get process logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset from the end of the log",
                        "name": "endOffset",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit of lines to get (0 will get all the lines till the end)",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Process Logs",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/process/start/{name}": {
            "post": {
                "description": "Starts the process if the state is not 'running' or 'pending'",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Start a process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Started Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/stop/{name}": {
            "patch": {
                "description": "Sends kill signal to the process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Stop a process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Stopped Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/processes": {
            "get": {
                "description": "Retrieves all the configured processes and their status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get all processes",
                "responses": {
                    "200": {
                        "description": "Processes Status",
                        "schema": {
                            "type": "object"
                        }
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
