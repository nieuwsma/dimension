# Readme

Prerequisites 
- Have docker installed

1. [SwaggerFile](api/swagger_training.yaml) - you can use [editor.swagger.io](editor.swagger.io) to view
1. Execute [./setupDeveloperEnvironment.full.sh](setupDeveloperEnvironment.full.sh) to stand up the webserver on port: 8080
1. Execute [./teardownDeveloperEnvironment.full.sh](teardownDeveloperEnvironment.full.sh) to tear down the webserver
1. `make image` to create the docker image
2. `make autoplayertest` to create the autoplayertest binary image
    3. `export ROUNDS=100 #sets rounds to 100; else default is 100,000`
    4. `./bin/autoplayertest #to execute`

