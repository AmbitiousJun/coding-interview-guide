// 通用的整型数结构, 方便表达递增操作
package integer

type I int

// Get 获取整型数值
func (i *I) Get() int { return (int)(*i) }

// GetAndIncr 获取值并自增
func (i *I) GetAndIncr() int {
	defer func() { (*i)++ }()
	return i.Get()
}

// IncrAndGet 自增并获取值
func (i *I) IncrAndGet() int {
	(*i)++
	return i.Get()
}
