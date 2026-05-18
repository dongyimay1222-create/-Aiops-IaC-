# -Aiops-IaC-
基于Aiops与IaC的智能自动化运维平台实践
AIOps-IaC-Platform/
│
├─ README.md                  # 项目概述、亮点、成果、使用说明
├─ setup_env.sh               # 系统初始化 + 软件依赖安装
├─ deploy_platform.sh         # 一键部署 AIOps 平台 + 高可用集群
├─ auto_heal.py               # Python/Shell 日志异常检测与闭环自愈
│
├─ ansible/                   # Ansible 部署和 IaC
│   ├─ hosts                  # 目标节点清单
│   ├─ playbook.yml           # 主 Playbook
│   └─ roles/
│       ├─ nginx/
│       │   └─ tasks/main.yml
│       ├─ keepalived/
│       │   └─ tasks/main.yml
│       └─ a_ops/
│           └─ tasks/main.yml
│
├─ salt/                      # SaltStack 状态文件，实现多节点同步
│   ├─ top.sls
│   ├─ common.sls
│   ├─ nginx.sls
│   ├─ keepalived.sls
│   └─ a_ops.sls
│
├─ templates/                 # 配置模板文件
│   ├─ nginx.conf.j2
│   ├─ keepalived.conf.j2
│   └─ a_ops_config.yml.j2
│
└─ docs/                      # 项目架构图、设计文档
    └─ architecture.svg
