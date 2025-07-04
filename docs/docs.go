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
        "/api/v1/auth/login": {
            "post": {
                "description": "Аутентифікація користувача та отримання токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Вхід користувача",
                "parameters": [
                    {
                        "description": "Дані для входу",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Користувач та токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Невірний логін або пароль",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/profile": {
            "get": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Отримати інформацію про поточного автентифікованого користувача",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Отримати профіль користувача",
                "responses": {
                    "200": {
                        "description": "Профіль користувача",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Користувача не знайдено",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Створення облікового запису користувача",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Реєстрація нового користувача",
                "parameters": [
                    {
                        "description": "Дані для реєстрації",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Статус успішної реєстрації",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Невірні дані",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/chat/id/{roomID}": {
            "get": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Отримати інформацію про кімнату за її ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Отримати кімнату по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID кімнати",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Кімната",
                        "schema": {
                            "$ref": "#/definitions/models.Room"
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Кімната не знайдена",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/chat/messages": {
            "post": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Відправка повідомлення в канал/room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Відправити повідомлення",
                "parameters": [
                    {
                        "description": "Дані для відправки повідомлення",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MessageRequest"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Помилка валідації або бізнес-логіки",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/chat/rooms": {
            "get": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Повертає список кімнат, в яких присутній авторизований користувач",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Отримати всі кімнати користувача",
                "responses": {
                    "200": {
                        "description": "Список кімнат",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Room"
                            }
                        }
                    },
                    "400": {
                        "description": "Не вдалося отримати кімнати",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Створення кімнати для спілкування",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Створення кімнати",
                "parameters": [
                    {
                        "description": "Дані для створення",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Помилка валідації або бізнес-логіки",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/chat/rooms/{roomID}/messages": {
            "get": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Повертає останні N повідомлень у кімнаті за roomID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Отримати останні повідомлення",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID кімнати",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Кількість повідомлень (за замовчуванням 20)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список повідомлень",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Message"
                            }
                        }
                    },
                    "400": {
                        "description": "Некоректний запит",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/chat/rooms/{roomName}": {
            "get": {
                "security": [
                    {
                        "UserTokenAuth": []
                    }
                ],
                "description": "Отримати інформацію про кімнату за її Name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Отримати кімнату по Name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name кімнати",
                        "name": "roomName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Кімната",
                        "schema": {
                            "$ref": "#/definitions/models.Room"
                        }
                    },
                    "401": {
                        "description": "Неавторизований доступ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Кімната не знайдена",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateRoomRequest": {
            "type": "object",
            "required": [
                "members",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "members": {
                    "description": "dive - перевірити кожен елемент масиву",
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.MessageRequest": {
            "type": "object",
            "required": [
                "content",
                "room_id"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                }
            }
        },
        "models.RegisterRequest": {
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
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "models.Room": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_message": {
                    "$ref": "#/definitions/models.Message"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_online": {
                    "type": "boolean"
                },
                "last_seen": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "UserTokenAuth": {
            "type": "apiKey",
            "name": "X-User-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.001.001",
	Host:             "127.0.0.1",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Shat API",
	Description:      "Boom boom — and into production.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
