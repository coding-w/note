package com.example.config;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.BeanPostProcessor;
import org.springframework.cglib.proxy.Enhancer;
import org.springframework.cglib.proxy.InvocationHandler;
import org.springframework.stereotype.Component;

import java.lang.reflect.Method;

/**
 * @author wangxiang
 * @description
 * @create 2025/5/5 18:07
 */
@Component
public class MyBeanPostProcessor implements BeanPostProcessor {
    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException {
        if (bean instanceof User) {
            User user = (User) bean;
            System.out.println("6. postProcessBeforeInitialization user = " + user);
        }
        return bean;
    }

    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException {
        if (beanName.equals("user")) {
            System.out.println("9. postProcessAfterInitialization user -> 对象方法开始增强");
            // // cglib 代理对象
            // Enhancer enhancer = new Enhancer();
            // // 设置需要增强的类
            // enhancer.setSuperclass(bean.getClass());
            // // 执行会调方法，增强方法
            // enhancer.setCallback(new InvocationHandler() {
            //     @Override
            //     public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
            //         //  执行目标方法
            //         return method.invoke(bean, args);
            //     }
            // });
            // return enhancer.create();
        }
        return bean;
    }
}
