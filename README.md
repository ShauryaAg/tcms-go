# tcms-go

> Note: Its still a work-in-progress, however the proof of concept is ready.

A test case management tool written in Go. Following all the SOLID properties.

This project is built in a way that anything can be plugged in, and replaced with something else.

The current **database** used is MongoDB, but can easily be replaced with any other database, even with a relational database.

As an example, we have used Chi as the router, but it can be replaced with Gorilla Mux easily (or any other router as one wishes).

### Why tho?
Why would anyone need to replace their router with something else?

Since the whole project is built in a way that any part of it can be replaced with something else, it is very easy to change the router, this is because the whole project is based on abstractions. Even the subrouter can be replaced with something else, that is, you can have multiple routers in the same project.

##### Ok, but WHY?
The point is not to have multiple routers in the same project, but combining multiple projects built using different frameworks into one at will, think of the whole project as a *modular monolith*.

![image](https://user-images.githubusercontent.com/31778302/179419422-3051e8be-62ba-48f0-a1da-095c657b2d5f.png)

Any number of services can be added or removed easily, and multiple microservices could be turned into a monolith. This can often save costs and runtime.