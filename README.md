Here, I have build apps which are in proper template

multiple unit testcase are simulated to check the code 

It consist of Docker Compose file
    1.For clientserver - deployments\serverclient.yml
    2.For calculator - deployments\calculator.yml

To run go files we need - go run .

To run docker compose files we need - docker-compose up --build

To make coverfile we need -  go test -v -coverprofile cover.out and go tool cover -html cover.out -o cover.html

About:
This repo consists of four apps
1. Heartbeat
2. Calculator
3. Reverse-number
4. Word-counter
