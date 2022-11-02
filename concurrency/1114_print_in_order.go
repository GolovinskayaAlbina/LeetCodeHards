package concurrency

type orderPrinter struct {
	sStart chan struct{}
	tStart chan struct{}
	res    chan struct{}
	result string
}

func NewOrderPrinter() *orderPrinter {
	return &orderPrinter{
		sStart: make(chan struct{}),
		tStart: make(chan struct{}),
		res:    make(chan struct{}),
		result: "",
	}
}

func (p *orderPrinter) First(s string) {
	p.result += s
	p.sStart <- struct{}{}
}
func (p *orderPrinter) Second(s string) {
	go func() {
		<-p.sStart
		p.result += s
		p.tStart <- struct{}{}
	}()
}
func (p *orderPrinter) Third(s string) {
	go func() {
		<-p.tStart
		p.result += s
		p.res <- struct{}{}
	}()
}

func (p *orderPrinter) GetResult() string {
	<-p.res
	return p.result
}
