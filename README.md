![security rating](./badges/securityRating.svg) ![vulnerabilities](./badges/vulnerabilities.svg)

# HackerEarth Hackathon Project for team SiteTech
This is all the code and documentation for our submission.

Team SiteTech

# Local Development Commands

## Install pre-commit
```
brew install pre-commit
```

## Run Application Using Doppler
```
go build -o bin/togo-read -v .
./bin/togo-read
```

## Run Tests
```
go test -v ./tests -coverprofile=./coverage.out -coverpkg ./...
go test ./tests -coverprofile=./tests/coverage.out -coverpkg ./... -json > ./tests/test-report.out
```

## See Coverage Report
```
go tool cover -html=./tests/coverage.out
```

## Run SonarQube
```
brew install sonar-scanner

sonar-scanner \
  -Dsonar.projectKey=togo-read-micro \
  -Dsonar.sources=. \
  -Dsonar.host.url=${SONAR_HOST_URL} \
  -Dsonar.login=${SONAR_LOGIN}
```

## Update SonarQube badges
```
curl ${badge_url} > ./badges/securityRating.svg
curl ${badge_url} > ./badges/vulnerabilities.svg
```