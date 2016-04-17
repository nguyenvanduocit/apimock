# API Mockup

Just another API server for mock your api up and run in 0 second

# Usage

Get server

```
go get github.com/nguyenvanduocit/apimock
```

Start server

```
apimock -ip=127.0.0.1 -port=8080
```

Then make request in any method, but remember to include 2 key:

| key                    | description                                  |
|------------------------|----------------------------------------------|
| expected_response      | Any text you want to recive in response body |
| expected_response_mime | Response mime type                           |
