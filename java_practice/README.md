- [Java](#java)
  - [Java dev environment setup](#java-dev-environment-setup)
  - [FAQ](#faq)
    - [Why use capital letter in String?](#why-use-capital-letter-in-string)
    - [How to declare constant variable?](#how-to-declare-constant-variable)
    - [How to change language of error message?](#how-to-change-language-of-error-message)
    - [main function? file name?](#main-function-file-name)
    - [Differences between "`extends`" and "`implements`](#differences-between-extends-and-implements)
    - [Thread/Concurrency related](#threadconcurrency-related)
      - [How to do thread-safe access?](#how-to-do-thread-safe-access)
      - [How to use Wait/Notify mechanism? (kind of Go's `sync.Cond`)](#how-to-use-waitnotify-mechanism-kind-of-gos-synccond)


## Java
### Java dev environment setup
- install [Java Extension Pack](https://code.visualstudio.com/docs/java/java-tutorial#_installing-extensions) for existing VS Code user

### FAQ
#### Why use capital letter in String?
- 小寫的是 primitive type，大小字母開頭的是 class type
- class type 擁有更多方法做更多事
- e.g. [Boolean](https://docs.oracle.com/javase/6/docs/api/java/lang/Boolean.html) have method to convert `boolean` to `String`, or `String` to `boolean`
- [ref stackoverflow](https://stackoverflow.com/a/4006311/8694937)

#### How to declare constant variable?
- use `final`

#### How to change language of error message?
> https://stackoverflow.com/questions/35581227/how-to-set-java-file-encoding-and-java-default-locale-in-linux-centos-server
- the java (jvm) is use default locale as default language
- change it via passing `-Duser.language=<lang>` to JVM command line
  - `java -Duser.language=en Main.class`
- or add `JAVA_TOOL_OPTIONS` to bashrc
  - `export JAVA_TOOL_OPTIONS="-Duser.language=en"`

#### main function? file name?

> https://www.geeksforgeeks.org/myth-file-name-class-name-java/
>
> https://stackoverflow.com/a/968360/8694937

- following java code will get compiling error (when doing `javac Trial.java`):
  - `error: class Geeks is public, should be declared in a file named Geeks.java`
```java
/***** File name: Trial.java ******/
public class Geeks {
    public static void main(String[] args)
    {
        System.out.println("Hello world");
    }
}
```
- because every single java file ***can only have one public class*** and must be declared in a file witch filename is same as class name.
- 而 compile 完之後，會製作很多 class file，就是每一個 class 的 byte code
- 使用 `java` 執行這些 file，就會去執行該 class 裡頭的 `public static void main(String[] args) {}` 來執行
  - 若被執行的 class file 中的 class 沒有定義 `public static void main(String[] args) {}`，會報錯：

```shell
Error: Main method not found in class FooTest, please define the main method as:
   public static void main(String[] args)
or a JavaFX application class must extend javafx.application.Application
```

#### Differences between "`extends`" and "`implements`
> https://www.geeksforgeeks.org/extends-vs-implements-in-java/
>
- 首先，這是在 java 中實作「繼承」的語法
- `extends`
  - 指定此 derived class 繼承自某個 base class
  - JAVA 不允許一個 class 繼承自多個 classes，故 `extends` 只允許一個
- `implements`
  - 用來實作 interface 用的
  - JAVA 中的 `interface` 是特別的 class，其中只包含抽象的方法定義 (abstract methods)
  - JAVA 中就允許 class 實作多個 `interface` 了
- So, a class can ***extend a class*** and ***can implement any number of interfaces*** simultaneously.
- `interface`
  - 然而，`interface` 可以 `extends` 多個 interfaces


#### Thread/Concurrency related
##### How to do thread-safe access?
- use `synchronized`, to ensure only one thread can access this critical section. e.g.
- 讓某方法同一時間只能被一條 thread 呼叫：
  ```java
  public synchronized int get() {
    ...
  }
  ```
- 讓某個物件同一時間只能被一條 thread 存取：
  ```java
  private Counter counter = ...
  public int increment {
    synchronized(counter) {
      int c = counter.get();
      c++; // or c--
      counter.set(c)
    }
  }
  ```

##### How to use Wait/Notify mechanism? (kind of Go's `sync.Cond`)
- `wait()` - 目前進入 critical section 的 thread 會歸還 lock，並且等著被喚醒
- `notify()` - Wakes up a single thread that is waiting on this object's monitor. 隨機選一條 thread，改變狀態成 BLOCKED ，其餘等待中的 threads 仍然是 WAITING 狀態
- `notifyAll()` - Wakes up all threads that are waiting on this object's monitor. 所有等待中的 threads 都變成 BLOCKED 狀態
- 使用時機取決於你希望某個資源