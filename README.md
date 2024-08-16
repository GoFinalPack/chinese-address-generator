## chinese-address-generator
中国地址生成器 - 三级地址 四级地址 随机生成完整地址

### 数据
数据在data文件夹中, 可以自己实现想要的相关逻辑

### 安装依赖

```shell
go get github.com/GoFinalPack/chinese-address-generator@v1.0.0
```

### 生成地址

```go
    g := chineseaddressgenerator.Generator{}
    g.Init()
    level1 := g.GenerateLevel1()  // 一级地址
    fmt.Println(level1)          // {"code": "230000", "region": "黑龙江省"}
    level2 := g.GenerateLevel2() // 二级地址
    fmt.Println(level2)          // {"code": "620100", "region": "甘肃省兰州市"}
	level3 := g.GenerateLevel3() // 三级地址
    fmt.Println(level3)          // {"code": "350205", "region": "福建省厦门市海沧区"}
    level4 := g.GenerateLevel4() // 四级地址 
    fmt.Println(level4)          // {"code": "310113111000", "region": "上海市市辖区宝山区高境镇"}
	fullAddress := g.FabricateFullAddress()  // 生成完整地址
    fmt.Println(fullAddress)     // {"code": "622926209000", "region": "甘肃省临夏回族自治州东乡族自治县五家乡1115号182室", "buildNo": 1115, "roomNo": 182}
```
P.S.: 生成规则:(001-1400)号(101-909)室

### 原项目地址

```
https://github.com/NiZerin/chinese-address-generator/tree/main   PHP

https://github.com/moonrailgun/chinese-address-generator     Node

```
### 关于贡献

基于MIT开源协议


