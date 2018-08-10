## Lab 2 - Java Basic Application (Continued)

- In this exercise you will create a simple Greeting application with Java. This lab assume you have no prior knowledge in Java, hence it is designed to show you how's the basic structure of Java main class looks like.  If you feel you already have a good understanding of it, you can feel free to skip this exercise.
- We encorage you to write the code without any IDE in order to grasp better understanding on how to compile the Java application by yourself.

Open the `Greeting.java` you created earlier, and do some changes as shown below:

```
public class Greeting {

  public static void main(String... args) {
    if (args.length == 0) {
      System.err.println("Usage: java Greeting <your_name>");
      System.err.println("e.g. : java Greeting Dude");
      System.exit(-1);
    }
    System.out.println("Hi " + args[0] + ", welcome to G2Lab Java MiniClass.");
  }
}
```
Recompile, and run again.
