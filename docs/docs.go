// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/add": {
            "get": {
                "description": "Take two numbers and add them together",
                "tags": [
                    "add"
                ],
                "summary": "Adding numbers",
                "parameters": [
                    {
                        "type": "number",
                        "description": "first number",
                        "name": "num1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "second number",
                        "name": "num2",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/division": {
            "get": {
                "description": "Divide two numbers, return the result",
                "tags": [
                    "division"
                ],
                "summary": "Divide numbers",
                "parameters": [
                    {
                        "type": "number",
                        "description": "first number",
                        "name": "num1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "second number",
                        "name": "num2",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/random": {
            "get": {
                "description": "Get random numbers",
                "tags": [
                    "random"
                ],
                "summary": "Get random numbers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Amount of random numbers to return",
                        "name": "Count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/substract": {
            "get": {
                "description": "Take two real numbers and return the result of subtracting them",
                "tags": [
                    "substract"
                ],
                "summary": "Substracting numbers",
                "parameters": [
                    {
                        "type": "number",
                        "description": "first number",
                        "name": "num1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "second number",
                        "name": "num2",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"result\": 4 }",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "User API documentation",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
