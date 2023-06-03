package entity

func GetDummyPokemons() []Pokemon {
	return []Pokemon{
		{
			Id:        1,
			Name:      "Pikachu",
			Category:  "Mouse",
			Height:    0.4,
			Weight:    6.0,
			Gender:    1,
			Abilities: "Static",
		},
		{
			Id:        2,
			Name:      "Charizard",
			Category:  "Flame",
			Height:    1.7,
			Weight:    90.5,
			Gender:    2,
			Abilities: "Blaze",
		},
		{
			Id:        3,
			Name:      "Bulbasaur",
			Category:  "Seed",
			Height:    0.7,
			Weight:    6.9,
			Gender:    0,
			Abilities: "Overgrow",
		},
	}
}
