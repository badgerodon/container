package container

// A Reversable is a container which can be reversed
type Reversable interface {
	Len() int
	Swap(i, j int)
}

// Reverse reverses a container
func Reverse(c Reversable) {
	sz := c.Len()
	for i := 0; i < sz/2; i++ {
		c.Swap(i, sz-i-1)
	}
}
