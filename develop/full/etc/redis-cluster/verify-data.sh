#!/bin/bash

echo "测试数据分片..."
echo "======================"

# 写入测试数据
for i in {1..100}; do
    key="key_${i}"
    value="value_${i}"
    redis-cli -c -p 7001 set $key $value > /dev/null
    echo "设置 $key = $value"
done

echo ""
echo "======================"
echo "检查数据分布："

# 检查各节点的键数量
for port in 7001 7002 7003; do
    count=$(redis-cli -c -p $port dbsize)
    echo "节点 $port: $count 个键"
done

echo ""
echo "查看具体数据分布："
redis-cli -c -p 7001 --scan --pattern "key_*" | while read key; do
    location=$(redis-cli -c -p 7001 cluster keyslot $key)
    echo "键: $key -> 槽位: $location"
done | head -20

echo ""
echo "检查槽位分配...(应该显示所有16384个槽位都已分配)"
redis-cli --cluster check 127.0.0.1:7001 | grep -A5 "slots"