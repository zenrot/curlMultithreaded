package antidbg

import (
	"os"
	"os/exec"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

func isDebuggerPresent() bool {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	isDebuggerPresent := kernel32.NewProc("IsDebuggerPresent")
	ret, _, _ := isDebuggerPresent.Call()
	return ret != 0
}

func checkRemoteDebuggerPresent() bool {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	check := kernel32.NewProc("CheckRemoteDebuggerPresent")
	var isDebugger uint32
	hProcess := windows.CurrentProcess() // хендл текущего процесса
	check.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&isDebugger)))
	return isDebugger != 0
}

func checkDbgTools() bool {
	out, err := exec.Command("tasklist").Output()
	if err != nil {
		return false
	}
	processes := string(out)
	if strings.Contains(processes, "ollydbg.exe") || strings.Contains(processes, "windbg.exe") {
		return true
	}
	return false
}

func checkExecutionTime() bool {
	start := time.Now()

	for i := 0; i < 1e7; i++ {
		_ = i * i
	}

	elapsed := time.Since(start)
	if elapsed > 500*time.Millisecond {
		return true
	}
	return false
}

func DebuggerSec() {
	go func() {
		if isDebuggerPresent() || checkRemoteDebuggerPresent() || checkExecutionTime() || checkDbgTools() {
			os.Exit(1)
		}
	}()
}
