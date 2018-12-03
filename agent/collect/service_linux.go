// +build linux

package collect

import (
	"regexp"
	"strings"
	"yulong-hids/agent/common"
)

type service struct {
	Caption   string
	Name      string
	PathName  string
	Started   bool
	StartMode string
	StartName string
}

// GetServiceInfo 获取服务列表
func GetServiceInfo()  (resultdata []map[string]string) {

	var serviceStr string=""
	serviceStr = common.Cmdexec("systemctl list-units --type service")
	//serviceStr = common.Cmdexec("chkconfig --list")
	if serviceStr != "" {

		serviceList := strings.Split(serviceStr, "\n")
		if len(serviceList) < 2 {
			return
		}
		for _, info := range serviceList[1 : len(serviceList)-1] {

			if !strings.Contains(info,"service") {
				continue
			}
			m := make(map[string]string)
			reg := regexp.MustCompile("\\s+")
			info = reg.ReplaceAllString(strings.TrimSpace(strings.Replace(info,"●","",-1)), " ")
			s := strings.Split(info, " ")
			if len(s) < 4 {
				continue
			}
			m["name"] = s[0]
			//"pathname": "启动命令，同command",
			//	"started": "当前启动状态",
			//	"startmode": "开机启动模式",
			//	"startname": "启动用户",
			//	"caption": "描述"
			//m["pathname"] = "pathname"
			m["started"] = s[3]

			m["caption"] = string([]rune(info)[strings.Index(info,s[4]):])


			resultdata = append(resultdata, m)

		}




	}	else 	{
		serviceStr = common.Cmdexec("chkconfig --list")
		serviceList := strings.Split(serviceStr, "\n")
		if len(serviceList) < 2 {
			return
		}
		for _, info := range serviceList[1 : len(serviceList)-1] {

			m := make(map[string]string)
			reg := regexp.MustCompile("\\s+")
			info = reg.ReplaceAllString(strings.TrimSpace(info), " ")
			s := strings.Split(info, " ")
			if len(s) < 6 {
				continue
			}
			m["name"] = s[0]
			//"pathname": "启动命令，同command",
			//	"started": "当前启动状态",
			//	"startmode": "开机启动模式",
			//	"startname": "启动用户",
			//	"caption": "描述"
			//m["pathname"] = "pathname"
			//m["started"] = "started"
			//m["startmode"] = "startmode"
			//m["startname"] = "startname"
			//m["caption"] = "caption"


			resultdata = append(resultdata, m)

		}

	}



	return resultdata
}
