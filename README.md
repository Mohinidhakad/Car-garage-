## Zopsmart-pro

# Step 1

As I am new to this technology i started with the que. "what is GO(GOLANG),GOFR and GORM?"
After taking the theoritical refrences from the artical->" https://medium.com/@mundhraumang.02?source=post_page-0ea41eaa6c62-  " and some other resources, I moved towords seeking some prectical implementations for that again i have to learn  not too much but basics of GO language such as packages , variables, functions etc.

# Step 2

The summarized steps based on our project 
1. Download and Install Go (https://golang.org/dl/)
2. Install Golang Editor
3. Check Go Installation(go version)
4. setup the project

# Step 3

Following the step for run the project using the following commands ->

1. Clone the Project
``` bash
   git clone <project_url>
   cd <project_directory>
```
2. Install Dependencies 
```bash
     go mod tidy
```
4. Create a Local MySQL Database
``` bash
   mysql -u root -p
```
4. Run the Program
``` bash
   go run main.go
```
5. Test with Postman 


# Step 4

I set up the project by creating a database named 'cargarage' with a table called 'Cars.' Subsequently, I established a connection to this database. Moving forward, I developed the CRUD methods—'Get,' 'Update,' 'Delete,' and 'Create'—within the *store file.
``` bash
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
```

# Step 5

# Diagram
1. Flowchart ![flowchat](https://github.com/Mohinidhakad/Zopsmart-pro/assets/99878435/893f3024-5764-447e-ae63-22b143c7a308)
2. Sequence Diagram ![flowchart1 drawio](https://github.com/Mohinidhakad/Zopsmart-pro/assets/99878435/317c43d2-30c6-4fd3-99ea-dade5780f33b)


    


                                                                        

                                                                  
