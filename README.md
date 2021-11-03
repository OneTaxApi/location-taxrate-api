# location-taxrate-api

Slightly modified from a demo application showing how to build rest api using golang. See git history for changes.

## Run Locally

To run, from the root of the repo

```
go run .
```

## Access the app 

The App can be extended to have more than one endpoints

All api endpoints are prefixed with `/api/v1`

To reach any endpoint use `baseurl:8080/api/v1/{endpoint}`

```text
Get taxrate by location
"/taxrate/location/{location}" 
```
