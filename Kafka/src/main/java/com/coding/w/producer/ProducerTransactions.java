package com.coding.w.producer;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;

import java.util.Properties;

public class ProducerTransactions {
    public static void main(String[] args) {
        // 1. 配置
        Properties props = new Properties();
        // 2. 连接集群
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka1:9093,kafka2:9094,kafka3:9095");
        // 3. 指定key和value序列化
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        props.put(ProducerConfig.TRANSACTIONAL_ID_CONFIG, "1");
        // 压缩
        props.put(ProducerConfig.COMPRESSION_TYPE_CONFIG, "snappy");
        // 4. kafka生产者对象
        KafkaProducer<String, String> kafkaProducer = new KafkaProducer<>(props);
        // 5. 初始化事物
        kafkaProducer.initTransactions();
        // 6. 开启事物
        kafkaProducer.beginTransaction();
        try {
            for (int i = 0; i < 5; i++) {
                kafkaProducer.send(new ProducerRecord<>("first", "kafka" + i));
            }
            kafkaProducer.commitTransaction();
        } catch (Exception e) {
            kafkaProducer.abortTransaction();
        } finally {
            kafkaProducer.close();
        }
    }

}
