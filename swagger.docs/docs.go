
package docs
//This file is generated automatically. Do not try to edit it manually.
func init() {
 ResourceListingJSON = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{basePath}}",
    "apis": [
        {
            "path": "/user",
            "description": "Friends management API"
        }
    ],
    "info": {
        "title": "Swagger Friends Management API",
        "description": "Swagger Friends Management API",
        "contact": "jingzhu.wang@chinasofti.com",
        "termsOfServiceUrl": "http://www.chinasofit.com",
        "license": "None",
        "licenseUrl": "#"
    }
}`
 APIDescriptionsJSON = map[string]string{"user":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{basePath}}",
    "resourcePath": "/user",
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "/user/friends",
            "description": "create a friend connection between two email addresses.",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "AddFriend",
                    "type": "api.output.Output",
                    "items": {},
                    "summary": "create a friend connection between two email addresses.",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "FriendInput",
                            "dataType": "api.input.FriendInput",
                            "type": "api.input.FriendInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/user/friend/block",
            "description": "block updates from an email address",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "BlockInput",
                    "type": "api.output.Output",
                    "items": {},
                    "summary": "block updates from an email address",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "BlockInput",
                            "dataType": "api.input.BlockInput",
                            "type": "api.input.BlockInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/user/friends/common",
            "description": "common friends list between two email addresses.",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "GetCommonFriends",
                    "type": "api.output.FriendListOutput",
                    "items": {},
                    "summary": "common friends list between two email addresses.",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "FriendInput",
                            "dataType": "api.input.FriendInput",
                            "type": "api.input.FriendInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.FriendListOutput"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/user/friend/list",
            "description": "friends list for an email address.",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "FriendList",
                    "type": "api.output.FriendListOutput",
                    "items": {},
                    "summary": "friends list for an email address.",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "EmailInput",
                            "dataType": "api.input.EmailInput",
                            "type": "api.input.EmailInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.FriendListOutput"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/user/friends/sender",
            "description": "all email addresses that can receive updates from an email address.",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "SenderFriends",
                    "type": "api.output.RecipientsOutput",
                    "items": {},
                    "summary": "all email addresses that can receive updates from an email address.",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "SenderInput",
                            "dataType": "api.input.SenderInput",
                            "type": "api.input.SenderInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.RecipientsOutput"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/user/friend/subscribe",
            "description": "subscribing to updates\" is NOT equivalent to \"adding a friend connection.",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "SubscribeFriend",
                    "type": "api.output.Output",
                    "items": {},
                    "summary": "subscribing to updates\" is NOT equivalent to \"adding a friend connection.",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "body",
                            "description": "SubscribeInput",
                            "dataType": "api.input.SubscribeInput",
                            "type": "api.input.SubscribeInput",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "success",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        },
                        {
                            "code": 400,
                            "message": "error",
                            "responseType": "object",
                            "responseModel": "api.output.Output"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        }
    ],
    "models": {
        "api.input.BlockInput": {
            "id": "api.input.BlockInput",
            "properties": {
                "requestor": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "target": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.input.EmailInput": {
            "id": "api.input.EmailInput",
            "properties": {
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.input.FriendInput": {
            "id": "api.input.FriendInput",
            "required": [
                "friends"
            ],
            "properties": {
                "friends": {
                    "type": "array",
                    "description": "friends desc",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                }
            }
        },
        "api.input.SenderInput": {
            "id": "api.input.SenderInput",
            "properties": {
                "sender": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "text": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.input.SubscribeInput": {
            "id": "api.input.SubscribeInput",
            "properties": {
                "requestor": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "target": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.output.FriendListOutput": {
            "id": "api.output.FriendListOutput",
            "properties": {
                "count": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "friends": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                },
                "success": {
                    "type": "bool",
                    "description": "desc",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.output.Output": {
            "id": "api.output.Output",
            "properties": {
                "success": {
                    "type": "bool",
                    "description": "desc",
                    "items": {},
                    "format": ""
                }
            }
        },
        "api.output.RecipientsOutput": {
            "id": "api.output.RecipientsOutput",
            "properties": {
                "recipients": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                },
                "success": {
                    "type": "bool",
                    "description": "desc",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
}
