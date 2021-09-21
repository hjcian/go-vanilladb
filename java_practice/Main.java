package java_practice;

import java.util.concurrent.TimeUnit;

class MyThread extends Thread {
    public static void main(String[] args) {
        System.out.println("main in thread");
    }

    @Override
    public void run() {
        while (true) {
            System.out.println("[thread] Run");
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (Exception e) {
                Thread.currentThread().interrupt();
            }
        }
    }
}

class Main {
    public static void main(String[] args) {
        System.out.println("Hello World");
        String name = "max";
        System.out.println(name.indexOf('a'));

        MyThread t1 = new MyThread();
        t1.start();
    }
}