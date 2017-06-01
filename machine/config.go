package machine

type TableConfig struct {
	Max_chairs int
}
func NewTableConfig() TableConfig{
	return TableConfig{2}
}