package middleware

import (
	"fmt"
	"testing"
)

type helloImpl struct {
}

func (h *helloImpl) SayTo(name string) string {
	return fmt.Sprintf("Hello_%s", name)
}

func logger(h Hello) Hello {
	return HelloFunc(func(name string) string {
		fmt.Println("logger before SayTo")
		res := h.SayTo(name)
		fmt.Println("logger after SayTo")
		return res
	})
}

func prefix(h Hello) Hello {
	return HelloFunc(func(name string) string {
		fmt.Printf("add prefix before %s\n", name)
		res := h.SayTo(name)
		res = fmt.Sprintf("prefix_%s", res)
		fmt.Printf("add prefix after %s\n", res)
		return res
	})
}

func subfix(h Hello) Hello {
	return HelloFunc(func(name string) string {
		fmt.Printf("add subfix before %s\n", name)
		res := h.SayTo(name)
		res = fmt.Sprintf("%s_subfix", res)
		fmt.Printf("add subfix after %s\n", res)
		return res
	})
}
func Test1(t *testing.T) {
	h := &helloImpl{}
	nh := WarpHello(h, logger, prefix, subfix)
	res := nh.SayTo("World")
	exp := "prefix_Hello_World_subfix"
	if exp != res {
		t.Errorf("exp %s, actual %s\n", exp, res)
	}
}
