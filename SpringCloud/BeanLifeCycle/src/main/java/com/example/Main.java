package com.example;

import com.example.config.User;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

/**
 * @author wangxiang
 * @description
 * @create 2025/5/5 18:25
 */
@SpringBootApplication
public class Main {
    public static void main(String[] args) {
        ConfigurableApplicationContext applicationContext = SpringApplication.run(Main.class, args);
        User bean = applicationContext.getBean(User.class);
        System.out.println(bean);
    }
}