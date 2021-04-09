## 剑指 Offer 05. 替换空格

#### https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
##### 题解
```
res := strings.Builder{}
res.WriteString("%20")
res.WriteByte()
```

```

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

 

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
```
