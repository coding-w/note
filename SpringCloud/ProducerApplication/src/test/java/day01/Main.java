package day01;

import java.net.InetAddress;
import java.net.NetworkInterface;
import java.util.Enumeration;

/**
 * @author wangxiang
 * @description
 * @create 2025/2/7 10:05
 */
public class Main {
    public static void main(String[] args) {
        try {
            System.out.println(InetAddress.getLocalHost().getHostAddress());
            Enumeration<NetworkInterface> interfaces = NetworkInterface.getNetworkInterfaces();
            while(interfaces.hasMoreElements()) {
                NetworkInterface networkInterface = interfaces.nextElement();
                if (!networkInterface.isLoopback() && networkInterface.isUp()) {
                    Enumeration<InetAddress> addresses = networkInterface.getInetAddresses();
                    while(addresses.hasMoreElements()) {
                        InetAddress address = addresses.nextElement();
                        if (!address.isLoopbackAddress() && !address.isLinkLocalAddress() && !(address instanceof java.net.Inet6Address)) {
                            System.out.println("网络接口: " + networkInterface.getDisplayName() + " 的IP地址是：" + address.getHostAddress());
                        }
                    }
                }
            }
        } catch (Exception e) {
            System.err.println("无法获取网络接口信息。");
            e.printStackTrace();
        }
    }
}
