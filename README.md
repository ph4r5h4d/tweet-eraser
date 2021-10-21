[![Tweet Eraser Build](https://github.com/ph4r5h4d/tweet-eraser/actions/workflows/cd.yaml/badge.svg)](https://github.com/ph4r5h4d/tweet-eraser/actions/workflows/cd.yaml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ph4r5h4d/tweet-eraser)](https://github.com/ph4r5h4d/tweet-eraser)
[![Docker pull](https://img.shields.io/docker/pulls/ph4r5h4d/tweet-eraser)](https://hub.docker.com/r/ph4r5h4d/tweet-eraser)
[![App release](https://img.shields.io/github/v/release/Ph4r5h4d/tweet-eraser)](https://github.com/ph4r5h4d/tweet-eraser/releases)
[![App License](https://img.shields.io/github/license/ph4r5h4d/tweet-eraser)](https://github.com/ph4r5h4d/tweet-eraser/blob/main/LICENSE)
# tweet-eraser
This small app helps you delete your old tweets

# Important notes
- This app is designed to run from your local machine (laptop, pc or ...). It will not store any information anywhere at all.
  - You are advised to check the source code and build the application yourself if you want to be entirely sure.
  - The builds and Docker images are solely there to make your life easier, but I highly advise you to build the application yourself.

- This Application is provided as-is. I will not be responsible for any possible problem that may happen to your data and account.
  I tried my best to make sure it won't harm anything, but Twitter may change its API or put limitations which
  I'll then try to react and make sure this app will continue to work for as long as possible, But there will be no guarantee.


## Requirements
Just download the app from the release page, and you should be good to go.

If you are crazy enough to build the app, you need to install the latest version of the Go and clone the repo, and just run:
```
go build -o te
``` 
And there should be a binary package ready to be used by you or any alien species using the same OS.

## Running the app
There are important parameters you need to pass to the app for the app to work correctly.  
Let's go over the parameters and see what they are and where you should obtain them.

| Command line parameter | description | how to retrieve |
|------------------------|-------------|-----------------|
|--file|The path to your Twitter archive|Just put it somewhere and then use the absolute path as the value|
|--authorization|Your account Bearer token|You need to retrieve it from a network request from your browser|
|--authToken|Your account auth token|You need to retrieve it from your Twitter cookie (auth_token)|
|--csrfToken|Your session CSRF token|You need to retrieve it from your Twitter cookie (ct0)|
|--offset|How many **days** to preserve your tweets|-|

A sample would be like the following:
```bash
./te --file=twitter.zip --offset=730 --authorizaion=YOUR_BEARER_TOKEN --authToken=YOUR_COOKIE_AUTH_TOKEN --csrfToken=YOUR_COOKIE_CSRF_TOKEN
```

Look at the `--offset=730`, this means all tweets older than 730 days (2 years) will be deleted.

## Using the Docker image
First, pull the image:
```bash
docker pull ph4r5h4d/tweet-eraser:latest
```
Well, put the file in a directory you want to run the command from and do the following:
```bash
docker run -v $(pwd)/YOURT_BACKUP.ZIP:/app/twitter.zip ph4r5h4d/tweet-eraser:latest --file=twitter.zip --offset=730 --authorizaion=YOUR_BEARER_TOKEN --authToken=YOUR_COOKIE_AUTH_TOKEN --csrfToken=YOUR_COOKIE_CSRF_TOKEN
```