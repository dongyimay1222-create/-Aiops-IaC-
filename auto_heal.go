package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// 定义需要监控的日志路径和核心服务
const LogFile = "/var/log/nginx/error.log"
var Services = []string{"nginx", "keepalived", "zabbix-agent"}

// checkLogs 实时流式读取日志并检测异常关键字
func checkLogs() int {
	file, err := os.Open(LogFile)
	if err != nil {
		return 0 // 如果日志文件不存在，直接返回
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorCount := 0

	// 流式扫描，避免文件过大撑爆内存
	for scanner.Scan() {
		line := scanner.Text()
		lineLower := strings.ToLower(line)
		// 匹配严重故障关键字
		if strings.Contains(lineLower, "error") || strings.Contains(lineLower, "critical") || strings.Contains(lineLower, "failed") {
			errorCount++
		}
	}
	return errorCount
}

// restartService 调用系统命令执行无状态服务自愈
func restartService(service string) {
	cmd := exec.Command("systemctl", "restart", service)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ERROR] 尝试重启服务 %s 失败: %v\n", service, err)
		return
	}
	fmt.Printf("[INFO] 服务 %s 已通过 Go 自愈引擎成功修复\n", service)
}

func main() {
	fmt.Printf("[%s] [INFO] Go云原生自愈守护进程开始日常巡检...\n", time.Now().Format("2006-01-02 15:04:05"))
	
	errors := checkLogs()
	if errors > 0 {
		fmt.Printf("[%s] [ALERT] 触发高并发自愈逻辑！检测到 %d 条系统严重错误日志。\n", time.Now().Format("2006-01-02 15:04:05"), errors)
		// 并发拉起所有故障服务，充分发挥 Go 的高并发优势
		for _, svc := range Services {
			go restartService(svc) 
		}
		// 挂起主线程，等待协程执行完毕（实际生产可以加入 sync.WaitGroup）
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("[INFO] 系统指标与日志一切正常。")
	}
}
