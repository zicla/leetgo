## leetgo
LeetCode in golang


LeetCode的go实现，笔者会按照LeetCode的题号顺序依次实现。

在刷LeetCode题目的途中会将常用的算法单独抽出来，
放在algorithm文件夹中，每个go文件都有一个简单的main函数用于测试。




### golang语法参考

#### 切片(数组)

```go
var identifier []type /* 声明 */
var slice1 []type = make([]type, len) /* 实例化方式1 */
slice1 := make([]type, len) /* 实例化方式2 */
slice1 := make([]type, length, capacity) /* 实例化方式3 */

s :=[] int {1,2,3 } /* 直接初始化 */
s := arr[startIndex:endIndex]  /* 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片 */
s := arr[:] 
s := arr[startIndex:]
s := arr[:endIndex] 
```


#### Map
```go
var map1 map[string]string /* 声明 */
map1 = make(map[string]string) /* 实例化 */
```


#### for循环
```go
for init; condition; post { }  /* 基本形式 */
for i:=0;i<10;i++{}

for condition { }  /*相当于while*/

```

