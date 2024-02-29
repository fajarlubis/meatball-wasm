package pdf

type Template int
type Orientation string

const (
	DefaultA4 Template = iota
	DefaultLetter
)

const (
	DefaultUnit string = "mm"

	Portrait  Orientation = "P"
	Landscape Orientation = "L"
)
