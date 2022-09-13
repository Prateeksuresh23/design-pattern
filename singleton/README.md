
1. Creational design pattermn
2. Only one instance of its kind exists
3. Provide single point of access to it

Where it will use?
1. There should be one DB instance use to make query
2. Single of logging variable

Other alternative?
We may have only one global instance of those. But then any function can change these and it will effect others also. Singleton pattern make sure there will be one instance of object at time and others caller can not update it.