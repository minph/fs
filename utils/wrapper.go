package utils

// Wrapper 包含错误类型的包装器
type Wrapper struct {
	Errs []error     // 错误信息指针切片
	Val  interface{} // 包装值
}

// NewWrapper 创建包装器
func NewWrapper(val interface{}) *Wrapper {
	return &Wrapper{
		Val: val,
	}
}

// AddError 错误信息指针切片中增加错误
func (w *Wrapper) AddError(errs ...error) *Wrapper {
	for _, e := range errs {
		if e != nil {
			w.Errs = append(w.Errs, e)
		}
	}
	return w
}

// GetError 判断没有错误
func (w *Wrapper) GetError() bool {
	if len(w.Errs) == 0 {
		return false
	}

	for _, err := range w.Errs {
		if err != nil {
			return true
		}
	}
	return false
}

// Check 检查错误
func (w *Wrapper) Check(call func(errs []error)) *Wrapper {
	if w.GetError() {
		call(w.Errs)
	}
	return w
}

// Unwrap 检查错误，获取包装值并提供默认值
// 当存在错误时返回默认值
func (w *Wrapper) Unwrap(call func(err []error), value interface{}) interface{} {
	if w.GetError() {
		call(w.Errs)
		return value
	}
	return w.Val
}
