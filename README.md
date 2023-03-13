![Go](https://img.shields.io/badge/Go-1.17-f21170?style=flat-square&logo=docker&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-3.3.2-f21170?style=flat-square&logo=docker&logoColor=white)

# Key Value & Search API

This API has two different duty, one of them is to create and fetch data to in-memory database The other one is fetch the data in the provided MongoDB collection and returns the results in the requested format.

## Table of Contents:

- [Getting Started](#getting-started)
    - [Requirements](#requirements)
  - [Building with Docker](#with-docker)
- [API Endpoints and Documentations](#api-endpoints-and-documentations)
  - [GET `/api/get`](#get-allkeyvaluepairs)
  - [POST `/api/set`](#post-keyvaluepair)
  - [GET `/api/get-all`](#get-keyvaluepairs)
- [Live URL](#live-url)
- [Contact Information](#contact-information)
- [License](#license)

<br/>

## Getting Started

<hr/>

### Requirements:

- Go v1.17 or higher -> [Go Installation Page](https://go.dev/dl/)

  Before starting the application, fork/download/clone this repo.

<hr/>

### Building with Docker

- To run the application on [localhost:8000](http://localhost:8000):

```
docker-compose up --build
```

<hr/>


## API Endpoints and Documentations

<hr/>

### POST `/api/create`

- Description: Creates new key-value pair.

#### Request Body:

``` 
{
    "key": "foo",
    "value": "bar"
}
```

#### Reponse:

```
{
    "key": "foo",
    "value": "bar"
}
```
Example bad request.
```
{
    "message": "There is an error in the requested data. Check the data. Data should be JSON."
}
```
| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 201    | Created| The key-value pair created |
| 400     | Bad Request       |    URL, JSON structure or request method is wrong |
<hr/>

### GET `/api/fetch?key=<foo>`

- Description: The key's value returns in response..


#### Request:

```
GET Request to '/api/get?key=foo' endpoint. //Foo means name of key, desired value can be entered.
```


#### Reponse:

```
{
    "key": "foo",
    "value": "bar",
}
```

If the key is not used:

```
{
    "message": "Bad Request."
}
```

| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success| The value of the key was searched |
| 400     | Bad Request       |   URL, JSON structure or request method is wrong |
<hr/>


### POST `/api/search`

- Description: Search and sum counts.
#### Request:

```
{
  "startDate": "2005-01-26",
  "endDate": "2018-02-02",
  "minCount": 1000,
  "maxCount": 3000
}
```

#### Reponse:

```
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "createdAt": "2017-01-28T01:22:14.398Z",
            "key": "TAKwGc6Jr4i8Z487",
            "totalCount": 2800
        },
        {
            "createdAt": "2017-01-27T08:19:14.135Z",
            "key": "NAeQ8eX7e5TEg7oH",
            "totalCount": 2900
        },
        {
            "createdAt": "2013-01-27T08:19:14.135Z",
            "key": "NAeQ8eX7e5TEg8oH",
            "totalCount": 2000
        }
    ]
}
```
If status code is 404.
```
{
    "code": 404,
    "msg": "No data found.",
    "records": null
}
```
| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success|Provided information about all key/value pairs |
| 404     | Bad Request       |  No data found |
<hr/>

### Live URL:

- [Live URL](https://kv-search-api.herokuapp.com/)

<hr/>

#### Author: İlker Rişvan

#### Github: ilkerrisvan

#### Email: ilkerrisvan@outlook.com

#### Date: March, 2023

## License

<hr/>

[MIT](https://choosealicense.com/licenses/mit/)
