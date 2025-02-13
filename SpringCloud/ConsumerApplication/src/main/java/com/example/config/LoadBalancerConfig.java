package com.example.config;

import com.alibaba.cloud.nacos.NacosDiscoveryProperties;
import com.alibaba.cloud.nacos.loadbalancer.NacosLoadBalancer;
import com.alibaba.nacos.api.config.annotation.NacosConfigListener;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.loadbalancer.core.RandomLoadBalancer;
import org.springframework.cloud.loadbalancer.core.ReactorLoadBalancer;
import org.springframework.cloud.loadbalancer.core.RoundRobinLoadBalancer;
import org.springframework.cloud.loadbalancer.core.ServiceInstanceListSupplier;
import org.springframework.cloud.loadbalancer.support.LoadBalancerClientFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;

@Configuration
public class LoadBalancerConfig {

    // @Bean
    // ReactorLoadBalancer<ServiceInstance> randomLoadBalancer(Environment environment,
    //                                                         LoadBalancerClientFactory loadBalancerClientFactory,
    //                                                         NacosDiscoveryProperties nacosDiscoveryProperties) {
    //     // System.out.println(nacosDiscoveryProperties);
    //     // String name = environment.getProperty("spring.application.name");
    //     String name = "producer-application";
    //     return new NacosLoadBalancer(loadBalancerClientFactory
    //             .getLazyProvider(name, ServiceInstanceListSupplier.class),
    //             name, nacosDiscoveryProperties);
    // }

    // @Bean
    // ReactorLoadBalancer<ServiceInstance> randomLoadBalancer(Environment environment,
    //                                                         LoadBalancerClientFactory loadBalancerClientFactory) {
    //     // String name = environment.getProperty("spring.application.name");
    //     String name = "producer-application";
    //     return new RandomLoadBalancer(loadBalancerClientFactory
    //             .getLazyProvider(name, ServiceInstanceListSupplier.class),
    //             name);
    // }

    @Bean
    ReactorLoadBalancer<ServiceInstance> roundRobinLoadBalancer(Environment environment,
                                                                LoadBalancerClientFactory loadBalancerClientFactory) {
        // String name = environment.getProperty("spring.application.name");
        String name = "producer-application";
        return new RoundRobinLoadBalancer(loadBalancerClientFactory
                .getLazyProvider(name, ServiceInstanceListSupplier.class),
                name);
    }
}
