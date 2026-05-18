# -Aiops-IaC-
基于Aiops与IaC的智能自动化运维平台实践
# 基于 AIOps 与 IaC 的智能自动化运维平台实践

## 项目概述
该项目结合 AIOps + 基础设施即代码 (IaC)，构建企业级高可用 Linux 集群，实现智能日志分析、指标预测、告警闭环与自动修复。  
目标是通过 AI Agent + Python/Shell 自动化脚本，提升运维效率并保障业务高可用性。

## 技术栈
- 操作系统与集群：OpenEuler / CentOS  
- 高可用架构：Nginx + LVS DR + Keepalived  
- 自动化部署与 IaC：Ansible Playbook + SaltStack  
- 智能运维与 AIOps：A-OPS 平台 + AI Agent / LLM  
- 多维可观测性：Zabbix + Prometheus  
- 脚本与自愈：Python / Shell  
- 版本控制与 CI/CD：Git

## 核心亮点
1. 高可用集群：VRRP + LVS + Keepalived，实现单点故障平滑切换  
2. 基础设施即代码：Ansible + SaltStack 支持多节点批量部署与无损滚动重启  
3. 智能自愈：AI Agent 结合 Python/Shell 脚本实现日志异常检测和自动修复  
4. 多维监控与告警：Zabbix + Prometheus，自定义指标，精细化告警策略  
5. 网络安全与策略：TCP/IP、DNS/HTTP、SELinux、安全隔离

## 快速部署
```bash
# 初始化环境
bash setup_env.sh

# 一键部署平台与高可用集群
bash deploy_platform.sh

# 自动化自愈脚本已配置 cron 定时运行
