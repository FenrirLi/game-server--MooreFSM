package machine

type TableConfig struct {
	MaxChairs int
	MaxRounds int
}
func NewTableConfig() TableConfig{
	return TableConfig{
		2,
		2,
	}
}