#!/bin/bash
# 一键部署高可用集群 + AIOps 平台

# 1. Ansible 批量部署
ansible-playbook -i ansible/hosts ansible/playbook.yml

# 2. SaltStack 多节点同步
salt '*' state.apply

# 3. 部署自动化自愈脚本
cp auto_heal.py /opt/auto_heal.py
chmod +x /opt/auto_heal.py
(crontab -l 2>/dev/null; echo "*/5 * * * * /usr/bin/python3 /opt/auto_heal.py >> /var/log/auto_heal.log 2>&1") | crontab -

echo "部署完成 ✅ 高可用集群 + AIOps 平台已启动"
