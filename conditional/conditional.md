
**golang 条件判断**

- if来申明
- if 条件判断中不需要加()

```
    if 布尔表达式{
        当表达式为true时执行
    }
```

**单一if判断**


```
    var a int = 10
	if a < 100 {
		fmt.Printf("a 小于 100\n")
	}
	fmt.Printf("a = %d\n",a)
```

**if else判断**

```
    var a int = 10
    if a < 5 {
    	fmt.Printf("a 小于 5\n")
    } else{
    	fmt.Printf("a 大于 5 \n")
    }
```

**if else if  else判断**
```
    a := 6
	if a<10 && a>5 {
		fmt.Println("5 > Value < 10 ")
	} else if a<5 {
		fmt.Println("Value < 5")
	} else {
		fmt.Println("value Error")
	}
```

**golang switch**

- switch 语句用于基于不同条件执行不同动作
- 每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止

**语法**

```
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

**判断选择**
```
    var d_type = "mem"
    switch  {
    	case d_type == "disk":
    		fmt.Println("Select disk")
    	case d_type == "mem":
    		fmt.Println("Select mem")
    	case d_type == "cpu":
    		fmt.Println("Select Cpu")
    	default:
    		fmt.Println("Success!")
    }

```

**计算器**

```
    func switchCalculator(a,b int,op string) int {
    	var result int

    	switch op {
    	case "+":
    		result = a + b
    	case "-":
    		result = a - b
    	case "*":
    		result = a * b
    	case "/":
    		result = a / b
    	default:
    		panic("unsupported operator:" + op)

    	}
    	return result
    }
```
