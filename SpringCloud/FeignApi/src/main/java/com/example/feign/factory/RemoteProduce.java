package com.example.feign.factory;

import com.example.feign.clients.Produce;
import org.springframework.cloud.openfeign.FallbackFactory;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/17 21:13
 */
public class RemoteProduce implements FallbackFactory<Produce> {
    @Override
    public Produce create(Throwable cause) {
        return new Produce() {

        };
    }
}
