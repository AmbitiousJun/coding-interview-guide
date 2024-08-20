# 猫狗队列

## 描述

宠物、狗和猫的类如下:

```go
type Pet interface {
	GetPetType() string
}

type Dog struct{}

func (d *Dog) GetPetType() string { return "dog" }

type Cat struct{}

func (c *Cat) GetPetType() string { return "cat" }
```

实现一种猫狗队列结构，要求如下：

- 用户可以调用 `add` 方法将 cat 类或 dog 类的实例放入队列中；
- 用户可以调用 `pollAll` 方法，将队列中所有的实例按照进队列的先后顺序依次弹出；
- 用户可以调用 `pollDog` 方法，将队列中 dog 类的实例按照进队列的先后顺序依次弹出；
- 用户可以调用 `pollCat` 方法，将队列中 cat 类的实例按照进队列的先后顺序依次弹出；
- 用户可以调用 `isEmpty` 方法，检查队列中是否还有 dog 或 cat 的实例；
- 用户可以调用 `isDogEmpty` 方法，检查队列中是否有 dog 类的实例；
- 用户可以调用 `isCatEmpty` 方法，检查队列中是否有 cat 类的实例。

## 总结

自定义一个 Pet 的包装类，包含时间戳字段。

在入队的时候，依次自增地将时间戳赋值到包装对象上，出队时以时间戳为依据即可。