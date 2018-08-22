# Oyster-Card

**Oyster-Card** represents a limited version of London's Oyster card system. 

## Prerequisites

To ensure the project successfully imported to your local, make sure that you're using these following stuffs:

* Java v8 or higher
* Maven v3 or higher

## Usage 

### Test

From within oyster-card folder, run `mvn test`.
You can check the test coverage by running `mvn jacoco:report` command. It will produce HTML reports in the `${project_dir}/target/site/jacoco/` folder.

### Build and Run

You can run the app with one of these following ways:
* `mvn spring-boot:run -Drun.jvmArguments="-Dspring.profiles.active=dev"` , or
* Build the package via `mvn clean package` command, subsequently run the app by using `java -jar -Dspring.profiles.active=dev target/oyster-card-1.0.0.jar` command.

## Authors

* Yauri Attamimi ([yauritux@gmail.com](mailto://yauritux@gmail.com)) - *Initial Committer* - [Profile](https://dzone.com/users/366249/yauritux.html)