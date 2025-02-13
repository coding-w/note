package t2;

/**
 * @author wangxiang
 * @description
 * @create 2025/2/12 16:42
 */
public class PrintNumDemo {
    public static Integer num = 0;

    public static void main(String[] args) {
        PrintOddNum oddNum = new PrintOddNum();
        PrintEvenNum evenNum = new PrintEvenNum();

        oddNum.start();  // 启动奇数线程
        evenNum.start(); // 启动偶数线程
    }
}

class PrintOddNum extends Thread {
    public void run() {
        synchronized (PrintNumDemo.class) {  // 使用类名来同步
            try {
                while (PrintNumDemo.num < 10) { // 控制打印的数值范围
                    // 当是偶数时，等待偶数线程打印
                    if (PrintNumDemo.num % 2 == 0) {
                        PrintNumDemo.class.wait();  // 等待偶数线程打印
                    } else {
                        System.out.println(PrintNumDemo.num);
                        PrintNumDemo.num++; // 打印后加 1
                        PrintNumDemo.class.notify(); // 通知偶数线程打印
                    }
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

class PrintEvenNum extends Thread {
    public void run() {
        synchronized (PrintNumDemo.class) {  // 使用类名来同步
            try {
                while (PrintNumDemo.num < 10) { // 控制打印的数值范围
                    // 当是奇数时，等待奇数线程打印
                    if (PrintNumDemo.num % 2 != 0) {
                        PrintNumDemo.class.wait();  // 等待奇数线程打印
                    } else {
                        System.out.println(PrintNumDemo.num);
                        PrintNumDemo.num++; // 打印后加 1
                        PrintNumDemo.class.notify(); // 通知奇数线程打印
                    }
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
