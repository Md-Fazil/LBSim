# SimpleLoadBalancer

## Description

This project is a load balancer simulator that was developed to understand the internal workings of load balancers and the various load balancing algorithms.
The program reverse proxies HTTP requests to healthy servers chosen based on the specified algorithm 
and conducts health checks of servers in the server pool at regular intervals.

The following load balancing algorithms are currently supported:
* Least Connection
* Round Robin

## Executing the program 
1. Clone the program to your local machine.

2. Execute the following command on your terminal:
``` go run main.go servers=${servers} port=${port} algorithm=${algorithm}```
* ${servers} is a **comma separated list** of server URLs to send requests to
* ${port} is the port to run the load balancer on
* ${algorithm} is either **RoundRobin** or **LeastConnection**
