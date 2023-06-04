package services

import (
	"fmt"
	"pokemon/entity"

	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
)

type PokemonDto struct {
	Id        int
	Name      string
	Category  string
	Height    float64
	Weight    float64
	GenderStr string
	Abilities string
}

func GetPokemons() []PokemonDto {
	var result []PokemonDto
	for _, pokemon := range entity.DummyData {

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

func GetPokemonById(id int) PokemonDto {

	var matchData = lo.Filter(entity.DummyData, func(data entity.Pokemon, index int) bool {
		return data.Id == id
	})

	var pokemonDto PokemonDto

	if matchData != nil {
		err := mapstructure.Decode(matchData[0], &pokemonDto)
		if err != nil {
			fmt.Println("Error:", err)
			return pokemonDto
		}

		pokemonDto.GenderStr = matchData[0].Gender.String()
	}

	return pokemonDto
}

func PostPokemon(pokemon entity.Pokemon) []PokemonDto {

	entity.DummyData = append(entity.DummyData, pokemon)

	var result []PokemonDto
	for _, pokemon := range entity.DummyData {

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

func PutPokemon() {
}

func DeletePokemon() {

}
