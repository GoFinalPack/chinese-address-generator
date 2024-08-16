package chinese_address_generator

import (
	"github.com/GoFinalPack/chinese-address-generator/utils"
	"os"
	"testing"
)

// TestReadLevel4 /**
func TestReadLevel4(t *testing.T) {
	// 创建一个测试文件
	file, err := os.CreateTemp("", "level4")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())

	// 写入测试数据
	_, err = file.WriteString(`{"code":"123","region":"Test Region"}`)
	if err != nil {
		t.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	// 调用 ReadLevel4 函数
	utils.ReadLevel4()
	// 检查 Level4 是否被正确填充
	if len(utils.Level4) < 0 {
		t.Errorf("Expected 1 Level4Data, got %d", len(utils.Level4))
	}
	if utils.Level4[0].Code != "110101001000" || utils.Level4[0].Region != "东华门街道" {
		t.Errorf("Expected Level4Data {{Code: \"123\", Region: \"Test Region\"}}, got %+v", utils.Level4[0])
	}
}

func TestReadLevel3(t *testing.T) {
	// 调用 ReadLevel3 函数
	utils.ReadLevel3()
	// 检查 Level3 是否被正确填充
	if len(utils.Level3) < 0 {
		t.Errorf("Expected 1 Level3Data, got %d", len(utils.Level3))
	}
	if utils.Level3[0].Code != "110000" || utils.Level3[0].Region != "北京市" {
		t.Errorf("Expected Level3Data {{Code: \"110101\", Region: \"东华门街道\"}}, got %+v", utils.Level3[0])
	}
}
