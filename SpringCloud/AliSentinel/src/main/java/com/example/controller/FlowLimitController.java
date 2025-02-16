package com.example.controller;

import com.alibaba.csp.sentinel.annotation.SentinelResource;
import com.example.feign.clients.RemoteProduceClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.io.BufferedReader;
import java.io.IOException;
import java.util.List;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/16 15:02
 */
@RestController
public class FlowLimitController {

    @Autowired
    private RemoteProduceClient remoteProduceClient;
    @Autowired
    private DiscoveryClient discoveryClient;

    @GetMapping("aaa")
    public void testServiceDiscovery() {
        List<ServiceInstance> instances = discoveryClient.getInstances("producer-application");
        if (instances.isEmpty()) {
            System.out.println("No instances found for producer-application");
        } else {
            for (ServiceInstance instance : instances) {
                System.out.println("Instance: " + instance.getUri());
            }
        }
    }


    @GetMapping("/testA")
    public String testA() {
        return "testA...";
    }


    @RequestMapping("/testB/**")
    public String testB(HttpServletRequest request) {
        System.out.println("params = " + request.getRequestURI());
        System.out.println("method = " + request.getMethod());
        System.out.println("query = " + request.getQueryString());
        // 获取请求体的 JSON 数据
        StringBuilder requestBody = new StringBuilder();
        try (BufferedReader reader = request.getReader()) {
            String line;
            while ((line = reader.readLine()) != null) {
                requestBody.append(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.out.println("requestBody = " + requestBody);
        return "testB...";
    }

    // @GetMapping("/testB/test/{params}")
    // public String testC(@PathVariable String params){
    //     System.out.println("params = " + params);
    //     return "testB...";
    // }


    @GetMapping("/query")
    public String query() {
        return "query...";
    }

    @PostMapping("/update")
    public String update() {
        return "update...";
    }


    @SentinelResource("hot")
    @GetMapping("/{id}")
    public String paramsLimit(@PathVariable Integer id) {
        System.out.println("id = " + id);
        return remoteProduceClient.get();
        // return "Asd";
    }

}