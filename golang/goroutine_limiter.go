package golang

// GLimiter is to limit the number of goroutines
type GLimiter struct {
	n int
	c chan struct{}
}

func NewGLimiter(n int) *GLimiter {
	return &GLimiter{
		n: n,
		c: make(chan struct{}, n),
	}
}

func (g *GLimiter) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}
