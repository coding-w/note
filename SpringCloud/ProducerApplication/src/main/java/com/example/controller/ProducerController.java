package com.example.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Random;


@RestController
@RequestMapping("produce")
public class ProducerController {

    @GetMapping
    public String hello(String name) {
        System.out.println("Hello " + name);
        return "Hello " + name;
    }

    @GetMapping("/get")
    public String get() {
        Random random = new Random();
        return "produce " + random.nextInt(100);
    }
}
