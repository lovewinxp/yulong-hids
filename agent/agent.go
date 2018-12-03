package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"yulong-hids/agent/client"
	"yulong-hids/daemon/common"
)

func main() {
	fmt.Println("test")
	if len(os.Args) <= 1 {
		fmt.Println("Usage: agent[.exe] ServerIP [debug]")
		fmt.Println("Example: agent 8.8.8.8 debug")
		return
	}
	out, err := common.CmdExec("lsmod| grep syshook_execve")
	//out, err := common.CmdExec(fmt.Sprintf("insmod %s/syshook_execve.ko", common.InstallPath))
	if strings.Contains(out, "syshook_execve") {
		log.Println("Insmod syshook_execve succeeded")
	} else {
		//log.Println("Insmod syshook_execve error, command output:", out)
		out, err = common.CmdExec(fmt.Sprintf("insmod %s/syshook_execve.ko", common.InstallPath))
		//fmt.Println(out)

		fmt.Println(fmt.Sprintf("insmod %s/syshook_execve.ko", common.InstallPath))
		fmt.Println(err)
	}
	var agent client.Agent
	agent.ServerNetLoc = os.Args[1]
	if len(os.Args) == 3 && os.Args[2] == "debug" {
		log.Println("DEBUG MODE")
		agent.IsDebug = true
	}
	agent.Run()
}
