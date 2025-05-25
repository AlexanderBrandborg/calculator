## dljkkl 

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
The REST API is documented by a postman collection, which can be found [here](https://martian-rocket-445009.postman.co/workspace/Calculator~d6272f3d-ab66-49e4-8c65-9de1ba785b57/collection/15879903-0da78f4b-7276-4d84-9bb2-9b59fe6a5279?action=share&creator=15879903&active-environment=15879903-22c16b80-88a0-4d49-8d37-b0271132a54a). 
In case that link becomes unavailable, a JSON export of the collection can be found at the root of this repository at `Calculator.postman_collection.json`.

The basic concept of the API is that it let's you build and evaluate arithmetic expressions. That is, any expression that can be constructed from integers and the +, -, * and / operators.

The design philosophy has been to not let the user create any illegal expressions, meaning they can't create an expressions with a missing operator argument, ex. '5 +' and they can't divide by zero. The benefit of this is that the expressions are always in a state, where they can be evaulated, which simiplifies the logic on the backend.

### Future Extensions
.

## Architecture
The application itself is currently deployed in firebase.
The architecture for the whole setup is showcased in the following:

![Uden titel](https://github.com/user-attachments/assets/158cd43c-7050-497c-ad1e-1595df6481df)

Basically, the go program defined in this repo is run in a singel docker container on Firebase's Cloud Run service.
To ensure that calculations are pesisted between container restarts/upgrades, all data is persisted in a Firebase Realtime Database, which is a noSQL service.

### Future Architecture
Given that I only spend a few days on this, and I'm not familiar with Firebase yet, there's still work to be done before this would be ready for the real world.
I've attempted to sketch out my thoughts in the following.

![Uden titel](https://github.com/user-attachments/assets/9eb38e5d-9753-4840-8077-c9ff61b14ec1)

#### Load Balancing
.

#### Cloud Run
.

#### Auth
.

#### Logging
.

#### Frontend
.


