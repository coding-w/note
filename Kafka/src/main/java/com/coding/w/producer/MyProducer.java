package com.coding.w.producer;

import org.apache.kafka.clients.producer.Callback;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.apache.kafka.common.serialization.StringSerializer;

import java.util.Properties;

public class MyProducer {
    public static void main(String[] args) throws InterruptedException {
        // 配置
        Properties props = new Properties();
        // 连接集群
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka1:9093,kafka2:9094,kafka3:9095");

        // 指定key和value序列化
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        // 关联自定义分区
        // props.put(ProducerConfig.PARTITIONER_CLASS_CONFIG, MyPartitioner.class.getName());
        // props.put(ProducerConfig.BATCH_SIZE_CONFIG, 16384);
        // props.put(ProducerConfig.LINGER_MS_CONFIG, 1);
        // kafka生产者对象
        KafkaProducer<String, String> kafkaProducer = new KafkaProducer<>(props);

        // 发送数据
        for (int i = 0; i < 500; i++) {
            kafkaProducer.send(new ProducerRecord<>("test2","Hello" + i), new Callback() {
                @Override
                public void onCompletion(RecordMetadata recordMetadata, Exception e) {
                    if (e == null) {
                        System.out.println("topic: " + recordMetadata.topic() + ", partition: " + recordMetadata.partition());
                    }
                }
            });
            Thread.sleep(10);
        }

        // 关闭资源
        kafkaProducer.close();
    }
}
