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
        "/opening": {
            "get": {
                "description": "Show a job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Show opening",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/opening.ShowOpeningResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Update opening",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening Identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Opening data to Update",
                        "name": "opening",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/opening.UpdateOpeningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/opening.UpdateOpeningResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Create opening",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/opening.CreateOpeningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/opening.CreateOpeningResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a job opening by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Delete opening",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/opening.DeleteOpeningResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/openings": {
            "get": {
                "description": "List all job openings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "List openings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/opening.ListOpeningsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/opening.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/recruiter": {
            "get": {
                "description": "Show a recruiter by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recruiters"
                ],
                "summary": "Show recruiter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recruiter identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ShowRecruiterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a recruiter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recruiters"
                ],
                "summary": "Update recruiter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recruiter Identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Recruiter data to Update",
                        "name": "opening",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recruiter.UpdateRecruiterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recruiter.UpdateRecruiterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new recruiter as a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recruiters"
                ],
                "summary": "Create recruiter",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recruiter.CreateRecruiterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recruiter.CreateRecruiterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a recruiter by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recruiters"
                ],
                "summary": "Delete recruiter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recruiter identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recruiter.DeleteRecruiterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/recruiters": {
            "get": {
                "description": "List all recruiters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recruiters"
                ],
                "summary": "List recruiter",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ListRecruitersResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/recruiter.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "opening.CreateOpeningRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "opening.CreateOpeningResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "opening.DeleteOpeningResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "opening.ErrorResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "opening.ListOpeningsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.OpeningResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "opening.ShowOpeningResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "opening.UpdateOpeningRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "opening.UpdateOpeningResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.CreateRecruiterRequest": {
            "type": "object",
            "properties": {
                "company_email": {
                    "type": "string"
                },
                "company_name": {
                    "type": "string"
                },
                "company_website": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "recruiter": {
                    "type": "boolean"
                }
            }
        },
        "recruiter.CreateRecruiterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.RecruitersResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.DeleteRecruiterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.RecruitersResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.ErrorResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.ListRecruitersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.RecruitersResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.ShowRecruiterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.RecruitersResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "recruiter.UpdateRecruiterRequest": {
            "type": "object",
            "properties": {
                "company_email": {
                    "type": "string"
                },
                "company_name": {
                    "type": "string"
                },
                "company_website": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "recruiter": {
                    "type": "boolean"
                }
            }
        },
        "recruiter.UpdateRecruiterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.RecruitersResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.OpeningResponse": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.RecruitersResponse": {
            "type": "object",
            "properties": {
                "DeletedAt": {
                    "type": "string"
                },
                "UpdatedAt": {
                    "type": "string"
                },
                "companyEmail": {
                    "type": "string"
                },
                "companyName": {
                    "type": "string"
                },
                "companyWebsite": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "type": "integer"
                },
                "recruiter": {
                    "type": "boolean"
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
