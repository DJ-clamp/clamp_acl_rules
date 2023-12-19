package main

import "fmt"

// 代理表接口
type NS interface {
	// HasRegexp() bool   // 查看正则匹配是否存在
	GeneratGroup() string //更加节点组信息生成拼接后的字符串
}

// 节点正则 表示用正则来表示的节点
type NodeRegexp string

//生成字符串 正则不需要带前缀 本身已经包含方法
func (n NodeRegexp) GeneratGroup() string {
	return fmt.Sprintf("%s", string(n))
}

// 节点名称 表示用字符串来表示的节点,通常情况下一个节点名称下回产生一组代理，所以在生成字符串时候需要添加[]关键字
type NodeName string

//生成字符串 包含有代理组所以需要带上前缀`[]`
func (n NodeName) GeneratGroup() string {
	return fmt.Sprintf("[]%s", string(n))
}

// 节点结构体
type Node[T any] struct {
	PROXY_GROUP T
}

// var PROXY_LOCALAREANETWORK  = NN(NS_LOCALAREANETWORK)

type NN interface {
	~string
	NS
}

const (
	PROXY_TYPE_SELECT       string     = "select"                                                                 // 手动选择
	PROXY_TYPE_URL_TEST     string     = "url-test"                                                               // 自动选择
	PROXY_TYPE_FALLBACK     string     = "fallback"                                                               // 自动退回
	PROXY_TYPE_LOAD_BALANCE string     = "load-balance"                                                           // 负载均衡
	PROXY_FILTER_HK         NodeRegexp = "(港|HK|hk|Hong Kong|HongKong|hongkong)"                                  // 正则HK
	PROXY_FILTER_JP         NodeRegexp = "(日本|川日|东京|大阪|泉日|埼玉|沪日|深日|[^-]日|JP|Japan)"                               // 正则JP
	PROXY_FILTER_US         NodeRegexp = "(美|波特兰|达拉斯|俄勒冈|凤凰城|费利蒙|硅谷|拉斯维加斯|洛杉矶|圣何塞|圣克拉拉|西雅图|芝加哥|US|United States)" // 正则US
	PROXY_FILTER_SG         NodeRegexp = "(新加坡|坡|狮城|SG|Singapore)"                                                // 正则sG
	PROXY_FILTER_TW         NodeRegexp = "(台|新北|彰化|TW|Taiwan)"                                                    // 正则TW
	PROXY_FILTER_KR         NodeRegexp = "(KR|Korea|KOR|首尔|韩|韓)"                                                  // 正则KR
	PROXY_FILTER_EU         NodeRegexp = "(英|法|德|阿姆斯特丹|荷兰|土耳其|比利时|瑞士)"                                            // 正则EU
	PROXY_FILTER_STREAMING  NodeRegexp = "(NF|奈飞|解锁|Netflix|NETFLIX|Media)"                                       // 正则Streaming
	PROXY_NETEASEMUSIC      NodeRegexp = "(网易|音乐|解锁|Music|NetEase)"                                               // 正则NETEASEMUSIC
	PROXY_FILTER_ALL        NodeRegexp = ".*"                                                                     // 正则ALL
)

var (
	//Node_name_sheet
	NS_LOCALAREANETWORK NodeName = "🎯 全球直连"
	NS_BANAD            NodeName = "🛑 广告拦截"
	NS_BANPROGRAM_AD    NodeName = "🍃 应用净化"
	NS_GOOGLEFCM        NodeName = "📢 谷歌FCM"
	NS_BING             NodeName = "Ⓜ️ 微软Bing"
	NS_ONEDRIVE         NodeName = "Ⓜ️ 微软云盘"
	NS_MICROSOFT        NodeName = "Ⓜ️ 微软服务"
	NS_APPLE            NodeName = "🍎 苹果服务"
	NS_TELEGRAM         NodeName = "📲 电报消息"
	NS_OPENAI           NodeName = "💬 OpenAi"
	NS_NETEASEMUSIC     NodeName = "🎶 网易音乐"
	NS_GAME_STORE       NodeName = "🎮 游戏平台"
	NS_YOUTUBE          NodeName = "📹 油管视频"
	NS_NETFLIX          NodeName = "🎥 奈飞视频"
	NS_BAHAMUT          NodeName = "📺 巴哈姆特"
	NS_BILIBILI         NodeName = "📺 哔哩哔哩"
	NS_CHINAMEDIA       NodeName = "🌏 国内媒体"
	NS_PROXYMEDIA       NodeName = "🌍 国外媒体"
	NS_FINAL            NodeName = "🐟 漏网之鱼"
	NS_HK               NodeName = "🇭🇰 香港节点"
	NS_JP               NodeName = "🇯🇵 日本节点"
	NS_US               NodeName = "🇺🇲 美国节点"
	NS_SG               NodeName = "🇸🇬 狮城节点"
	NS_TW               NodeName = "🇨🇳 台湾节点"
	NS_KR               NodeName = "🇰🇷 韩国节点"
	NS_EU               NodeName = "🌍 欧洲及中东地区"
	NS_STREAMING        NodeName = "🎥 流媒体节点"
	NS_PROXYGFWLIST     NodeName = "🚀 节点选择"
	NS_PROXY            NodeName = "🚀 手动切换"
	NS_DIRECT           NodeName = "DIRECT"
	NS_REJECT           NodeName = "REJECT"
)
