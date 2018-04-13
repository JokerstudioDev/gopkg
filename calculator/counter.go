package calculator

type Counter struct {
	Count int
}

func NewCounter(count int) *Counter {
	if count == 0 {
		panic("count equal 0")
	}
	return &Counter{Count: count}
}

func (c *Counter) Increase() {
	c.Count++
}
