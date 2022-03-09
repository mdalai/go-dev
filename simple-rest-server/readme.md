
Build a simple GO REST API with gorilla/mux following this tutorial https://carlosmv.hashnode.dev/building-a-simple-rest-api-in-go-with-gorillamux


```sh

mkdir simple-rest-server && cd simple-rest-server

# my module name is example/user/gorestapi
go mod init example/user/gorestapi

go get -u github.com/gorilla/mux


# run 
go run main.py

# if all files are in the project root, has to run with
go run .

```