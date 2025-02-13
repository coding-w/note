package com.example.nacosdiscovery;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/20 11:40
 */
import java.security.*;
import java.io.*;
import java.util.Base64;

public class KeyPairGeneratorExample {
    public static void main(String[] args) {
        try {
            // 创建一个 KeyPairGenerator 实例，指定算法为 RSA
            KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");

            // 初始化密钥生成器，密钥大小设为 2048 位
            keyPairGenerator.initialize(8);

            // 生成密钥对
            KeyPair keyPair = keyPairGenerator.generateKeyPair();

            // 获取公钥和私钥
            PublicKey publicKey = keyPair.getPublic();
            PrivateKey privateKey = keyPair.getPrivate();

            // 将公钥和私钥编码为 Base64 格式，方便存储或传输
            String encodedPublicKey = Base64.getEncoder().encodeToString(publicKey.getEncoded());
            String encodedPrivateKey = Base64.getEncoder().encodeToString(privateKey.getEncoded());



            System.out.println("密钥对生成成功，公钥和私钥已保存到文件");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

