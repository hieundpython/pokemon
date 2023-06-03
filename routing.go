package main

import (
	"net/http"
	"pokemon/services"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string
}

func InitAppRouting(r *gin.Engine) {
	r.GET("/pokemons", getPokemons)
	r.GET("/pokemons/:id", getPokemonById)
	r.POST("/pokemons", postPokemon)
	r.PUT("/pokemons/:id", putPokemon)
	r.DELETE("/pokemons/:id", deletePokemon)

}

func getPokemons(c *gin.Context) {
	pokemons := services.GetPokemons()
	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonById(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, Data{Message: id})
}

func postPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data{Message: "post pokemon"})
}

func putPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data{Message: "put pokemon"})
}

func deletePokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data{Message: "delete pokemon"})
}
