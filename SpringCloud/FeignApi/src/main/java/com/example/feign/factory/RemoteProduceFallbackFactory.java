package com.example.feign.factory;

import com.example.feign.clients.RemoteProduceClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.openfeign.FallbackFactory;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/17 21:13
 */
public class RemoteProduceFallbackFactory implements FallbackFactory<RemoteProduceClient> {

    Logger logger = LoggerFactory.getLogger(RemoteProduceFallbackFactory.class);


    @Override
    public RemoteProduceClient create(Throwable cause) {
        return new RemoteProduceClient() {
            @Override
            public String get() {
                logger.error("get 错误", cause);
                return "请稍后";
            }
        };
    }
}
