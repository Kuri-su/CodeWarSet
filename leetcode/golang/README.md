# LeetCode 详细题解索引

思路清晰， 字符画并茂

* 42 Trapping Rain Water (接雨水)
    * Tag
        * Stack(栈)
    * 解法
        * [暴力解法](stack/42.go)
        * [栈解法](stack/42_stack.go)
        * 左右指针解法 // TODO
        * 动态规划解法 // TODO 

* 71 Simplify Path (简化路径)
    * Tag
        * Stack
    * 解法
        * [栈解法](stack/71.go)
        
* 215 Kth Largest Element in an Array (数组中的第K个最大元素)
    * Tag
        * 分治(快排+二分查找)
        * 堆
    * 解法
        * [分治](heap/215_divide.go) (空间复杂度待优化)
        * [堆](heap/215_heap.go) 
    
* 224 Basic Calculator (简单计算器)
    * Tag
        * Stack(栈)
    * 解法
        * [不太好的解法](stack/224_bad.go)
        * [巧妙的利用栈的解法](stack/224_good.go)
   
* 316 Remove Duplicate Letters (去除重复字母)
    * Tag
        * Stack(栈)
    * 解法
        * [利用 栈 和 计数器 ,使用 贪心思想](stack/316.go)
        * 贪心法(每次拼凑都检查修改前和修改后的)

* 378 Kth Smallest Element in a Sorted Matrix (有序矩阵中第K小的元素)
    * Tag
        * Heap(堆)
        * 二分查找
    * 解法
        * [堆](heap/378_heap.go)
        * [二分查找](heap/378_binary_search.go)(待完成)
        * [O(n) X+Y](heap/378_cube.go)(待完成)
        
* 456 132 Pattern (132模式)
    * Tag
        * Stack(栈)
    * 解法
        * [利用 栈 处理此问题](stack/456.go)       
        
* 591 Tag Validator (标签验证器)
    * Tag
        * Stack(栈)
    * 解法
        * [栈解法](stack/591.go)
        