## Zopsmart-pro
# Step 1
As I am new to this technology i started with the que. "what is GO(GOLANG),GOFR and GORM?"
After taking the theoritical refrences from the artical->" https://medium.com/@mundhraumang.02?source=post_page-----0ea41eaa6c62--------------------------------  " and some other resources, I moved towords seeking some prectical implementations for that again i have to learn  not too much but basics of GO language such as packages , variables, functions etc.

# Step 2
The summarized steps based on our project 
1. Download and Install Go (https://golang.org/dl/)
2. Set Up Go Environment in VS Code
3. Install Golang Compiler
4. Check Go Installation(go version)


# Step 3
I set up the project by creating a database named 'cargarage' with a table called 'Cars.' Subsequently, I established a connection to this database. Moving forward, I developed the CRUD methods—'Get,' 'Update,' 'Delete,' and 'Create'—within the *store file.
        Zopsmart-pro/
        ├── configs/
        │     └── .env
        │
        ├── handler/
        │     └── handler.go
        │
        ├── model/
        │     └── model.go
        │
        ├── store/
        │     ├── store.go
        │── testcase
        │    ├── store_test.go
        │    ├── handler_test.go
        │       
        ├── main.go
        ├── go.mod

# Step 4

Following the step for run the project using the following commands ->

1. Clone the project
2. Initialize Go Module
   ```bash   
   go mod init Zopsmart-pro
3. Install Dependencies
   ```bash   
   go mod tidy or go get -u github.com/example/library
4. Setup the mysql store
   ```bash
      go get github.com/DATA61/sqlmock
      or
      go get -u github.com/stretchr/testify
5. Run Project
   ```bash
    go run main.go

    
    


                                                                        

                                                                  
