// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Giovanny Massuia",
            "url": "giovannymassuia.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/link": {
            "get": {
                "description": "Get all links",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Get all links",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Link"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create short link from long link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Create short link",
                "parameters": [
                    {
                        "description": "Long URL",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CrateShortLinkRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateShortLinkResponse"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/{shortUrl}"
                            }
                        }
                    }
                }
            }
        },
        "/{shortLink}": {
            "get": {
                "description": "Get long link from short link",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Get long link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short link",
                        "name": "shortLink",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Redirect to long link",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Link not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Link": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "longUrl": {
                    "type": "string"
                },
                "shortUrl": {
                    "type": "string"
                }
            }
        },
        "dto.CrateShortLinkRequest": {
            "type": "object",
            "properties": {
                "long_url": {
                    "type": "string"
                }
            }
        },
        "dto.CreateShortLinkResponse": {
            "type": "object",
            "properties": {
                "short_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Mini URL API",
	Description:      "This is a simple API to create short links",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
