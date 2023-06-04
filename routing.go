package main

import (
	"fmt"
	"net/http"
	"pokemon/entity"
	"pokemon/services"
	"strconv"

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

	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pokemon := services.GetPokemonById(num)

	c.IndentedJSON(http.StatusOK, pokemon)
}

func postPokemon(c *gin.Context) {
	var newPokemon entity.Pokemon
	if err := c.BindJSON(&newPokemon); err != nil {
		return
	}

	pokemons := services.PostPokemon(newPokemon)

	c.IndentedJSON(http.StatusOK, pokemons)
}

func putPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data{Message: "put pokemon"})
}

func deletePokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data{Message: "delete pokemon"})
}
