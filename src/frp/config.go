package frp

import (
	"bytes"
	"fmt"
	"go-frpc/src/utils"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetIniFilePath() string {
	// 获取配置文件路径
	workspace, _ := filepath.Abs("")
	frpcpath := workspace + "/frpc.ini"
	_, err := os.Stat(frpcpath)
	if err != nil {
		return ""
	}
	return frpcpath
}

func FullContent() (string, error) {
	filename := GetIniFilePath()
	println(filename)
	println(filename)
	//读取文本内容到文本框中
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error : %s\n", err)
		return "", err
	}
	return string(bytes), nil
}

func SetContent(content string) {
	filename := GetIniFilePath()
	ioutil.WriteFile(filename, []byte(content), 0777)
}

func GetSection(section string) string {
	filename := GetIniFilePath()
	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	commonSection, sectionErr := cfg.GetSection(section)
	var lines bytes.Buffer
	if sectionErr != nil {
		fmt.Printf("Fail to load section: %v", sectionErr)
	} else {
		for _, s := range commonSection.KeyStrings() {
			lines.WriteString(fmt.Sprintln(s, "=", commonSection.Key(s).String()))
		}
	}

	return lines.String()
}

func GetSections() ([]string, error) {
	filename := GetIniFilePath()
	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return nil, err
	}
	return cfg.SectionStrings(), nil
}

func SaveSection(section string, content string) {
	filename := GetIniFilePath()
	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 删除原数据
	commonSection, _ := cfg.GetSection(section)
	for _, s := range commonSection.KeyStrings() {
		commonSection.DeleteKey(s)
	}
	// 获取新数据设置到section中
	newcfg, err := ini.Load([]byte(content))
	newSection, _ := newcfg.GetSection(ini.DefaultSection)
	for _, s := range newSection.KeyStrings() {
		commonSection.NewKey(s, newSection.Key(s).Value())
	}
	cfg.SaveTo(filename)
	utils.SendNotifiction(
		section +
			"服务模块已更新，请重启frpc使其生效")
}
