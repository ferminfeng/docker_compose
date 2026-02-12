# 创建测试表
```azure
docker exec -it full_scylla_single cqlsh -ucassandra -pcassandra

CREATE KEYSPACE IF NOT EXISTS message_local_26_db WITH replication = {
'class': 'SimpleStrategy',
'replication_factor': 1 } AND durable_writes = true;
```

# 查看所有键空间（能看到 message_local_26_db 即为成功）
```azure
DESCRIBE KEYSPACES;
```

# 或精准查询该键空间的详情
```azure
DESCRIBE KEYSPACE message_local_26_db;
```