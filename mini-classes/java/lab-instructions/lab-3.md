## Lab 2 - Java Basic Application (Continued)

- In this exercise you will add some JavaDoc comments into your application, subsequently generate the API doc by using `javadoc` command. 
- We encorage you to write the code without any IDE in order to grasp better understanding on how to compile the Java application by yourself.

Open the `Greeting.java` you created earlier, and do some changes as shown below:

```
/**
 * Class Greeting, respond with a warm greeting to user.
 * @author G2Lab
 * @version 1.0.0-RC1
 * @since 0.0.1
 * @see java.lang.Object
 */
public class Greeting {

  /**
   * main starting point
   * @param args shell parameters
   */
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

Create the Doc API by using `javadoc` command. e.g. : `javadoc Greeting.java`
