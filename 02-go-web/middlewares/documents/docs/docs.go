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
            "name": "API Support",
            "email": "random@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/movies": {
            "get": {
                "description": "Get all movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Get movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseBodyGetMovies"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseBodyGetMovies"
                        }
                    }
                }
            }
        },
        "/movies/{title}": {
            "get": {
                "description": "Get a movie by title",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Get movie by title",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie title",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseBodyGetMovieByTitle"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseBodyGetMovieByTitle"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseBodyGetMovieByTitle"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.MovieJSON": {
            "type": "object",
            "properties": {
                "genre": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "handlers.ResponseBodyGetMovieByTitle": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handlers.MovieJSON"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.ResponseBodyGetMovies": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.MovieJSON"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Movies API",
	Description:      "This is a sample server for Movies API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
