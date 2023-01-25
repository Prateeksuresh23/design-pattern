## SOLID Principle

```
S - Single Responsibility
O - Open / Closed Principle
L - Liskov Substitution Principle
I - Interface Segmented Principle
D - Dependency Inversion Principle
```

### Advantages:
These principles helps in following:
1. Easy to mainatain
2. Easy to understand
3. Avoid duplicate code
4. Reduce complexity
5. Flexible software

## Details:
### Single Responsibility Principle:

A class should have only 1 reason to change

```
Marker Entity: 

class Marker{
    String color,name;
    int price,year;
    public Marker(int price, String color, String name){
        this.name = name;
        this.color = color;
        this.price = price;
    }
}
```
Marker and invoice 'has a' relationships.
```
class Invoice{
    private Marker marker;
    private int quantity;
    public Invoice(Marker marker, int quantity){
        this.marker = marker;
        this.quantity = quantity;
    }

    public int calculateTotal(){
        return marker.price * this.quantity;
    }

    public void printInvoice(){
        // print logic
    }

    public void saveToDB(){
        // DB logic
    }

}
```

This example is not following single responsibility principle as class Invoice will be changed due to multiple reason like change in calculateTotal() logic (ex: include GST also) or change in DB logic (ex: change to save it in files )

Changes to follow S principle:

```
class Invoice{
    private Marker marker;
    private int quantity;
    public Invoice(Marker marker, int quantity){
        this.marker = marker;
        this.quantity = quantity;
    } 

    public int calculateTotal(){
        return marker.price * this.quantity;
    }
}

class InvoiceDao{
    private Invoice invoice;
    public InvoiceDao(Invoice invoice){
        this.invoice = invoice;
    }
    public void saveToDB(){
        // DB logic
    }
}


class InvoicePrinter{
    private Invoice invoice;
    public InvoiceDao(Invoice invoice){
        this.invoice = invoice;
    }
    public void printInvoice(){
        // print logic
    }
}
```

This principle help in maintaing and understanding code.

### Open / Closed principle:
It states that open for extension but closed for modification

Example:
```
class InvoiceDao{
    private Invoice invoice;
    public InvoiceDao(Invoice invoice){
        this.invoice = invoice;
    }
    public void saveToDB(){
        // DB logic
    }
}
```

Suppose above class is tested and running in production. In future there is requirement to add capability to save data in file also.

```
interface InvoiceDao{
    public void save()
}

class DbInvoiceDao implements InvoiceDao{
    @override
    public void save(){
        // logic to save in DB
    }
}

class FileInvoiceDao implements InvoiceDao{
    @override
    public void save(){
        // logic to save in file system
    }
}
```

### Liskov Substitution Principle

If class B is subtype of class A then we should be able replace object of A with B without breaking the behaviour of program.

Subclass should extend the capability of parent class but not narrow it down

Example:

```
interface Bike{
    public void turnOnEng();
    public void increaseSpeed();
}

class MotorCycle implements Bike{
    public void turnOnEng(){
        this.isEngineOn = true;
    }

    public void increaseSpeed(){
        this.speed = this.speed+10;
    }
}

class Bicycle implements Bike{
    public void turnOnEng(){
        throw new AssertionError("there is no engine");
    }

    public void increaseSpeed(){
        this.speed = this.speed+10;
    }
}
```

### Interface segmented principle

Interface should be such that client should not implement unnecessary function they do not need

Example:

Not following:
```
interface ResturantEmployee{
    public void takeOrder();
    public void cookFood();
    public void washDishes();
}

class Waiter implements ResturantEmployee{
    public void cookFood(){
        // Not my job
    }
    public void takeOrder(){
        System.out.println("taking order");
    }
    public void washDishes(){
        // Not my job
     }
}
```

Chages:
```
interface WaiterInterface{
    public void takeOrder();
    public void serveCustomer;
}

interface cookInterface{
    public void cookFood();
}

interface Helper{
    public void washDishes();
}
```

### Dependency Inversion principle

Classes should depend on interfaces rather than conrete classes

Example:
Interface Keyboard --> two classes WiredKeyBoard and BluetoothKeyboard
Interface Mouse    --> two classes WiredMouse and BluetoothMouse

```
class Mac{
    private final WiredKeyBoard keyboard;
    private final BluetoothMouse mouse;
    public Mac(){
        keyboard = new WiredKeyBoard();
        mouse = new BluetoothMouse();
    }
}
```

Class mac can have only wirekeyboard and blutoothmouse

changes:
```
class Mac{
    private final KeyBoard keyboard;
    private final Mouse mouse;
    public Mac(KeyBoard keyboard, Mouse mouse){
        keyboard = this.keyboard;
        mouse = this.mouse;
    }
}
```

Reference:
https://www.youtube.com/watch?v=XI7zep97c-Y&list=PL6W8uoQQ2c61X_9e6Net0WdYZidm7zooW
