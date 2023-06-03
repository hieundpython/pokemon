Project give function for add and search pokemon over transport internet 
dataset: https://github.com/fanzeyi/pokemon.json

Architecture
- Golang 
- PostgreSql
- Elasticsearch
- Swagger
- Monolithic app
- Gin gonic


Api: using rest api

GET: pokemons/{id}
GET: pokemons 
POST: pokemons
PUT: pokemons/{id}
DELETE: pokemons/{id}


TODO:
- Using dataset for show pokemon list
- Write unit test for project
- Migration dataset into database using background job
- DevOps: Using docker container, deploy into AWS or Google Engine