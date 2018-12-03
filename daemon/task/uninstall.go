package task

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
	"yulong-hids/daemon/common"
)

// UnInstallALL 卸载
func UnInstallALL() {
	common.KillAgent()

	if runtime.GOOS == "windows" {
		common.CmdExec("net stop pro")
		// common.CmdExec("net stop npf")
		common.CmdExec("sc delete pro")
		// common.CmdExec("sc delete npf")
	} else {
		time.Sleep(time.Second * 1)
		out, err := common.CmdExec("lsmod| grep syshook_execve")
		//out, err := common.CmdExec(fmt.Sprintf("insmod %s/syshook_execve.ko", common.InstallPath))
		if strings.Contains(out, "syshook_execve") {
			out, err = common.CmdExec("rmmod syshook_execve")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(out)
				log.Println("rm syshook_execve succeeded")
			}
		}
		common.CmdExec("rm -rf "+common.InstallPath)
		common.CmdExec("rm -rf /tmp/daemon")
		fmt.Println("Delete "+common.InstallPath+" successful")
	}
	if err := common.Service.Uninstall(); err != nil {
		log.Println("Uninstall yulong-hids error:", err.Error())
	}
	log.Println("Uninstall completed")
	os.Exit(1)
}

func ClearALL() {

	if runtime.GOOS == "windows" {
		common.CmdExec("net stop pro")
		// common.CmdExec("net stop npf")
		common.CmdExec("sc delete pro")
		// common.CmdExec("sc delete npf")
	} else {


		out, err := common.CmdExec( "ps -ef | grep '"+common.InstallPath+"agent' | grep -v 'grep' | awk '{print $2\" \"$3}'" )

		if len(out)>0 {
			s:=strings.Split(strings.Replace(out, "\n", "", -1)," ")
			common.CmdExec("kill "+s[0])

			log.Println("Kill Agent successful")
			//common.CmdExec("rmmod syshook_execve")
			time.Sleep(time.Second * 1)
			out, err := common.CmdExec("lsmod| grep syshook_execve")
			//out, err := common.CmdExec(fmt.Sprintf("insmod %s/syshook_execve.ko", common.InstallPath))
			if strings.Contains(out, "syshook_execve") {
				out, err = common.CmdExec("rmmod syshook_execve")
				if err != nil {
					fmt.Println(err)
				} else {

					log.Println("rm syshook_execve succeeded")
				}
			}


			common.CmdExec("kill "+s[1])
			log.Println("Kill Daemon successful")
			common.CmdExec("rm -rf "+common.InstallPath)
			log.Println("Delete "+common.InstallPath+" successful")
		} else {

			fmt.Println(err)
		}



	}
	if err := common.Service.Uninstall(); err != nil {
		log.Println("Uninstall yulong-hids error:", err.Error())
	}
	log.Println("Uninstall completed")
	os.Exit(1)
}