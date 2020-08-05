# Mocking with gomock in table-driven tests
In this tutorial I'll present our way of tackling the problem of writing precise, roboust and explicit unit tests with gomock package. Bear in mind, that's not *the only way* to do it.  


## Gomock
There are many ways to mock behaviour in golang, the basic one being just explicitly writing the returns you need. You can also save yourself some work and use a package for it. We chose [gomock](https://github.com/golang/mock) because it:
-   lets us precisely set the functions' invokations
-   is actively maintained 

### Handling the code generation
`mockgen` is a tool which generates gomocks based on given input interface. It has two modes: reflect and 
source. From user perspective, the main differences are:
-   source mode lets you generate unexported interfaces
-   reflect mode can be used in go:generate annotations
-   relefct mode 



## Table-driven tests