package main

import (
	"fmt"
	"strings"
)

// 设置分组表
type PROXY_GROUP string

// 普通节点组
var (
	PROXIES_REGULAR      = []any{NS_HK, NS_TW, NS_SG, NS_JP, NS_US, NS_KR, NS_EU, NS_PROXY, NS_DIRECT}
	PROXIES_MINI         = []any{NS_DIRECT, NS_HK, NS_TW, NS_SG, NS_JP, NS_PROXY}
	PROXIES_NETFLIX      = []any{NS_NETFLIX, NS_HK, NS_TW, NS_SG, NS_JP, NS_US, NS_KR, NS_EU, NS_PROXY, NS_DIRECT}
	PROXIES_BAHAMUT      = []any{NS_TW, NS_PROXY, NS_PROXYGFWLIST, NS_DIRECT}
	PROXIES_BILIBILI     = []any{NS_LOCALAREANETWORK, NS_TW, NS_HK}
	PROXIES_NETEASEMUSIC = []any{NS_DIRECT, NS_PROXYGFWLIST, PROXY_NETEASEMUSIC}
	PROXIES_DIRECT       = []any{NS_DIRECT, NS_PROXYGFWLIST}
	PROXIES_REJECT       = []any{NS_REJECT, NS_DIRECT}
)

// 生成代理组
func Generate_PROXY_GROUP(node_name NodeName, node_proxy_type string, prxoy_sections ...any) string {
	var node_group string
	if node_name == "" || len(prxoy_sections) <= 0 {
		fmt.Println("error null node or sections")
		return ""
	}
	var arr []string
	for _, section := range prxoy_sections {

		arr = append(arr, Get_Groups(section))
	}
	tmp := strings.Join(arr, "`")
	node_group = fmt.Sprintf("%s`%s`%s", node_name, node_proxy_type, tmp)
	fmt.Println(node_group)
	return node_group
}

func Get_Groups(section any) string {
	var tmp string

	switch section := section.(type) {
	case NodeName:
		tmp = section.GeneratGroup()
	case NodeRegexp:
		tmp = section.GeneratGroup()
	default:
		return ""
	}
	return tmp

}

var proxy_group_lists = []PROXY_GROUP{
	// "🚀 节点选择`select`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_PROXYGFWLIST, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🚀 手动切换`select`.*",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_PROXY, PROXY_TYPE_SELECT, []any{PROXY_FILTER_ALL}...)),
	// "📲 电报消息`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_TELEGRAM, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "💬 OpenAi`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_OPENAI, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "📹 油管视频`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_YOUTUBE, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🎥 奈飞视频`select`[]🎥 奈飞节点`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_NETFLIX, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "📺 巴哈姆特`select`[]🇨🇳 台湾节点`[]🚀 节点选择`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_BAHAMUT, PROXY_TYPE_SELECT, PROXIES_BAHAMUT...)),
	// "📺 哔哩哔哩`select`[]🎯 全球直连`[]🇨🇳 台湾节点`[]🇭🇰 香港节点",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_BILIBILI, PROXY_TYPE_SELECT, PROXIES_BILIBILI...)),
	// "🌍 国外媒体`select`[]🚀 节点选择`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_PROXYMEDIA, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🌏 国内媒体`select`[]DIRECT`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_CHINAMEDIA, PROXY_TYPE_SELECT, PROXIES_MINI...)),
	// "📢 谷歌FCM`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_GOOGLEFCM, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "Ⓜ️ 微软Bing`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_BING, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "Ⓜ️ 微软云盘`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_ONEDRIVE, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "Ⓜ️ 微软服务`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_MICROSOFT, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🍎 苹果服务`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_APPLE, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🎮 游戏平台`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_GAME_STORE, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🎶 网易音乐`select`[]DIRECT`[]🚀 节点选择`(网易|音乐|解锁|Music|NetEase)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_NETEASEMUSIC, PROXY_TYPE_SELECT, PROXIES_NETEASEMUSIC...)),
	// "🎯 全球直连`select`[]DIRECT`[]🚀 节点选择",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_LOCALAREANETWORK, PROXY_TYPE_SELECT, PROXIES_DIRECT...)),
	// "🛑 广告拦截`select`[]REJECT`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_BANAD, PROXY_TYPE_SELECT, PROXIES_REJECT...)),
	// "🍃 应用净化`select`[]REJECT`[]DIRECT",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_BANPROGRAM_AD, PROXY_TYPE_SELECT, PROXIES_REJECT...)),
	// "🐟 漏网之鱼`select`[]🚀 节点选择`[]DIRECT`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_FINAL, PROXY_TYPE_SELECT, PROXIES_REGULAR...)),
	// "🇭🇰 香港节点`select`(港|HK|hk|Hong Kong|HongKong|hongkong)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_HK, PROXY_TYPE_SELECT, []any{PROXY_FILTER_HK}...)),
	// "🇯🇵 日本节点`select`(日本|川日|东京|大阪|泉日|埼玉|沪日|深日|[^-]日|JP|Japan)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_JP, PROXY_TYPE_SELECT, []any{PROXY_FILTER_JP}...)),
	// "🇺🇲 美国节点`select`(美|波特兰|达拉斯|俄勒冈|凤凰城|费利蒙|硅谷|拉斯维加斯|洛杉矶|圣何塞|圣克拉拉|西雅图|芝加哥|US|United States)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_US, PROXY_TYPE_SELECT, []any{PROXY_FILTER_US}...)),
	// "🇸🇬 狮城节点`select`(新加坡|坡|狮城|SG|Singapore)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_SG, PROXY_TYPE_SELECT, []any{PROXY_FILTER_SG}...)),
	// "🇨🇳 台湾节点`select`(台|新北|彰化|TW|Taiwan)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_TW, PROXY_TYPE_SELECT, []any{PROXY_FILTER_TW}...)),
	// "🇰🇷 韩国节点`select`(KR|Korea|KOR|首尔|韩|韓)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_KR, PROXY_TYPE_SELECT, []any{PROXY_FILTER_KR}...)),
	// "🌍 欧洲及中东地区`select`(英|法|德|阿姆斯特丹|荷兰|土耳其|比利时|瑞士)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_EU, PROXY_TYPE_SELECT, []any{PROXY_FILTER_EU}...)),
	// "🎥 奈飞节点`select`(NF|奈飞|解锁|Netflix|NETFLIX|Media)",
	PROXY_GROUP(Generate_PROXY_GROUP(NS_STREAMING, PROXY_TYPE_SELECT, []any{PROXY_FILTER_STREAMING}...)),
}
