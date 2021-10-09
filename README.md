<h1 align="center">✨Instagram APIs ✨</h1>
<p align="center">
<img src="https://img.shields.io/badge/Made%20with-Go-blue.svg"/>
</p>

<span align="center">

#### The task given by Appointy completed to develop APIs for a basic version of Instagram.<br/>
 Create an User<br/>
 Get User by Id<br/>
 Create a Post<br/>
 Get a Post by Id<br/>
 Get all Posts by an User with Pagination<br/>

</span>

## Setup:
#### STEP-1 : Install and Set Up MongoDB
```
            https://docs.mongodb.com/v5.0/installation/
```
#### STEP-2 : Clone this repository
```
            git clone https://github.com/aayush1607/instagram_api.git
            cd instagram-api
```
#### STEP-3 : Run local MongoDB instance
```
            mongo
```
#### STEP-4 : Download Dependencies
```
            go mod download
```
#### STEP-5 : Finally Run the Server
```
            go run main.go
```
#### STEP-6 : Then simply navigate in your browser or test the above mentioned APIs on Postman or CURL

## APIs :
1. Create an User  
    ***Method Allowed-POST***  
     ***URL***
    ``` http://localhost:8080/users ```  
    ***JSON Request Body***
    ```json
    {
        "Name":"aayush",
        "Email":"auc1607@gmail.com",
        "Password": "aayush@1234"
    }
    ```
    ***JSON Response Body***
    ```json
    {
        "InsertedID": "61617acff3b743f10bb0923a"
    }
    ```
2. Get User By id  
    ***Method Allowed-GET***      
    ***URL***
    ``` http://localhost:8080/users/?id=61617acff3b743f10bb0923a ```  
    ***JSON Response Body***
    ```json
    {
        "_id": "61617acff3b743f10bb0923a",
        "name": "aayush",
        "email": "auc1607@gmail.com",
        "password": "$2a$10$NI5Ub2L2jIePsJu8ljOjBeK.LtfdTj5OoaWV.G7XqU3zuhAF/inCS"
    }
    ```
3. Create a Post  
    ***Method Allowed-POST***       
     ***URL***
    ``` http://localhost:8080/posts ```  
    ***JSON Request Body***
    ```json
    {
        "User": "61615ad9d9a12f1927083b7f",
        "Caption": "Hello Instagram 4",
        "Image_url": "www.insta.com",
        "Timestamp": ""
    }
    ```
    ***JSON Response Body***
    ```json
    {
        "InsertedID": "61616540d9fbb9baedcff879"
    }
    ```    
4. Get a Post by id  
    ***Method Allowed-GET***      
    ***URL***
    ``` http://localhost:8080/posts/?id=61616540d9fbb9baedcff879 ```  
    ***JSON Response Body***
    ```json
    {
        "_id": "61616540d9fbb9baedcff879",
        "user": "61615ad9d9a12f1927083b7f",
        "caption": "Hello Instagram 4",
        "image_url": "www.insta.com",
        "timestamp": "2021-10-09T09:47:44.93Z"
    }
    ```
5. Get all Posts by user id with pagination  
    (If limit and page parameters are no provided then by default they will be taken as 1)  
    ***Method Allowed-GET***      
    ***URL***
    ``` http://localhost:8080/posts/users/?id=61615ad9d9a12f1927083b7f&limit=2&page=1 ```  
    ***JSON Response Body***
    ```json
    {
        "posts": [
            {
                "_id": "6161616aa66d684e5bb38fbc",
                "user": "61615ad9d9a12f1927083b7f",
                "caption": "Hello Instagram",
                "image_url": "www.insta.com",
                "timestamp": "0001-01-01T00:00:00Z"
            },
            {
                "_id": "616161e885e8e66798b7ed78",
                "user": "61615ad9d9a12f1927083b7f",
                "caption": "Hello Instagram 2",
                "image_url": "www.insta.com",
                "timestamp": "2021-10-09T09:33:28.758Z"
            }
        ],
        "total": 4,
        "page": 1,
        "last_page": 2,
        "limit": 2
    }
    ```

## > [Postman API collection link](https://bit.ly/3BrUVqR) 

## Packages Used
1. [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.4.0) to work with mongodb database
2. [Bcrpyt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) to hash passwords
2. [net/http](https://pkg.go.dev/std) for response request and response writer
3. [encoding/json](https://pkg.go.dev/std) for encoding and decoding from json to type struct and vice versa
4. [time](https://pkg.go.dev/std) for saving current timestamps if not provided by frontend interface

## Languages & Tools Used 
<p align='center'>
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />&nbsp;&nbsp;
<img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" />&nbsp;&nbsp;
<img src="https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=Postman&logoColor=white" />&nbsp;&nbsp;
<img src="https://img.shields.io/badge/Git-F05032?style=for-the-badge&logo=git&logoColor=white" />&nbsp;&nbsp;
<img src="https://img.shields.io/badge/Visual_Studio_Code-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white" />&nbsp;&nbsp;
</p>



