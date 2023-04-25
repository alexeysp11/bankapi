# Todo 

## How to test rest api: 

1. Create several http clients and send requests to the server simultaniously and 
sequentially. 
The following input parameters could vary: 
- the number of http clients (this could be calculated taking into an account the number of 
atms, branches, eftpos, active users of web and mobile app, cities and countries); 
- number of requests, that each of the clients make in a period of time;
- delay between requests (for each client and in general). 
how to evaluate if the server works correctly: 
- no dead-locks on the server while it's processing multiple requests; 
- each client should finally get response on its request (response time should be less than 
timeout value); 
- load balancer should work well (it could be implemented in a way that allows to process 
requests from particular city/country separately, or divide users by their names 
alphabetically; some sort of a database also could be implemented on a load balancer). 

Potential issues: 
- what to do, if a client sends two requests in a very short time. 

In order to implement test http client, use `client.go` (this means that `client.go` should be able to get url of a server either from config file or from `Init()` function). 
