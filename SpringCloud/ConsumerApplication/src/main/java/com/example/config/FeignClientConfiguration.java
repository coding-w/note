package com.example.config;

import feign.Logger;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/17 15:10
 */
@Configuration
public class FeignClientConfiguration {
    @Bean
    Logger.Level feignLoggerLevel() {
        return Logger.Level.FULL; // 设置为 FULL，记录详细日志
    }
}
