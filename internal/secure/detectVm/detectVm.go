package detectVm

import (
	"hedgedcurl/internal/secure/secfmt"
	"net"
	"os"
)

func detectVMByMAC() {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		mac := iface.HardwareAddr.String()
		if mac == "" {
			continue
		}
		if mac[:8] == secfmt.Sprintf("RVVYU1VEXA==") || mac[:8] == secfmt.Sprintf("RV9SVixIV0s=") || mac[:8] == secfmt.Sprintf("RVdSVl9IV0U=") {
			os.Exit(1)
		}
	}
}

func detectVMByFile() {
	//ищем VirtualBox драйвер
	if _, err := os.Stat(secfmt.Sprintf("WhwRFUARCRMcFlkWCBtNDAtZHxAaCx0FGy0LEwIA")); err == nil {
		data, _ := os.ReadFile(secfmt.Sprintf("WhwRFUARCRMcFlkWCBtNDAtZHxAaCx0FGy0LEwIA"))
		if string(data) == secfmt.Sprintf("IwYaEhoTCTAAHQ==") {
			os.Exit(1)
		}
	}
}

func CheckVM() {
	go detectVMByFile()
	go detectVMByMAC()
}
