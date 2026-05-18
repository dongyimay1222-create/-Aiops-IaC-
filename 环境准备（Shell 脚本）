#!/bin/bash
# setup_env.sh
# 自动化初始化环境：安装必要软件、配置防火墙和用户

# 更新系统
sudo dnf update -y || sudo yum update -y

# 安装必要组件
sudo dnf install -y python3 python3-pip git vim curl wget || sudo yum install -y python3 python3-pip git vim curl wget

# 安装 Ansible 和 SaltStack
pip3 install --user ansible salt

# 配置防火墙（允许 HTTP/Nginx, Zabbix, Prometheus）
sudo firewall-cmd --add-service=http --permanent
sudo firewall-cmd --add-service=https --permanent
sudo firewall-cmd --add-port=10050/tcp --permanent  # Zabbix agent
sudo firewall-cmd --reload

echo "环境初始化完成"
