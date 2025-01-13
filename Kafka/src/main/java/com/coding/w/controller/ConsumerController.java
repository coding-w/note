package com.coding.w.controller;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/13 14:28
 */
@RestController
public class ConsumerController {

    @KafkaListener(topics = "aaa")
    public void consumerTopic(String msg) {
        System.out.println("收到消息：" + msg);
    }
}
