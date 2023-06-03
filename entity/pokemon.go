package entity

type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Other  Gender = 3
)

func (g Gender) String() string {
	names := map[Gender]string{
		Male:   "Male",
		Female: "Female",
		Other:  "Other",
	}
	if name, ok := names[g]; ok {
		return name
	}
	return "Unknown"
}

type Pokemon struct {
	Name      string
	Category  string
	Height    float64
	Weight    float64
	Gender    Gender
	Abilities string
}
