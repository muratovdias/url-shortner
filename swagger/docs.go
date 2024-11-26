// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/api/v1/shortener": {
            "get": {
                "description": "Метод возвращает список всех коротких ссылок с их данными (оригинальный URL, алиас и срок действия).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Короткие ссылки"
                ],
                "summary": "Получение списка всех коротких ссылок",
                "responses": {
                    "200": {
                        "description": "Список коротких ссылок",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/src_server_http_v1.UrlsListResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Метод принимает URL и возвращает короткую ссылку.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Короткие ссылки"
                ],
                "summary": "Создание короткой ссылки",
                "parameters": [
                    {
                        "description": "Данные для сокращения URL",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/src_server_http_v1.UrlShortenerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешный ответ с созданной короткой ссылкой",
                        "schema": {
                            "$ref": "#/definitions/src_server_http_v1.UrlShortenerResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос или данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/{link}": {
            "get": {
                "description": "Метод перенаправляет пользователя на оригинальный URL, связанный с указанной короткой ссылкой.",
                "tags": [
                    "Короткие ссылки"
                ],
                "summary": "Перенаправление на оригинальный URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Алиас короткой ссылки",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Успешное перенаправление"
                    },
                    "400": {
                        "description": "Короткая ссылка истекла или недействительна",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Оригинальный URL не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Запрос на удаление короткой ссылки.",
                "tags": [
                    "Удаление"
                ],
                "summary": "Удаления короткой ссылки.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Алиас короткой ссылки",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное удаление"
                    },
                    "400": {
                        "description": "Некорректный алиас",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Короткая ссылка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/{link}/stats": {
            "get": {
                "description": "Метод возвращает статистику по указанной короткой ссылке (количество переходов и последнее время доступа).",
                "tags": [
                    "Короткие ссылки"
                ],
                "summary": "Получение статистики короткой ссылки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Алиас короткой ссылки",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ со статистикой",
                        "schema": {
                            "$ref": "#/definitions/src_server_http_v1.urlStatsResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный алиас",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Короткая ссылка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "src_server_http_v1.UrlShortenerRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "src_server_http_v1.UrlShortenerResponse": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "expire_time": {
                    "type": "string"
                }
            }
        },
        "src_server_http_v1.UrlsListResponse": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "expires": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "src_server_http_v1.urlStatsResponse": {
            "type": "object",
            "properties": {
                "clicks": {
                    "type": "integer"
                },
                "last_access": {
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
	Title:            "Url-Shortener API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
