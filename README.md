#cashier-TW-homework

###Installation
确保你安装了golang后执行以下安装命令
```bash
go get github.com/jas0ns/cashier-TW-homework
go install github.com/jas0ns/cashier-TW-homework
```

###Usage
运行题目指定测试用例，请cd到源码目录后
(默认为`$GOPATH/src/github.com/jas0ns/cashier-TW-homework/`)运行
```bash
go test
```

请参考单元测试代码`cashier_test.go`中的`func Example1~5()`使用该模块

###Notice(or Bug)
当为一个商品添加多种优惠的时候，有冲突的优惠需要按优先顺序依次添加<br>
例: 需要先添加"买二赠一"，再添加"95折"，才能正确的使其只享受优先级高的优惠"买二赠一"


###BTW
初学golang，望面试官手下留情
