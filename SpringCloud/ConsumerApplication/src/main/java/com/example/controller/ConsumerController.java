package com.example.controller;

import com.example.feign.ConsumerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;


@RestController
public class ConsumerController {

    @Autowired
    private RestTemplate restTemplate;

    @Autowired
    private ConsumerService consumerService;

    @GetMapping("consume")
    // http://localhost:8003/consume?name=
    public String consume(String name) {
        // String url = "http://producer-application/produce?name=";
        // return restTemplate.getForObject(url+ name, String.class);
        return consumerService.produce(name);
    }
}
