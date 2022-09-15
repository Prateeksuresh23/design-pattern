# Inheritance
The idea behind inheritance is that we can create new classes that are built upon existing classes. When we inherit from an existing class, we can reuse methods and fields of the parent class. Moreover, we can add new methods and fields in our current class also.

## Java
Encapsulation role:
1. Public methods of the parent class are also public in the subclass
2. Methods declared protected in a superclass must either be protected or public in subclasses; they cannot be private
3. Methods declared private are not inherited into the child class at all, so there is no rule for them

```
class Human {
    public int height;
    public int weight;
    public int age;
     
    public void Breathing() {
       System.out.println("Human is breathing");
    }
     
    public void Eating() {
       System.out.println("Human is eating");
    }
 }
 
public class Doctor extends Human {
    public String nameOfHospital;
     
    public void Working() {
       System.out.println("A Human who is a Doctor, is saving lives.");
    }
 }
```

## Golang
Go does not support inheritance, instead it does composition. This means, putting together all the fields. So, when I define a “Doctor” struct, I can simple put together fields from Human struct and add more

```
type Human struct {
    weight int
    height int
    age int
}

func(h *Human) breathing() {
    fmt.Println("A Human is Breathing")
}

type Doctor struct {
    Human
    nameOfHospital string
}

func(d *Doctor) working() {
    fmt.Println("A doctor is saving lives")
}

func main() {
    doctorA := Doctor{
        Human: Human{175, 70, 22},
        nameOfHospital: "HospitalA",
    }

    doctorA.Human.breathing()
    doctorA.working()
}

// OUTPUT
> A Human is Breathing
> A doctor is saving lives
```

We can also name the inherited struct if we want.
```
type Doctor struct {
    human Human
    nameOfHospital string
}
```