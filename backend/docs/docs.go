// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/users": {
            "get": {
                "description": "list user entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List user entities",
                "operationId": "list-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User entity",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get user by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a user entity by ID",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "update user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a user entity by ID",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User entity",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "get user by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a user entity by ID",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Playlist": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PlaylistQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PlaylistEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PlaylistEdges": {
            "type": "object",
            "properties": {
                "owner": {
                    "description": "Owner holds the value of the owner edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                },
                "playlistVideos": {
                    "description": "PlaylistVideos holds the value of the playlist_videos edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.PlaylistVideo"
                    }
                }
            }
        },
        "ent.PlaylistVideo": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PlaylistVideoQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PlaylistVideoEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PlaylistVideoEdges": {
            "type": "object",
            "properties": {
                "playlists": {
                    "description": "Playlists holds the value of the playlists edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Playlist"
                },
                "resolution": {
                    "description": "Resolution holds the value of the resolution edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Resolution"
                },
                "video": {
                    "description": "Video holds the value of the video edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Video"
                }
            }
        },
        "ent.Resolution": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ResolutionQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ResolutionEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "resolution": {
                    "description": "Resolution holds the value of the \"resolution\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ResolutionEdges": {
            "type": "object",
            "properties": {
                "playlistVideos": {
                    "description": "PlaylistVideos holds the value of the playlist_videos edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.PlaylistVideo"
                    }
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "Age holds the value of the \"age\" field.",
                    "type": "integer"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.UserEdges"
                },
                "email": {
                    "description": "Email holds the value of the \"email\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "password": {
                    "description": "Password holds the value of the \"password\" field.",
                    "type": "string"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "playlists": {
                    "description": "Playlists holds the value of the playlists edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Playlist"
                    }
                },
                "videos": {
                    "description": "Videos holds the value of the videos edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Video"
                    }
                }
            }
        },
        "ent.Video": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the VideoQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.VideoEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "video_title": {
                    "description": "VideoTitle holds the value of the \"video_title\" field.",
                    "type": "string"
                }
            }
        },
        "ent.VideoEdges": {
            "type": "object",
            "properties": {
                "owner": {
                    "description": "Owner holds the value of the owner edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                },
                "playlistVideos": {
                    "description": "PlaylistVideos holds the value of the playlist_videos edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.PlaylistVideo"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API",
	Description: "This is a sample server for SUT SE 2563",
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
	swag.Register(swag.Name, &s{})
}