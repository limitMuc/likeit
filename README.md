LikeIt
=========

LikeIt 是一个统计Go项目里广受人吐槽的```if err != nil```的工具，这个错误处理方式虽然饱受诟病，但在Go2没出来前，还是好好接受吧。


目录结构
--------
```
├── bin     # 可执行程序
│   ├── likeit
│   └── likeit.exe
├── example     # 用法举例
├── go.mod
├── main.go
├── pkg
│   ├── file.go
│   ├── file_test.go
│   ├── tree.go
│   └── tree_test.go
└── README.md
```



用法
-----

windows
```
likeit.exe -p absolute_path
# e.g.: likeit.exe -p F:\go\src\likeit
```
![](https://img-blog.csdnimg.cn/db6901290ecc43968645dbe30bb53d4b.png)


linux
```
likeit -p absolute_path
# e.g.: likeit -p /workspace/go/src/likeit
```

![](https://img-blog.csdnimg.cn/f9b5406c7fe9408a89629dc09bea8cee.png)
