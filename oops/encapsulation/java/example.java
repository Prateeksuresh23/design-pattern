//A Java class which is a fully encapsulated class.  
//It has a private data member and getter and setter methods.  
package com.encapsulation; 

public class Student{  
    //private data member  
    private String name;
    /*
    we cannot modify college of a student because the attribute is private and 
    there is no method that will help us modify it
    */ 
    private String college = "IIT BHU";  

    // method to get a student's name 
    public String getName(){  
        return name;  
    }  

    // method to set a student's name 
    public void setName(String name){  
    this.name=name; 
    }
     
}