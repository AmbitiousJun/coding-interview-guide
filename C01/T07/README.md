# 生成窗口最大值数组

## 描述

有一个整型数组 `arr` 和一个大小为 `w` 的窗口从数组的最左边滑到最右边，窗口每次向右边滑一个位置。例如，数组为 [4, 3, 5, 4, 3, 3, 6, 7]，窗口大小为 3 时：

```go
[4 3 5] 4 3 3 6 7     // 窗口最大值为 5

4 [3 5 4] 3 3 6 7     // 窗口最大值为 5

4 3 [5 4 3] 3 6 7     // 窗口最大值为 5

4 3 5 [4 3 3] 6 7     // 窗口最大值为 4

4 3 5 4 [3 3 6] 7     // 窗口最大值为 6

4 3 5 4 3 [3 6 7]     // 窗口最大值为 7
```

如果数组长度为 `n`，窗口大小为 `w`，则一共产生 `n-w+1` 个窗口的最大值。

请实现一个函数。

- 输入：整型数组 `arr`，窗口大小为 `w`。
- 输出：一个长度为 `n-w+1` 的数组 `res`，`res[i]` 表示每一种窗口状态下的最大值。

以本题为例，结果应该返回 {5,5,5,4,6,7}。

## 总结

本题可以借助一个双端队列 `qmax` 来实现一个时间复杂度为 `O(n)` 的解法。

只需要在遍历数组 `arr` 的时候，实时将 `qmax` 维护成一个从队首到队尾降序排列的状态即可，维护时需要注意以下细节：

1. 每遍历到一个元素，就加入 `qmax` 队尾中，在加入前，要循环判断队尾元素是否不大于当前元素，一直弹出直到满足条件后再加入
2. 每遍历到一个元素，要检查一下 `qmax` 队首是否已经超出窗口之外，需要弹出进行移除
