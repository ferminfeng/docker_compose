#!/bin/bash

echo "========== Redis集群验证开始 =========="
echo ""

# 检查所有节点
PORTS=(7001 7002 7003 7004 7005 7006)

for port in "${PORTS[@]}"; do
    echo "检查节点 127.0.0.1:$port"
    echo "------------------------"

    # 测试连接
    if redis-cli -p $port ping | grep -q "PONG"; then
        echo "✓ 连接正常"

        # 获取角色
        role=$(redis-cli -p $port info replication | grep "role:" | cut -d: -f2)
        echo "  角色: $role"

        # 获取集群节点数
        if [ "$role" = "master" ]; then
            keys=$(redis-cli -c -p $port dbsize)
            echo "  键数量: $keys"
        fi

        # 查看节点信息
        echo "  集群节点信息:"
        redis-cli -c -p $port cluster nodes | grep "myself" | head -1
    else
        echo "✗ 连接失败"
    fi
    echo ""
done

echo "========== 集群完整性检查 =========="
redis-cli --cluster check 127.0.0.1:7001 --cluster-search-multiple-owners