# Encapsulation
In short, it is controlling access to information of an object. For example, a human might not expose his/her age to other objects.

## Java
In java we can control it using access modifier:
### 1.public:
Methods and attributes that use the public modifier can be accessed within your current class and by all other classes.

### 2.protected:
Attributes and methods with the access modifier protected can be accessed within the class, by all other classes within the same package, and by all subclasses within the same or other packages.

### 3.private:
This is the most restrictive and commonly used access modifier. A private attribute or method can only be accessed within the same class. Subclasses or any other classes within the same or a different package can’t access a private attribute or method.

### 4.No modifier(default):
What if we do not provide any modifier at all. In such cases a language-specific default is used. In java it is only within the package.

## Golang
There are no access modifiers in Go. All we have are exported and unexported fields. Any method or field whose name start with an uppercase letter are considered exported in Go, while lowercase are unexported.An exported field or function can be accessed “directly” from other packages, while unexported cannot be.

NOTE: Unexported fields cannot be accessed directly from other packages. This does not mean that we cannot access them at all. Here also we can create getter & setter method similar to java.
