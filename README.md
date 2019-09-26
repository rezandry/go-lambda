# go-lambda
Skeleton for golang in AWS lambda

To debug in local:
1. run lambda in specific port with this syntax 
```_LAMBDA_SERVER_PORT=8001 go run main.go```
2. invoke the main.go via debug.go with this syntax
``` go run debug.go 8001 ```