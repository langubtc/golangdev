**golang循环**

- for来定义一个循环函数
- 可以执行指定次数的循环
- 有三种表现形式

**语法**
```
    for init; condition; post { }
```

```
    for condition { }
```

```
    for { }
```


**标准循环**

```
    	for a :=0 ; a < 20; a++ {
    		fmt.Printf("当前循环值为%d\n",a)
    	}
```

**循环变量**

```
    var a int = 10
    	var b int = 20
    	for a < b {
    		a++
    		fmt.Printf("a 的值为%d\n",a)
    	}
```

**循环一个列表**

```
    numberList  := [6]int {1,2,3,4,5,6}
    	for i,x:= range numberList {
    		fmt.Printf("当前值为%d,索引为%d\n",x,i)
    	}
```