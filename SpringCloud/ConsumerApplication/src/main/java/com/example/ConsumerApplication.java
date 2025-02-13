package com.example;

import com.example.config.FeignClientConfiguration;
import com.example.config.LoadBalancerConfig;
import feign.Logger;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.cloud.loadbalancer.annotation.LoadBalancerClients;
import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;

@SpringBootApplication
@EnableDiscoveryClient
@EnableFeignClients
// org.springframework.cloud.loadbalancer.core.RoundRobinLoadBalancer
// com.alibaba.cloud.nacos.loadbalancer.NacosLoadBalancer
// @LoadBalancerClient(name = "org.springframework.cloud.loadbalancer.core.RoundRobinLoadBalancer")
@LoadBalancerClients(defaultConfiguration = LoadBalancerConfig.class)
public class ConsumerApplication {
    public static void main(String[] args) {
        SpringApplication.run(ConsumerApplication.class, args);
    }

    @Bean
    @LoadBalanced
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }


}