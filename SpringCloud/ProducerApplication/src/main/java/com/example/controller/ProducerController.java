package com.example.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
public class ProducerController {

    @GetMapping("produce")
    public String hello(String name) {
        System.out.println("Hello " + name);
        return "Hello " + name;
    }
}
