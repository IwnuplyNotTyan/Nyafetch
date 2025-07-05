package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/muesli/termenv"
)

func getDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown"
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(line[12:], "\"")
		}
	}
	return "Unknown"
}

func getKernelVersion() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

func getWM() string {
	out, err := exec.Command("ps", "-e", "-o", "comm=").Output()
	if err != nil {
		return "Unknown"
	}

	processList := strings.Split(string(out), "\n")
	wmList := []string{"i3", "openbox", "kwin_x11", "mutter", "xfwm4", "metacity", "compiz", "fluxbox", "bspwm", "awesome", "dwm"}

	for _, proc := range processList {
		trimmedProc := strings.TrimSpace(proc)
		for _, wm := range wmList {
			if trimmedProc == wm {
				return wm
			}
		}
	}
	return "Unknown"
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	hn := termenv.String("hn ").Foreground(termenv.ANSIBlue)
	kr := termenv.String("kr ").Foreground(termenv.ANSIRed)
	os := termenv.String("os ").Foreground(termenv.ANSIGreen)
	wm := termenv.String("wm ").Foreground(termenv.ANSICyan)

	fmt.Println("  ／l、    ", os, getDistro())
	fmt.Println("（ﾟ､ ｡ ７  ", wm, getWM())
	fmt.Println(" l  ~ ヽ   ", kr, getKernelVersion())
	fmt.Println(" じしf_,)ノ", hn, hostname)
}
