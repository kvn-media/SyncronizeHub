# SyncronizeHub

SyncronizeHub is a Golang backend application designed for handling synchronization and data flow between PostgreSQL and Firebase.

## Project Structure

```
syncronizehub/
|-- configs/
|-- delivery/
|   |-- controllers/
|   |-- middleware/
|-- managers/
|-- models/
|-- repository/
|-- usecase/
|-- utils/
|   |-- authenticator/
|-- go.mod
|-- main.go
```

## Dependencies

- `github.com/jinzhu/gorm`: GORM is used as the Object Relational Mapper for PostgreSQL integration.
- `github.com/lib/pq`: The PostgreSQL driver for GORM.
- `github.com/gorilla/mux`: Gorilla Mux is utilized for routing and handling HTTP requests.
- `firebase-admin-go`: Firebase Admin SDK for handling authentication and data synchronization.

## Getting Started

1. Clone the repository.
2. Set up your PostgreSQL database and configure the connection string in the code.
3. Initialize Firebase Admin SDK with your project credentials.
4. Run the application.

## Features

- User registration and login via RESTful API.
- CRUD operations for flow data.
- Integration with PostgreSQL for data storage.
- Firebase authentication and data synchronization.

## API Endpoints

- `/register`: User registration endpoint.
- `/login`: User login endpoint.
- `/flowdata/insert`: Insert flow data.
- `/flowdata/update`: Update flow data.
- `/flowdata/delete`: Delete flow data.

## Security Measures

- Hashing and salting user passwords.
- Sanitizing user inputs to prevent SQL injection attacks.
- Regular security updates for OS, database, and Firebase SDK.
- HTTPS communication.

## Contributors

- Muhammad Kevin

Feel free to contribute and make SyncronizeHub even more awesome!

```

**Topology Diagram**

```
                        +-------------+
                        | PostgreSQL  |
                        |    Database |
                        +------+------+
                               |
            +------------------v------------------+
            |               SyncronizeHub           |
            |   +----------------+   +------------+ |
            |   |   Controllers  |   | Middleware | |
            |   +----------------+   +------------+ |
            |   |   Use Cases    |   |   Utils    | |
            |   +----------------+   +------------+ |
            |         |                    |        |
            +---------|--------------------|--------+
                      v                    v
            +-------------------+ +------------------+
            |      GORM         | |  Firebase Admin  |
            | PostgreSQL Driver | |      SDK         |
            +-------------------+ +------------------+
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This README file includes instructions on installing the required dependencies using the `go get` command. Make sure to adjust the installation commands as needed, and it's always a good idea to use the latest stable versions of the dependencies.