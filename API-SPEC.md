# API Specification

All API must use this authentication
Request:
- Header:
    - X-Api-Key: ``` "your secret api key" ```

SUMMARY:
- Users:                            Line 19
- User_course:                      Line 170
- User_details:                     Line 286
- Courses:                          Line 442
- Modules:                          Line 594
- Module_submission:                Line 738
- Module_articles:                  Line 870
- Answers:                          Line 990
- Question:                         Line 1124

## Users
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
        "name" : "string",
        "username" : "string", // unique
        "email" : "string", // unique
        "password" : "string",
        "role" : "integer", // enum(1,2)
        "email_verified_at" : "timestamp", // timestamp
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
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
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

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
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
- Body:
``` json
    {
        "name" : "string",
        "username" : "string", // unique
        "email" : "string", // unique
        "password" : "string",
        "role" : "integer", // enum(1,2)
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
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
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
                "id" : "integer", // primary key
                "name" : "string",
                "username" : "string", // unique
                "email" : "string", // unique
                "password" : "string",
                "role" : "integer", // enum(1,2)
                "email_verified_at" : "timestamp", // timestamp
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
```
------------------------------
Delete Users
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/users/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## User_course
------------------------------
## Create User_course
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/user_course```
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
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
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
## Update User_course
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
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
## List User_course
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/user_course```
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
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## User_details
------------------------------
## Create User_details
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/user_details```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "user_id" : "integer", // foreign key1
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(1,2,3)
        "address" : "string",
        "birthdate" : "date", // date
        "image" : "string",
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
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string",
            "description" : "string"
        }
    }
```
------------------------------
## Get User_details
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/user_details/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json    
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string",
            "description" : "string"
        }
    }
```
------------------------------
## Update User_details
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/user_details/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "user_id" : "integer", // foreign key1
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(1,2,3)
        "address" : "string",
        "birthdate" : "date", // date
        "image" : "string",
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
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string",
            "description" : "string"
        }
    }
```
------------------------------
## List User_details
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/user_details```
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
                "user_id" : "integer", // foreign key1
                "phone" : "string",
                "gender" : "integer", // enum(1,2)
                "type_of_disability" : "integer", // enum(1,2,3)
                "address" : "string",
                "birthdate" : "date", // date
                "image" : "string",
                "description" : "string"
            },
        ]
    }
```
------------------------------
## Delete User_details
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/user_details/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
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

## Module_submissions
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
## Previous Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{code}/submissions/{submissionId}/previous```
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
## Module_articles
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
## User_Submissions
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
- Endpoint: ```/api/answers```
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
- Endpoint: ```/api/answers/{id}```
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
- Endpoint: ```/api/answers/{id}```
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
- Endpoint: ```/api/answers```
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
- Endpoint: ```/api/answers/{id}```
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
- Endpoint: ```/api/questions```
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
- Endpoint: ```/api/questions/{id}```
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
- Endpoint: ```/api/questions/{id}```
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
- Endpoint: ```/api/questions```
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
- Endpoint: ```/api/questions/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```