package chinese_address_generator

import (
	"fmt"
	chineseaddressgenerator "github.com/GoFinalPack/chinese-address-generator"
	"testing"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/8/15
 * @Desc:
 * @Project: chinese-address-generator
 */

func TestGenerateLevel1(t *testing.T) {
	g := chineseaddressgenerator.Generator{}
	g.Init()
	level1 := g.GenerateLevel1()
	fmt.Println(level1)
}

func TestGenerateLevel2(t *testing.T) {
	g := chineseaddressgenerator.Generator{}
	g.Init()
	level2 := g.GenerateLevel2()
	fmt.Println(level2)
}

func TestGenerateLevel3(t *testing.T) {
	g := chineseaddressgenerator.Generator{}
	g.Init()
	level3 := g.GenerateLevel3()
	fmt.Println(level3)
}

func TestGenerateLevel4(t *testing.T) {
	g := chineseaddressgenerator.Generator{}
	g.Init()
	level4 := g.GenerateLevel4()
	fmt.Println(level4)
}

func TestFabricateFullAddress(t *testing.T) {
	g := chineseaddressgenerator.Generator{}
	g.Init()
	level1 := g.GenerateLevel1()
	fmt.Println(level1)
	level2 := g.GenerateLevel2()
	fmt.Println(level2)
	level3 := g.GenerateLevel3()
	fmt.Println(level3)
	level4 := g.GenerateLevel4()
	fmt.Println(level4)
	fullAddress := g.FabricateFullAddress()
	fmt.Println(fullAddress)
}
