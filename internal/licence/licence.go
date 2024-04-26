package licence

type Copyright struct {
	Year   int
	Holder string
}

type Licence struct {
	Type string
	Copyright
}
