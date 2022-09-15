# Polymorphism
Polymorphism is the ability of an object to take on many forms. In other words, in polymorphism behaviour may change according to the objects.

### Example:
Imagine we have a class Shape with two attributes: area and perimeter. The class also has two methods calculateArea() and calculatePerimeter(). Way we calculate area for different shapes is different. For a square it is side*side and for a circle it is pi * radius * radius.This means that a single calculateArea() method will not be enough to calculate the area for any shape we provide. This is where polymorphism comes into picture

## Java

```
// Online Java Compiler
// Use this editor to write, compile and run your Java code online

interface shape{
    public double area();
    public double perimeter();
}

class square implements shape{
    int side;
    public square(int side){
        this.side = side;
    }
    public double area(){
        return side*side;
    }
    public double perimeter(){
        return 4*side;
    }
}
class rectangle implements shape{
    int l,b;

    public rectangle(int l,int b){
        this.l = l;
        this.b = b;
    }
    public double area(){
        return l*b;
    }
    public double perimeter(){
        return 2*(l+b);
    }
}
class HelloWorld {
    public static void main(String[] args) {
       rectangle r = new rectangle(2,4);
       square s = new square(2);
       measure(r);
       measure(s);
    }
    public static void measure(shape s){
        System.out.println(s.area());
        System.out.println(s.perimeter());
    }
}

/* Output:
8.0
12.0
4.0
8.0
*/
```

## Golang

```
package main

import (
    "fmt"
    "math"
)

// Our interface with the two methods
type Shape interface {
    area() float64
    perimeter() float64
}

type Square struct {
    side float64
}

type Circle struct {
    radius float64
}

func (s Sqaure) area() float64 {
    return r.side * r.side
}

func (s Square) perimeter() float64 {
    return 4 * s.side
}

func (c Cirlce) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func measure(s Shape) {
    fmt.Println(s)
    fmt.Println(s.area())
    fmt.Println(s.perimeter())
}

func main() {
    s := Sqaure{side: 3}
    c := Circle{radius: 5}

    measure(s) // This will use functions for the square
    measure(c) // This will use functions for the circle
}

```