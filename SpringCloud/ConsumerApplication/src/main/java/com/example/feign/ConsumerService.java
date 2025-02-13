package com.example.feign;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/15 14:58
 */
@Service
@FeignClient("producer-application")
// @LoadBalancerClient(name = "org.springframework.cloud.loadbalancer.core.RoundRobinLoadBalancer")
public interface ConsumerService {
    @GetMapping("/produce")
    String produce(@RequestParam("name") String name);
}
