# MYCovidAPI

This is a port of my other MYCovidAPI written in Python. See my [blog post](https://blog.izzudinhafiz.com) on motivation for porting this over to Go.

## Introduction
JSON [API Site](https://mycovidapi.izzudinhafiz.com) for Malaysian Covid statistics hosted at https://mycovidapi.izzudinhafiz.com.

The project builds upon the Covid and MySejahtera dataset by [MoH Malaysia](https://github.com/MoH-Malaysia/covid19-public) as well as the vaccination and registration data by [CITF](https://github.com/CITF-Malaysia/citf-public).

The intent is to allow any developer to easily develop dashboards using the comprehensive data supplied with a JSON friendly interface.

## Documentation
Documentation is hosted on the API itself and can be found [here](https://mycovidapi.izzudinhafiz.com/docs/index.html). There are examples there and further down this document.

## Overview
The current _V1_ API presents all the raw data from MoH as is with no normalization or bucketing. This is to allow developers full access to the underlying data.

The API uses [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format dates `YYYY-MM-DD`.

The API renames some of the fields that was given by MoH to be more consistent and self-descriptive. You can generally base the field name with those in MoH dataset.

Generally for each endpoint:
- For most endpoints, there are `/state/` and `/country/` sub-endpoints
- For endpoints having state data, you need to include `state_id` as below.
- Omitting both `start_date` and `end_date` will return all the data.
- Omitting `start_date` will return data from the start of the dataset.
- Omitting `end_date` will return data up to the current date.

The API uses standardized abbreviation for Malaysian state that is based on this [Wikipedia table](https://en.wikipedia.org/wiki/States_and_federal_territories_of_Malaysia#States).

|    State Name     | State Code |
| :---------------: | :--------: |
|       Johor       |    JHR     |
|       Kedah       |    KDH     |
|     Kelantan      |    KTN     |
|      Melaka       |    MLK     |
|  Negeri Sembilan  |    NSN     |
|      Pahang       |    PHG     |
|   Pulau Pinang    |    PNG     |
|       Perak       |    PRK     |
|      Perlis       |    PLS     |
|       Sabah       |    SBH     |
|      Sarawak      |    SWK     |
|     Selangor      |    SGR     |
|    Terengganu     |    TRG     |
| W.P. Kuala Lumpur |    KUL     |
|    W.P. Labuan    |    LBN     |
|  W.P. Putrajaya   |    PJY     |



## Authentication
No authentication is required. Everyone is free to use this API.

## Rate Limit
There is no rate limit currently, however, we expect users to use the API respectfully. We reserve the right to limit API calls if we detect excessive call rate.

Developer should also cache the data as the data is updated only once a day.


# Examples

1. How do I get a state's new Covid cases for a particular date?

	You can use the same start and end date in the API query.

	E.g. `/api/v1/state/cases?state_id=KUL&start_date=2021-01-01&end_date=2021-01-01`

2. How do I get the nationwide new Covid cases for a particular date range?

	E.g. `/api/v1/country/cases?start_date=2021-01-01&end_date=2021-01-01`

3. How do I get all the data for new Covid cases since the start?

	You can omit the `start_date` and `end_date` in the queries

	E.g. (state) `/api/v1/state/cases?state_id=KUL`

	E.g. (nation) `/api/v1/country/cases`

# Local Setup for Development

## Clone the repo locally
`git clone https://github.com/izzudinhafiz/go-mycovidapi.git`

## Download the required Go modules
```bash
go mod download
```

## Seeding/updating the database

```bash
go run main.go ingest
```

## Start the application
```bash
go run main.go serve
```

# Deploy using Docker Compose

## Clone the repo locally
`git clone https://github.com/izzudinhafiz/go-mycovidapi.git`

```bash
cd go-mycovidapi
touch .env
```

## Write a .env file
```bash
# Dont change the host if you're running this from docker-compose. Else, point it to your database host address
POSTGRES_HOST=covid_db
POSTGRES_USER=postgres
# Make sure you change this password to something strong
POSTGRES_PASSWORD=password
# This does not change the Postgres port address of the spawned Postgres service, this is only for the API server to know which port to connect to
POSTGRES_PORT=5432
# This is the table that it'll create
POSTGRES_DB=mycovidapi
# This is port address of the server. If you change this, ensure you change the port in the docker-compose file
SERVER_PORT=8000
GIN_MODE=release
# Optionally, you can set TZ for easier logging.
TZ=Asia/Kuala_Lumpur
```

## Setup nginx
Use your preferred text editor to create an nginx site file

```bash
sudo nano /etc/nginx/sites-available/mycovidapi.izzudinhafiz.com
```

Create this file:

```nginx
server {
	server_name mycovidapi.izzudinhafiz.com;
	root /home/user/go-mycovidapi/;

	location / {
		proxy_pass http://127.0.0.1:8000;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	}
}
```

Link it to `sites-enabled` and allow nginx firewall access
```bash
sudo ln -s /etc/nginx/sites-available/mycovidapi.izzudinhafiz.com /etc/nginx/sites-enabled

sudo ufw allow 'Nginx Full'
```

## Restart nginx and start the docker instance
```bash
sudo systemctl restart nginx
docker-compose up --build -d
```

# Sources

[Open data on COVID-19 in Malaysia](https://github.com/MoH-Malaysia/covid19-public) - [MoH Malaysia](https://github.com/MoH-Malaysia)

[Open data on Malaysia's National Covid-​19 Immunisation Programme](https://github.com/CITF-Malaysia/citf-public) - [CITF-Malaysia](https://github.com/CITF-Malaysia)

[JSON time-series of coronavirus cases](https://github.com/pomber/covid19) - [pomber](https://github.com/pomber)


# Roadmap
- [ ] Integrate [CITF](https://github.com/CITF-Malaysia/citf-public) data
- [ ] Add consolidated and bucketed data