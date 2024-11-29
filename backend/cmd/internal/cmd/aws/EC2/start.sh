#!/bin/bash

echo "Hello World" > /tmp/hello.txt
# 直接使用 chpasswd 修改密码
echo "root:nissan@123" | chpasswd
# 如果上面成功了则提示修改成功
if [ $? -eq 0 ]; then
    echo "修改密码成功" > /tmp/hello.txt
fi
# 修改ssh配置允许root登录
if grep -q "^#PermitRootLogin" /etc/ssh/sshd_config; then
    # 如果存在但被注释，取消注释并修改
    sed -i 's/^#PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config
elif grep -q "^PermitRootLogin" /etc/ssh/sshd_config; then
    # 如果存在且未注释，直接修改
    sed -i 's/^PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config
else
    # 如果不存在，添加
    echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
fi

# 修改 PasswordAuthentication
if grep -q "^#PasswordAuthentication" /etc/ssh/sshd_config; then
    # 如果存在但被注释，取消注释并修改
    sed -i 's/^#PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config
elif grep -q "^PasswordAuthentication" /etc/ssh/sshd_config; then
    # 如果存在且未注释，直接修改
    sed -i 's/^PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config
else
    # 如果不存在，添加
    echo "PasswordAuthentication yes" >> /etc/ssh/sshd_config
fi

# 重启sshd服务
systemctl restart ssh && systemctl restart sshd