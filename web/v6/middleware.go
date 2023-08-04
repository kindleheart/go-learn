package v6

// Middleware 函数式的责任链模式(洋葱模式)
type Middleware func(next HandleFunc) HandleFunc
