#!/bin/bash
# verify-replication.sh

echo "验证主从复制关系..."
echo "======================"

# 获取所有主节点
masters=$(redis-cli -c -p 7001 cluster nodes | grep "master" | grep -v "fail" | awk '{print $1, $2}')

while read -r master_id master_addr; do
    echo "主节点: $master_addr ($master_id)"

    # 查找该主节点的从节点
    slaves=$(redis-cli -c -p 7001 cluster nodes | grep "slave" | grep "$master_id")

    if [ -n "$slaves" ]; then
        echo "  从节点:"
        echo "$slaves" | while read slave; do
            echo "    - $(echo $slave | awk '{print $2}')"
        done
    else
        echo "  警告: 没有从节点!"
    fi
    echo ""
done <<< "$masters"

echo "测试主从同步..."
# 在主节点写入数据
redis-cli -c -p 7001 set replication_test "master_value"

# 检查从节点是否同步
for port in 7004 7005 7006; do
    value=$(redis-cli -c -p $port get replication_test 2>/dev/null)
    if [ "$value" = "\"master_value\"" ]; then
        echo "✓ 从节点 $port 同步正常"
    else
        echo "✗ 从节点 $port 同步异常"
    fi
done