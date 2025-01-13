package com.coding.w.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/13 14:12
 */
@RestController
public class ProduceController {

    @Autowired
    KafkaTemplate<String, String> kafkaTemplate;

    @GetMapping("kafka")
    public String data(String msg) {
        kafkaTemplate.send("aaa", msg);
        return "success";
    }
}
