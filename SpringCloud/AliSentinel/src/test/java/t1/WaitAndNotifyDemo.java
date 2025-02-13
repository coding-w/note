package t1;

/**
 * @author wangxiang
 * @description
 * @create 2025/2/12 16:39
 */
class MyThread extends Thread {

    public void run() {
        synchronized (this) {
            System.out.println("before notify");
            System.out.println("notify " + Thread.currentThread().getName());
            notify();
            System.out.println("after notify");
        }
    }
}

public class WaitAndNotifyDemo {
    public static void main(String[] args) throws InterruptedException {
        MyThread myThread = new MyThread();
        synchronized (myThread) {
            try {
                myThread.start();
                // 主线程睡眠3s
                Thread.sleep(3000);
                System.out.println("main " + Thread.currentThread().getName());
                System.out.println("before wait");
                // 阻塞主线程
                myThread.wait();
                System.out.println("after wait");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}