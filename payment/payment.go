package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		p := Payment{
			From:   "Wile E. Coyote",
			To:     "ACME",
			Amount: 123,
		}
	*/

	p := NewPayment("Wile E. Coyote", "ACME", 123)
	p.Execute()
	p.Execute()
}

func (p *Payment) Execute() {
	/*
		p.once.Do(func() {
			p.execute(time.Now())
		})
	*/
	p.do()
}

func (p *Payment) execute(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("%s: %s -> [%d¢] -> %s\n", ts, p.From, p.Amount, p.To)
}

func NewPayment(from, to string, amount int) *Payment {
	p := Payment{
		From:   from,
		To:     to,
		Amount: amount,
	}

	p.do = sync.OnceFunc(func() {
		p.execute(time.Now())
	})

	return &p
}

type Payment struct {
	From   string
	To     string
	Amount int // ¢

	// once sync.Once
	do func()
}
