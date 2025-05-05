package com.example.config;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.BeanFactory;
import org.springframework.beans.factory.BeanFactoryAware;
import org.springframework.beans.factory.BeanNameAware;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;

/**
 * @author wangxiang
 * @description
 * @create 2025/5/5 18:27
 */
@Component
public class User implements BeanNameAware, BeanFactoryAware, ApplicationContextAware, InitializingBean {

    private String name;

    public User(){
        System.out.println("1. 调用了构造函数");
    }

    @Value("test")
    public void setName(String name){
        this.name = name;
        System.out.println("2. 依赖注入");
    }

    @Override
    public void setBeanName(String name) {

        System.out.println("3. setBeanName 执行了 " + name);

    }

    @Override
    public void setBeanFactory(BeanFactory beanFactory) throws BeansException {
        System.out.println("4. setBeanFactory 执行了 " + beanFactory.getClass().getName());
    }



    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        System.out.println("5. setApplicationContext 执行了 " + applicationContext.getApplicationName());
    }

    @PostConstruct
    public void init(){
        System.out.println("7. init() 方法执行了");
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        System.out.println("8. afterPropertiesSet() 方法执行了");
    }

    @PreDestroy
    public void destroy(){
        System.out.println("10. destroy() 销毁方法执行了");
    }
}
