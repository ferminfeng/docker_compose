# 可以在宿主机设置一下hosts
```azure
vi /etc/hosts

127.0.0.1 host.docker.internal
```

# 创建集群
```azure
docker exec -it yh_redis_6479 redis-cli --cluster create host.docker.internal:6479 host.docker.internal:6480 host.docker.internal:6481 --cluster-replicas 0
```


# 验证(方式一)
```azure
docker exec -it yh_redis_6479 redis-cli cluster nodes

docker exec -it yh_redis_6479 redis-cli cluster info
```

# 验证(方式二)

- 验证集群状态
```azure
sh ./etc/redis-cluster/verify-cluster.sh
```

- 验证数据写入、查看数据分布
```azure
sh ./etc/redis-cluster/verify-data.sh
```

- 验证主从复制关系
- - 由于创建集群时 cluster-replicas=0 未启用从节点，可忽略本条
```azure
sh ./etc/redis-cluster/verify-master-slave.sh
```

# 测试数据读写（验证集群可用）
```azure
- 写入数据
    docker exec -it yh_redis_6479 redis-cli -c -p 6379 SET test "im-cluster-success"
    
- 读取数据（会自动路由到对应节点）
    docker exec -it yh_redis_6479 redis-cli -c -p 6379 GET test

```

# 修复集群配置不一致问题
```azure
docker exec -it yh_redis_6479 redis-cli --cluster fix 127.0.0.1:6379
```