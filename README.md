# Mock GO

This is very simple golang http mock server based on gorilla/mux.

## Install

```sh
go get -u github.com/rukavina/mockgo
```
## Configuration

It's expected to have `config.json`  in the same folder as executable. Eg:

```json
{
    "description": "Test mock definitions",
    "endpoints":[
        {
            "request":{
                "path": "/hello"
            },
            "response":{
                "bodyFileName": "responses/hello.json"
            } 
        },
        {
            "request":{
                "path": "/hello2"
            },
            "response":{
                "body": "{\"text\": \"Hello\",\"name\": \"World\"}",
                "headers": {
                    "Content-Type":"application/json"
                }
            } 
        }        
    ]
}
```

The `request.path` defines the url at which the mock serves responses. It's gorilla mux'es path definition. Request routing can be further specified via _request_ fields like: `scheme`,`host`,`method`,`query`,`headers`.

Response data can be either defined inline via `body` or as external file in `bodyFileName`.