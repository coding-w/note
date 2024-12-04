cp -i ./redis.conf ./7002/conf/ ./7001/conf


redis-cli --cluster create 192.168.3.222:7000 192.168.3.222:7001 \
192.168.3.222:7002 192.168.3.222:7003 192.168.3.222:7004 192.168.3.222:7005 \
--cluster-replicas 1

firewall-cmd --zone=public --add-port=7000-7010/tcp --permanent


rm -rf ./7000/data/* ./7001/data/* ./7002/data/* ./7003/data/* ./7004/data/* ./7005/data/*
