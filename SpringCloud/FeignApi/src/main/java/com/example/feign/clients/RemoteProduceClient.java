package com.example.feign.clients;

import com.example.feign.config.FeignClientConfiguration;
import com.example.feign.factory.RemoteProduceFallbackFactory;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/17 21:05
 */
@FeignClient(contextId = "remoteProduceClient", name = "producer-application",
        configuration = FeignClientConfiguration.class,
        fallbackFactory = RemoteProduceFallbackFactory.class)
public interface RemoteProduceClient {
    @GetMapping("/produce/get")
    String get();
}
