package definitionsrpc

import "fmt"

//Call call
type Call struct {
}

//FirstCall FirstCall
func (c *Call) FirstCall(a, b *int) (err error) {
	i := 10
	*b = i
	fmt.Println("Received: ", *a)
	fmt.Println("Sending: ", *b)
	return
}
