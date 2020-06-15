# Sonarqube Badges

[![](https://images.microbadger.com/badges/image/ph1p/sonar-badges.svg)](https://hub.docker.com/repository/docker/ph1p/sonar-badges)
[![](https://images.microbadger.com/badges/version/ph1p/sonar-badges.svg)](https://hub.docker.com/repository/docker/ph1p/sonar-badges)

This small project offers you a way to get your sonarqube badges out of your secured sonarqube instance.

### Environment variables

Set these variables:

- USERNAME (default: "")
- PASSWORD (default: "")
- SERVER_URL (default: "")

I would recommend that you create an additional user within your instance and grant this user permissions.

##### Optional

- ENABLE_REQUEST_LOG (default: false")
- PORT (defaul: 8080)

### Howto

- download **docker-compose.yml**
- set environment variables
- run `docker-compose up -d`
- open `http://localhost:8080/KEY/METRIC.svg`

**METRIC:** a metric key from the list down below

**KEY:** the project key inside your sonarqube instance

#### Metric Types

- bugs
- code_smells
- coverage
- duplicated_lines_density
- ncloc
- sqale_rating
- alert_status
- reliability_rating
- security_rating
- sqale_index
- vulnerabilities
