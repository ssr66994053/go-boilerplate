package middleware

// Hello 目标接口 单方法接口
type Hello interface {
	SayTo(name string) string
}

// HelloFunc adapter func 适配目标接口的方法
type HelloFunc func(string) string

// SayTo HelloFunc实现SayTo达到适配的目的 由此提供出代理原始方法的能力
func (fn HelloFunc) SayTo(name string) string {
	// 调用自身
	return fn(name)
}

// HelloMiddleware 中间件 类型  包裹接口
type HelloMiddleware func(Hello) Hello

// WarpHello 链式嵌套中间件
func WarpHello(h Hello, m ...HelloMiddleware) Hello {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
