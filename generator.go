package chinese_address_generator

import (
	"chinese-address-generator/utils"
	"math/rand"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/8/15
 * @Desc:
 * @Project: chinese-address-generator
 */

type Generator struct {
	Level3 []utils.Level3Data
	Level4 []utils.Level4Data
}

func (g *Generator) Init() {
	// 判断 Level3 是否为空
	if len(g.Level3) == 0 {
		// 读取 data 目录下的 level3.json
		utils.ReadLevel3()
		g.Level3 = utils.Level3
	}

	// 判断 Level4 是否为空
	if len(g.Level4) == 0 {
		// 读取 data 目录下的 level4.txt
		utils.ReadLevel4()
		g.Level4 = utils.Level4
	}
}

func (g *Generator) getRandomProvince() utils.Level3Data {
	// 获取随机省份
	if g.Level3 == nil || len(g.Level3) == 0 {
		// 读取 data 目录下的 level3.json
		utils.ReadLevel3()
		g.Level3 = utils.Level3
	}
	// 从 Level3 中随机获取一个省份
	return g.Level3[rand.Intn(len(g.Level3))]
}

func (g *Generator) GenerateLevel1() string {
	province := g.getRandomProvince()
	return province.ToJSON()
}

func (g *Generator) GenerateLevel2() string {
	// 获取有县级数据的省份
	province := utils.GetProvinceWithChildren()
	// 从 province 中随机获取一个县级数据
	level2 := province[rand.Intn(len(province))]

	return level2.RegionToJSON()
}

func (g *Generator) GenerateLevel3() string {
	// 获取有县级数据的省份
	province := utils.GetProvinceWithCityChildren()
	// 从 province 中随机获取一个县级数据
	level3 := province[rand.Intn(len(province))]

	return level3.RegionCityToJSON()
}

func (g *Generator) GenerateLevel4() string {
	province := utils.GetProvinceWithTownship()
	level4 := province[rand.Intn(len(province))]
	return level4.RegionTownShipToJSON()
}

func (g *Generator) FabricateFullAddress() string {
	province := utils.GetProvinceWithTownship()
	level4 := province[rand.Intn(len(province))]
	return level4.RegionFullAddressToJSON()
}
