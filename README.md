## Crony :clock530:

* All the flexibility and power of Cron as a Service.
* Simple REST protocol, integrating with a web application in a easy and straightforward way.
* No more wasting time building and managing scheduling infrastructure.

## Basic Concepts
Crony works by calling back to your application via HTTP GET according to a schedule constructed by you or your application.

## Setup
Env vars
```bash
export DATASTORE_URL="postgresql://postgres@localhost/crony?sslmode=disable"
export PORT=3000
```

```sh
mkdir -p $GOPATH/src/github.com/rafaeljesus
cd $GOPATH/src/github.com/rafaeljesus
git clone https://github.com/rafaeljesus/crony.git
cd crony
make all
```

## Running server
```
./dist/crony
# => Starting Crony at port 3000
```

## Authentication
This API does not ship with an authentication layer. You **should not** expose the API to the Internet. This API should be deployed behind a firewall, only your application servers should be allowed to send requests to the API.

## API Endpoints
- [`GET` /health](#get-health) - Get application health
- [`GET` /events](#get-events) - Get a list of scheduled events
- [`POST` /events](#post-events) - Create a event
- [`GET` /events/:id](#get-eventsid) - Get a single event
- [`DELETE` /events/:id](#delete-eventsid) - Delete a event
- [`PATCH` /events/:id](#patch-eventsid) - Update a event

### API Documentation
#### `GET` `/events`
Get a list of available events.
- Method: `GET`
- Endpoint: `/events`
- Responses:
    * 200 OK
    ```json
    [
       {
          "id":1,
          "url":"your-api/job",
          "expression": "0 5 * * * *",
          "status": "active",
          "max_retries": 2,
          "retry_timeout": 3,
          "created_at": "2016-12-10T14:02:37.064641296-02:00",
          "updated_at": "2016-12-10T14:02:37.064641296-02:00"
       }
    ]
    ```
    - `id` is the id of the event.
    - `url`: is the url callback to called.
    - `expression`: is cron expression format.
    - `status`: tell if the event is active or paused.
    - `max_retries`: the number of attempts to send event.
    - `retry_timeout`: is the retry timeout.

#### `POST` `/events`
Create a new event.
- Method: `POST`
- Endpoint: `/events`
- Input:
    The `Content-Type` HTTP header should be set to `application/json`

    ```json
   {
      "url":"your-api/job",
      "expression": "0 5 * * * *",
      "status": "active",
      "max_retries": 2,
      "retry_timeout": 3,
   }
    ```
- Responses:
    * 201 Created
    ```json
   {
      "url":"your-api/job",
      "expression": "0 5 * * * *",
      "status": "active",
      "max_retries": 2,
      "retry_timeout": 3,
      "updated_at": "2016-12-10T14:02:37.064641296-02:00",
      "created_at": "2016-12-10T14:02:37.064641296-02:00"
   }
    ```
    * 422 Unprocessable entity:
    ```json
    {
      "status":"invalid_event",
      "message":"<reason>"
    }
    ```
    * 400 Bad Request
    ```json
    {
      "status":"invalid_json",
      "message":"Cannot decode the given JSON payload"
    }
    ```
    Common reasons:
    - the event job already scheduled. The `message` will be `Event already exists`
    - the expression must be crontab format.
    - the retry must be between `0` and `10`
    - the status must be `active` or `incative`

#### `GET` `/events/:id`
Get a specific event.
- Method: `GET`
- Endpoint: `/events/:id`
- Responses:
    * 200 OK
    ```json
   {
      "url":"your-api/job",
      "expression": "0 5 * * * *",
      "status": "active",
      "max_retries": 2,
      "retry_timeout": 3,
      "updated_at": "2016-12-10T14:02:37.064641296-02:00",
      "created_at": "2016-12-10T14:02:37.064641296-02:00"
   }
    ```
    * 404 Not Found
    ```json
    {
      "status":"event_not_found",
      "message":"The event was not found"
    }
    ```

#### `DELETE` `/events/:id`
Remove a scheduled event.
- Method: `DELETE`
- Endpoint: `/events/:id`
- Responses:
    * 200 OK
    ```json
    {
      "status":"event_deleted",
      "message":"The event was successfully deleted"
    }
    ```
    * 404 Not Found
    ```json
    {
      "status":"event_not_found",
      "message":"The event was not found"
    }
    ```

#### `PATCH` `/events/:id`
Update a event.
- Method: `PATCH`
- Endpoint: `/events/:id`
- Input:
    The `Content-Type` HTTP header should be set to `application/json`

    ```json
   {
      "expression": "0 2 * * * *"
   }
    ```
- Responses:
    * 200 OK
    ```json
   {
      "url":"your-api/job",
      "expression": "0 2 * * * *",
      "status": "active",
      "max_retries": 2,
      "retry_timeout": 3,
      "updated_at": "2016-12-10T14:02:37.064641296-02:00",
      "created_at": "2016-12-10T14:02:37.064641296-02:00"
   }
    ```
    * 404 Not Found
    ```json
    {
      "status":"event_not_found",
      "message":"The event was not found"
    }
    ```
    * 422 Unprocessable entity:
    ```json
    {
      "status":"invalid_json",
      "message":"Cannot decode the given JSON payload"
    }
    ```
    * 400 Bad Request
    ```json
    {
      "status":"invalid_event",
      "message":"<reason>"
    }
    ```

## Cron Format
The cron expression format allowed is:

|Field name| Mandatory?|Allowed values|Allowed special characters|
|:--|:--|:--|:--|
|Seconds      | Yes        | 0-59            | * / , -|
|Minutes      | Yes        | 0-59            | * / , -|
|Hours        | Yes        | 0-23            | * / , -|
|Day of month | Yes        | 1-31            | * / , - ?|
|Month        | Yes        | 1-12 or JAN-DEC | * / , -|
|Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?|
more details about expression format [here](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format)

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges
[![CircleCI](https://circleci.com/gh/rafaeljesus/crony.svg?style=svg)](https://circleci.com/gh/rafaeljesus/crony)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/crony)](https://goreportcard.com/report/github.com/rafaeljesus/crony)
[![](https://images.microbadger.com/badges/image/rafaeljesus/crony.svg)](https://microbadger.com/images/rafaeljesus/crony "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/crony.svg)](https://microbadger.com/images/rafaeljesus/crony "Get your own version badge on microbadger.com")
