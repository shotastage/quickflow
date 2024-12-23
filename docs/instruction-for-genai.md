# Clean Architecture Code Generation Instructions

When generating a new model (Entity) struct, also generate the following files. Unless specified otherwise, implement full CRUD functionality for the model's data. Additionally, include the file path as a comment at the top of each generated file. Follow the file structure below:

1. **Model (Entity)**  
   - Path: `internal/domain/[ModelName]/[ModelName].go`  
   - Purpose: Defines the core domain entity structure, representing the data model used throughout the application.
   - Add the following comment at the top of the file:
     ```go
     // File: internal/domain/[ModelName]/[ModelName].go
     ```

2. **Service**  
   - Path: `internal/application/[ModelName]/[ModelName]_service.go`  
   - Purpose: Contains business logic for handling the model. This service coordinates between the repository and any other necessary components.
   - Add the following comment at the top of the file:
     ```go
     // File: internal/application/[ModelName]/[ModelName]_service.go
     ```

3. **Repository**  
   - Path: `internal/infrastructure/repository/[ModelName]_repository.go`  
   - Purpose: Responsible for database interactions. This file contains the CRUD operations and any other database-related logic specific to the model.
   - Add the following comment at the top of the file:
     ```go
     // File: internal/infrastructure/repository/[ModelName]_repository.go
     ```

4. **Handler**  
   - Path: `internal/interfaces/httpserver/handler/[ModelName]_handler.go`  
   - Purpose: Handles incoming HTTP requests related to the model. It manages routing and links the service layer with the web interface.
   - Add the following comment at the top of the file:
     ```go
     // File: internal/interfaces/httpserver/handler/[ModelName]_handler.go
     ```
