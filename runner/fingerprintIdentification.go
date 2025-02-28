package runner

import (
	"fmt"
	"strings"
)

type finger struct {
	Name    string
	Type    string
	Keyword []string
}

var (
	urlType     = "urlType"
	bodyType    = "bodyType"
	headerType  = "header"
	faviconType = "faviconType"
	regexType   = "regex"

	// 指纹库规则
	fingers = map[int]*finger{
		10:  {Name: "华天OA", Type: bodyType, Keyword: []string{"htoa/image/comm/BLUE/top_logo.png"}},
		11:  {Name: "华天OA", Type: headerType, Keyword: []string{"x-webobjects-loadaverage"}},
		20:  {Name: "党建系统", Type: bodyType, Keyword: []string{"/dj_web/loginAction!checkLogin.action"}},
		30:  {Name: "蓝凌OA", Type: bodyType, Keyword: []string{"/sys/ui/js/LUI.js"}},
		31:  {Name: "蓝凌OA", Type: bodyType, Keyword: []string{"login_single_horizontal/js/jquery.js"}},
		40:  {Name: "O2OA", Type: bodyType, Keyword: []string{"<title>O2OA</title>"}},
		41:  {Name: "O2OA", Type: bodyType, Keyword: []string{"o2_core/o2.min.js"}},
		50:  {Name: "通达OA", Type: bodyType, Keyword: []string{"tongda.ico"}},
		51:  {Name: "通达OA", Type: bodyType, Keyword: []string{"请联系OA软件开发商寻求解决办法或下载360安全卫士查杀"}},
		52:  {Name: "通达OA", Type: bodyType, Keyword: []string{"/static/templates/20"}},
		60:  {Name: "万户网络-ezOFFICE", Type: bodyType, Keyword: []string{"defaultroot/scripts"}},
		61:  {Name: "万户网络-ezOFFICE", Type: bodyType, Keyword: []string{"btn_get_account_Click"}},
		70:  {Name: "用友OA", Type: urlType, Keyword: []string{"yyoa/index.jsp"}},
		71:  {Name: "用友OA", Type: bodyType, Keyword: []string{"用友U8"}},
		80:  {Name: "泛微OA", Type: bodyType, Keyword: []string{"theme/ecology"}},
		81:  {Name: "泛微OA", Type: bodyType, Keyword: []string{"images/loginmode"}},
		82:  {Name: "泛微OA", Type: bodyType, Keyword: []string{"wui/theme"}},
		83:  {Name: "泛微OA", Type: bodyType, Keyword: []string{"wui/common/page"}},
		84:  {Name: "泛微OA", Type: urlType, Keyword: []string{"wui/index.html"}},
		90:  {Name: "致远OA", Type: bodyType, Keyword: []string{"seeyon/common"}},
		91:  {Name: "致远OA", Type: bodyType, Keyword: []string{"seeyon/main/login"}},
		92:  {Name: "致远OA", Type: urlType, Keyword: []string{"seeyon/index.jsp"}},
		100: {Name: "帆软报表", Type: bodyType, Keyword: []string{"isSupportForgetPwd"}},
		101: {Name: "帆软报表", Type: bodyType, Keyword: []string{"FineReport"}},
		102: {Name: "帆软报表", Type: urlType, Keyword: []string{"Report"}},
		103: {Name: "帆软报表", Type: urlType, Keyword: []string{"/ReportServer"}},
		110: {Name: "极限OA", Type: bodyType, Keyword: []string{"templates/default/index.css"}},
		111: {Name: "极限OA", Type: bodyType, Keyword: []string{"images/sohuu.ico"}},
		120: {Name: "金蝶", Type: bodyType, Keyword: []string{"kingdee.com"}},
		121: {Name: "金蝶", Type: bodyType, Keyword: []string{"apusic.com"}},
		122: {Name: "金蝶EAS", Type: bodyType, Keyword: []string{"eassso/common"}},
		123: {Name: "金蝶EAS", Type: bodyType, Keyword: []string{"portalClientHelper.jsp"}},
		124: {Name: "金蝶EAS", Type: urlType, Keyword: []string{"eassso/login"}},
		130: {Name: "源天OA", Type: bodyType, Keyword: []string{"vmain/login.jsp"}},
		131: {Name: "源天OA", Type: bodyType, Keyword: []string{"vmain/browser.min.js"}},
		132: {Name: "源天OA", Type: urlType, Keyword: []string{"vmain/login.jsp"}},
		133: {Name: "源天OA", Type: bodyType, Keyword: []string{"vthemes/common/login.css"}},
		140: {Name: "红帆", Type: urlType, Keyword: []string{"ioffice/Portal"}},
		141: {Name: "红帆", Type: bodyType, Keyword: []string{"iOffice/Portal"}},
		142: {Name: "红帆", Type: bodyType, Keyword: []string{"iOffice.ne"}},
		150: {Name: "金和", Type: bodyType, Keyword: []string{"C6/WebResource.axd"}},
		151: {Name: "金和", Type: urlType, Keyword: []string{"JHSoft.Web.Login"}},
		152: {Name: "金和", Type: bodyType, Keyword: []string{"JHsoft.UI.Lib"}},
		160: {Name: "致翔", Type: bodyType, Keyword: []string{"document.getElementById('TPassword')"}},
		161: {Name: "致翔", Type: bodyType, Keyword: []string{"hex_md5(\"\" + aa + \"\")"}},
		170: {Name: "信呼协同办公系统", Type: bodyType, Keyword: []string{"信呼协同办公系统"}},
		171: {Name: "信呼协同办公系统", Type: bodyType, Keyword: []string{"webmain/login/loginscript.js"}},
		172: {Name: "信呼协同办公系统", Type: bodyType, Keyword: []string{"mode/plugin/jquery-rockmodels.js"}},
		180: {Name: "智明协同", Type: bodyType, Keyword: []string{"progid:DXImageTransform.Microsoft.AlphaImageLoader(src='/content/images/login_background.jpg"}},
		181: {Name: "智明协同", Type: bodyType, Keyword: []string{"OA智能办公系统"}},
		190: {Name: "一米OA", Type: bodyType, Keyword: []string{"skin/lte/login.css"}},
		191: {Name: "一米OA", Type: urlType, Keyword: []string{"/wap/index.jsp"}},
		192: {Name: "一米OA", Type: bodyType, Keyword: []string{"skin/lte/images"}},
		200: {Name: "新点OA", Type: bodyType, Keyword: []string{"EpointTextBox"}},
		201: {Name: "新点OA", Type: bodyType, Keyword: []string{"jwebui/js/dest"}},
		202: {Name: "weblogic", Type: bodyType, Keyword: []string{"Error 404--Not Found"}},
		210: {Name: "瑞数waf", Type: bodyType, Keyword: []string{"if($_ts.lcd)$_ts.lcd()"}},
		211: {Name: "瑞数waf", Type: bodyType, Keyword: []string{"$_ts=window['$_ts']"}},
		212: {Name: "瑞数waf", Type: bodyType, Keyword: []string{"412 Precondition Failed"}},
		213: {Name: "SpringBoot", Type: bodyType, Keyword: []string{"Whitelabel Error Page"}},
		220: {Name: "Nexus", Type: bodyType, Keyword: []string{"Nexus Repository Manage"}},
		221: {Name: "Nexus", Type: bodyType, Keyword: []string{"progressMessage('Initializing"}},
		230: {Name: "Grafana", Type: bodyType, Keyword: []string{"window.grafanaBootData"}},
		231: {Name: "Grafana", Type: bodyType, Keyword: []string{"grafana-app"}},
		240: {Name: "zabbix", Type: bodyType, Keyword: []string{"Zabbix SIA"}},
		241: {Name: "zabbix", Type: bodyType, Keyword: []string{"assets/img/apple-touch-icon"}},
		250: {Name: "phpcms", Type: bodyType, Keyword: []string{"<input type=\"hidden\" name=\"typeid\" value=\""}},
		251: {Name: "phpcms", Type: bodyType, Keyword: []string{"<input type=\"hidden\" name=\"m\" value=\"search\""}},
		260: {Name: "easySite", Type: bodyType, Keyword: []string{"\"easysiteHiddenDiv\""}},
		261: {Name: "SpringBoot", Type: bodyType, Keyword: []string{"Whitelabel Error Page"}},
		262: {Name: "SpringBoot", Type: bodyType, Keyword: []string{"\"status\":401,\"error\":\"Unauthorized\",\"message\":\"No message available\""}},
	}
)

func fingerprintIdentification(headers map[string][]string, body []byte, URL, LastURL string) []string {
	var result []string
	for key, value := range fingers {
		// 匹配URL
		if value.Type == urlType && strings.Contains(strings.ToLower(URL), strings.ToLower(value.Keyword[0])) {
			result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
		}

		// 匹配URL-2
		if value.Type == urlType && strings.Contains(strings.ToLower(LastURL), strings.ToLower(value.Keyword[0])) {
			result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
		}

		// 匹配body
		if value.Type == bodyType && strings.Contains(strings.ToLower(string(body)), strings.ToLower(value.Keyword[0])) {
			result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
		}

		// 匹配header
		if value.Type == headerType {
			for k, v := range headers {
				if strings.Contains(strings.ToLower(k+": "+strings.Join(v, " ")), strings.ToLower(value.Keyword[0])) {
					result = append(result, fmt.Sprintf("%d.%s", key, value.Name))
				}
			}
		}
	}
	return result
}
