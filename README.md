# Calculator 
This project implements a simple calculator app, which can be interacted with through a REST interface.
The backend is coded in golang.

## Running locally
To run this locally:
1. Ensure golang is installed on your machine
2. Clone this repo to your machine
3. cd to `/main`
4. execute `go run .`
5. The service will now be running.
 
The service runs on port 8080 by default. If you wish to configure the service to run on another port, simply set your `PORT` environment variable to another port number.

## API
The basic concept of the API is that it let's you build and evaluate arithmetic expressions. That is, any expression that can be constructed from integers and the +, -, * and / operators.
The design philosophy has been to not let the user create any illegal expressions, meaning they can't create an expressions with a missing operator arguments, ex. '5 +' and they can't divide by zero. The benefit of this is that the expressions are always in a state, where they can be evaulated, which simiplifies the logic on the backend.

The REST API is documented by a postman collection, a JSON export of which can be found at the root this repository at `Calculator.postman_collection.json`.
To run the collection, you will need to create two variables in your postman environment `calculationId` and `baseUrl`.

`calculationId` holds the id of the calculation currently being worked on. It is automatically set, when you run the `Initialize Calculation` endpoint.

* If you wish to run the requests in the collection against your local service set `baseUrl=http://localhost:8080` in your environment.
* If you wish to run the requests in the collection against the currently deployed service set `baseUrl=https://helloworld-676480202164.europe-north1.run.app` in your environment.

### Future Extensions
The API is quite simple right now, but could be extended in the following ways given more time:
* Support for floats as operator arguments
* List all calculations assosicated with a user
* Support for unary operators like square root and pow
* Support for subexpressions, such as  `5 + (5 + 5) * 5 `. This could be implemented by allowing operators to take calculation ids as input in addition to integers. So if you have a calculation `(5 * 5)` with id `1234`, then it could be used as a sub-expression in the former expression as `(5 + id=1234 * 5)`.
* Support for changing individual elements in the expression. Should be relativly simple, as the expressions are already build from individual elements in a list, so the user just needs to give the index of the section the want to change.

## Architecture
The application itself is currently deployed in firebase at: https://helloworld-676480202164.europe-north1.run.app.
The architecture for the whole setup is showcased in the following:

![Uden titel](https://github.com/user-attachments/assets/06702f09-3e5c-4d73-bf1a-535de0f8d44f)

Basically, the go program defined in this repo is run in a single docker container on Firebase's Cloud Run service.
To ensure that calculations are pesisted between container restarts/upgrades, all data is persisted in a Firebase Realtime Database, which is a noSQL service.

### Future Architecture
Given that I only spent a few days on this, and I'm not familiar with Firebase yet, there's still work to be done before this would be ready for the real world.
I've attempted to sketch out my thoughts in the following.

![Uden titel](https://github.com/user-attachments/assets/9eb38e5d-9753-4840-8077-c9ff61b14ec1)

#### Load Balancing
In front of Cloud Run I would install a load balancer, such as nginx.

Adding a load balancer will allow me to distribute requests across multiple instances of the service, even though the user is just hitting the same URL. 
In addition a load balancer can add some protection against DDos attacks, such as by rate limiting requests coming from particular IPs.

#### Cloud Run
Within Cloud Run I can scale the number of service instances according to need and thus accomodate a larger load. Whether we need to scale up could be gauged by the CPU load across instances, though this would only work for scaling if the service turns out to be CPU bounded.

#### Auth & Users
The service is currently not set up with any authentication and therefore doesn't have any concept of a user.
Authentication could be implemented by creating a regular sign-up flow with email/password, storing the user's, hashed, credentials in a DB, and then providing a jwt access token to go along with their requests when they log in.
Alternativly we could use OpenID, and thus authenticate a user through some third-party like Google. This would of course require our users to have an account with the third-party already.

Either way, with a way of authenticating our users, we can associate each calculation with a particular user and provide authorization so that users can only read/write on calculations they own.
The permission model would be simple, if you created it you can read or write to it and no one else can.
This is quite restrictive, so we could later on extend the system to allow for sharing of calculations between users, but that would require a more complicated permission model,

#### Logging
A log management service such as Humio (now Falcon LogScale) will ensure that we can monitor our entire architecture from a single service. 
With this we can create dashboards to monitor deployment health and set up alerts to beep us if there are critical production issues.
From experience, it's also really nice to be able to search through the logs, when trying to debug something like a customer issue.

#### Frontend
Most people don't like to interact with computers through a REST API, so a frontend would be a nice addition.
Given that all of the logic is in the backend, the frontend could just be a simple page built on html/css/js that funnels requests to the backend.
