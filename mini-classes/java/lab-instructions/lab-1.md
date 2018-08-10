## Lab 1 - Java Basic Application

- In this exercise you will create a simple Greeting application with Java. This lab assume you have no prior knowledge in Java, hence it is designed to show you how's the basic structure of Java main class looks like.  If you feel you already have a good understanding of it, you can feel free to skip this exercise.
- We encorage you to write the code without any IDE in order to grasp better understanding on how to compile the Java application by yourself.

**Part 1 - Workspace Preparation**

1. Before you start any lab on this mini-class, we suggest you to create a new directory to be used as your lab workspace. e.g. : `mkdir java_mini_class`, or name it whatever you'd like to.
2. Enter the directory you've just created on the aforementioned step. 
3. Create a new folder inside it, name it `lab-1`, then `cd` into it.


**Part 2 - Write Code**

From within `lab-1` folder, create a **java** source file with the following code:

```
public class Greeting {

  public static void main(String... args) {
    System.out.println("Hi.., welcome to G2Lab Java MiniClass.");
}
```
  
**Part 3 - Run the Program**
 
Compile using `javac` and run using `java` from your terminal. 

**Part 4 - Explore Further**

Explore with various options for `javac` command. For instance, you can create your source code file under `src` folder, and try to compile/produce the output class inside another folder like `bin`.
