 # Packform Web App Purpose

Creating a dashboard for monitoring orders on packform.

# Installing / Getting started

The app is composed of two major elements: client app (Vuejs) and server app (Go). 
To start the client app, use the command below.

```shell
npm run serve
```

To start the server app, use the command below.
```shell
go build backend/* && ./app
```

After executing both commands, client app will start on port 8080 and server app on port 8888.

# Initial Configuration

Before running client and server apps, postgreSQL and mongoDB need to be available running on localhost, listening to port 5432 and 27017 respectively.





