# 由两个栈组成的队列

## 描述

编写一个类, 用两个栈实现队列, 支持队列的基本操作（add、poll、peek）。

## 总结

维护两个栈, 一个是输入栈, 一个是输出栈。

只要输出栈为空, 就一次性将输入栈的元素都搬到输出栈中。

必须满足两个条件, 才能使得在栈间转移数据时, 能够不破坏队列的特性:

1. stackPush 必须 **一次性地** 将所有数据都转移到 stackPop 中

2. 只有当 stackPop **为空** 时, 才能进行数据转移操作