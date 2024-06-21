package ipfs

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func InitIPFS() error {
	fmt.Println("[IPFS] Init")
	// installを開始する
	err := InstallIPFS()
	if err != nil {
		fmt.Println("[IPFS] Install error:", err)
		return err
	}
	return nil
}

func StartDaemon() error {
	if isIPFSDaemonRunning() {
		fmt.Println("[IPFS] Daemon is already running.")
		return nil
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "start", ipfsExecutablePath, "daemon")
	} else {
		cmd = exec.Command("nohup", ipfsExecutablePath, "daemon", "&")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("[IPFS] Starting daemon...")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("[IPFS] Failed to start daemon: %v", err)
	}

	fmt.Println("[IPFS] Daemon started.")
	return nil
}

func isIPFSDaemonRunning() bool {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("pgrep", "ipfs")
	}
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	if runtime.GOOS == "windows" {
		return strings.Contains(string(output), "ipfs.exe")
	} else {
		return len(strings.TrimSpace(string(output))) > 0
	}
}

func Upload(filePath string) (string, error) {
	fmt.Println("[IPFS] Upload:", filePath)

	cmd := exec.Command(ipfsExecutablePath, "add", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("[IPFS] Add error: %v, %s", err, string(output))
	}
	fmt.Println("[IPFS] Add output:", string(output))

	re := regexp.MustCompile(`added (\S+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return "", fmt.Errorf("[IPFS] Unexpected output: %s", output)
	}
	return matches[1], nil
}

func Download(cid, filePath string) error {
	fmt.Println("[IPFS] Download:", cid, filePath)

	cmd := exec.Command(ipfsExecutablePath, "get", cid, "-o", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("[IPFS] Get error: %v, %s", err, string(output))
	}

	fmt.Println("[IPFS] Get output:", string(output))

	return nil
}
