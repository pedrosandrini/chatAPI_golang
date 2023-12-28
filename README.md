# ChatGPT API with Golang, DDD, and Clean Architecture

Welcome to my ChatGPT API project! This API is designed to seamlessly integrate with ChatGPT, offering a powerful and efficient solution for incorporating advanced natural language processing into your applications. Developed using the Go programming language (Golang) and following the principles of Domain-Driven Design (DDD) and Clean Architecture, this project provides a robust, scalable, and maintainable foundation for your language-related endeavors.

## Features

1. **Golang Implementation:** Leverage the efficiency and performance benefits of Golang for a high-quality API.

2. **ChatGPT Integration:** Seamlessly connect to the ChatGPT language model for advanced natural language processing.

3. **Domain-Driven Design (DDD):** Structured around core business domains for better understanding and maintainability.

4. **Clean Architecture:** Follows Clean Architecture patterns for modularity and adaptability to changes without sacrificing scalability.

5. **Scalability and Performance:** Designed to handle increased loads while maintaining optimal performance.

6. **Flexibility and Extensibility:** Easily customizable and extensible to meet specific project requirements.

7. **Documentation:** Thoroughly documented API endpoints, codebase, and architectural decisions for easy onboarding and collaboration.

## Getting Started

To get started with the ChatGPT API, follow the steps below:

1. **Obtain OpenAI Authorization Key:**

   - Get your OpenAI API key at [https://platform.openai.com/api-keys](https://platform.openai.com/api-keys).
   - Note: The account generating the API key must be created within the last 3 months for the OpenAI trial period to work.

2. **Configure API Key:**

   - Copy the obtained key and paste it into the `.env` file under the `OPENAI_API_KEY` field.

3. **Set Initial Chat Context:**

   - Define the desired chat context in the `.env` file under the `INITIAL_CHAT_MESSAGE` field.

4. **Start Services:**

   - In the root folder, run `docker compose up -d` to initiate the services.

5. **Access Golang Service:**

   - Execute `docker exec -it chatservice_app bash` to access the Golang service.

6. **Run the Web Server:**

   - Execute `go run cmd/chatservice/main.go` to start the web server.

7. **Test the API:**
   - Use REST requests in Insomnia, Postman, or the REST Client extension in VSCode to test the API.

## License

This project is licensed under the [MIT License](LICENSE).
