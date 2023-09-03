# FocusQueue Web Application

FocusQueue is a web-based task management application built using Go, MySQL, and Gorilla Mux. It allows users to create, complete, and manage their to-do list with real-time updates of daily request counts. This project also demonstrates effective logging practices.

## Features

- User-friendly web interface for managing to-do lists.
- Real-time request count display using Server-Sent Events.
- MySQL database integration for data persistence.
- Structured logging for monitoring server activities.
- Written in the Go programming language for performance and scalability.

## Getting Started

Follow these steps to set up and run the Go-Todo application:

1. **Prerequisites:** Ensure you have Go and MySQL installed on your system.

2. **Clone the Repository:** Clone this repository to your local machine:

   ```bash
   git clone https://github.com/nihankhan/go-todo.git
   cd go-todo
   
1. Database Configuration: Configure your MySQL database credentials in the config/config.go file.

2. Run the Application: Start the Go server:

go run main.go

3. Access the Application: Open your web browser and visit http://localhost:8080 to start using the Go-Todo application.

 Usage

    Add new tasks by entering task descriptions and clicking "Submit."
    Mark tasks as completed or delete them as needed.
    Observe real-time request counts in the application header.

Logging

All server activities are logged for monitoring and debugging purposes. Logs are stored in the server.log file.

Contributing

Contributions are welcome! Feel free to open issues, submit pull requests, or suggest improvements.

License

This project is licensed under the MIT License - see the LICENSE file for details.


Feel free to customize this README file to provide more specific information about your project. Be sure to replace placeholders like `yourusername` and add any additional sections or details that are relevant to your project.
    

