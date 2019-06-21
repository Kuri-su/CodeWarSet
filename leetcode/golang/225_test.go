package golang

import (
	"fmt"
	"testing"
)

func TestMyStack_Pop(t *testing.T) {
	c := Constructor()
	c.Push(1)
	c.Push(2)
	fmt.Println(c.Top())
	fmt.Println(c.Pop())
	fmt.Println(c.Empty())

}
