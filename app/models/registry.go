package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Post{}},
		{Model: Section{}},
		{Model: Category{}},
	}
}
