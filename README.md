# Technical Test for Deporvillage

## Project structure
The project structure is defined using the hexagonal architecture

* The application layer contains the configuration and the potential "usecase" (not applied for this application).
* The domain layer contains the application business logic 
* The infrastructure contains all the necessary code to execute the server and save the info in a file.
* The "cmd" folder contains the main.go to execute the application.

##App configuration
Change the file app.json to change the application configuration

- **Port**: server port
- **Host**: server host
- **MaxClientConnections**: Number of maximum clients can connect to the server at the same time.
- **ConnectionType**: The connection type of the server (tcp, upd, etc).
- **SkuLogPath**: The path of the sku logs

##Makefile
Use the `make` command to install and execute the application, you can find more info using the command `make help`