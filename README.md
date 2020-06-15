# Sonarqube Badges

### Environment variables

- USERNAME (default: "")
- PASSWORD (default: "")
- SERVER_URL (default: "")
- ENABLE_REQUEST_LOG (default: false")
- PORT (defaul: 8080)

### Howto

#### URL

```bash
$SERVER_URL:$PORT/[project_key]/[metric].svg
```

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
