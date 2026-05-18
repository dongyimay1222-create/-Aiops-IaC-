#!/usr/bin/env python3
import re, subprocess

LOG_FILES = ["/var/log/nginx/error.log", "/var/log/a_ops.log"]
SERVICES = ["nginx", "keepalived", "a_ops"]

def scan_logs():
    alerts = []
    for f in LOG_FILES:
        with open(f) as log:
            for line in log:
                if re.search(r"(ERROR|CRITICAL|FAIL)", line):
                    alerts.append(line.strip())
    return alerts

def self_heal():
    errors = scan_logs()
    if errors:
        print(f"[ALERT] 检测到 {len(errors)} 条异常日志")
        for svc in SERVICES:
            subprocess.run(["systemctl", "restart", svc])
            print(f"[INFO] {svc} 已重启")
    else:
        print("[INFO] 日志正常，无需操作")

if __name__ == "__main__":
    self_heal()
