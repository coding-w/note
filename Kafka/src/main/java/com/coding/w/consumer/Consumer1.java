package com.coding.w.consumer;

import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.serialization.StringDeserializer;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;


public class Consumer1 {
    public static void main(String[] args) throws Exception {

        // 0. 配置
        Properties props =  getProperties();
        // 1. 创建消费者
        KafkaConsumer<String, String> kafkaConsumer = new KafkaConsumer<>(props);
        // 2. 定义主题 first
        List<String> topics = new ArrayList<>();
        topics.add("test");
        kafkaConsumer.subscribe(topics);
        // 3. 消费数据
        while (true) {
            ConsumerRecords<String, String> consumerRecords = kafkaConsumer.poll(Duration.ofSeconds(1));
            for (ConsumerRecord<String, String> item : consumerRecords) {
                System.out.println(item);
            }
        }
    }

    private static Properties getProperties() {
        Properties props = new Properties();
        // 连接
        props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka1:9093,kafka2:9094,kafka3:9095");
        // 反序列化
        props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class.getName());
        props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class.getName());
        // 消费者组名
        props.put(ConsumerConfig.GROUP_ID_CONFIG, "test");
        return props;
    }


}
