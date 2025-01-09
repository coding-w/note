rm -rf ./kafka_data1/* ./kafka_data2/* ./kafka_data3/* ./zookeeper1/* ./zookeeper2/* ./zookeeper3/*

mkdir ./kafka_data1 ./kafka_data2 ./kafka_data3 ./zookeeper1 ./zookeeper2 ./zookeeper3

chmod -R 777 ./kafka_data1 ./kafka_data2 ./kafka_data3 ./zookeeper1 ./zookeeper2 ./zookeeper3

kafka-topics.sh --list --bootstrap-server kafka1:9093
