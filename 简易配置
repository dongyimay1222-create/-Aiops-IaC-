1. 配置hosts
[webservers]
10.0.0.2
10.0.0.3

[lb]
10.0.0.1

2. 配置playbook.yml
---
- name: 高可用集群部署
  hosts: all
  become: true
  roles:
    - nginx
    - keepalived
    - zabbix

3. 配置roles/keepalived/tasks/main.yml
---
- name: 安装 Keepalived
  package:
    name: keepalived
    state: present

- name: 拷贝 Keepalived 配置
  template:
    src: keepalived.conf.j2
    dest: /etc/keepalived/keepalived.conf
  notify:
    - Restart Keepalived

- name: Restart Keepalived
  service:
    name: keepalived
    state: restarted

4. 配置roles/zabbix/tasks/main.yml
---
- name: 安装 Zabbix agent
  package:
    name: zabbix-agent
    state: present

- name: 启动 Zabbix agent
  service:
    name: zabbix-agent
    state: started
    enabled: yes

5. SaltStack 状态文件（多节点同步 & IaC）
salt/top.sls
base:
  '*':
    - common
    - nginx
    - keepalived
    - zabbix

salt/nginx.sls
nginx:
  pkg.installed
  service.running:
    - enable: True

salt/keepalived.sls
keepalived:
  pkg.installed
  service.running:
    - enable: True

salt/zabbix.sls
zabbix-agent:
  pkg.installed
  service.running:
    - enable: True

使用 salt '*' state.apply 即可在所有节点同步配置，实现批量部署。

6. Python 自愈脚本（日志异常检测与自动修复）
#!/usr/bin/env python3
# auto_heal.py
# 实现日志异常检测 + 自动重启服务

import os
import subprocess
import re

LOG_FILE = "/var/log/nginx/error.log"
SERVICES = ["nginx", "keepalived", "zabbix-agent"]

def check_logs():
    """检查日志中是否出现 ERROR 或 CRITICAL"""
    with open(LOG_FILE, "r") as f:
        logs = f.read()
    errors = re.findall(r"(ERROR|CRITICAL)", logs)
    return errors

def restart_service(service):
    """重启指定服务"""
    subprocess.run(["systemctl", "restart", service])
    print(f"[INFO] {service} 已重启")

if __name__ == "__main__":
    errors = check_logs()
    if errors:
        print(f"[ALERT] 检测到 {len(errors)} 条错误日志，启动自愈流程")
        for svc in SERVICES:
            restart_service(svc)
    else:
        print("[INFO] 日志正常，无需操作")

7. 一键执行脚本
#!/bin/bash
# deploy_all.sh
# 一键部署环境 + 配置 + 自愈脚本

# 1. 环境准备
bash setup_env.sh

# 2. Ansible 批量部署
ansible-playbook -i ansible/hosts ansible/playbook.yml

# 3. SaltStack 状态同步
salt '*' state.apply

# 4. 部署 Python 自动化脚本
mkdir -p /opt
cp auto_heal.py /opt/auto_heal.py
chmod +x /opt/auto_heal.py

# 5. 配置 cron 定时自愈
(crontab -l 2>/dev/null; echo "*/5 * * * * /usr/bin/python3 /opt/auto_heal.py >> /var/log/auto_heal.log 2>&1") | crontab -

echo "部署完成 高可用集群 + 自动化运维平台已启动"
