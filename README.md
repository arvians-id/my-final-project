# Engineering 12 Final Project - Teenager

## Collaborators
  1. Irfan Kurniawan - Product Manager - FE2110451 ([Irfan858](https://github.com/Irfan858))
  2. Mohd Ryan Obillah - Frontend Developer - FE2163553 ([obillahh](https://github.com/obillahh))
  3. Umbu Theofilus Dendimara - Backend Developer - BE2144668 ([Rendydinar](https://github.com/Rendydinar))
  4. Widdy Arfiansyah - Backend Developer - BE2108902 ([arvians-id](https://github.com/arvians-id))
  5. Muhammad Abid Fajar - Backend Developer - BE2219863 ([abid313](https://github.com/abid313))
  6. Rudiansyah Wijaya Pratama - Backend Developer - BE2242991 ([Reezyx](https://github.com/Reezyx)) 

# API Specification

All API must use this authentication
Request:
- Header:
  - X-Api-Key: ``` "your secret api key" ```

SUMMARY:
- [Users](#users)                               ```(11/11) 100%```
- [User_course](#user-course)                   ```(5/5) 100%```
- [Courses](#courses)                           ```(7/7) 100%```
- [Module_submissions](#module-submissions)     ```(8/8) 100%```
- [Module_articles](#module-articles)           ```(7/7) 100%```
- [User_Submissions](#user-submissions)		    ```(4/4) 100%```
- [Answers](#answers)                           ```(5/5) 100%```
- [Questions](#questions)                       ```(5/5) 100%```

There are a total of ```52``` APIs


## users
------------------------------
## Create Users
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users ```
- Header:
  - Content-Type: ``` application/json ```
  - Accept: ``` application/json ```
- Body:
``` json
    {
        "name": "string",
        "username": "string",
        "email": "string",
        "password": "string",
        "role": "integer", // enum (1, 2)
        "phone": "string",
        "gender": "integer", // enum (1, 2)
        "type_of_disability": "integer", // enum (0, 1, 2)
        "birthdate": "date"
    }
```
Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary
            "name": "string",
            "username": "string", // unique
            "email": "string", // unique
            "password": "string",
            "role": "integer", // enum(1, 2)
            "phone": "string",
            "gender": "integer", // enum (1, 2)
            "type_of_disability": "integer", // enum (0, 1, 2)
            "birthdate": "string",
            "email_verification": "timestamp",
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    }
```
------------------------------
## Users Login
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users/login ```
- Header:
  - Content-Type: ``` application/json ```
  - Accept: ``` application/json ```
- Body:
``` json
    {
        "email": "string",
        "password": "string"
    }
```
Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary
            "name": "string",
            "username": "string", // unique
            "email": "string", // unique
            "role": "integer", // enum(1, 2)
            "gender": "integer",
            "type_of_disability" : "integer" // enum
        }
    }
```
------------------------------
## Get User Status
------------------------------
Request:
- Method: ```GET ```
- Endpoint: ```/api/userstatus ```
- Header:
  - Content-Type: ``` application/json ```
  - Accept: ``` application/json ```
- Body:
``` json
    {
        "id": "integer"
    }
```
Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary
            "name": "string",
            "username": "string", // unique
            "email": "string", // unique
            "role": "integer", // enum(1, 2)
            "gender": "integer",
            "type_of_disability" : "integer", // enum
            "address" : "string",
            "birthdate" : "string",
            "image" : "string",
            "description" : "string"
        }
    }
```
------------------------------
## User Logout
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users/logout ```
- Header:
  - Content-Type: ``` application/json ```
  - Accept: ``` application/json ```

Response:
``` json
    {   
        "code" : "number",
        "status" : "string"
    }
```
------------------------------
## Get Users
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users/{id}```
- Header:
  - Accept: ```application/json```
- Variable:
  - id: ```integer```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary
            "name": "string",
            "username": "string", // unique
            "role": "integer", // enum(1, 2)
            "phone": "string",
            "gender": "integer", // enum (1, 2)
            "type_of_disability": "integer", // enum (0, 1, 2)
            "address": "string",
            "birthdate": "string",
            "image": "string",
            "description": "string"
        }
    }
```
------------------------------
## Update Users
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/users/{id}```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Variable:
  - id: ```integer```
- Body:
``` json
    {
        "name": "string",
        "username": "string", // unique
        "role": "integer", // enum (1, 2)
        "phone": "string",
        "gender": "integer", // enum (1, 2)
        "type_of_disability": "integer", // enum (0, 1, 2)
        "address": "string",
        "birthdate": "string",
        "image": "string",
        "description": "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary
            "name": "string",
            "username": "string", // unique
            "role": "integer", // enum(1, 2)
            "phone": "string",
            "gender": "string", // enum (1, 2)
            "type_of_disability": "integer", // enum(0, 1, 2)
            "address": "string",
            "birthdate": "string",
            "image": "string",
            "description": "string"
        }
    }
```
------------------------------
## Update Users Role
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/users/roleupdate/{id}```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Variable:
  - id: ```integer```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability": "integer", // enum(0,1,2)
            "address": "string",
            "birthdate": "string",
            "image": "string",
            "description": "string"
        }
    }
```
------------------------------
## List Users
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users```
- Header:
  - Accept: ```application/json```
- Query Param:
  - size : ```number```
  - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id": "integer", // primary
                "name": "string",
                "username": "string", // unique
                "role": "integer", // enum(1, 2)
                "phone": "string",
                "gender": "integer", // enum(1, 2)
                "type_of_disability": "integer", // enum(0, 1, 2)
                "address": "string",
                "birthdate": "string",
                "image": "string",
                "description": "string"
            },
        ]
    }
```
------------------------------
## Delete Users
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/users/{id}```
- Header:
  - Accept: ```application/json```
- Variable:
  - id: ```integer```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
------------------------------
## List User Submission
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users/submissions```
- Header:
  - Accept: ```application/json```
- Query Param:
  - limit : ```number``` ```optional``` ```default = all list```

Response:
``` json
    {
        "id_module_submission": "integer",
        "name_course": "string",
        "name_module_submission": "string",
        "grade": "integer",
        "file": "string"
    }
```
------------------------------
## Verify Email User
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users/verify?email=&signature=```
- Header:
  - Accept: ```application/json```
- Query Param:
  - signature : ```string``` ```required```
  - email : ```string``` ```required```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : null
    }
```
## User course
------------------------------
## Create User_course
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/usercourse```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "user_id" : "integer", // foreign key1
        "course_id" : "integer" // foreign key2
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
```
------------------------------
## Get User_course
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/usercourse/{id}/{course}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
```
------------------------------
## List User_course
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/usercourse```
- Header:
  - Accept: ```application/json```
- Query Param:
  - size : ```number```
  - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "user_id" : "integer", // foreign key1
                "course_id" : "integer" // foreign key2
            },
        ]
    }
```
------------------------------
## Delete User_course
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/usercourse/{user_id}/{course_id}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
------------------------------
## Find All Course By User Logged In
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/usercourse/courses```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id_course": "integer", // Primary Key
            "course_name": "string",
            "course_code": "string",
            "course_class": "string"
        }
    }
```
## Courses
------------------------------
## Create Courses
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "name" : "string",
        "code_course" : "string", // unique
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    } 
```  
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "is_active" : "boolean",
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## Get All courses
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses```
- Header:
  - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```
    
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "code_course" : "string", // unique
                "class" : "string",
                "tools" : "string", // longtext
                "about" : "string", // longtext
                "is_active" : "boolean",
                "description" : "string", // longtext
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp 
            },
        ]
    }
```
------------------------------
## Get Courses
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "is_active" : "boolean",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## Update Courses
------------------------------
Request:
- Method: ```PATCH```
- Endpoint: ```/api/courses/{code}```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "name" : "string",
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "is_active" : "boolean",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## List Courses
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses?status=true&limit=1```
- Header:
  - Accept: ```application/json```
- Query Param:
  - status : ```boolean``` ```optional``` ```default = true```
  - limit : ```number``` ```optional``` ```default = all list```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "code_course" : "string", // unique
                "class" : "string",
                "tools" : "string", // longtext
                "about" : "string", // longtext
                "description" : "string", // longtext
                "is_active" : "boolean",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp 
            },
        ]
    }
```
------------------------------
## Delete Courses
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/courses/{code}```
- Header:
  - Accept: ```application/json```
- Query Param:
  - code : ```string```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : null
    }
```
------------------------------
## Deactivate Courses
------------------------------
Request:
- Method: ```PATCH```
- Endpoint: ```/api/courses/{code}/status```
- Header:
  - Accept: ```application/json```
- Query Param:
  - code : ```string```
- Body:
``` json
    {
        "is_active" : "boolean",
    }
```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : null
    }
```
------------------------------
## List Users In Course Detail
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/users```
- Query Param:
  - code : ```string```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id_user": "integer", // primary key
            "user_name": "string",
            "user_username": "string",
            "user_email": "string"
        }
    }
```
## Module submissions
------------------------------
## Create Module_submissions
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses/{code}/submissions```
- Query Param:
  - code : ```string```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "name": "string",
        "description": "string",
        "deadline": "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id": "integer", // foreign key
            "name": "string",
            "description": "string",
            "deadline": "string"
        }
    }
```
------------------------------
## Get Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
```
------------------------------
## Update Module_submissions
------------------------------
Request:
- Method: ```PATCH```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "file" : "string",
        "type" : "string",
        "max_size" : "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
```
------------------------------
## List Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions```
- Query Param:
  - code : ```string```
- Header:
  - Accept: ```application/json```
- Query Param:
  - size : ```number```
  - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "file" : "string",
                "type" : "string",
                "max_size" : "integer"
            },
        ]
    }
```
------------------------------
## Delete Module_submissions
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
------------------------------
## Next Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}/next```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": {
            "id": "integer", // primary key
            "code_course": "string",
        }
    }
```
------------------------------
## Previous Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}/previous```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": {
            "id": "integer", // primary key
            "code_course": "string",
        }
    }
```
------------------------------
## List User Module_submissions In Teacher
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}/get```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": {
            "id_user_submission": "integer", // primary key
            "user_name": "string",
            "module_submission_name": "string",
            "file": "string"
        }
    }
```
## Module articles
------------------------------
## Create Module_articles
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses/{code}/articles```
- Query Param:
  - code : ```string```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "name": "string",
        "content": "string",
        "estimate": "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary key
            "course_id": "integer", // foreign key 
            "name": "string",
            "content": "string",
            "estimate": "integer"
        }
    }
```
------------------------------
## Get Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/articles/{articleId}```
- Query Param:
  - code : ```string```
  - articleId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id": "integer", // primary key
            "course_id": "integer", // foreign key 
            "name": "string",
            "content": "string",
            "estimate": "integer"
        }
    }
```
------------------------------
## Update Module_articles
------------------------------
Request:
- Method: ```PATCH```
- Endpoint: ```/api/courses/{code}/articles/{articleId}```
- Query Param:
  - code : ```string```
  - articleId : ```number```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "name": "string",
        "content": "string",
        "estimate": "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "course_id": "integer", // foreign key 
            "name": "string",
            "content": "string",
            "estimate": "integer"
        }
    }
```
------------------------------
## List Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/articles```
- Query Param:
  - code : ```string```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "course_id": "integer", // foreign key 
                "name": "string",
                "content": "string",
                "estimate": "integer"
            },
        ]
    }
```
------------------------------
## Delete Module_articles
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/courses/{code}/articles/{articleId}```
- Query Param:
  - code : ```string```
  - articleId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": null
    }
```
------------------------------
## Next Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/articles/{articleId}/next```
- Query Param:
  - code : ```string```
  - articleId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": {
            "id": "integer", // primary key
            "code_course": "string",
        }
    }
```
------------------------------
## Previous Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/articles/{articleId}/previous```
- Query Param:
  - code : ```string```
  - articleId : ```number```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data": {
            "id": "integer", // primary key
            "code_course": "string",
        }
    }
```
## User Submissions
------------------------------
## Submit File
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses/:code/submissions/:submissionId/user-submit```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
- Header:
  - Content-Type: ```multipart/form-data```
- Body:
``` json
    {
        "file" : "string",
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "module_submission_id" : "integer", //foreign key2
            "file" : "string"
        }
    }
```
------------------------------
## Update Grade
------------------------------
Request:
- Method: ```PATCH```
- Endpoint: ```/api/courses/:code/submissions/:submissionId/user-submit/:userSubmissionId```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
  - userSubmissionId : ```number```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "grade" : "integer",
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : null
    }
```
------------------------------
## Get User Submission
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/:code/submissions/:submissionId/user-submit/:userSubmissionId```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
  - userSubmissionId : ```number```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "module_submission_id" : "integer", //foreign key2
            "file" : "string",
            "grade" : "integer"
        }
    }
```
------------------------------
## Download User Submission
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses/:code/submissions/:submissionId/user-submit/:userSubmissionId/download```
- Query Param:
  - code : ```string```
  - submissionId : ```number```
  - userSubmissionId : ```number```
- Header:
  - Content-Type: ```{mimetype}```
  - Content-Disposition: ```{attachment; filename=file}```

## Answers
------------------------------
## Create Answers
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/answers/create```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Get Answers
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/answers/by-user/{userId}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Update Answers
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/answers/update/{answerId}```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## List Answers
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/answers/all```
- Header:
  - Accept: ```application/json```
- Query Param:
  - size : ```number```
  - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "question_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
```
------------------------------
## Delete Answers
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/answers/{answerId}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Questions
------------------------------
## Create Questions
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/questions/create```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Get Questions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/questions/by-user/{userId}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Update Questions
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/questions/update/{questionId}```
- Header:
  - Content-Type: ```application/json```
  - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## List Questions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/questions/all```
- Header:
  - Accept: ```application/json```
- Query Param:
  - size : ```number```
  - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "title" : "string",
                "tags" : "string",
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
```
------------------------------
## Delete Questions
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/questions/{questionId}```
- Header:
  - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```