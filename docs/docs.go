// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth_user": {
            "post": {
                "description": "AuthUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "AuthUser",
                "parameters": [
                    {
                        "description": "AuthUser",
                        "name": "AuthUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/check_results": {
            "post": {
                "description": "CheckResults",
                "produces": [
                    "application/json"
                ],
                "summary": "CheckResults",
                "parameters": [
                    {
                        "description": "CheckResultsRequest",
                        "name": "CheckResultsRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CheckResultsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CheckResultsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/create_user": {
            "post": {
                "description": "CreateUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "CreateUser",
                        "name": "CreateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_tasks_from_test": {
            "post": {
                "description": "GetTasksFromTest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetTasksFromTest",
                "parameters": [
                    {
                        "description": "GetTasksFromTest",
                        "name": "GetTasksFromTest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GetTasksFromTestsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetTasksFromTestsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_tests": {
            "get": {
                "description": "GetTests",
                "produces": [
                    "application/json"
                ],
                "summary": "GetTests",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetTestsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/insert_task": {
            "post": {
                "description": "InsertTask",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "InsertTask",
                "parameters": [
                    {
                        "description": "InsertTest",
                        "name": "InsertTest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InsertTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InsertTaskResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/insert_test": {
            "post": {
                "description": "InsertTest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "InsertTest",
                "parameters": [
                    {
                        "description": "InsertTest",
                        "name": "InsertTest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InsertTestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InsertTestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/send_answers": {
            "post": {
                "description": "SendAnswers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SendAnswers",
                "parameters": [
                    {
                        "description": "SendAnswers",
                        "name": "SendAnswers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SendAnswersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SendAnswersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Result": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "string"
                },
                "grade": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "maxGrade": {
                    "type": "integer"
                },
                "start": {
                    "type": "string"
                },
                "studentID": {
                    "type": "integer"
                },
                "testID": {
                    "type": "integer"
                }
            }
        },
        "dto.Task": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "maxGrade": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "testID": {
                    "type": "integer"
                }
            }
        },
        "dto.Test": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                }
            }
        },
        "models.Answer": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "taskID": {
                    "type": "integer"
                }
            }
        },
        "models.AuthUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.AuthUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.BadRequestResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                }
            }
        },
        "models.CheckResultsRequest": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.CheckResultsResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Result"
                    }
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.GetTasksFromTestsRequest": {
            "type": "object",
            "properties": {
                "test_id": {
                    "type": "integer"
                }
            }
        },
        "models.GetTasksFromTestsResponse": {
            "type": "object",
            "properties": {
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Task"
                    }
                }
            }
        },
        "models.GetTestsResponse": {
            "type": "object",
            "properties": {
                "tests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Test"
                    }
                }
            }
        },
        "models.InsertTaskRequest": {
            "type": "object",
            "properties": {
                "task": {
                    "$ref": "#/definitions/dto.Task"
                }
            }
        },
        "models.InsertTaskResponse": {
            "type": "object",
            "properties": {
                "task_id": {
                    "type": "integer"
                }
            }
        },
        "models.InsertTestRequest": {
            "type": "object",
            "properties": {
                "test": {
                    "$ref": "#/definitions/dto.Test"
                }
            }
        },
        "models.InsertTestResponse": {
            "type": "object",
            "properties": {
                "test_id": {
                    "type": "integer"
                }
            }
        },
        "models.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                }
            }
        },
        "models.SendAnswersRequest": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "testID": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.SendAnswersResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.Result"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
