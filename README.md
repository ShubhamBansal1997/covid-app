# Covid-19 API 

Covid-19 API to get covid-19 stats in your state and country

* [Global Requisites](#global-requisites)
* [Install, Configure & Run](#install-configure--run)
* [List of Routes](#list-of-routes)

# Global Requisites

* go (>= 1.16.0)
* mongodb
* redis

```
Ingress Service 🙌 
Manages all the webhook API(s) 
```

# Contents

* [Global Requisites](#global-requisites)
* [App Structure](#app-structure)
* [Install, Configure & Run](#install-configure--run)
* [List of Routes](#list-of-routes)

# Global Requisites

* node (>= 10.5.0)
* tsc (>= 3.0.1)
* typescript (>= 3.0.1)
* mongoose (>= 3.6.2)
* redis

# Install, Configure & Run

Below mentioned are the steps to install, configure & run in your platform/distributions.

```bash
# Clone the repo.
git clone https://github.com/ShubhamBansal997/covid-app.git;

# Goto the cloned project folder.
cd covid-app;

# Install go dependencies.
go mod tidy

# To Start Development
export MONGO_URL="mongodb://localhost:27017/myFirstDatabase"
export COVID_API="https://data.covid19india.org/v4/min/data.min.json"
export GEOLOCATION_API="http://api.positionstack.com/v1/reverse?access_key=ACCESS_KEY&query="
export REDIS_HOST="localhost"
export REDIS_PORT=":6739"
export REDIS_PASSWORD="REDIS_PASSWORD"
go run main.go

# To Build
go build -o bin/covid-app -v .

# To Build Docs
swag init
```

# List of Routes

```sh
# API Routes:

+--------+-------------------------+
  Method | URI
+--------+-------------------------+
  GET    | /
  POST   | /fetch-data
  POST   | /get-data?lat=12.12&long=12.12
+--------+-------------------------+
```