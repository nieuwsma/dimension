package domain

type Rules struct {
	Rules map[int]Rule
}
type Rule struct {
	Name        string
	Description string
}
