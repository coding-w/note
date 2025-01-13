package com.coding.w.consumer;

import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.clients.consumer.OffsetAndTimestamp;
import org.apache.kafka.common.TopicPartition;
import org.apache.kafka.common.serialization.StringDeserializer;

import java.time.Duration;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Properties;
import java.util.Set;


public class ConsumerSeekTime {
    public static void main(String[] args) throws Exception {

        // 0. 配置
        Properties props =  getProperties();
        // 1. 创建消费者
        KafkaConsumer<String, String> kafkaConsumer = new KafkaConsumer<>(props);
        // 2. 定义主题 test
        List<String> topics = new ArrayList<>();
        topics.add("test");
        kafkaConsumer.subscribe(topics);
        // 保证分区分配方案
        Set<TopicPartition> assignment = new HashSet<>();
        while (assignment.isEmpty()) {
            kafkaConsumer.poll(Duration.ofSeconds(1));
            // 获取消费者分区分配信息（有了分区分配信息才能开始消费）
            assignment = kafkaConsumer.assignment();
        }
        // 把时间转换对应offset
        HashMap<TopicPartition, Long> partitionLongHashMap = new HashMap<>();
        for (TopicPartition topicPartition : assignment) {
            partitionLongHashMap.put(topicPartition, System.currentTimeMillis() - 1 * 24 * 60 * 60 * 1000);
        }
        Map<TopicPartition, OffsetAndTimestamp> offsetAndTimestampMap = kafkaConsumer.offsetsForTimes(partitionLongHashMap);

        for (TopicPartition topicPartition : assignment) {
            OffsetAndTimestamp offsetAndTimestamp = offsetAndTimestampMap.get(topicPartition);
            kafkaConsumer.seek(topicPartition, offsetAndTimestamp.offset());
        }

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
