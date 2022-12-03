# Golang Password Generator.

As a member of the cybersecurity community. It is clear to me that one of the most basic, yet most important principles of cybersecurity is not being followed. The principle in mind is the usage of strong and original passwords for each account. According to a survey conducted by Google, as many as 65% of people reuse passwords. This practice of reusing passwords is partly the reason that almost 22% of date breaches in the corporate world are caused by compromised credentials any of these cases where cases of the same credentials being used in multiple places in the company or as credentials by the employees on personal accounts. For this reason, I'm looking to create a tool to make it easier for those who reuse passwords now to use strong, unique passwords.

# Requirements.

To be able to aid people that reuse passwords use unique strong passwords, my solution will need to be able to meet several requirements;

1. The solution should be easy to use.
2. The solution should be quick to use.
3. The solution should generate unique passwords.
4. The solution should generate password that are secure.
5. The solutions should be able to be ran on any modern computer operating system .
6. The solution should be able to run on a wide variety of hardware .

# Use Case Example.

1. Program is executed.
2. The program should **print** _“How many passwords do you want to generate?”._
3. The user then proceeds to give a number any other input will be returned with _"Please provide correct value for number of passwords”._
4. If the user did put in a number the generator will provide the user with text stating which password this is and then the password it generated for it.
5. The program will continue to generate passwords until the user has provided the number of passwords they want to generate.

# Why Golang?

Golang is efficient, scalable and easy to use. It is a language that is used by many companies and is a popular language for many developers. It allows me to create a program that is easy to use and easy to understand. Go code can be compiled on any modern computer operating system allowing my password generator to be used by anyone with a computer and without any additional software except for Golang. Using Golang I don't need to rely on an interpreter to run my code unlike other languages such as python or ruby. Golang uses a compiler which is particularly fast allowing me to generate passwords quickly. Go has native support for concurrent operations this means that I can generate passwords in parallel and the program will not be slowed down by the use of multiple threads.

# Chosen Programming Paradigms.

In my code I utilize **procedural programming** which is a programming paradigm that uses a series of steps to solve a problem. In the case of my password generator I use a series of statements to determine which password to generate. Each statement is a step in the process of generating a password. The statements are executed in the order they are written in the code. This is known as procedural programming.

My code also utilizes **functional programming** as I use a series of functions to determine which part of the password to generate. This is known as functional programming. using functional programming I can generate a password in parallel and the program will not be slowed down by the use of multiple threads. It also allows me to establish a clear and logical structure to my code making it easier to modify and maintain int the future.

# Design Decisions.

The design decisions I made for my password generator are as follows: The use of global variables and functions, The minimization of usage of packages.

the usage of global variables in this project is to ensure that the program is as fast as possible. And to ensure that the program is easy to read and understand this also helps to ensure that the program is easy to modify and extend if needed. throughout the program I use global variables to store the number of passwords to generate, the length of the password, the number of upper case and lower case characters and the number of numbers in the password. Adding each section and characteristic of the password generation to its own variable allows me to easily modify the password generation process. This can be seen by looking at the code. If we refer back to the section of code from line numbers 24-33 we can see all the variables are used to hold the characteristics to generate the password. If in the future or for a specific use case I need to change the characteristics of the password I can easily change the values of the variables and the program will generate the correct password to match those new characteristics. Using the variables that are present in my code we can really easily identify the characteristics of the password that we will be generating.

During the design process I made sure to use as minimal amount of packages and dependencies as possible. This was important because the program was intended to be usable by anyone with a computer and without any additional software except for Golang. Having too many packages and dependencies would make the program compile slowly and be more demanding to to run wile also making it more difficult to maintain the program.

The small amount of packages I did use caused me to use less memory and CPU time than I would have if I had used more packages. This was because the packages I used were small and simple and therefore the program would run faster and more efficiently. This made the program more resilient to cyber security threats as it creates a way smaller attack vector. As seen on lines 14-21 the entire program only uses six packages. All the packages used are well known well maintained well documented packages this helps to ensure the security of the program. Which is crucial as the program is intended to generate cryptographically secure passwords.

# Programming concepts.

The programming concepts I would use in my password generator if I where to start over would be as follows: Abstraction & Inheritance. Both of these programming concepts are part of Object Oriented Design and are used in many of the most popular programming languages. Abstraction is used to hide the implementation details of a program and inheritance is used to create a new class that inherits from an existing class.

Abstraction as a concept allows me to hide the implementation details of a program allowing me to create more easily maintainable programs. Not only are these programs more maintainable they are also easier to validate as a user this ensures that the user is able to validate the program and understand what it does before they trust it to generate them cryptographically secure passwords. Abstraction also makes it easier for other developers to add new features to the program, this helps to ensure the serviceability of the program past the initial release and the support of the initial development team.

Inheritance plays a very crucial part in the object oriented design workflow of programming as without it the different levels of abstraction would cause nothing but confusion and unnecessary fragmentation to the code. Inherence allows us to create classes that inherent information from pre existing classes this means that information or code only needs to be modified and stored in only location but it will still be able to be used as a global variable for the other classes in the program.

# Bibliography.

Anon (2016) 22 Percent of Data Breaches Are Caused by Compromised Credentials [online]. Available from: https://www.esecurityplanet.com/networks/data-breaches-are-caused-by-compromised-credentials/ (Accessed 3 August 2022).

Zurkus, K. (2019) Google Survey Finds Two in Three Users Reuse Passwords [online]. Available from: https://www.infosecurity-magazine.com/news/google-survey-finds-two-users/ (Accessed 3 August 2022).
