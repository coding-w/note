# Bean 生命周期

Bean的创建过程步骤比较多，这里结合代码一起汇总一下

1. 通过 BeanDefinition 获取 bean 的定义信息

2. 调用构造函数实例化 bean

3. Bean 的依赖注入

4. 处理 Aware 接口（BeanNameAware、BeanFactoryAware、ApplicationContextAware）

5. Bean 的前置处理器 BeanPostProcessor - 前置（postProcessBeforeInitialization）

6. 初始化方法（InitializingBean、init-method）

7. Bean 的后置处理器 BeanPostProcessor - 后置（postProcessAfterInitialization）

8. 销毁 bean（DisposableBean、destroy-method）

牢记执行流程，流程如下

![](../../images/springcloud/Bean生命周期.jpg)


结合代码查看Bean的创建过程
```java
// User.java Bean 创建
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
```

```java
// MyBeanPostProcessor.java bean 后置处理器
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
```

```java
// Main.java 启动类
@SpringBootApplication
public class Main {
    public static void main(String[] args) {
        ConfigurableApplicationContext applicationContext = SpringApplication.run(Main.class, args);
        User bean = applicationContext.getBean(User.class);
        System.out.println(bean);
    }
}
```


```text
1. 调用了构造函数
2. 依赖注入
3. setBeanName 执行了 user
4. setBeanFactory 执行了 org.springframework.beans.factory.support.DefaultListableBeanFactory
5. setApplicationContext 执行了 
6. postProcessBeforeInitialization user = com.example.config.User@688d411b
7. init() 方法执行了
8. afterPropertiesSet() 方法执行了
9. postProcessAfterInitialization user -> 对象方法开始增强
10. destroy() 销毁方法执行了
```