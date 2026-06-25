package swagger

const SwaggerJSON = `{
  "swagger": "2.0",
  "info": {
    "title": "proto/sso/sso.proto",
    "version": "version not set"
  },
  "tags": [
    {
      "name": "Auth"
    },
    {
      "name": "Marketing"
    }
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/auth/login": {
      "post": {
        "operationId": "Auth_Login",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/authLoginResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/authLoginRequest"
            }
          }
        ],
        "tags": [
          "Auth"
        ]
      }
    },
    "/v1/auth/register": {
      "post": {
        "operationId": "Auth_Register",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/authRegisterResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/authRegisterRequest"
            }
          }
        ],
        "tags": [
          "Auth"
        ]
      }
    },
    "/v1/marketing/documents": {
      "get": {
        "operationId": "Marketing_ListDocuments",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingListDocumentsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "limit",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "offset",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "filterLabelSlug",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "filterType",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Marketing"
        ]
      },
      "post": {
        "operationId": "Marketing_CreateDocument",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingCreateDocumentResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/marketingCreateDocumentRequest"
            }
          }
        ],
        "tags": [
          "Marketing"
        ]
      }
    },
    "/v1/marketing/documents/{id}": {
      "get": {
        "operationId": "Marketing_GetDocument",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingGetDocumentResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "slug",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "Marketing"
        ]
      },
      "delete": {
        "operationId": "Marketing_DeleteDocument",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingDeleteDocumentResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Marketing"
        ]
      },
      "put": {
        "operationId": "Marketing_UpdateDocument",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingUpdateDocumentResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/MarketingUpdateDocumentBody"
            }
          }
        ],
        "tags": [
          "Marketing"
        ]
      }
    },
    "/v1/marketing/labels": {
      "get": {
        "operationId": "Marketing_ListLabels",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingListLabelsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "tags": [
          "Marketing"
        ]
      },
      "post": {
        "summary": "Labels",
        "operationId": "Marketing_CreateLabel",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingCreateLabelResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/marketingCreateLabelRequest"
            }
          }
        ],
        "tags": [
          "Marketing"
        ]
      }
    },
    "/v1/marketing/labels/{id}": {
      "delete": {
        "operationId": "Marketing_DeleteLabel",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/marketingDeleteLabelResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Marketing"
        ]
      }
    }
  },
  "definitions": {
    "MarketingUpdateDocumentBody": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        },
        "contentJson": {
          "type": "string"
        },
        "labelIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "authLoginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "appId": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "authLoginResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "authRegisterRequest": {
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
    "authRegisterResponse": {
      "type": "object",
      "properties": {
        "userId": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "marketingCreateDocumentRequest": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "contentJson": {
          "type": "string"
        },
        "labelIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "marketingCreateDocumentResponse": {
      "type": "object",
      "properties": {
        "document": {
          "$ref": "#/definitions/marketingDocument"
        }
      }
    },
    "marketingCreateLabelRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "color": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      }
    },
    "marketingCreateLabelResponse": {
      "type": "object",
      "properties": {
        "label": {
          "$ref": "#/definitions/marketingLabel"
        }
      }
    },
    "marketingDeleteDocumentResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        }
      }
    },
    "marketingDeleteLabelResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        }
      }
    },
    "marketingDocument": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "title": "e.g. \"text\", \"draw_board\""
        },
        "contentJson": {
          "type": "string",
          "title": "JSON string representing document content"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/marketingLabel"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "marketingGetDocumentResponse": {
      "type": "object",
      "properties": {
        "document": {
          "$ref": "#/definitions/marketingDocument"
        }
      }
    },
    "marketingLabel": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "color": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      },
      "title": "Models"
    },
    "marketingListDocumentsResponse": {
      "type": "object",
      "properties": {
        "documents": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/marketingDocument"
          }
        },
        "totalCount": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "marketingListLabelsResponse": {
      "type": "object",
      "properties": {
        "labels": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/marketingLabel"
          }
        }
      }
    },
    "marketingUpdateDocumentResponse": {
      "type": "object",
      "properties": {
        "document": {
          "$ref": "#/definitions/marketingDocument"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        }
      },
      "additionalProperties": {}
    },
    "rpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    }
  }
}`
