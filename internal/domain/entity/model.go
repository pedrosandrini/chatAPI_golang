package entity

type Model struct {
	Name      string // Modelo do chat que sera usado (3, 3.5, 4 etc..)
	MaxTokens int
}

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

func (m *Model) GetMaxtokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
