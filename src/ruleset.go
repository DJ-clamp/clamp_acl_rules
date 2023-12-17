package main

import "strings"

// 规则片段结构
type Rule struct {
	Proxy_name string // 代理规则名称
	Url        string // 片段导入地址
}
type Rules []string // 规则列表
// 规则设置表
type PROXY_OPTION struct {
	Ruleset     Rules         `ini:"ruleset,,allowshadow"`
	PROXY_GROUP []PROXY_GROUP `ini:"custom_proxy_group,,allowshadow"`
	// 自定义内建设置
	ENABLE_RULE_GENERATOR    bool `ini:"enable_rule_generator"`
	OVERWRITE_ORIGINAL_RULES bool `ini:"overwrite_original_rules"`
}

//规则表 数组形式
var rule_lists = []Rule{
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/LocalAreaNetwork.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/UnBan.list"},
	{"🛑 广告拦截", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanAD.list"},
	{"🍃 应用净化", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanProgramAD.list"},
	{"📢 谷歌FCM", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/GoogleFCM.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/GoogleCN.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/SteamCN.list"},
	{"Ⓜ️ 微软Bing", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Bing.list"},
	{"Ⓜ️ 微软云盘", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/OneDrive.list"},
	{"Ⓜ️ 微软服务", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Microsoft.list"},
	{"🍎 苹果服务", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Apple.list"},
	{"📲 电报消息", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Telegram.list"},
	{"💬 OpenAi", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/OpenAi.list"},
	{"🎶 网易音乐", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/NetEaseMusic.list"},
	{"🎮 游戏平台", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Epic.list"},
	{"🎮 游戏平台", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Origin.list"},
	{"🎮 游戏平台", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Sony.list"},
	{"🎮 游戏平台", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Steam.list"},
	{"🎮 游戏平台", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Nintendo.list"},
	{"📹 油管视频", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/YouTube.list"},
	{"🎥 奈飞视频", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Netflix.list"},
	{"📺 巴哈姆特", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Bahamut.list"},
	{"📺 哔哩哔哩", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/BilibiliHMT.list"},
	{"📺 哔哩哔哩", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Bilibili.list"},
	{"🌏 国内媒体", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaMedia.list"},
	{"🌍 国外媒体", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyMedia.list"},
	{"🚀 节点选择", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyGFWlist.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaDomain.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaCompanyIp.list"},
	{"🎯 全球直连", "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Download.list"},
	{"🎯 全球直连", "[]GEOIP,CN"},
	{"🐟 漏网之鱼", "[]FINAL"},
}

// 设置分组表
type PROXY_GROUP string

var proxy_group_lists = []PROXY_GROUP{
	// "🚀 手动切换`select`.*",
	// "📲 电报消息`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	// "💬 OpenAi`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	// "📹 油管视频`select`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	// "🎥 奈飞视频`select`[]🎥 奈飞节点`[]🚀 节点选择`[]🇸🇬 狮城节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	// "📺 巴哈姆特`select`[]🇨🇳 台湾节点`[]🚀 节点选择`[]🚀 手动切换`[]DIRECT",
	// "📺 哔哩哔哩`select`[]🎯 全球直连`[]🇨🇳 台湾节点`[]🇭🇰 香港节点",
	// "🌍 国外媒体`select`[]🚀 节点选择`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换`[]DIRECT",
	// "🌏 国内媒体`select`[]DIRECT`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🚀 手动切换",
	// "📢 谷歌FCM`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "Ⓜ️ 微软Bing`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "Ⓜ️ 微软云盘`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "Ⓜ️ 微软服务`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "🍎 苹果服务`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "🎮 游戏平台`select`[]DIRECT`[]🚀 节点选择`[]🇺🇲 美国节点`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "🎶 网易音乐`select`[]DIRECT`[]🚀 节点选择`(网易|音乐|解锁|Music|NetEase)",
	// "🎯 全球直连`select`[]DIRECT`[]🚀 节点选择",
	// "🛑 广告拦截`select`[]REJECT`[]DIRECT",
	// "🍃 应用净化`select`[]REJECT`[]DIRECT",
	// "🐟 漏网之鱼`select`[]🚀 节点选择`[]DIRECT`[]🇭🇰 香港节点`[]🇨🇳 台湾节点`[]🇸🇬 狮城节点`[]🇯🇵 日本节点`[]🇺🇲 美国节点`[]🇰🇷 韩国节点`[]🚀 手动切换",
	// "🇭🇰 香港节点`select`(港|HK|hk|Hong Kong|HongKong|hongkong)",
	// "🇯🇵 日本节点`select`(日本|川日|东京|大阪|泉日|埼玉|沪日|深日|[^-]日|JP|Japan)",
	// "🇺🇲 美国节点`select`(美|波特兰|达拉斯|俄勒冈|凤凰城|费利蒙|硅谷|拉斯维加斯|洛杉矶|圣何塞|圣克拉拉|西雅图|芝加哥|US|United States)",
	// "🇸🇬 狮城节点`select`(新加坡|坡|狮城|SG|Singapore)",
	// "🇨🇳 台湾节点`select`(台|新北|彰化|TW|Taiwan)",
	// "🇰🇷 韩国节点`select`(KR|Korea|KOR|首尔|韩|韓)",
	// "🌍 欧洲及中东地区`select`(英|法|德|阿姆斯特丹|荷兰|土耳其|比利时|瑞士)",
	// "🎥 奈飞节点`select`(NF|奈飞|解锁|Netflix|NETFLIX|Media)",
}

func (r Rule) GenerateRule() string {
	return strings.Join([]string{r.Proxy_name, r.Url}, ",")
}

// 从已知规则表中查询规则并导出字符串，如果查询不到则返回空
func (r Rule) GetRule(proxy_name string) string {
	if len(rule_lists) <= 0 {
		return ""

	}
	for _, rule_list := range rule_lists {
		if rule_list.Proxy_name == proxy_name {
			return rule_list.GenerateRule()
		}
	}
	return ""
}

// 从已知规则表中查询所有的规则并直接导出数组形式，如果查询不到直接返回空数组
func (r Rule) GetRules() []string {
	if len(rule_lists) <= 0 {
		return []string{}

	}
	var tmp_array []string
	for _, rule_list := range rule_lists {
		tmp_array = append(tmp_array, rule_list.GenerateRule())

	}
	return tmp_array
}
func (r *Rule) SetRule(node string, url string) string {

	return strings.Join([]string{node, url}, ",")
}

func FetchOption() *PROXY_OPTION {
	var r = new(Rule)
	var rules = Rules(r.GetRules())
	var ruleSet = new(PROXY_OPTION)
	ruleSet.Ruleset = rules
	ruleSet.ENABLE_RULE_GENERATOR = true
	ruleSet.OVERWRITE_ORIGINAL_RULES = true
	ruleSet.PROXY_GROUP = proxy_group_lists
	return ruleSet
}
