package main

import (
	"fmt"
	"strings"
)

// 规则片段结构
type Rule struct {
	Proxy_name NodeName // 代理规则名称
	Url        string   // 片段导入地址
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

// 规则表 数组形式
var rule_lists = []Rule{
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/LocalAreaNetwork.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/UnBan.list"},
	{NS_BANAD, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanAD.list"},
	{NS_BANPROGRAM_AD, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/BanProgramAD.list"},
	{NS_GOOGLEFCM, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/GoogleFCM.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/GoogleCN.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/SteamCN.list"},
	{NS_BING, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Bing.list"},
	{NS_ONEDRIVE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/OneDrive.list"},
	{NS_MICROSOFT, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Microsoft.list"},
	{NS_APPLE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Apple.list"},
	{NS_TELEGRAM, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Telegram.list"},
	{NS_OPENAI, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/OpenAi.list"},
	{NS_NETEASEMUSIC, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/NetEaseMusic.list"},
	{NS_GAME_STORE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Epic.list"},
	{NS_GAME_STORE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Origin.list"},
	{NS_GAME_STORE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Sony.list"},
	{NS_GAME_STORE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Steam.list"},
	{NS_GAME_STORE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Nintendo.list"},
	{NS_YOUTUBE, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/YouTube.list"},
	{NS_NETFLIX, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Netflix.list"},
	{NS_BAHAMUT, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Bahamut.list"},
	{NS_BILIBILI, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/BilibiliHMT.list"},
	{NS_BILIBILI, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Ruleset/Bilibili.list"},
	{NS_CHINAMEDIA, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaMedia.list"},
	{NS_PROXYMEDIA, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyMedia.list"},
	{NS_PROXYGFWLIST, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ProxyGFWlist.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaDomain.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/ChinaCompanyIp.list"},
	{NS_LOCALAREANETWORK, "https://raw.githubusercontent.com/ACL4SSR/ACL4SSR/master/Clash/Download.list"},
	{NS_LOCALAREANETWORK, "[]GEOIP,CN"},
	{NS_FINAL, "[]FINAL"},
}

// 生成ruleset字符串
func GenerateRuleSetContent[T any](nodename T) string {
	// _, ok := nodename.(string)
	return fmt.Sprintf("%s", nodename)
}

func (r Rule) GenerateRule() string {
	return strings.Join([]string{GenerateRuleSetContent(r.Proxy_name), r.Url}, ",")
}

// 从已知规则表中查询规则并导出字符串，如果查询不到则返回空
func (r Rule) GetRule(proxy_name NodeName) string {
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
