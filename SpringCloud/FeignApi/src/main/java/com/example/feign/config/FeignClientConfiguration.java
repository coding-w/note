package com.example.feign.config;

import feign.Logger;
import org.springframework.context.annotation.Bean;
public class FeignClientConfiguration {
    @Bean
    Logger.Level feignLoggerLevel() {
        return Logger.Level.FULL; // 设置为 FULL，记录详细日志
    }
}
