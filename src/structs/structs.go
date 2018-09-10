package structs

type Story struct {
	Title   string
	Story   []string
	Options []OptionType
}

type OptionType struct {
	Text string
	Arc  string
}
