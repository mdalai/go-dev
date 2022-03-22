
Converting simple-rest-server project into structure that aligns with openapi codegen. 


```sh
# get medicines
curl -X GET 'http://localhost:8118/api/v1/medicines'

# get medicine by name
curl -X GET 'http://localhost:8118/api/v1/medicines/Banz'

# post medicine and return all medicines
curl -X POST -H 'Content-Type: application/json' -d '{"name": "New item", "unit": "nnl"}' 'http://localhost:8118/api/v1/medicines'


```