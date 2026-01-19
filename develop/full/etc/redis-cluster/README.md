# 创建docker-compose.yml

###################################### docker-compose内容 ##########################################
version: "3.9"

services:
  redis-6479:
    image: redis:7.2
    container_name: redis-6479
    command: >
      redis-server --port 6379
                   --bind 0.0.0.0
                   --protected-mode no
                   --cluster-enabled yes
                   --cluster-config-file nodes.conf
                   --cluster-node-timeout 5000
                   --cluster-announce-ip host.docker.internal
                   --cluster-announce-port 6479
                   --cluster-announce-bus-port 16479
                   --dir /data
                   --appendonly yes
    ports:
      - "6479:6379"
      - "16479:16379"
    volumes:
      - ./data/6479:/data

  redis-6480:
    image: redis:7.2
    container_name: redis-6480
    command: >
      redis-server --port 6379
                   --bind 0.0.0.0
                   --protected-mode no
                   --cluster-enabled yes
                   --cluster-config-file nodes.conf
                   --cluster-node-timeout 5000
                   --cluster-announce-ip host.docker.internal
                   --cluster-announce-port 6480
                   --cluster-announce-bus-port 16480
                   --dir /data
                   --appendonly yes
    ports:
      - "6480:6379"
      - "16480:16379"
    volumes:
      - ./data/6480:/data

  redis-6481:
    image: redis:7.2
    container_name: redis-6481
    command: >
      redis-server --port 6379
                   --bind 0.0.0.0
                   --protected-mode no
                   --cluster-enabled yes
                   --cluster-config-file nodes.conf
                   --cluster-node-timeout 5000
                   --cluster-announce-ip host.docker.internal
                   --cluster-announce-port 6481
                   --cluster-announce-bus-port 16481
                   --dir /data
                   --appendonly yes
    ports:
      - "6481:6379"
      - "16481:16379"
    volumes:
      - ./data/6481:/data


	
###################################### docker-compose内容 ##########################################



# 当前目录下创建redis集群映射数据目录
mkdir data, data\6479, data\6480, data\6481

# 创建并启动容器
docker-compose up -d

# 检查容器状态（显示 Up 即为正常）
docker ps | findstr redis-

# 创建集群
 # docker exec -it redis-6479 redis-cli --cluster create redis-6479:6379 redis-6480:6379 redis-6481:6379 --cluster-replicas 0
 
 # docker exec -it redis-6479 redis-cli --cluster create 127.0.0.1:6479 127.0.0.1:6480 127.0.0.1:6481 --cluster-replicas 0
 
 # windows->
	docker exec -it redis-6479 redis-cli --cluster create host.docker.internal:6479 host.docker.internal:6480 host.docker.internal:6481 --cluster-replicas 0



# 验证集群状态
docker exec -it redis-6479 redis-cli cluster nodes

docker exec -it redis-6479 redis-cli cluster info

		# cluster_state:ok
		
# 测试数据读写（验证集群可用）
		# 写入数据
		docker exec -it redis-6479 redis-cli -c -p 6479 SET test "im-cluster-success"

		# 读取数据（会自动路由到对应节点）
		docker exec -it redis-6479 redis-cli -c -p 6479 GET test
