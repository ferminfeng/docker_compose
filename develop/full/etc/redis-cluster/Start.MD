redis-cli --cluster create 172.20.10.7:7001 172.20.10.7:7002 172.20.10.7:7003 172.20.10.7:7004 172.20.10.7:7005 172.20.10.7:7006 --cluster-replicas 1


redis-cli --cluster create 172.20.10.7:6479 172.20.10.7:6480 172.20.10.7:6481 --cluster-replicas 0

redis-cli --cluster create 192.168.0.36:6479 192.168.0.36:6480 192.168.0.36:6481 --cluster-replicas 0


