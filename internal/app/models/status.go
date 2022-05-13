package models

type (
	StatusModel struct {
		Model
	}
)

func (StatusModel) GetAll() (interface{}, error) {
	all := TodoStatusStrings()

	return all, nil
}
