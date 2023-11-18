# go语言笔记

# 一:简要概述

go语言是Google开发的一门用于服务器编程的,具有很好并发优势的一门编程语言

它具有以下特点

1. 语法级别的并发支持
2. 编译型语言
3. 运行速度快,占有内存低,反正比Java快就是了
4. 编写者一定是个特立独行的家伙(请忽略这一句)

# 二:入门

跟Java一样,go也需要一个编译器和运行器,然而你只需要安装一个安装包,这两者都有了,很简单不是吗

在这一章里,我们要做的事有

1. 搜索go官方网站,下载go安装包,并安装

2. 设置两个环境变量

   1. `GOROOT`  设置为go的安装路径

   2. `GO111MODULE` 设置为`on` 

      > 或许你可以设置那个啥,嗯GOPATH 把它设置为go代码编写的工作空间
      >
      > GO111MODULE和GOPATH 的配置是互斥来着,以后再讲解吧

3. 编写一个go文件,输出hello吧

   ```
   package main
   
   import "fmt"
   
   func main(){
   	 fmt.Print("你好，世界")
   }
   
   ```
   
      
   
4. 编译这个文件

   `go build xxxx`

   编译后会生成一个EXE文件,仅限Windows

5. 运行这个文件,命令行敲入文件就行了

**几个要点**

1. package是包名,将来可以用`包名.方法名/属性名`访问到对应的数据,一般来说,包名要和所属的文件夹名一样

   > 不这么做也可以,go编译器不理这事,但是你的项目经理可能不会容许这样的写法
   >
   > 另外,main包是个例外,main包表示这个包下面可以有个main方法,程序的入口就是在这里啦

2. `import "fmt"` 是导包,不过你可以忽略这个,一般的IDE会自动导,GO语言提供的工具也会帮你导

3. ```go
   func main(){
   	 fmt.Print("你好，世界")
   }
   ```

   这个是一个最简单的方法了,以后你还可以看到各种千奇百怪的方法,做好准备吧

# 三:代码结构

一般go语言的代码结构是这样的

```
root
|-------src
|----------------controller/xx.go,xxx.go
|-------bin   
bin是输出路径
```

一个go文件的代码结构是

最顶层是包

包中间的变量是包级别变量,生存周期和包一样

包里面还有函数

函数内再有变量,生存周期是函数生存周期

# 四: 基础语法

## 4.1 命名规则

必须以字母或者下划线开头,大小写严格区分,不能和关键字,预定义名字雷同

命名和访问权限挂钩

如果首字母小写,那么仅在同包内,可以访问

如果首字母大写,那么在不同包,加上`包名.`还是可以访问到

建议短命的变量,起个短平快的名字,活的长的变量,再起个丰富一点的名字

建议用驼峰命名法命名

> 建议,你可以不准守,如果你真这么做的话,小心项目经理42米长的大刀

## 4.2 声明

但凡编程语言,谈到声明,无非就是那几个

1. 声明变量 语法`var xx type`  或者是var xx 表达式

   > 第二个声明语法的变量类型,go内部将用表达式推导
   >
   
2. 声明方法  语法

   ```
   func xxxx(parameter1 type1,parameter2 type2....)(returnType1,returnType1)
   ```

   举例子

   ```
   func Method(a int,b int)(string,int){
   	return fmt.Sprintf("%d%d",a,b),a+b
   }
   ```

3. 声明常量

   ```
   const pi = 3.141594
   const (
   e = 2.71828182845904523536028747135266249775724709369995957496696763
   pi = 3.14159265358979323846264338327950288419716939937510582097494459
   )
   ```

4. 声明类型

   ```
   type Celsius float64
   ```

   类型是什么呢?再往下看

### 4.2.1 零值

变量仅声明,不赋值,一般默认给0值

数字零值:0

布尔零值:false

引用(slice,函数,chan,map,接口....)零值:nil

结构体零值:就是一个空的结构体,只不过里面字段全都是零值

字符串零值:""

## 4.3 赋值

对一个已经定义的变量赋值很简单

`变量名=xx表达式`

如果想一步到位,用一句代码实现声明和赋值,也可以

`变量名:=表达式`

这样go内部会自动推导变量名的类型

如果变量名之前已经声明的话,上面那个表达式亦可以使用,不过就有点多余了

## 4.4强转赋值

某些类型，可以强制转换成其他类型

使用这个语法

`a:=type(xx)`

示例

`var _ io.Writer = (*bytes.Buffer)(nil)`

将nil强转为(*bytes.Buffer)类型

## 4.4 类型

在go语言中,类型的作用就是给基本数据类型换个马甲,让冰冷冷的基本数据类型,换个活法

比如说之前你想声明一个int类型变量,代表摄氏温度

那么你可以

```
var Celsius float64 //这个变量是摄氏温度
```

有了类型后,你可以给`float64`换个马甲,叫做`Celsius`

```
type Celsius float64 // 摄氏温度
```

那么你再声明摄氏温度的变量时,就可以

```
var Temperature Celsius //这个变量是摄氏温度
```

其实个人感觉纯属多此一举了,因为不看好这个语法,就介绍到这吧

## 4.5 循环语法

go语言只有一种循环方式,就是`for`

循环中最常见的是,容器的循环,以数据循环为例

```go
	var arr=[]string{"1","2","3"}
	//传统的循环方式
	for i:=0;i<len(arr) ;i++  {
		fmt.Println(arr[i])
	}
	//很6的循环方式
	for _,value:=range arr  {
		fmt.Println(value)
	}
```

关于那个很6的循环方式,有几个需要注意的

1. range 是一个关键字,后面跟着容器,可以返回两个值,一个是索引,一个是索引对应的值的副本

2. range返回的值是乱序的,据说是编写者故意为之,所以是很6的循环

3. go语言变量声明必须使用,但是也有的需求是不想要某个变量,咋办,机制的编写者给了一个特殊的变量`_`

   该变量允许你不使用它,作为变量垃圾桶再好不过,可惜一行只能用一个

   

## 4.6 IF语法

   `if`语法也很传统,除了一个

```
if err := xx.xxxx(xxx); err != nil {
		ctx.API.SetError(err)
		return
	}
```

这个语法允许你在if判断之前,执行一个预语句

 另外,if语句不需要括号,下面的执行语句才需要括号包起来

## 4.7 defer 语法

defer后面的表达式可以延迟到该代码块执行完毕后再执行

类似于java的finally代码块

语法

```
defer sum()
```

## 4.8 x.(type)语法（类型选择语法）

举例子

```
func Sprint(x interface{}) string {
	switch x := x.(type) {
	case string:
		return x
	case int:
		return strconv.Itoa(x)
		// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}
```

看上去和其他语言并没有什么大区别，实际上，它还是有点区别的

1. switch 后面可以不带任何值，默认为true，这种情况，case表示可以为条件表达式
2. x.(type) 可以将x的真正类型返回，但是这种写法只能用在switch语句中

## 4.9 类型断言（`t := i.(T)`）

形如`t := i.(T)`

就是断言i是T类型，并且将值拷贝到T类型的变量t上

相当于强转

## 4.x 其他需要注意的

1. go语言语句不需要分号结尾
2. go语句对于大小写括号的位置,有着谜一般的执着,你不能随便放
3. 变量一旦声明就必须使用,否则编译运行甩你一脸错误信息
4. 无异常语法,想告诉别人方法执行有误,一般方法要返回一个error类型的数据
5. 一些显而易见的常识我就不普及了,如果你还没学至少一门语言,请出门左转找C语言

## GO关键字

 用go关键字+函数名可以开启一个协程（类似线程）来并发执行这个方法

有1点特殊的

1. 如果用于产生协程的主协程很快结束，则子协程可能会执行不到

   例如

   ```
   func HelloWorld() {
   	fmt.Println("Hello world goroutine")
   }
   
   func main() {
   	go HelloWorld() // 开启一个新的并发运行
   	fmt.Println("我后面才输出来")
   }
   
   ```

   运行几次，很大概率会没调用到HelloWorld方法

   解决办法：

   1. 主协程睡眠

   2. 用通道，主协程go后，从通道里面拿数据，此举动会堵塞主协程，直至拿到数据

      而go协程里面的方法，用通道往里面塞数据

      > 一般是协程执行完毕后，再塞数据。不然协程执行到一半，主协程拍屁股跑了。子协程直接结束
      >
      > 不过这他喵跟非协程调用有毛线区别？？？!!!

      如果是在web环境下，就不需要考虑这个问题，因为主协程始终是存活的



# 五:基础数据类型

## 5.1 整型

go提供和很多位的整形

带符号的:`int8、int16、int32和int64`

不带符号的`uint8、uint16、uint32和uint64`

还有我们经常用的int和uint

但是这两个数据类型的位数在不同的硬件平台上有区别,可能是32或64又或者其他

>整形还有一个运算符,和其他编程语言雷同,不解释
>
>还有位操作运算符,但是感觉不太用到,忽略

## 5.2 浮点数

go提供两种浮点数

`float32`和`float64`



## 5.3 布尔类型

类型名:`bool`

只有两个值`true` `false`

## 5.4 字符串

类型名:string

go语言的字符串是可不变的

可以用内建的len(str)方法计算字符串的**字节**长度

> 字符长度这个算不出来

也可以用str[index]拿到第index个**字节**的数据

> 字节和字符不一样。不能相当然的认为str[index]能拿到第index个字符的数据
>
> 因为对于某些非ascll码的字符，会编码成多个字节。比如说中文

其实换一种角度,字符串就是字符数组啦

字符串可以和字符串用"+"联合起来

字符串可以用`>`  `<`   `==`  来比较

比较是将每一位字符用字符串自然編碼比较

## 5.5 常量

关注声明就行,其他太过简单,不解释

## 5.6 数组

定义数组

```
var a [3]int
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}  //第二个元素填充零值
q := [...]int{1, 2, 3}  //忽略长度,长度go会自己计算
```

数组不可变,所以用的地方不多,不解释

数组运行"="的原理是两个数组的元素一个个比较,全部相等后才认为数组相等

`!=`同理

数组作为参数传参,函数接受到的数组是原始数组的拷贝

## 5.7 切片

长度可变的数组,一般叫做`Slice`


> 那么不可变的数组叫做啥？就是array

有两个属性

1. 长度  对应切片存放的元素数量
2. 容量  实际分配的元素数量

**声明语法**

```go
var list []string
//使用make函数创建一个长度为5的切片
a := make([]int, 5)  
//使用make函数创建一个长度为0,容量为5的切片
b := make([]int, 0, 5)
```

**添加数据**

```
list = append(list,$元素)
//不会耗费大量系统资源,尽管使用此方法扩充切片长度
```

数组和切片操作都有一个方便的语法,用于取子集

假设有一个切片list

list[i:j]就是取第i个元素到第j个元素的切片作为子集

list[1:]取第i个元素到末尾元素的切片作为子集

list[:i]取首个元素到第i个元素的切片作为子集

list[:]取整个切片

切片一般不能比较,除非自己实现方法一个个比较

切片的nil值代表一个长度,容量为0的切片

所以建议,如果想判断一个切片是否为空,最好判断切片的长度

**内存泄露**

如果你将一个数组进行切片，那么切片完的数组其实还是和原数组共享同一套内存块的。  
这可能会导致内存泄露。特别是，如果你以为切片操作是重新生成一个新的切片的话。。。  
即使你将原先的切片解除引用，只要子切片还存在，原先切片的内存就不会被回收。  

**cap**
cap是取数组分配的长度。  但是它从第一个元素开始算起    
之前我们提到slice切出来的新slice，其实还是和原slice共享同一个内存区块的。  
那么理论上，新和原slice的cap要一致。因为拢共就分配那么多内存块出来。  
然而，实际并非如此。新的始终要小点。因为它始终从第一个元素算起。  
假设A 有5个元素。B 取A 从第2个元素以后的元素。  
那么A的cap应该是5.（它的第一个元素还在）       
B的cap是4（A的第一个元素被排除掉了）      

**新特色**
来自2013发布的version 1.2。      
现在reslice可以接受第三个参数作为切出来的容量cap了.    
然而，由于众所周知的原因（cap总是从第一个元素算起）  
实际的cap总是最后一个index（cap）减去第一个index.  
比如说A有4个元素.  
则B[1:4:4] 的cap则是3  


## 5.8 map

这是一个存储键值对的容器.和其他语言的map类型一模一样

> 嗯,key存放的类型必须是能用`==`运算符的
>
> value的话,基本所有类型都OK

不一样的

就是神奇的声明方式了

```
m:=make(map[string]int)  //声明一个空map,key
m:=map[string]int{
	"index":1,
	"age":12,
}
m:=map[string]int{}
```

还有神奇的访问元素方式

> 以上面的m变量为例子

```
//可以访问到map里面"index"里面的value
value:=m["index"]
```

删除语法是

```go
delete(m, $key$)
```

检查元素是否存在

```
elem,has:=m["index"]
//如果存在，has则为true。
```



对了,map元素存放的不是变量,所以下面语法有编译错误

```
&m["index"]
```

> 根源是map会扩充,导致原本的地址偏移,所以取的地址也会偏移
>
> 于是同理,切片元素也不能取地址





## 5.9 结构体

啥是结构体

你把它认为是java中被阉割的实体类就行了

因为它只有字段,没有成员方法,没有构造函数

由于go里面只有这个阉割的实体类-结构体

所以在go的web开发中,一般,只能用结构体来映射数据库的表

讲了这么多,来说说**怎么定义一个结构体吧**

```
type Employee struct {
    ID int
    Name string
    Address string
    DoB time.Time
    Position string
    Salary int
    ManagerID int
}
```

然后**声明结构体变量**的语法是

```
var emp Employee
a:=model.Demo{ID: md.ID}
```

很简单,但是别忘了,结构体没有零值,你声明出来的结构体,直接就可以

**取结构体的属性**

```
emp.ID
```

> 这点倒和java同步了

只不过这个空的结构体,属性全都是零值

下面是结构体的几个特性

1. 结构体指针也可以通过`.`操作符访问到结构体的属性

   > 所以emp.ID 和 (&emp).ID 是等价的

2. 结构体类型里面的属性不能是它本身,所以一下代码编译失败

   ```
   type tree struct {
       value int
       left, right tree
   }
   ```

   但是可以包含它自身类型的指针

   ```
   type tree struct {
       value int
       left, right *tree
   }
   ```

   所以二叉树,链表什么的,还是可以用go语言来实现的

### 5.11 结构体内嵌(类似继承)

使用结构体内嵌，可以使嵌入的那个结构体，拥有其他结构体的字段和方法

首先要定义被嵌入的结构体，一般的结构体就是了

```
type User struct {
	UserName string
	Password string
	ID int
	Sex string
}
func (u *User)SetUserName(name string)  {
	u.UserName = name
}
```

然后就是那个被嵌入的结构体了

```
type Student struct {
	User
	StdNo int
	Class string
	Grade string
}
```

这个被嵌入的结构体`Student`就拥有了`User`所有属性，还可以把它自己作为接受者，调用User的方法

但是嵌入和被嵌入并不是真正意义上的继承关系，至少如果参数是User的话，你是无法传Student进去的

## 5.12 结构体Tag

语法：

```

```





## 5.10 json

go语言倒是没有把json数据类型作为内置的数据类型

只不过提供了几个方法便于结构体和json互转罢了

> 比起java好像更好点?

涉及到的两个方法如下

```go
//结构体转json字节数组
data,err:=json.Marshal(m)
//MarshalIndent 这个输出的json有格式化
//json字节数组转结构体
var movie  []Movie
err:=json.Unmarshal([]byte(jsonStr),&movie)
```

涉及到序列化和反序列化的知识,就还要讲到字段的序列化和反序列化

默认情况下,序列化和反序列化用的字段名肯定是结构体的字段名,如果要覆写的话

需要再结构体的字段上面写`tag`

比如说

```
type Movie struct {
	Id int
	Title string
	//结构体的json序列化和反序列化时,参照的字段是released
	Year int `json:"released"`
	//如果Color字段为零值,比如说false,则序列化成json时不会给出这个字段
	Color bool `json:"color,omitempty"`
	Actors []string
}
```

如上,多的就不解释了

另外,go语言提供的json序列化方法略渣,所以一般企业开发会使用easyjson

据说速度秒甩官方几条街

> 为什么速度这么快,据说是序列化和反序列化是静态生成的
>
> 不是通过反射

## 5.11 格式化输出的标识符

浮点数  `%g`  `%f`



## 5.12 通道 chan

这个通道主要是用于多线程的通讯的

 **声明语法 **

`ch := make(chan int)`

 **向通道发送值 **

```
ch <- v    // 将 v 发送至信道 ch。
```

 **从 ch 接收值并赋予 v **

`v := <-ch `

当通道作为参数时，有几点需要注意

```
func fetch(url string, ch chan<- string) {
	//TODO ....
}
```

像这样声明的函数的通道形参，只能用于接受数据，而不能拿数据

如果`ch chan<- string`改为` ch<- chan string `  则是可读，但不可写

## 5.13 字符

golang内部有有两种类型，来表示字符

一种是byte，通常用于表示ascil字符

一种是rune,通常用于表示unicode或utf-8字符

> 其他的字符类型可能还有，不做保证

golang内建字符编码是UTF-8，在这个编码下，一个汉字，含3个字节



# 六:函数

## 6.1 函数的声明

```
func name(param1 type1,param2 type2) returnType {
	//body
	return xx
}
```

```
func name(param1 type1,param2 type2) (returnType1,returnType2) {
	//body
	return xx,xx
}
```

```
func name(param1 type1,param2 type2) (result1 returnType1,result2 returnType2) {
	//body
	result1=xx
	result=xx
	return 
}
//此函数声明没有函数体，因为此函数不是go语言编写的
func Sleep(d Duration)
```

就这几种,下面还有一个变态型的,这里先不讲

想必你也发现了,函数的返回值是可以多个的

而且还可以将返回值设置为变量,设置好变量的值,就不用按顺序一个个return了

这点倒是比`java`好多了

**注意:除了切片,指针,map,function,Chanel类型外,所有类型调用时传参都是拷贝值过去**

## 6.2 函数值

这是一个新的特性

在go语言中,函数是一种类型,可以赋值给变量

```
func main() {
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS")) // "WNT"
	fmt.Println(strings.Map(add1, "Admix")) // "Benjy"
}

func add1(r rune) rune {
	return r + 1
}
```

如上,就是把函数值作为变量,传递给`strings.Map`函数

作为函数值类型的函数参数,它需要这么定义和使用

```
func Map(mapping func(rune) rune, s string)  {
		r := mapping(xxxx)
}
```

> 可以用于模板设计模式

函数值的几个特性

1. 函数类型的零值默认是nil

2. 函数值可以和nil比较,但是函数值之间不能比较

3. 函数值内部会保存函数上一次执行后的局部变量,不过要想使用它不太容易

   因为所谓局部变量,都是在内部先声明后,再使用的,于是每次执行后,局部变量都会被初始化,这样怎么看到效果?

   接下来当然要举个反例啦

   ```
   func squares() func()  {
   	var x int
   	x++
   	return func()   // 下面的f变量,指的就是这个{
   		x++
   		fmt.Println(x)
   	}
   }
   func main() {
   	//接受一个匿名函数的函数值,这个匿名函数值就是squares函数执行后返回的
   	f := squares()
   	f()
   	f()
   	f()
   	f()
   }
   ```

   > go里面用闭包来实现函数值，go程序员也把函数值称为闭包

   

## 6.3 匿名函数

就是没有名字的函数,一般用在接受参数是函数值的情况,你不想`定义一个函数,再把函数值作为参数`

就可以直接把函数写在实参上

上面我们提到了一个形参是函数值的函数

```
func Map(mapping func(rune) rune, s string)  {
		r := mapping(xxxx)
}
```

现在我们要调用他,可以酱紫

```go
strings.Map(
    func(r rune) rune {
		return r + 1 
	}, 
"HAL-9000")
```

像不像java的匿名内部类?

不过java的写法复杂多了

## 6.4  可变参数

意义:省略

语法和调用方式

```
func main() {
	sum(1,2,3,4,5,6)
	//以下语法很神奇哦
	var arr=[]int{1,2,3,4,5}
	sum(arr...)
}
func sum(arr ...int){
	var sum int
	for _,v:=range arr  {
		sum+=v
	}
	fmt.Println(sum)
}
```

## 6.5 方法

方法就是一类变态版的函数

这种函数,再也不是无所依靠了,它可以通过某个类型的实例用`.`来访问

类似于java中对象的成员方法

java成员方法可以访问到成员对象,用`this`,大家都知道吧

go的函数也有这种访问方式

不瞎比比了,亮代码

```
type User struct{
	ID int
	UserName string
	Password string
}
// user又被称为接受者,推荐使用这种指针方法
func (user *User) SetId(id int){
	user.ID = id
}

func (user User) SetId(id int){
	user.ID = id
}

func main() {
	var user User
	//本来实际上真*写法是  (&user).SetId(3)
	//下面的写法实际隐式了这个&，这是一种语法糖
	user.SetId(3)
	fmt.Println(user)
}
```

另外一提,方法的接收者,不仅仅可以是结构体类型,也可以是切片,字符串等类型,只要不是interface和指针类型

所以以下代码是无法通过编译的

```
type P *int
func (P) f() {
	/* ... */ 
}
```

> 推荐使用上面说的指针方法
>
> 因为接受者说白了,就是隐式的多加一个形参,形参传值默认是拷贝值的,所以用指针可以只拷贝指针,避免整个值的拷贝

> 一般正式程序内,只要有一个指针方法存在,同类型的其他方法也需要是指针方法,当然,go语言不会做限制,只不过约定成俗是酱紫的





# 七:特性

## 6.1 模板输出

我们用的最多的就是字符串的模板输出

```
fmt.Sprintf("您要删除的记录[%s]不存在", id)
```

但很明显,不够强大

如果我要更复杂的输出怎么办

go语言提供了一个将变量值填充到html和文本的这么一个便捷方法

怎么使用?

```go
const templ = `{{.Username}} friend:
{{range .Friend}}----------------------------------------
UaseName: {{.Username}}
ID: {{.ID}}
Password: {{.Password}}
Age:{{.Age}}
NextAge:{{.Age|nextAge}}
{{end}}`

var report = template.Must(template.New("User").
	Funcs(template.FuncMap{"nextAge": nextAge}).
	Parse(templ))
func main() {
	var user  = demo.User{
		ID:12,
		Username:"张三",
		Password:"李四",
		Age:123,
	}
	var friend = user
	user.Friend[0] = &friend
	if err := report.Execute(os.Stdout, user); err != nil {
		log.Fatal(err)
	}
}
func nextAge(age int32)int32{
	return age+1
}
```

就是这样子

用的地方不多,自行了解吧.

此程序最后会基于模板生成字符串



# 八：接口

go语言的接口和java有点相似

目标都是使结构体/对象的方法实现某个抽象方法，以便在其他方法传参时能够传接口类型，实现动态传参

所以关键点有两个

1. 必有一个接口，定义所有抽象方法

   ```
   type Write interface {
   	WriteByte(b byte)
   	Fresh()
   }
   ```

2. 必有至少一个结构体，里面的方法全都实现了抽象方法

   ```
   type FileWrite struct{
   }
   func (write *FileWrite)WriteByte(b byte){
   	//implement super method
   }
   func (write *FileWrite)Fresh(){
   	//implement super method
   }
   ```

有什么好处

1. 我现在可以再定义一个方法，形参是接口类型，实际调用时我可以传多个结构体类型的实参，只要它的方法实现了接口所有的抽象方法。

   ````
   func demo(w mode.Write)  {
   	w.WriteByte(1)
   }
   ````

   调用方式

   ```
   var f  mode.FileWrite
   demo(&f)
   ```

   顺带一提，这个有点重要，你应该有所疑问，为什么demo方法调用时需要传FileWrite类型的地址呢？

   **这是因为，我们只是为FileWrite的指针类型实现了Write接口，但是FileWrite类型并没有实现Write接口接口，所以，你以为FileWrite是Write子类，实际上，FileWrite指针类型才是Write子类**

   > 子类的说法只是方便解释而已，go语言没有子类的概念，只是相似而已

2. 父类参数可以接受子类对象

   如果把Write认为父类，FileWrite认为子类，那么可以

   ```
   var w mode.Write
   var f  mode.FileWrite
   w =&f
   ```

   还是必须用地址赋值

   大概就是实现了多态的一点小玩意

## 8.1 接口内嵌

一个接口可以包含另一个接口的方法，称为接口内嵌

请类比java的接口继承

过于简单，不解释

## 8.2 接口值

什么叫接口值，假如说我现在声明了一个接口类型

```
//io.Writer  是接口类型，自己去看源码
var w io.Writer
```

这里的`w`就是接口值

接口值有两部分组成，一个是具体的类型（动态类型），一个是类型的值（动态值）。

然而严格来说，go语言并没有类型值，所谓具体类型，准确的讲，叫类型的描述符，里面包含类型的名称和方法等等。

现在让我们为这个接口值赋值

```
w = os.Stdout
```

这里面有隐藏操作，具体的讲，就是

1. `os.Stdout`指针的类型描述符被设置给接口值的动态类型，

2. 接口值里面的动态值拷贝了`os.Stdout`

几个特性

1. 接口值可比较，`==`的条件是两个接口值都是nil或者说动态类型都一样，动态值都相等

   > 如果动态类型本身不可以比较，比如说切片类型，那么拿对应的两个接口值进行比较就是报错
   >
   > 所以接口值的比较，还是谨慎点





## 8.x interfact{} 类型

类似于java的Object类型，interfact{} 类型是可以代表所有类型的，但是这个类型可是没有任何方法的

作用

1. 作为函数形参，可以接受所有数据类型的参数
2. 作为变量，可以接受所有类型的参数

# 九：反射

跟java语言一样。go语言的反射可以做到

1. 给一个interface{}类型，获知它是什么类型

   ```
   t := reflect.TypeOf(3) // a reflect.Type
   fmt.Println(t) // "int"
   ```

2. 给一个interface{}类型，获知它是数据类型种类

   ```
   v := reflect.ValueOf(3)
   kind:=v.Kind()
   ```

   虽然经过结构体和`type`关键字，go数据类型可以无限扩充，但是对于数据类型种类，却很明确的知道只有这几种

   1. Bool
   2. String 
   3.  所有數字類型的基礎類型; 
   4. Array 和 Struct 對應的聚合類型;
   5.  Chan,
   6. Func,
   7.  Ptr, 
   8. Slice, 
   9.  Map ; 
   10. 接口類型; 
   11. 表示空值的無效類型

3. 获取结构体类型的字段信息

   ```
   var u User
   	ref:=reflect.TypeOf(u)
   	fieldCount:=ref.NumField()
   	for i:=0;i<fieldCount ;i++  {
   		field:=ref.Field(i)
   		fmt.Println("字段名："+field.Name)
   		//field还可以获取字段的tag，相当于字段的注解了
   	}
   ```

4. 获取方法的就不说了。。

> 当然啦，go语言的反射和java一样，执行速度慢~~~~~



# 十一：冷知识

1. switch语法中，`{` 和前面的表达式必须在前一行

2. switch语法中，如果switch不带表达式，默认是true

   所以这个代码。。。

   ```
   func alwaysFalse() bool {
     return false
   }
   
   func main() {
     switch alwaysFalse()
     {
     case true:
       println("真")
     case false:
       println("假")
     }
   }
   ```

   执行结果是true。

   这件事告诉我们，左大括号换行的家伙都是异端

   另外，别琢磨这个代码了，没人会写这种坑爹的代码的
   
3. 比较好用，如果map的value类型不重要，比如说做查表法时，建议map的value用结构体类型

   结构体类型做value类型可以节省内存空间

   想判断key是否存在，可以用`_, has := mapVariable[key]`

   拿到has则可以判断
   
4. golang的字符串的读写不是线程安全的，所以

   假设有两个协程

   一个协程读

   一个协程写

   读的那个协程就有可能读到一些乱七八糟的数据

   1. 包括，字符串后面突然出现GOGC?

   2. 包括，原本只能写apple和plane，但是读出了appl这样的数据，数据被截断了

      究其原因，是因为golang字符串底层的数据结构是一个结构体

      包含 

      一个指向原始字符串的指针

      一个保存字符串长度的变量

      而字符串赋值，其实就是要给这两个变量赋值

      就有可能要么指针赋值了，但是长度还没有

      于是当我们读取的时候，按旧的长度读取就会表示为截断（假设旧的长度比较短）

      当然有可能旧的长度比较长，但是长度还在允许的范围内，对于多出的那些长度，会填充GOGC以及乱七八糟的字符进去

      如果空串赋值的话，指针会指向空

      所以打印时还可能会报空指针异常

# 十二:常见方法和调用方式

//TODO 

# 十三：常见工具和使用

## 13. 1 XORM根据表生成结构体

XORM是某人写的一个开源的，go语言专属的ORM框架，既然是ORM框架了，就涉及到数据库表和Model的转化，这里的Model，在go语言叫做结构体

好了，问题来了，怎么使用XORM提供的工具从数据表生成结构体？

1. 使用go get命令下载XORM，并让其自动生成XORM工具的可执行文件，生成的文件大致是`xorm.exe`

   >`go get github.com/go-xorm/cmd/xorm`

2. 将下载下来的源代码中，`templates`文件夹复制到`xorm.exe`文件夹下面

3. 执行命令

   ```
   xorm reverse mysql userName:password@(ip:port)/dataBase?charset=utf8 templates\goxorm savePath tablename
   ```

   解释

   tablename 是表名，也可以使用通配符来表示多张表



## 13.2 easyjson生成结构体json帮助文件

命令很简单

`easyjson -all goPath`

执行完毕后会在goPath上面生成一个xxxxxeasyjson.go的文件

> -all 得加这个参数才行，否则会生成一个辣鸡文件

# 十四 企业开发规范

## 14.1 接受者必要性

在企业开发中，一般的写函数的go文件，都建议加一个接受者，接受者在本go文件内定义，可以不带任何属性

这样做的好处

1. go程序只有包的概念，没有类的概念，为了区分包里面的不同go文件，需要用接受者来区分。

   > 只可惜变量名不行，在做常量时，只能统一写在一个包内了

2. 可读性好，当然前提是接受者命名要恰当

3. 作为隐藏的第三个参数用



# 十五:go的测试

1. 单元测试

你可以为go的一个函数编写单元测试，单元测试必须要用一个新文件，文件名必须以`_test.go`结尾

然后里面要写一个测试方法

例如

```
func Test_batchDetail_ProcessLack(t *testing.T) {
}
```

注意：方法名必须以Test开头



写完之后，来试试运行这个测试方法

2. 性能测试





# 十六:go的正则表达式

主要用的包`regexp`

想知道怎么使用，以及有什么功能，可以在该项目的源码中找到

# 十七: Garbage collection for GO
参考: [here](https://tip.golang.org/doc/gc-guide)
## 应用场景
golang的垃圾收集器策略和java的类似。  
一般只是单纯的值分配的话，内存会放在栈上。由于编译器能预先确定这些内存什么时候能被释放。所以它其实比GC那种要高效的多。  
但是反过来，如果编译器不能确定这些值什么时候释放，那么这些内存的分配就会逃逸到推上。这个时候，就是由垃圾收集器处理的时候了。  
关于哪些值会逃逸，哪些不会，取决于这些值工作的代码上下文，以及golang自己的逃逸算法。
see here for more information: [here](https://tip.golang.org/doc/gc-guide#Eliminating_heap_allocations)  
## 关键字解释
**object** 指内存，里面存储一个或者多个golang的值  

**Pointer**： 一个地址，指向一个object里面的任何值。值可以是自然值，也可以是golang内置的类型，比如说切片，字符串，map等  

**根指针**：根指针是指程序中一定使用到的指针。比如说全局变量或者本地变量    

**扫描**：从根指针（在程序中一定会使用的指针）开始，扫描那些还在使用中的对象并一一标记。 

> 这个扫描概念对所有跟踪式计数的GC都适用，不同之处在于GC发现那些存活的内存后的做法。
其中一个做法是迁移，把一些存活的内存迁移到另一个内存区块，然后更新这些内存的指针。
然而GC for Golang并不是这种迁移式的垃圾收集器

**边际成本：**
经济学术语Marginal cost. 指增加多一个产物，随之增加的额外成本

**稳态：**
稳态是一种服务器状态，在这个状态下
1. 分配新内存的速率是固定的，这个速率跟新内存是否live无关。也并不意味着如果新内存是活跃的，则活跃的堆内存会变大。
2. 服务器的对象数量每次看都差不多，指针数量也差不多。
假设工作负载是固定的，那么一个golang程序经过一段时间的运行后，应该会趋近于稳态。
而典型的golang应用程序工作负载应该变化不大。
所以说稳态是典型的golang应用程序应该达到的一种状态

## GC的类型
GC有几种类型，比如说引用计数式。但是golang是一个跟踪式计数的GC.
它的工作流程是
1. 扫描：
2. 大清除
3. 关闭
流程从1，2，3步反复的循环  
进一步的思考：
1. 其中扫描和大清除看似可以同时完成，其实不行。一遍扫描一遍清除可能会错误的清掉还在活跃的内存（仅仅只是因为还没扫描到那个地址而已）
2. 当没有什么工作需要它的时候，有可能GC根本就不工作

## GC 的代价
GC只会在两个地方花费代价
1. CPU占用
2. 内存占用

### 内存占用
GC的内存开销主要由三种内存块决定
1. 确认存活的内存块，其大小由前一次GC决定
2. 新创建的内存块，其大小不固定，在当前的GC周期内创建
3. 元数据，虽然大小和前一次的GC占用成正比，但是相对来说还是比较小的
### CPU占用
GC的CPU占用由固定成本和与堆大小成比例的边际成本构成
> 

值得注意的是，在GC的三个阶段中，清除过程要比扫描和标记的过程更耗费资源。
它的代价随着堆的变大而变大。理由很简单，它需要找到那些不再活动的内存块并释放。
但是在当今的GC for Golang 实现中，它实际要比扫描和标记要快得多的。

### 相互作用
现在，GC的内存占用和CPU占用之前是如何互相影响的？

#### GC的频率
GC的频率可以影响CPU和内存的占用。
如果频率高：更少的内存占用，但是CPU的波动会更经常
如果频率低：更多的内存占用，但是总的来说，CPU的成本会变低

#### GOGC
这决定了CPU占用和内存占用的平衡。它通过几个因素起作用
1. Target heap memory : 限定的一个堆内存大小，GC的目标是在当前堆大小达到目标堆大小之前，就完成一次周期。（但是这并不是一个强制约束，有其他情况会影响到GC 的行为）
3. Total heap size：当前堆的大小（大小由前一次GC决定）+新分配的堆内存大小

可以用一个公式来表达它们两者之间的关系
`Target heap memory = Live heap + (Live heap + GC roots) * GOGC / 100`

从这个公式可以看出。
如果GOGC的值设定得越小，目标堆就越小，从而GC发生的频率就会上升。（CPU成本升高，内存开销变小）
如果GOGC的值设定的越大，目标堆就越大，从而GC 发生的频率就会下降（CPU成本降低，内存开销变大）
当然可以把GOCO关闭，这样意味着GC FOR GOLANG 永远不会运行了
历史沿革  
GOGC是在1.18之后才包含了GC root的，在此之前它只包含live heap.

栈的影响：
一般来说，栈所占的内存都太小，堆内存还是在GC处理中占主导地位。但是如果应用程序含有非常多的Goroutines. 那么GC 可能会做出糟糕的判断。

如何设置GOGC.
很简单，可以在环境变量里面设置，也可以在程序内调用`SetGCPercent(runtime package)`

#### memory limit  
尽管GOGC已经为CPU和内存的平衡做出了不可磨灭的奉献。  
但是一旦遇到内存使用peak的时候，还是过于僵硬。  
考虑到如果实际机器的内存十分的小，而内存使用peak如果一旦超过了这个值，将直接带来内存溢出的致命错误。因此.  
Golang在1.19推出了一个全新的参数.  
memory limit, 它用于直接限制应用程序能够使用的内存。  
而GoGC会基于memory limit来调整total heap size。从而尽量避免内存的使用超过memory limit.  
具体的算法就不介绍了。

缺点：
如果memory limit设置的比较少（基本上小于peak memory），则go gc会频繁的启动来使内存尽量不超过这个limit 

实际上，它的设置是确定了一个GC启动频率的最小值，通过设置它，尽量维持设定好的memory limit

尽管如果，如果我们的参数得当的话
1. 程序安全的在限定的memory limit内运行
2. 更有效率的利用资源

特别的一点是：即使我们把GOGC设置为关闭，Memory Limit也可以继续生效。
其中理由就不好说了，可能这种情况下GOGC还会运行，或许有其他机制去保证内存的限制。


#### 最大GC所占用的CPU时间
GO并不保证memory limit一定生效，尤其是如果你把memory limit设置的非常小。
这样是为了避免程序卡死，因为足够小的限制迫使GC运行时十分频繁，导致CPU 被占满。
因为达到这个目的，Go内部维护了一个变量，为GO GC占用CPU时间的最大值。
GOGC会尽量满足这个时间，但是带来的后果GO GC会延迟执行，同时memory limit会被超过。

注：当前这个最大值被设置约为50%。
为什么呢？这个值的设定情况是，假设应用程序在一个内存充足的机器内运行，但是memory limit被错误地设定为特别低，那么应用程序最差的情况，也就是慢两倍。不至于全卡死


### 浅谈调优
1. 当程序运行的执行环境完全由你控制的时候，自由的去设定memory limit吧



























# X:末尾 BUG合集

1. 对某个资源执行go get -u xxxx 时报

   `..\xxxx\lint.go:23:2: module xxxx found (v0.0.0-20200823
   205832-c024452afbcd), but does not contain package xxxx`

   言下之意自然是
   
   我们找到了xxxx 模块，但是在这个模块下，没有找到xxxx
   
   这张情况，可能包被其他人删了，你需要找到替换包
   
2. go mod tidy 报

   `module declares its path as: xxxx	but was required as: @#@￥#%￥`

   这个原因是依赖包下载完后，发现依赖包go.mod的module声明为xxxx

   但是项目里面导入的包路径是`@#@￥#%￥` 

   两者不匹配，解决办法是执行

   `go mod edit -replace @#@￥#%￥=xxxx@latest`

   将地址重新指向

3. idea莫名报

   ```
   go list -m: can't compute 'all' using the vendor directory
   	(Use -mod=mod or -mod=readonly to bypass.)
   ```

   