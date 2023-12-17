package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Config string

const (
	CUSTOM string = "custom"
)

// 自定义处理规则的生成器
func main() {
	fmt.Println("Generator to ini file for proxy tools")
	cfg := ini.Empty(ini.LoadOptions{
		AllowShadows:              true,
		SpaceBeforeInlineComment:  false,
		PreserveSurroundedQuote:   false,
		UnescapeValueDoubleQuotes: false,
	})
	ini.PrettyFormat = false //不做对齐处理
	cfg.NewSections([]string{(CUSTOM)}...)
	content, err := cfg.GetSection((CUSTOM))
	if err != nil {
		fmt.Println(err)
		return
	}
	content.Comment = "这是一个测试片段,暂时无法使用"
	err = content.ReflectFrom(FetchOption())
	if err != nil {
		fmt.Println(err)
		return
	}
	cc, err := cfg.GetSection("")
	if err != nil {
		fmt.Println(err)
		return
	}
	cc.SetBody("asdasdd")
	fmt.Println(cc.Body())
	if err = cfg.SaveTo("../Clash/dst/demo.ini"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done!")

	// cfg, err := ini.ShadowLoad("../Clash/dst/demo.ini")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// content, err := cfg.GetSection(CUSTOM)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// keys := content.Keys()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(keys[1])
}
