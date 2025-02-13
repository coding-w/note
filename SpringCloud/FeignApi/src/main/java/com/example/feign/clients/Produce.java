package com.example.feign.clients;

import com.example.feign.config.FeignClientConfiguration;
import org.springframework.cloud.openfeign.FeignClient;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/17 21:05
 */
@FeignClient(value = "producer-application", configuration = FeignClientConfiguration.class)
public interface Produce {
    //  ...
}
