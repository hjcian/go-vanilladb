# database from scratch
- [database from scratch](#database-from-scratch)
  - [Java dev environment setup](#java-dev-environment-setup)
  - [FAQ](#faq)
    - [Why use capital letter in String?](#why-use-capital-letter-in-string)
    - [How to declare constant variable?](#how-to-declare-constant-variable)
    - [How to change language of error message?](#how-to-change-language-of-error-message)
    - [main function? file name?](#main-function-file-name)
    - [Differences between "`extends`" and "`implements`](#differences-between-extends-and-implements)
  - [Learning Materials](#learning-materials)
    - [清大課程-雲端資料庫](#清大課程-雲端資料庫)
    - [Golang Learning |【Go 夜读】#117 详解 rosedb 及存储模型](#golang-learning-go-夜读117-详解-rosedb-及存储模型)

## Java dev environment setup
- install [Java Extension Pack](https://code.visualstudio.com/docs/java/java-tutorial#_installing-extensions) for existing VS Code user

## FAQ
### Why use capital letter in String?
- 小寫的是 primitive type，大小字母開頭的是 class type
- class type 擁有更多方法做更多事
- e.g. [Boolean](https://docs.oracle.com/javase/6/docs/api/java/lang/Boolean.html) have method to convert `boolean` to `String`, or `String` to `boolean`
- [ref stackoverflow](https://stackoverflow.com/a/4006311/8694937)

### How to declare constant variable?
- use `final`

### How to change language of error message?
> https://stackoverflow.com/questions/35581227/how-to-set-java-file-encoding-and-java-default-locale-in-linux-centos-server
- the java (jvm) is use default locale as default language
- change it via passing `-Duser.language=<lang>` to JVM command line
  - `java -Duser.language=en Main.class`
- or add `JAVA_TOOL_OPTIONS` to bashrc
  - `export JAVA_TOOL_OPTIONS="-Duser.language=en"`

### main function? file name?

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

### Differences between "`extends`" and "`implements`
> https://www.geeksforgeeks.org/extends-vs-implements-in-java/
>
- 首先，這是在 java 中實作「繼承」的語法
- `extends`
  - 指定此 derived class 繼承自某個 parent class
...

## Learning Materials

### 清大課程-雲端資料庫
- CS 471000 Introduction to Database Systems - Implementation, architectural design, and trade-offs.
- https://nthu-datalab.github.io/db/

### Golang Learning |【Go 夜读】#117 详解 rosedb 及存储模型
- https://www.youtube.com/watch?v=Knh7EYfVIKs
- https://talkgo.org/t/topic/2386
