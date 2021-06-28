package ops
type Operators interface {
	Generate(int, int) string
	Degenerate(string) (int, int, error)
}
