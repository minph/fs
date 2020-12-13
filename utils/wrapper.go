package utils

// Wrapper 包含错误类型的包装器
type Wrapper struct {
	Err []error     // 错误信息指针切片
	Val interface{} // 包装值
}

// NewWrapper 创建包装器
func NewWrapper(val interface{}, err error) *Wrapper {
	return &Wrapper{
		Err: []error{err},
		Val: val,
	}
}

// AddError 错误信息指针切片中增加错误
func (w *Wrapper) AddError(err ...error) *Wrapper {
	for _, e := range err {
		w.Err = append(w.Err, e)
	}
	return w
}

// GetError 判断没有错误
func (w *Wrapper) GetError() bool {
	return len(w.Err) != 0
}

// Check 检查错误
func (w *Wrapper) Check(call func(err []error)) *Wrapper {
	if w.GetError() {
		call(w.Err)
	}
	return w
}

// Unwrap 检查错误，获取包装值并提供默认值
// 当存在错误时返回默认值
func (w *Wrapper) Unwrap(call func(err []error), value interface{}) interface{} {
	if w.GetError() {
		call(w.Err)
		return value
	}
	return w.Val
}
