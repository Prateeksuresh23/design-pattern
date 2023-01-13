## Strategy pattern


![Alt text](example1.png)

In above example there is code duplicate as heavyVeh and sportsVeh uses same drive logic.

To solve this problem we can use strategy pattern

![Alt text](example2.png)

Using this pattern if in future any other new class require some new drive logic then new drive class can be added which implements DriveInterface().