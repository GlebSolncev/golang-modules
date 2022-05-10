package models

type (
	Status struct {
		Model
	}
)

func (Status) GetAll() (interface{}, error) {
	all := TodoStatusStrings()

	return all, nil
}
