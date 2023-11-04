package goroutines

type routines struct{}

var Routines routines

// 如果挂起的协程过多，那么它也会造成内存泄漏的现象。尽管它并不是内存泄漏。
//但是在它挂起的那段时间内，它所占用的资源一直没法回收
//这里不会有代码示例说明
