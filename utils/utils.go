package utils

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type RegionEntity struct {
	Code          string         `json:"code"`
	Region        string         `json:"region"`
	RegionEntitys []RegionEntity `json:"regionEntitys"` // 县级市
}

type Level3Data struct {
	Code          string         `json:"code"`
	Region        string         `json:"region"`
	RegionEntitys []RegionEntity `json:"regionEntitys"` // 地级市，有可能没有
}

type Level4Data struct {
	Code   string `json:"code"`
	Region string `json:"region"`
}

var Level3 []Level3Data
var Level4 []Level4Data
var level4Map map[string][]RegionEntity

//go:embed data/level3.json
var level3File embed.FS

//go:embed data/level4.txt
var level4File embed.FS

// ReadLevel3 读取 level3.json 文件并解析为 Level3 列表
func ReadLevel3() {
	dataBytes, err := level3File.ReadFile("data/level3.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = json.Unmarshal(dataBytes, &Level3)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
}

// ReadLevel4 读取 level4.txt 文件并解析为 Level4 列表
func ReadLevel4() {
	data, err := level4File.ReadFile("data/level4.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			fmt.Printf("invalid line format: %s\n", line)
			continue
		}

		// 创建 Level4Data 实例
		data := Level4Data{
			Code:   fields[0],
			Region: fields[1],
		}
		Level4 = append(Level4, data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func PreprocessLevel4() {
	level4Map = make(map[string][]RegionEntity)
	for _, level4 := range Level4 {
		// 使用前缀作为键，存储对应的 RegionEntity
		prefix := level4.Code[:6]
		level4Map[prefix] = append(level4Map[prefix], RegionEntity{
			Code:   level4.Code,
			Region: level4.Region,
		})
	}
}

// ToJSON 返回 Level3Data 的 JSON 字符串
func (l *Level3Data) ToJSON() string {
	return fmt.Sprintf(`{"code": "%s", "region": "%s"}`, l.Code, l.Region)
}

// RegionToJSON 返回随机选取的子 RegionEntity 的 JSON 字符串
func (l *Level3Data) RegionToJSON() string {
	regionEntity := l.GetRandomChildren()
	return fmt.Sprintf(`{"code": "%s", "region": "%s"}`, regionEntity.Code, l.Region+regionEntity.Region)
}

// RegionCityToJSON 返回随机选取的子 RegionEntity 及其子级的 JSON 字符串
func (l *Level3Data) RegionCityToJSON() string {
	regionEntity := l.GetRandomChildren()
	city := regionEntity.GetRandomRegionEntity()
	return fmt.Sprintf(`{"code": "%s", "region": "%s"}`, city.Code, l.Region+regionEntity.Region+city.Region)
}
func (l *Level3Data) RegionTownShipToJSON() string {
	regionEntity, city, township, err := l.GetRandomRegionCityTownship()
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return fmt.Sprintf(`{"code": "%s", "region": "%s"}`, township.Code, l.Region+regionEntity.Region+city.Region+township.Region)
}

func (l *Level3Data) RegionFullAddressToJSON() string {
	regionEntity, city, township, err := l.GetRandomRegionCityTownship()
	// 生成随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成 buildNo 范围: 001-1400
	buildNo := r.Intn(1400) + 1                      // 生成 1-1400 的数字
	formattedBuildNo := fmt.Sprintf("%03d", buildNo) // 格式化为 001-1400

	// 生成 roomNo 范围: 101-909
	roomNo := r.Intn(809) + 101                    // 生成 101-909 的数字
	formattedRoomNo := fmt.Sprintf("%03d", roomNo) // 格式化为 101-909

	if err != nil {
		return ""
	}
	return fmt.Sprintf(`{"code": "%s", "region": "%s", "buildNo": %d, "roomNo": %d}`, township.Code, l.Region+regionEntity.Region+city.Region+township.Region+formattedBuildNo+"号"+formattedRoomNo+"室", buildNo, roomNo)

}

func (l *Level3Data) GetRandomRegionCityTownship() (RegionEntity, RegionEntity, RegionEntity, error) {
	regionEntity := l.GetRandomChildren()
	city := regionEntity.GetRandomRegionEntity()
	township := city.GetRandomRegionEntity()
	return regionEntity, city, township, nil
}

// HasChildren 检测 Level3Data 是否有市级数据
func (l *Level3Data) HasChildren() bool {
	return len(l.RegionEntitys) > 0
}

// GetRandomChildren 随机选取一个子 RegionEntity
func (l *Level3Data) GetRandomChildren() RegionEntity {
	return l.RegionEntitys[rand.Intn(len(l.RegionEntitys))]
}

// GetRandomRegionEntity 随机选取一个子 RegionEntity
func (r *RegionEntity) GetRandomRegionEntity() RegionEntity {
	if len(r.RegionEntitys) == 0 {
		return RegionEntity{}
	}
	return r.RegionEntitys[rand.Intn(len(r.RegionEntitys))]
}

// GetRandomTownship 随机选取一个子街道
func (l *Level3Data) GetRandomTownship() {
	city := l.GetRandomChildren()
	fmt.Println(city)
}

// GetProvinceWithChildren 获取包含市级数据的省份
func GetProvinceWithChildren() []Level3Data {
	var provinces []Level3Data
	for _, province := range Level3 {
		if province.HasChildren() {
			provinces = append(provinces, province)
		}
	}
	return provinces
}

// HasSubRegions 检测 RegionEntity 是否有子区域
func HasSubRegions(regionEntity RegionEntity) bool {
	if len(regionEntity.RegionEntitys) > 0 {
		return true
	}

	// 递归检查子区域
	for _, subRegion := range regionEntity.RegionEntitys {
		if HasSubRegions(subRegion) {
			return true
		}
	}

	return false
}

// GetProvinceWithCityChildren 获取包含县级市数据的省份
// GetProvinceWithCityChildren 获取有市级数据且包含子级区域的省份
func GetProvinceWithCityChildren() []Level3Data {
	var provinces []Level3Data
	for _, province := range Level3 {
		if province.HasChildren() {
			hasCityWithSubRegions := false
			for _, city := range province.RegionEntitys {
				if HasSubRegions(city) {
					hasCityWithSubRegions = true
					break
				}
			}
			if hasCityWithSubRegions {
				provinces = append(provinces, province)
			}
		}
	}
	return provinces
}

// GetProvinceWithTownship 获取有县级市数据的省份，并将 Level4 数据拼接进去
func GetProvinceWithTownship() []Level3Data {
	// 先预处理 level4 数据
	PreprocessLevel4()
	provinces := GetProvinceWithCityChildren()
	var result []Level3Data

	for i := range provinces {
		province := &provinces[i]
		hasLevel4Data := false

		for j := range province.RegionEntitys {
			city := &province.RegionEntitys[j]

			if len(city.RegionEntitys) == 0 {
				// 如果 city 没有子区域，则跳过
				continue
			}

			for k := range city.RegionEntitys {
				area := &city.RegionEntitys[k]
				if regions, found := level4Map[area.Code]; found {
					// 将匹配到的 Level4 数据添加到 area 的 RegionEntitys 列表中
					area.RegionEntitys = append(area.RegionEntitys, regions...)
					hasLevel4Data = true
				}
			}
		}

		// 只添加有 Level4 数据的省份
		if hasLevel4Data {
			result = append(result, *province)
		}
	}

	return result
}
