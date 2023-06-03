package services

import (
	"fmt"
	"pokemon/entity"

	"github.com/mitchellh/mapstructure"
)

type PokemonDto struct {
	Name      string
	Category  string
	Height    float64
	Weight    float64
	GenderStr string
	Abilities string
}

func GetPokemons() []PokemonDto {
	dummyData := entity.GetDummyPokemons()

	var result []PokemonDto
	for _, pokemon := range dummyData {

		var pokemonDto PokemonDto
		err := mapstructure.Decode(pokemon, &pokemonDto)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		pokemonDto.GenderStr = pokemon.Gender.String()
		result = append(result, pokemonDto)
	}

	return result
}
