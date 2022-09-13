### OOP:
#### Formal definition:
    Object-oriented programming (OOP) is a programming paradigm based on the concept of “objects”, which can contain data, in the form of fields (often known as attributes or properties), and code, in the form of procedures (often known as methods).
#### Simpler definition:
     In OOP everything in the world can be classified into classes- for example, you and I are objects of the class “human”. Every Human has two types of features: state (fields: height,weight etc) and behavior (methods: breathing,eating etc). Similarly, an Object in the OOP paradigm has attributes and methods

### Object
This section describe how object look in java and golang using above 'human' example, 

#### JAVA example:

```
class Human{
    //fields/state
    int weight;
    int height;
    int age;

    //behaviour/method
    int breath() {
            printf("A human of age %d is breathing", age);
    }
}

int main(){
    Human h = new Human("65","165","25");
    System.out.println(h.breath());
}

```

#### GOLANG example:
In Go, data and functions are kept separate. There are no classes in Go, just structs which can only store data and do not have any methods. But we can achieve methods using receivers (notice the () block right after func).

```

type Human struct{
    Height  int
    Weight  int
    Age     int
}

func (h *Human) Breath(){
    fmt.Print("")
}
func main() {
    me := Human{ Height: 175, Weight: 70, Age: 22 }
    me.Breath()
}

```