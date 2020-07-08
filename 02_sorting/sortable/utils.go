package sortable

type Sortable interface {
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
}