# python 入门
# 一：Python简介

1. 它是一门编程语言  

2. 它也使用自动内存管理（我为什么要说也呢？）

3. 语法优美，易用！各种标准库支持，可以运行在多个平台上！

   > 想不到吧。java（此处应有表情包）

4. 没了（小声比比，目前AI首选编程语言）

# 二：别瞎比比了，亮代码

Python提供了两种写程序的手段，

1. 一种命令行直接怼，很简单，右键打开cmd，敲入python。输入

   `print("你好，再见，拜拜啦")`，回车。搞定

2. 另一种就是写一个py脚本，把代码写进去，然后借助IDE或者命令行继续怼

   `python.exe C:/xxxx/hello.py`

   简单粗暴
   
   

# 三：语法特点

## 3.1 通用

1. 和go语言一个尿性，语句后面不需要加分号。而是用换行符来标识一个语句

   甚至变量也不用声明，直接拿来用。

   > python内心OS : go语言，java，你们都弱爆了
   
   > 如果要用多行来表示一个语句，则在行与行之间用`\`来连接，或者把多行用括号、花括号等包裹起来
   
2. 不用`{}`来表示语句块，而是用缩进来，缩进的行数不用考虑，但是相同的语句必须要有相同的缩进量。

   如果同个语句块，使用不同的缩进量，则解释器无法识别一个完整的语句块。

   这里指的语句块，包括循环，方法体

   > 不正确的行缩进会报`IndentationError`

3. Python区分大小写

## 3.1 注释

1. 一切以`#`开头的行都被Python解释器认为是注释

2. 多行注释的话，行首和行尾用`'''` 或者`"""` 包裹起来

3. 还有一种多行注释,行首和行尾用`"""`包裹起来 ，可以被解释器识别到，并通过一个特殊的执行语句获取到

   ```
   def double(num):
       """我要搞事情"""
       return 2 * num
   ```

   对这个函数执行这个语句

   ```
   print(double.__doc__)
   ```

   可以打印那句注释

   > 猜测Python也有所谓函数值的说法，`__doc__`就是函数值的属性咯
   >
   > 

## 3.2 循环

**for循环**

语法

```
for val in sequence:
	## Do something with val
```

举例

```python
for i in range(1, 11):
    print(i)
    if i == 5:
        break
```

遍历集合的方法

```
genre = ['pop', 'rock', 'jazz']

# iterate over the list using index
for i in range(len(genre)):
	print("I like", genre[i])
```

值得一提的是，循环后面还可以接一个`else`语句块，用于循环结束时，会调用到

但是如果循环中用break退出了，则不会执行到该语句块

使用方法

```
a = range(1, 10)
for i in a:
    print("i is [{}]".format(i))
else:
    print("end!")
```

**while循环**

语法

```
while expression:
	//to do something
```

不解释

哦，对了，while循环后面依然还可以跟着一个else语句块，意为while条件不满足时执行此语句块

其实包含两层意思啦，

1. 就是第一次进入while循环时条件不满足，立即执行else语句块

2. 就是while结束循环之时，亦是条件不满足之时，此时也会执行else语句块。

但是和上面的for一样，break也会跳过else语句块



## 3.3 变量

想必你也知道了，Python根本不存在声明变量(深的javaScript的精髓)。想用直接赋值

```
number = 10
number = 1.1

```

听说这不是叫赋值，只是Python把值的引用给变量而已



## 3.4 常量

常量一般在模块里面定义。

这里的模块，暂且认为是一个py文件吧。这个文件里面可以写常量，然后通过import语句导入到主文件里面，最后使用它们。

首先我们创建一个存放常量的模块文件，命名为`constant.py`

```
PI = 3.13
NAME = "你爸爸"
PASSWORD = "不搞淑敏"
```

然后创建一个主文件，随便命名吧。

然后里面写上

```
import constant
print(constant.PI)
```

就酱，没啦。

不过虽然说是常量，但是实际上，你想改这个常量也是OK，解释器不会阻止你。

额鹅鹅鹅，所以，怎么说呢？约定俗成吧。。。哈哈哈哈（尬笑）

## 3.5 字面量

这个东西就是你直接写一段字面量上去。Python会认为这些字面量是什么类型，其实绝大多数编程语言都会对字面量进行类型判断的。

不然你以为在JAVA定义一个字符串常量，编译器是如何得知这个常量是字符串的。

偏题了，所以说，Python也有这个大众功能。举栗子

```
a = 0b1010  # 二进制字面量
b = 100  # 十进制字面量
c = 0o310  # 八进制字面量
d = 0x12c  # 十六进制字面量

# 浮点数字面量
float_1 = 10.5
float_2 = 1.5e2

# 复杂字面量，估计是复数
x = 3.14j
```

值得注意的是，你将这些字面量用print函数打印出来，它们都会变成十进制







## 3.4 数据类型

在Python里面，数据类型是类，而数据类型的值，就是类的实例。

然后就是介绍Python有哪些数据类型了，其实很容易就能想到

### 3.4.1 数据类型

数字类型，有`int`  `float`  `complex`（复数）

> 注意：float最多只有15位小数，如果小数点超出，则小数部分会被截断

复数的声明需要注明一下哈

```
 c = 1+2j
```

其他都很简单。得益于Python不用声明变量，所以直接用赋值语句定义一个变量并存储对应的值。

Python会自己做类型推断

然后介绍一下怎么表示各种进制的数字

| 进制     | 前缀         |
| :------- | :----------- |
| 二进制   | '0b' or '0B' |
| 八进制   | '0o' or '0O' |
| 十六进制 | '0x' or '0X' |

接下来说一下数字计算的通病，这是所有编程语言都不能绕过的

什么通病：

就是计算机其实无法准确的存储小数

所以对于小数的计算：`1.1 + 2.2`是无法给出正确答案：`3.3`

如果一定要算出上面的结果，可以使用Python的`decimal`模块

使用如下

```
from decimal import Decimal as D
print(D('1.1') + D('2.2'))
```



### 3.4.2 列表

也就是java中的集合了，但是神奇的一点的是，Python内的集合的元素不要求是同种类型。

也就是说，你往里面塞数字，字符串都是OK的

代码示例

```
set = [1, "two", 3, 4.00]
print(set)
print(set[2])
print(set[3])
set[2] = "fuck"
set[3] = "shift"
for i in set:
    print(i)
```

同go语言一样，可以使用`:`来截取一段列表的子集出来

代码示例

```
set = [1, "two", 3, 4.00]
print(set)
print(set[1:3])
```

另外，和绝大多数编程语言一样，只包含头，不包含尾

**列表里面可以包含列表**

```
n_list = ["Happy", [2,0,1,5]]

# Output: a  取第一个元素，即字符串的第二个字符
print(n_list[0][1])    

# Output: 5 取第二个元素，即列表的的第4个字符
print(n_list[1][3])
```

**列表索引允许负数**

效果是往后索引

**列表里面追加数据**

```
odd = [1, 3, 5]
odd.append(7)  # 添加一个元素7
odd.extend([9, 11, 13])  # 添加一个列表进列表里面
odd = odd + [9, 7, 5] # 添加一个列表进列表里面
strArr = ["re"] * 3  
```

**列表插队**

```
odd = [1, 9]
odd.insert(1,3)
# 最后是 [1, 3, 9]
```

**列表删除**

```
my_list = ['p','r','o','b','l','e','m']
del my_list[2]  # 删除第二个元素
del my_list[1:5]  
del my_list  # 整个列表删除，不能再访问
# remove()  pop()
# clear() 清空
```

方法一览

| [**append（）** - 在列表末尾添加一个元素](https://www.programiz.com/python-programming/methods/list/append) |
| ------------------------------------------------------------ |
| [**extend（）** - 将列表的所有元素添加到另一个列表中](https://www.programiz.com/python-programming/methods/list/extend) |
| [**insert（）** - 在定义的索引处插入一个项目](https://www.programiz.com/python-programming/methods/list/insert) |
| [**remove（）** - 从列表中删除项目](https://www.programiz.com/python-programming/methods/list/remove) |
| [**pop（）** - 删除并返回给定索引处的元素](https://www.programiz.com/python-programming/methods/list/pop) |
| [**clear（）** - 从列表中删除所有项目](https://www.programiz.com/python-programming/methods/list/clear) |
| [**index（）** - 返回第一个匹配项的索引](https://www.programiz.com/python-programming/methods/list/index) |
| [**count（）** - 返回作为参数传递的项数的计数](https://www.programiz.com/python-programming/methods/list/count) |
| [**sort（）** - 按升序对列表中的项目进行排序](https://www.programiz.com/python-programming/methods/list/sort) |
| [**reverse（）** - 反转列表中项目的顺序](https://www.programiz.com/python-programming/methods/list/reverse) |
| [**copy（）** - 返回列表的浅表副本](https://www.programiz.com/python-programming/methods/list/copy) |

**建立List的另一种方式**

示例

`pow2 = [2 ** x for x in range(10)]`

结果会创建如下列表

```
[1, 2, 4, 8, 16, 32, 64, 128, 256, 512]
```

其实就相当于下面这个程序

```
pow2 = []
for x in range(10):
   pow2.append(2 ** x)
```

一些比较特立独行的例子

```
pow2 = [2 ** x for x in range(10) if x > 5]
结果：`[64, 128, 256, 512]`
[x for x in range(20) if x % 2 == 1]
结果：`[1, 3, 5, 7, 9, 11, 13, 15, 17, 19]`
[x+y for x in ['Python ','C '] for y in ['Language','Programming']] 
结果：`['Python Language', 'Python Programming', 'C Language', 'C Programming']`
```

**判断元素是否在列表内**

用in关键字

```
my_list = ['p','r','o','b','l','e','m']

print('p' in my_list)
```

**遍历列表**

```
for fruit in ['apple','banana','mango']:
    print("I like",fruit)
```



### 3.4. Set（集）

同样和上面的三个，都是容器，不过这个容器存放的元素是无序的。

定义方法

```
a = {1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 7, 1}
# 下面的代码运算失败，因为Set无序，用索引访问也就没有意义了
# print(a[1])
print(a)
```

可以看到Set的两个特性

1. 唯一性，也叫不可重复性，如果添加相同的元素，但是这个元素已经存在了，那么本次添加是无效的，，不过也不会报错

2. 无序性，当然打印看不出来，但是如果想用索引访问Set，那是白日做梦，痴心妄想。癞蛤蟆。。

   > （对不起，偏题了）
   
   
   
   



### 3.4. 字典（Dictionary）

其实更像是Map。教程介绍说，这是一个字典数据类型，数据结构是`key-value`

一般用于存放大量数据,key和value可以是任意类型。

虽说是字典，但是这个字典啊，可以随意更改，像任人打扮的小姑娘一样。

不瞎比比了，亮代码

```
a = {1: 'this is one', "key": 1, True: False, None: "fuck", "key": "重复的key"}
print("a[1] is", a[1])
print("a[key]", a["key"])
print("a[key]", a["key"])
print("a[True]", a[True])
print("a[None]", a[None])
# runtime error KeyError: 'x' 因为key不存在，就会报这个错误
# print("not exist", a["not exist"])
# 像这种情况可以用get方法 ,如果不存在，是none
print("not exist", a.get("not exist"))
a[1] = 33

print("changed a[1] is", a[1])


```

特性：没有，与Map雷同

我觉得不应该叫它字典，它明明就是一个不支持查无操作的Map啊！

支持的方法(由自身调用)

1. pop(key)  删除一个key-value 返回该value
2. `popitem()`随机删除一个key，并返回它的value
3. clear()  清空字段
4. del 删除这个字典等

**通过反人类的循环语法创建字典**

```
squares = {x: x*x for x in range(6)}

# Output: {0: 0, 1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
print(squares)
```



**迭代字典**

```
squares = {1: 1, 3: 9, 5: 25, 7: 49, 9: 81}
for i in squares:
    print(squares[i])
```



**使用in  或者  not in 操作符判断key是否存在**

```
squares = {1: 1, 3: 9, 5: 25, 7: 49, 9: 81}
print(1 in squares)
```







### 3.4.3 Tuple（元祖）

这其实跟列表的同父异母，失散多年的兄弟。

唯一不同的是，元祖里面的数据不可以更改，这意味着如果你想改变里面的某个元素的值，是没有办法的。

基于这个特性，元祖通常比列表要快那么一丢丢。且大多被用在写保护上

话说了这么多，怎么定义一个元祖呢？

```
tupleInstance = (1, 2, "this is tuple", "nothing")
# 以下代码运行时报错  'tuple' object does not support item assignment
# tupleInstance[1] = 1
print(tupleInstance)
# 以下不用括号来建立元祖
my_tuple = 3, 4.6, "dog"
# 由于上面这个特性和其他定义变量的形式冲突，所以如果你想用上面这种形式建立元祖
my_tuple = ("hello",)
my_tuple = "hello",
需要在后面加`,`，不然的话，编译器会错认类型是string
```

元祖支持`+`  `*`操作符

示例

```
a = ((1, 2, 3) + (4, 5, 6))
# 上面将新增一个元祖：(1, 2, 3, 4, 5, 6)
b = ("Repeat",) * 3
# 上面将新增一个元祖：(('Repeat', 'Repeat', 'Repeat')
```

元祖的优点

1. 迭代比list快
2. 元祖可以作为字典的Key,List不可以
3. 

### 3.4.5 空类型

Python用None来表示空，也即未创建字段。

使用示例

```
drink = "beer"
food = None


def bar(menu):
    if menu == drink:
        print("please wait ,beer preparingg")
    else:
        print(food)


bar(drink)
bar(food)
```

对了,None不能连接字符串。

> java：Python你丫的弱爆了

### 3.4.6 布尔类型

Python的布尔类型是首字母大写的，也就是`True`和`False`

而且学py的人还可以回忆下当年学C的风光。

因为在py里面，非0值和0值`None也是0值`也被认为是True和False

说白了，这两者是等价的。

所以，这个算式OJBK的

`a = True + 1`

以上

### 3.4.7 python的字符串

Python的字符串加了很多特技

简单的 `a="我是字符串"`

这个就不用多说了。

关键是其他写法，6的一批

1. `multiline_str = """This is a multiline string with more than one line code."""`

   这个是多行的字符串,你可以写多行的字符串，就问你6不6

   和多行注释谜之相似

2. `unicode = u"\u00dcnic\u00f6de"`  这就是定义了一个Unicode的字符串

   前面的`u`就是为了标识而生的。

3. `raw_string = r"raw \n string"`  这定义了一个原始的字符串，一切按原始数据存放，不管什么换行符什么鬼的。

4. 字符串还支持`*`运算符了，初见时惊为天人

   怎么说，例如啊例如

   ```
   x = "global"
   x = x * 2
   print(x)
   ```

   打印结果为`globalglobal`

   佩服佩服，这种设计真是够了，剑走偏锋，就要跟其他语言区分开来

另外，Python字符串和GO语言一样，都可以用`:`截取一部分字符串作为子集

还可以通过下标访问，不过不能更改

这个就不演示了，超简单

Python的字符串其实相当于元祖了，元祖能用方法，它一般也能用

**如何在字符串中存储`""`**

1. 使用转义字符

   > a = '''He said, "What's there?"'''

2. 使用三引号

   > ```
   > a = 'He said, "What\'s there?"'
   > 
   > a = "He said, \"What's there?\""
   > ```

3. 

### 3.4.8 数组

Python也有数组，不过相对来说，它的存在感薄弱。

如何创建它呢？

```
import array as arr
a = arr.array('d', [1.1, 3.5, 4.5])
print(a)
```

数组终于像其他编程语言一样，只能存储特定类型的元素了

怎么说，唔，你得在前面写上一个标识符，标识这个数组，存的是什么元素。

比如说上面的`d`，就是告示解释器，这个数组存double类型的数据

除了`d`外，你还可以用其他标识符，来看一看吧

| Code  | C 语言类型     | Python Type | 最小字节 |
| :---- | :------------- | :---------- | :------- |
| `'b'` | signed char    | int         | 1        |
| `'B'` | unsigned char  | int         | 1        |
| `'u'` | Py_UNICODE     | Unicode     | 2        |
| `'h'` | signed short   | int         | 2        |
| `'H'` | unsigned short | int         | 2        |
| `'i'` | signed int     | int         | 2        |
| `'I'` | unsigned int   | int         | 2        |
| `'l'` | signed long    | int         | 4        |
| `'L'` | unsigned long  | int         | 4        |
| `'f'` | float          | float       | 4        |
| `'d'` | double         | float       | 8        |

数组长度可扩充，其他的，自己扩展吧

一般不建议使用数组，唔。教程说的



### 3.4.9 矩阵

矩阵是一种二维数据，里面的数据排列成列和行

Python内置类型没有矩阵，不过我们可以自己造一个。。。

用嵌套List的方式

```
A = [[1, 4, 5], 
    [-5, 8, 9]]
```

可以视它是二行三列的矩阵

对矩阵的一些使用

```
A = [[1, 4, 5, 12], 
    [-5, 8, 9, 0],
    [-6, 7, 11, 19]]

print("A =", A) # 打印整个矩阵
print("A[1] =", A[1])      # 打印矩阵的第二行
print("A[1][2] =", A[1][2])   # 打印矩阵的第二行的第三个元素
print("A[0][-1] =", A[0][-1])   # 打印矩阵的第一行，最末列的元素
# 以下代码选择矩阵每行的第三个元素，形成新的list
column = [];        # empty list
for row in A:
  column.append(row[2])   

print("3rd column =", column)
```

**转置矩阵**

所谓转置矩阵，就是把原来矩阵的行，变成列，把原本的列，变成行

```
[
[12,7],
[4 ,5],
[3 ,8]
]
```

转置后，变成

```
[12,4,3] 
[7,5,8]
```

**矩阵乘法**

两个矩阵相乘，得左边矩阵的列数和右边矩阵的行数匹配，相乘才有意义





**用NumPy开源包创建矩阵**

虽然用嵌套list可以创建矩阵，然而这也太原始了，让我为你介绍新的工具

`NumPy`

这货是一个科学计算的开源包，建立N维的矩阵不在话下

不过你得先下载它，唔，关于怎么下载嘛，看下面的软件包下载吧

**用法1：创建多维数组**

```
A = np.array([[1, 2, 3], [3, 4, 5]])
print(A)

A = np.array([[1.1, 2, 3], [3, 4, 5]])  # Array of floats
print(A)

A = np.array([[1, 2, 3], [3, 4, 5]], dtype=complex)  # 里面的元素是复数
print(A)
```

**创建元素是0的多维数组**

```
zeors_array = np.zeros( (2, 3) )
```

这将生成如下的矩阵

```
 [[0. 0. 0.]
  [0. 0. 0.]]
```

**创建元素是1的矩阵**

```
ones_array = np.ones( (1, 5), dtype=np.int32 ) 
```






### 3.4.x其他

可以使用type函数来判断这个值是什么类型的，还可以用`isinstance`函数来判断给定值是否是指定类型

代码举例

```
a = 3
print(type(a))
print(isinstance(a, int))
```

输出

```
<class 'int'>
True
```

**通用容器方法**

这里的通用容器指的是一维数组和衍生品

它们都支持

1. 求并集  把两个集合用`|`  得到新的集合，其实就是两个集合啦

2. 求交集，把两个集合用`&` 得到新的交集

3. 求差集，用`-`联合两个集合

   ```
   A = {1, 2, 3, 4, 5}
   B = {4, 5, 6, 7, 8}
   print(A - B)
   ```

   结果生成A集合里面，不属于交集部分的集合。也就是`{1, 2, 3}`

4. 求对称差，把两个集合用`^`得到的新集合

   结果是两个集合，不包括交集部分

5. 



## 3. 类型转换

类型转换分为两种，一种显式，一种隐式

### 3.1 显式类型转换（强转）

和绝大多数编程语言接轨，强转的话，也是把值放到类型+括号里面

比如说

```
a = 5
print(5)
print(float(5))
print(int(5.99999))
```

本来是int类型，要把它转成float，就是如此简单

本来是float类型，要转成int，就会丢掉小数，别妄想什么四舍五入，tan90~

当然，强转也是有限制的

```
print(int("sdfgdh"))
```

你就不能把字符串转成数字

比起其他编程语言，Python有一点厉害的，就是允许字符串直接强转成数字，eg`int("124")`

其他的强转，还有容器的，直接贴教程代码

```
a = [1,2,3]
set(a)  
# 其实等同于：set([1,2,3])
# 集合set转元祖
tuple({5,6,7})
# 字符串转列表
list('hello')
dict([[1,2],[3,4]])
{1: 2, 3: 4}
```

```
# 两个Set作为一个大Set里面的元素，然后强转为字典
dict([[1,2],[3,4]])
{1: 2, 3: 4}
```

字典的转化要类似于这样格式，才可以转化成功。

不过，我觉得没人会用这个特性搞一个字典，可读性不好，又没什么实际作用。

**有趣的强转**

`eval(arg)`

参数可以接受一个表达式，返回值就是表达式的值

比如

`num = eval("1+2")`

### 3.2 隐藏类型转换

其实和绝大多数语言一样，在某些运算中，算式的结构会自动转换成参与运算中类型更广的哪一个

eg

```
# 这是一个int类型
num_int = 123
# 这是一个float类型
num_flo = 1.23
# 两者相加，结果的类型会隐式转换成float
num_new = num_int + num_flo
```

 

## 4：导入

Python也需要导入语法来导入来自其他Python文件的内容

这里的导入，指的是导入模块，模块即可以是Python自带的，也可以是自己写的

比如说，Python有一个内置的模块是`math.py`，你想输出这个模块里面的pi常量，就需要导入

程序如下

```
import math
print(math.pi)
```

其实所谓模块，只不过是一个Python的文件罢了

还可以**只导入模块里面几个常量或者方法**

```
from math import pi
# from math import *  这样的话，是将math里面的所有符号导入进来主模块的命名空间
print(pi)
```



只有在`sys.path`能看到的目录里面的Python文件，才可以用导入语法导入进来，其他位置的Python文件不可以导入

当然我们可以往sys.path添加的说

**导入并且给模块重命名**

代码示例

```

import math as m
print("The value of pi is", m.pi)
```

**重复导入的话，解释器会只导入一次**

顺带导入时的执行语句，也只执行一次

如果想在程序进行中，再次导入之前已经导入的语句的话，可以执行这个方法

```
imp.reload(my_module)
```

当然，执行语句也会再执行一次哦



## 5：算术运算符

介绍几个Python独有的，其他的懒得介绍，和其他编程语言一样。

1. `**` 指数运算。举例`2**6`，用数学语言描写就是`2^6` 

2. `//` 取整，或是叫地板除，比如说`6//4`，本来按照数学的除法来说，答案应该是`1.5`

   不过用地板除，结果就是1了

> 关于地板除译法，有一个小插曲，想象一栋两层楼，楼上住着2，楼下住着1，中间住着1.5
>
> 地板除将会得到楼下的数字，而不是中间或是楼上。
>
> 同理，还有一个属于叫天花板除。。你懂得

3. 逻辑运算符，这里的`&`和`|`以及`!`被`and` 和`or`以及`not`篡位了，其他用法一样

   除了传统的`!=`  `==`仍然支持,`&& ||`已经不用来进行逻辑判断了，改用`and`和`or`了

   另外，除了支持传统的`!=`  `==`外，还支持`not is`来进行类似的判断。

   不过，前者是用于判断值的，后者是用于判断地址的

4. 位运算符
| 运算符 | 含义     | 举例子                     |
| ------ | -------- | -------------------------- |
| ＆     | 按位AND  | x＆y = 0（`0000 0000`）    |
| \|     | 按位OR   | x \| y = 14（`0000 1110`） |
| ~      | 按位NOT  | ~x = -11（`1111 0101`）    |
| ^      | 按位异或 | x ^ y = 14（`0000 1110`）  |
| >>     | 按位右移 | x >> 2 = 2（`0000 0010`）  |
| <<     | 按位左移 | x << 2 = 40（`0010 1000`） |

5. 标识运算符

   `is`和`is not`代表两个标识运算符，标识运算符用于确定运算的两个变量，是否用的是同一块内存。

   该运算符和`=`运算符不能等同，后者仅仅判断字面量是否相等。

   有时候，`=`表达式为true的两个变量，用于`is`运算符。结果不一定也为True

   举例子

   ```
   a = 5
   b = 5
   print(a is b)
   # 结果为True 说明a和b使用的是同一块内存
   s1 = "hello"
   s2 = "hello"
   print(s1 is not s2)
   # 结果为False，说明s1和s2使用的是同一块内存
   x3 = [1, 2, 3]
   y3 = [1, 2, 3]
   print(x3 is y3)
   # 结果为False，说明x3和y3使用的不是同一块内存
   ```

6. 成员关系运算符

   这个就比较高端了，可以判断某个元素是否在容器里面。

   语法是

   `element in container `    `element` 是否在`container`里面

   `element not in container `    `element` 是否不在`container`里面

   比较特殊的是字典容器,只能判断key是否存在

   举例子

   ```
   s = "hello,world"
   print("hello" in s)
   container = (1, 2, 3, 4, 5)
   print(6 not in container)
   dictionary = {"key": "value", "index": 3}
   print("value" in dictionary)
   print("key" in dictionary)
   ```

## 6：函数

函数在Python中，也是一个对象，可以用值来接受它

```
def printHello():
    print("Hello")     
a = printHello()

# Output: Hello
a
```

这句代码定义了一个函数`printHello`，并用a来接受它

然后是一个`a`就执行了`printHello`函数里面的内容

## 7：判断表达式

语法

```
if expression:
    //to do something
```

使用举例

```\
a = 3
if a == 3:
    print("OJBK")
```

再来

```
if expression：
    //to do something
else：
	//to do something
```

举例

```
a = 4
if a == 3:
    print("OJBK")
else:
    print("No Ojbk")
```

其他的用法懒得解释了，相信你有幸看到我的笔记的时候，应该已经浸淫编程语言多年了吧。这么简单的举一反三就不用说了吧。



## 8:pass空操作

这是一个新语法，这个语法可以单独使用，它使用之后没有任何效果

> 要不把它命名为水溅跃吧。笑~

那它还有个卵子用！？摔

且慢，存在即合理，其实哪怕是空操作，也有存在的意义呢？

你知道吗？Python不允许函数体是空的，但是我们定义一个函数后，想后期再实现咋办。

这个时候直接用pass，函数体就不是空的，而且执行后也没有任何东西。

完美

使用举例

```
def hahahah():
    pass

# 这是一个混眼熟的代码，请无视它
print(1)
```



## 9 函数

语法：

```
def function_name(parameters):
	"""docstring"""
	statement(s)
```

参数值不随缘，必填的

什么，你问函数的返回值是什么？

返回值是随缘的，主要看一下情况

1. 没有return 语句，默认返回一个None

2. 只有return语句，也是返回一个None

3. return expression 返回这个表达式的值

4. 可以返回多个返回值。

   其实看成只返回一个列表也是可以的，得益于Python的列表可以包含万象，所以函数体只需返回一个列表，

   就相当于支持了多返回值了。

   当然接受的话，也是按照列表的方式来接受的

   **函数参数**
   
   函数参数可以指定一个默认值。指定了默认值的参数在调用时是可选的
   
   例子
   
   ```
   def methon4(name=3):
       print("name is [{}]".format(name))
   methon4()
   ```
   
   如果默认参数后面还有函数参数，则后面的函数参数也必须设置默认值。
   
   嗯，这是动态语言的通病
   
   **手动指定形参位置**
   
   以前的编程语言，形参和实参的顺序必须一致，现在嘛，Python可以手动指定形参顺序了
   
   亮代码
   
   ```
   def methon4(name=3):
       print("name is [{}]".format(name))
   
   
   methon4(name="safs")
   ```
   
   > 讲真，只是些语法糖而已
   
   **可变参数**
   
   我就知道有这货
   
   亮代码
   
   ```
   def methon5(*names):
       for name in names:
           print("hello,{}".format(name))
   
   
   methon5("李大喵", "张全碳", "训悟空")
   ```
   
   > `*`的另一种用法——标记可变参数
   
   **关键字参数**
   
   在函数定义上，你可以在函数的形参上使用关键字参数。
   
   该参数允许调用这传入零个或多个带参数名的参数，而在函数内，这个传来的参数会变成dict，也就是字段
   
   举例
   
   ```
   def showParameter(**kw):
       print("parameter:", kw)
   
   
   showParameter(name="fuck", value="cao")
   ```
   
   
   
   
   
   **匿名函数/Lambda函数**
   
   由于在Python中，匿名函数用Lambda定义，所以才有这个别称
   
   代码举例
   
   ```
   double = lambda x: x * 2
   print(double(2))
   ```
   
   未完待续。。。
   
   
   
   

# 四：命令

Python一旦下载下来安装，并添加了环境变量，你就可以在任何地方打开cmd，敲入python来执行Python命令

支持的命令有

`quit()`  退出Python命令行   

# 五：常用的函数

## 5.1输出

输出我们已经用过了，就是简单的`print()`函数，但是其他还有很多用法

首先隆重介绍下print函数的定义

```
def print(self, *args, sep=' ', end='\n', file=None)
```

1. self  别问我，我也不知道，后面可能回补上
2. *args 要打印的数据
3. 分隔字符，如果args传入的是多个数据的话，输出时就会用这个分隔符分隔
4. end 输出结束，再补上一个什么鬼的结束符，默认是回车
5. file  输出目的地，默认是`sys.stdout`

使用举例

```
print(1, 2, 3, 4)
# Output: 1 2 3 4

print(1, 2, 3, 4, sep='*')
# Output: 1*2*3*4

print(1, 2, 3, 4, sep='#', end='&')
# Output: 1#2#3#4&
```

**格式化输出**

这是字符串的一个方法。

调用方式1

```
'The value of x is {} and y is {}'.format(x,y)
```

其中，`{}`作为占位符。

调用方式2

模板字符串里面还可以写顺序，如

`'The value of x is {1} and y is {0}'.format('x','y')`

则输出

`The value of x is y and y is x`

调用方式

`s = '我是{name},今年{age}'.format(name='李大喵', age=13)`

输出`我是李大喵,今年13`

## 5.2 输入

输入也很简单，一个函数调用即可

`result = input(arg)`

arg参数是可选的，它仅作为输入的提示，显示在控制台上。

输入的结果是字符串

## 5.3 range函数

可以用它来生成一端数字

用法：range(1,10)

生成1~10的一段数字

但是它其实不会把数字全部生成在内存中，而是在获取的时候再生成

这个更多的是搭配循环来食用

具体请查看循环的语法哈

## 5.4 len(arg)函数

不解释，仅记录

## 5.5 dir(arg)

参数传模块名（注意，不是字符串）

这个函数可以列出模块所有命名过的对象名称

如变量，函数名，也会顺带把模块内置的一些名称变量名和方法名列出来

比如

```
['__builtins__', '__cached__', '__doc__', '__file__', '__loader__', '__name__', '__package__', '__spec__', 'variable']
```

形如`__xxxxx__`都是模块内置的一些名称，比如说`__name__`存储的，是模块名，可以用`模块名.__name`__

来访问他，最末尾的，就是我们在那个模块定义的一个变量名`variable`



对了，如果不传参数的话，默认是得到本模块所有定义过的名称



## 5.6 type()函数

得到该变量是什么类型的函数

## 5.7 isinstance()函数

得到该变量是否属于此类型的函数



## 5.8 拷贝函数

拷贝函数可以浅拷贝和深拷贝一个对象

拷贝函数需要导入`copy`模块

使用示例

```
import copy

old_list = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
new_list = copy.deepcopy(old_list)

old_list[1][0] = 'BB'

print("Old list:", old_list)
print("New list:", new_list)
```



```
import copy

old_list = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
new_list = copy.copy(old_list)

old_list[1][1] = 'AA'

print("Old list:", old_list)
print("New list:", new_list)
```





# 六： 相关警告输出

1. `shadows built-in name`  内置名称，避免歧义，不要用内置名称作为变量名

# 七：命名空间

首先说明一点，在Python中，所有东西都叫对象

Python对于简单的`a=3`是这么理解的。

1. `3`是存在内存中的一个对象
2. `a`是和内存中的`3`关联起来的一个名称,这个名称无关类型，因此可以用于命名所有类型

这是很简单的介绍。

下面的重点是

**Python不会创建重复的对象**

这意味着，如果你

用a定义了一个a变量`a=3`

用b定义了一个变量`b=3`

这两个变量指向的内存都是一样。

这个隆重介绍下Python的取地址函数:`id(arg)`

实验如下

```
a = 3
b = 3
print(id(a))
print(id(b))
print(id(3))
```

结果都是`1396663424`

这说明，至始至终，这个程序只在内存创建一个`3`对象

更有趣的用法

```
a = 2  # 创建了一个2对象，a指向这个对象

print('id(a) =', id(a))
a = a+1  # 创建了一个3对象，a指向这个对象

print('id(a) =', id(a))  
 
print('id(3) =', id(3))  # 打印之前创建在内存里面的3对象的地址
b = 2

print('id(2) =', id(2))  # 打印之前创建在内存里面的2对象的地址
```

以上只是开胃小菜，接下来介绍命名空间

何谓命名空间

简单来说，就是名称的集合，用于隔离相同名称。

当我们启动Python的解释器时，会创建属于Python自己的命名空间，里面包含了Python可用的所有内置名称。

而这就是为什么我们能使用各种内置函数`id()`  `print()`的原因

每个模块都会创建属于自己的命名空间。所以，你可以在不同的Python文件定义两个相同名称的变量，而他们不会冲突

每个函数也会创建自己的命名空间，里面包含这个函数里面所有名称的定义

如果函数和模块有同样的变量名，则函数里面修改那个相同的变量的值，不会导致模块的那个变量名遭到修改。

也就是说，函数对本模块的变量名只有Read权限，没有修改权限，你要修改的话，也影响不到外部的变量

有一个关键字，可以让函数对模块里的变量有查看和修改权限。

这个关键字是`global`

使用举例

```
def func():
    global a
    a = 233


a = 10
func()
print(a)
```

但是只能用于这种场景，如果是函数内嵌套函数，想使用global来修改外部函数的变量。。

还是洗洗。。还是有手段的，用`nonlocal`在嵌套函数里面标记上一层函数的某个变量，则在嵌套函数的修改

对外部函数也是可见的

## 7.1模块

在Python中，模块被认为是一个Python文件，想在其他Python文件中使用Python模块，可以使用import语法

当解释器遇到import语句，会去创建import对应模块的命名空间，并且执行该模块里面，可以执行的语句。

这会导致一种情况

如果后一个模块修改了前一个模块的全局变量

那么在两个模块都导入的主模块，主模块访问的变量值是已被修改的了

不过，重复导入的话，始终只执行一次执行语句

# 八：关键字

1. `global` 用这个关键字标识的变量

   如果处在全局变量的位置上，则毫无卵用

   如果处在局部变量上，这个局部变量摇身一变会变成全局变量

   即使这个变量名，在此之间，并没有在全局变量上声明过

   
   
   
   
# 九：安装开源包

类比`java`的`maven`和`go`的`go get `

Python也与时俱进的提出了pip思想,它是Python的软件包安装程序

一般来说，高版本的Python已经包含了这个东西，我们不需要再安装它了

蛤，它可以安装托管在Python包索引的开源包

使用示例

`python -m pip install --user numpy scipy matplotlib ipython jupyter pandas sympy nose`

它可以安装`numpy `开源包



简单的功能有

1. `pip install name` 根据软件包的名称来安装软件包，需要网络连接

   如果本地已经下载了安装包，可以`pip install path`来安装它

2. 显示你安装了哪些软件包

   ```
   pip show --files SomePackage
   ```

3. 看看你哪些软件包已经outdated

   `pip list --outdated`

4. 升级包

   ```
   pip install --upgrade SomePackage
   ```

   一个叫做`SomePackage`的包得到升级啦

5. 卸载包

   ```
   pip uninstall SomePackage
   ```

# 十：文件系统操作

Python怎么能少了文件操作呢？

## 10.1 文件操作

### 10.1.1 打开文件

在Python中，打开一个文件的代码如下所示

```
f = open("test.txt")
```

这将返回一个文件对象，也叫句柄

默认是以文本方式打开的，且是只读，当然我们可以在打开文件时指定

指定的方式如下

```
f = open("test.txt",'w')
```

用一个`w`来标识打开的文件用于写，如果文件不存在，会创建该文件，如果文件存在，则会覆盖文件

其他的标识符有

| Mode | Description                                                  |
| :--- | :----------------------------------------------------------- |
| 'r'  | 已只读方式打开文件，这是默认选项                             |
| 'w'  | 以写方式打开文件，如果文件存在，则覆盖，如果文件不存在，则创建该文件 |
| 'x'  | 以独占方式创建该文件。如果文件存在，抛出异常                 |
| 'a'  | 以写方式打开该文件，所做的修改将追加在文件末尾，而不会覆盖   |
| 't'  | 以文本模式打开文件                                           |
| 'b'  | 以二进制方式打开文件                                         |
| '+'  | 以读写方式打开文件                                           |

在文本方式打开文件时，强烈建议制定一个文件编码，不然会依据平台的默认编码打开文件。

如何制定

```
f = open("test.txt",mode = 'r',encoding = 'utf-8')
```

### 10.1.2 如何关闭文件

很简单，用句柄对象的CLose方法

```
f = open("test.txt",encoding = 'utf-8')
# perform file operations
f.close()
```

更好的写法是

```
try:
   f = open("test.txt",encoding = 'utf-8')
   # perform file operations
finally:
   f.close()
```

`finally:`块里面的代码可以保证一定会执行，所以不必担心中途遇到异常，程序退出，但是资源没释放

终极写法是

```
with open("test.txt",encoding = 'utf-8') as f:
   # perform file operations
```

你完全不用写关闭文件的代码，Python会自动帮你关闭

### 10.1.3 如何写入文件

打开文件时，需要指定好正确的打开方式

```
with open("test.txt",'w',encoding = 'utf-8') as f:
   f.write("my first file\n")
   f.write("This file\n\n")
   f.write("contains three lines\n")
```

   蛮糟糕的一点是，换行符需要你自己输入。。

### 10.1.4 如何读取文件

```
>>> f = open("test.txt",'r',encoding = 'utf-8')
>>> f.read(4)    # 读取首四个字符

>>> f.read(4)    # 再读取下面四个字符


>>> f.read()     # 一次性全部读取


>>> f.read()  # 如果文件已经被读取完，再次调用它只会返回空串
```

其他方法

1. `tell()`    告诉你当前读取到第几个字符了
2. `seek(0)` 更改读取的位置，这里0,就是从0开始读取了
3. `readline()`



中级读取文件大法：循环法

```
 for line in f:
     print(line, end = '')
     #  end = ''  可以避免再打印一个换行符，因为文件里面本身就存在换行符了
结果是：
This is my first file
This file
contains three lines
```

## 10.2 文件夹操作

直接列出方法（os模块）

1. `getcwd()`  得到当前的工作文件夹

2. `chdir(str)` 输入一个文件夹路径以变更工作目录

3. `listdir(str)`列出指定文件夹里面的文件和文件夹列表

   如果不传参数，默认是工作目录

4. mkdir(str) 创建文件夹

5. rename(original，newName)  能变更文件夹名字

6. remove()删除文件

7. `rmdir()`删除一个空目录

   > 目录必须为空，不为空删除不了

8. `shutil.rmtree()`  这个不是OS模块的方法，但是它可以删除非空的文件夹




# 十一：Python里面的异常机制

Python的异常有两种

1. 一种是语法错误或解析错误，类似于编译错误
2. 一种是运行时错误，被称为异常，Python里面存在许多这种内置异常

## 11.1 内置异常

Python有多少种内置异常呢？

调用这个方法一看究竟

`locals()['__builtins__']`

Python有哪些常见的内置异常呢？

| 异常名                | 异常                                                         |
| :-------------------- | :----------------------------------------------------------- |
| AssertionError        | y由于`assert` 语句断言失败                                   |
| AttributeError        | 属性赋值或者访问时失败                                       |
| EOFError              | 读到预料之外的文件结尾                                       |
| FloatingPointError    | 浮点数运算失败                                               |
| GeneratorExit         | 调用自动生成的close函数失败                                  |
| ImportError           | 不能找到导入的模块                                           |
| IndexError            | 索引越界                                                     |
| KeyError              | 字典中不存在该key                                            |
| KeyboardInterrupt     | 用户键入中断键，Ctrl+C或者del键时会触发该错误                |
| MemoryError           | 操作耗尽内存啦                                               |
| NameError             | 当变量在本地和全局范围没有找到时引发                         |
| NotImplementedError   | 从抽象方法抛出                                               |
| OSError               | 系统操作导致系统相关错误时引发                               |
| OverflowError         | 当运算的结果太大导致无法显示引发                             |
| ReferenceError        | Raised when a weak reference proxy is used to access a garbage collected referent. |
| RuntimeError          | 当错误不属于任何一种类型时引发                               |
| StopIteration         | 调用next方法，但是已经没有更多的迭代项，由此引发             |
| SyntaxError           | 遇到语法错误由解析器引发                                     |
| IndentationError      | 不正确的缩进语法                                             |
| TabError              | 缩进由不一致的tab和空格引发                                  |
| SystemError           | 解释器遇到内部错误                                           |
| SystemExit            | 由sys.exit()函数抛出                                         |
| TypeError             | 当函数或者操作应用不正确的类型引发                           |
| UnboundLocalError     | Raised when a reference is made to a local variable in a function or method, but no value has been bound to that variable. |
| UnicodeError          | 当遇到Unicode相关的，解码和编码操作时引发                    |
| UnicodeEncodeError    | 在编码期间发生与Unicode相关的错误时引发。                    |
| UnicodeDecodeError    | 在解码期间发生与Unicode相关的错误时引发。                    |
| UnicodeTranslateError | Raised when a Unicode-related error occurs during translating. |
| ValueError            | 不恰当的函数值，比如强转传入的值。                           |
| ZeroDivisionError     | 当除法或模的操作，第二个操作数为0时引发                      |

除了以上系统内置的异常，我们还可以定义自己的异常

## 11.2 用户自定义的异常

自定义异常，需要创建新类，该类必须是`Exception`的子类

扔出代码

```
class CustomerException(Exception):
    pass


try:
    raise CustomerException("fuck 辣妹儿")
except:
    print("异常发生了哦")
finally:
    print("扫尾工作")
```







## 11.3 处理异常

在Python里面，异常可以由自己处理，也可以一级级的向上抛出，直至解释器打印出异常，并中断程序的执行

怎么处理异常，使用try...except.....来捕获并处理特定异常

代码示例

```
randomList = ['a', 0, 2]

for entry in randomList:
    try:
        print("即将除与", entry)
        r = 1 / int(entry)
        break
    except:
        print("Oops!", sys.exc_info()[0], "occured.")
        print("下一个被除数.")
        print()
    finally:
    	# 释放资源用，总是会执行到
    	pass
print("The reciprocal of", entry, "is", r)
```

可以预料的是，这个程序会出现两个异常

1. `a` 不能转成数字，报的是`ValueError`
2. 不能除与0，报的是`ZeroDivisionError`

在这个程序中，我们调用了` sys.exc_info()[0]`代码来打印错误类型，其实在里面，我们还可以打印更多信息

因为`sys.exc_info()`  返回的是一个列表之类的东西。

经验证，

第零个元素就是错误类型

第一个元素就是引发错误的原因。

第二个元素是内存跟踪（不确定）

**处理异常的方式-升级版**

```
randomList = ['a', 0, 2]

for entry in randomList:
    try:
        print("即将除与", entry)
        r = 1 / int(entry)
        break
    except (ValueError, ZeroDivisionError):
        print("Oops!", sys.exc_info()[0], "occured.")
        print("下一个被除数.")
        print()
    except:
        print("其他异常捕捉", sys.exc_info()[0])
print("The reciprocal of", entry, "is", r)
```

明确捕捉特定的异常。

**手动抛出异常**

代码示例

```

try:
    raise ValueError("手动写BUG,就问你怕不怕")
    # raise KeyboardInterrupt
except (ValueError, ZeroDivisionError):
    print("Oops!", sys.exc_info()[0], "occured.")
    print("下一个被除数.")
    print()
except:
    print("其他异常捕捉", sys.exc_info()[0])
```



# 十二：Python对象

一个对象有两个特征

1. 属性
2. 行为

Python支持面向对象编程的几个特点

1. 继承
2. 封装
3. 多态

## 12. 1 构建类

要想创建一个对象，首先要创建类

在Python里面，类的代码大致如下

```
class Parrot:

	"this is class doc desc "
	
    species = "bird"

    def __init__(self, name, age):
        self.name = name
        self.age = age

    pass
```

这就创建了一个类，里面有一个类属性`species`

还有一个构造函数`__init__`，里面传入两个参数，作为对象的属性

> 对，你没看错，对象的属性是在构造函数才定义的，和java的有所不同
>
> 莫名觉得这才是正确的方式。。

> 当我们定义一个类，就会创建与该类有相同名称的类对象，并分配一个类独有的命名空间。
>
> 然后这个类对象允许我们访问类属性（类名.类属性访问），以及实例化该类的对象

使用类的方法创建一个对象

```

a = Parrot(name="小鹦鹉", age=14)
# 直接打印对象属性
print(a.name)
# 拿到对象所属的类，再拿到类属性
print(a.__class__.species)

a.temp = "这是一个临时对象属性"
```

**可以在对象创建之后，再动态定义对象的属性**

> 你知道吗，当你在用`对象.属性名`访问对象的属性时，其实内部在用对象的`__dict__`字典中搜索有该属性名的值。
>
> 换句话说
>
> a.temp
>
> 等价于
>
> `a.__dict__["temp"]`

上面在创建类的过程中，顺带介绍了属性，还有方法没介绍

```
class Parrot:
    species = "bird"

    def __init__(self, name, age):
        self.name = name
        self.age = age

    def sing(self, song):
        print("一个名为", self.name, "的小鸟正在引吭高歌：", song)

    pass


a = Parrot(name="小鹦鹉", age=14)
print(a.name)
a.sing("鸡你太美")
```

可以看到特殊的一点是，方法的第一个参数始终是self，但是实际上我们没有传入这个参数，或者说它不是必传的

其实这个参数，就是调用这个方法的对象，默认对象调用方法时，自动传入的。

也就是说

`a.sing("鸡你太美")`等同于`Parrot.sing(a,"鸡你太美")`

有点类似于golang接受者的概念,只不过接受者的位置不一样罢了

对了`self`可以起其他名称，不过不建议，除非你想报复社会

**删除对象，删除属性**

可以用del（它真是万能的方法。。）

值得注意的是，del的对象和属性，并不会在内存中直接删除，而是存活一段时间，等GC删除



## 12.2 继承类

亮出代码

this parent class

```
class bird:

    def __init__(self):
        print("鸟准备好了")

    def sing(self, song):
        print("鸟在叫，" + song + ",人坏掉")

    def fly(self):
        pass

    pass
```

this is child class

```

class penguin(bird):
    def __init__(self):
     	super().__init__()
        print("企鹅蓄势待发")

    def fly(self):
        print("阿勒，企鹅这么飞来着")
```

use case

```
p = penguin()
# call parent class methon
p.sing(song="哼哼哼")
# call child class methon
p.fly()
```

几个特点

1. Python所有类都继承自object
2. `isinstance()`用于判断第一个参数（对象）是不是第二个参数（类）的类对象或者子类对象
3. `issubclass()`用于判断第一个参数是否是第二个参数的子类
4. Python里面的类允许继承多个（它有多个老爸）

## 12.3 封装（访问权限符）

对于Python类里面的某些属性，我们可以在属性名前面加前缀`__`来标识它只能在对象内部使用

不能在外部访问，也不能在外部修改

代码示例

```
class person:
    def __init__(self):
        self.__wife = "xt"

    def showWife(self):
        print("my wife is {}".format(self.__wife))

    def changeWife(self, wife):
        self.__wife = wife


p = person()
# this code will raise error
# print("try access wife of p", p.__wife)
print("try change wife of p")
p.__wife = "lzp"
p.showWife()
print("real change wife ")
p.changeWife("not object found !")
p.showWife()
```

## 12.4 多态

Python变量可是号称万物都可容纳，接受一个区区的类对象有何不可。

直接给代码

```
class parrot:
    def fly(self):
        print("look I can Fly,you see see you")

    def swim(self):
        print("look ,I ...I .. can't swim")


class penguin:
    def fly(self):
        print("I can't fly ,because I is pengui")

    def swim(self):
        print("I can swim ,you see you")


# 通用方法，参数可接受一切鸟类
def fly(bird):
    bird.fly()


p = parrot()
fly(p)
p = penguin()
fly(p)

```

## 12.5 重载方法，重载操作符

额，我想说的重载不是简单的重载父类方法，而是重载内置函数

比如，在类里面，我们可以重载`__str_`来改变默认的对象字符串表现形式

如果你不重载这个方法，打印对象，会展现这个鬼东西

```
__main__.point object at 0x00A54C50
```

多难看啊，你说是不

代码示例

```
class point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    # 重载字符串表现形式方法，print(instance)可见效果
    def __str__(self):
        return "self x is {} y is {}".format(self.x, self.y)

    # 重载+运算符，将两个对象+起来，将得到一个新的对象
    def __add__(self, other):
        return point(self.x + other.x, self.y + other.x)


p1 = point(1, 2)
p2 = point(3, 4)
print(p1)
print(p1 + p2)

```

其他操作符重载函数

| Operator            | Expression | Internally          |
| :------------------ | :--------- | :------------------ |
| Addition            | p1 + p2    | p1.__add__(p2)      |
| Subtraction         | p1 - p2    | p1.__sub__(p2)      |
| Multiplication      | p1 * p2    | p1.__mul__(p2)      |
| Power               | p1 ** p2   | p1.__pow__(p2)      |
| Division            | p1 / p2    | p1.__truediv__(p2)  |
| Floor Division      | p1 // p2   | p1.__floordiv__(p2) |
| Remainder (modulo)  | p1 % p2    | p1.__mod__(p2)      |
| Bitwise Left Shift  | p1 << p2   | p1.__lshift__(p2)   |
| Bitwise Right Shift | p1 >> p2   | p1.__rshift__(p2)   |
| Bitwise AND         | p1 & p2    | p1.__and__(p2)      |
| Bitwise OR          | p1 \| p2   | p1.__or__(p2)       |
| Bitwise XOR         | p1 ^ p2    | p1.__xor__(p2)      |
| Bitwise NOT         | ~p1        | p1.__invert__()     |

你以为这样就没了，年轻！还有逻辑运算符重载啊

```
class Point:
    def __init__(self, x = 0, y = 0):
        self.x = x
        self.y = y
    
    def __str__(self):
        return "({0},{1})".format(self.x,self.y)
    
    def __lt__(self,other):
        self_mag = (self.x ** 2) + (self.y ** 2)
        other_mag = (other.x ** 2) + (other.y ** 2)
        return self_mag < other_mag
Point(1,1) < Point(-2,-3)
```

再来一个逻辑运算符重载函数

| Operator                 | Expression | Internally    |
| :----------------------- | :--------- | :------------ |
| Less than                | p1 < p2    | p1.__lt__(p2) |
| Less than or equal to    | p1 <= p2   | p1.__le__(p2) |
| Equal to                 | p1 == p2   | p1.__eq__(p2) |
| Not equal to             | p1 != p2   | p1.__ne__(p2) |
| Greater than             | p1 > p2    | p1.__gt__(p2) |
| Greater than or equal to | p1 >= p2   | p1.__ge__(p2) |

## 12.6 类的高级特性

这个特性可以让

你在访问类的属性时，实际是访问类的Get方法

你在设置类的属性时，实际是访问类的Set方法

show code

```
class user:
    def __init__(self):
        self.name = ""

    def setName(self, name):
        print("即将设置", "name值")
        self._name = name

    def getName(self):
        print("即将获取", "  name值")
        return self._name

    name = property(getName, setName)


u = user()
u.name = "王兰花"
print(u.name)
```

实际结果如下

```
即将设置 name值
即将设置 name值
即将获取   name值
王兰花
```

更高级的写法

```
class user:
    def __init__(self):
        self.name = ""

    @property
    def name(self):
        print("即将获取", "  name值")
        return self._name

    @name.setter
    def name(self, name):
        print("即将设置", "name值")
        self._name = name


u = user()
u.name = "王兰花"
print(u.name)
```





# 十二：迭代器

迭代器其实是一种设计范式。

它主要提供两个或三个方法，`next`和`iter`，或者还有`hasNext`

在Python实现迭代器，可以使用最原始的迭代器思想，一个个next，直到结束。

或者也可以将迭代器用在for循环上，与`java`的高级for循环一样

不比比了，亮代码

```
class PowTwo:
    """Class to implement an iterator
    of powers of two"""
	# 创建迭代器
    def __init__(self, max = 0):
        self.max = max
	# 得到迭代器自身。并重置迭代次数
    def __iter__(self):
        self.n = 0
        return self
	# 获取下一个元素
    def __next__(self):
        if self.n <= self.max:
            result = 2 ** self.n
            self.n += 1
            return result
            # 如果没有更多元素，抛出此异常
        else:
            raise StopIteration
```

一般的用法

```
>>> a = PowTwo(4)
>>> i = iter(a)
>>> next(i)
1
>>> next(i)
2
>>> next(i)
.....
```

学霸的用法

```
 for i in PowTwo(5):
	   print(i)
```

**生成器，定义迭代器的另一种方式**

在Python函数中，如果出现了`yield`语句，则这个函数可以作为生成器使用。

定义代码如下

```
def my_gen():
    n = 1
    print('This is printed first')
    # Generator function contains yield statements
    yield n

    n += 1
    print('This is printed second')
    yield n

    n += 1
    print('This is printed at last')
    yield n
```

> 在每一个调用这个生成器的next方法时，实际上就进入了这个方法体

> 这个方法体会在执行起来，直至遇到`yield`语句，就会把该语句后面的值丢出去，并暂停执行

> 下次再执行`next`方法时，代码会从上一次中断的地方开始。

这个函数并不能直接调用，正确的使用方法是

```
a = my_gen()
next(a)
next(a)
next(a)
```

或者更高级用法

```
for i in my_gen():
    print(i)
```

解释一下：

`my_gen()`这个语句会创建一个迭代器对象。里面默认含有next方法

```
def _next_(slef):
	.....
```

所以你可以这么用`next(a)`   或是`for i in my_gen()`  或者`a.__next__()`

用这张方式创建的一个迭代器对象，只能迭代一次，想再迭代一次，需要再调用这个代码`my_gen()`

**生成器也有表达式**

总体来说，类似于列表的表达式，只不过`[]`用`()`代替

举例子

```
my_list = [1, 3, 6, 10]
# 用列表表达式创建了一个新的列表
[x**2 for x in my_list]
# 用生成器表达式创建了一个生成器
a = (x**2 for x in my_list)
next(a)
next(a)
next(a)

```

当生成器作为参数传递给函数时，`()`可以忽略

比如说

```
sum(x**2 for x in my_list)
```

**使用生成器创建迭代对象的好处**

1. 简单，代码量小

   弄一个传统的迭代器，你得弄一个类，实现两个方法

   但是用生成器，你一个方法就搞定了

   > 虽然它并不能像真正的方法一样调用



# 十三：闭包

动态语言怎么可能没有这货。

所谓闭包，一般是酱紫的

1. 一个函数里面定义了嵌套函数
2. 嵌套函数里面用到了外部函数的变量
3. 外部函数将嵌套函数当做返回值返回出去

这其实也是实现闭包的手段。

Python的闭包有两个特点

1. 即使你把外部函数用del删除了，闭包依然可以正常使用
2. 闭包函数的值会记住外部函数的值。

实现闭包的代码如下

```
def external(msg):
    def internal():
        print(msg)
    return internal


internal = external("我胡汉三又回来了")
del external
internal()

```

所以函数对象都有`__closure__`属性，如果这个函数是一个闭包，那么它的`__closure__`属性是一个容器。

而容器的第一个元素就是闭包的值

# 十四：装饰器

以往我们在java使用装饰器，只能增强某个类的某个方法的功能。

现在到了Python，可以用装饰器来装饰方法了，这主要是通过闭包来实现的

实现代码

```
def log(func):
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper

def helloWorld():
    print("hello,wordl")
```

使用方式

```
a = log(helloWorld)
a()
```

这将装饰helloWorld这个方法。

使其在调用`helloWorld`方法之前，先打印调用的是什么函数

里面有几个特殊的点

1. `*args`是一个可变参数，可接受零到无限个参数值

2. `**kw`是可以关键字参数，可接受零到无穷个带参数名的参数值

   结合两个点，便可以将一个函数的参数，传递到任意函数的参数上

   好像也只能用在这里了。。。ORZ

> 或许你也有跟我一样有一个疑问，认为装饰器没必要使用闭包，
>
> 装饰器函数内调用装饰代码，然后返回被装饰的函数或者执行它
>
> 嗯，这样确实更直观，不过呢，有一个问题，就是上面的`*args, **kw`特性就用不了了



**装饰器改进版**

是不是觉得使用装饰器调用方法有点不习惯。

你得调用另一个方法，拿到闭包函数，再调用闭包函数，好不习惯。

别急，这里有**装饰器高级版**

```
def log(func):
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper

@log
def helloWorld(name):
    print("dict",name)
    print("hello,wordl")


helloWorld("你个渣渣")
```

其实就是在被装饰的方法上，加上`@+装饰函数名`

这样在就可以在调用原始的函数上，增强这个函数的功能了。

通过这个语法，还可以为这个函数装饰多个功能。

最后其实类似于拦截器，程序必须先执行所有的装饰器后，再执行真正的函数

## 十五：断言

在Python中使用断言，来假设你的程序能按预期进行，如果不能按预期进行，就会产生一个`AssertionError`

使用语法如下

```
assert <condition>
assert <condition>,<error message>
```

举例

```
def avg(marks):
    assert len(marks) != 0
    return sum(marks)/len(marks)

mark1 = []
print("Average of mark1:",avg(mark1))
```



```
def avg(marks):
    assert len(marks) != 0,"List is empty."
    return sum(marks)/len(marks)
    
 
mark1 = []
print("Average of mark1:",avg(mark1))
```

我觉得很不错哦

关键是断言失败会告诉你在哪地方发生了断言失败。

有了它，貌似不用再写单元测试了耶。（逃~







