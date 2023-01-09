## SOLID Principle

```
S - Single Responsibility
O - Open / Closed Principle
L - Liskov Sunstitution Principle
I - Interface Segmented Principle
D - Dependency Inversion Principle
```

### Advantages:
These principles helps in following:
1. Easy to mainatain
2. Easy to understand

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