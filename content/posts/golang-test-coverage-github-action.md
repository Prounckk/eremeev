+++
title = "Golang & GitHub Actions: Test coverage and show badge"
date = "2022-11-02"
author = "prounckk"
tags = ["golang"]
keywords = ["golang", "benchmark", "test", "coverage", "github actions"]
description = ""
showFullContent = false
+++

The tests are great. Period.  
I guess i could stop the blog post right here. There is nothing to talk about. But we, devs, love to make things more complicated, right? We like dashboards, data, numbers and graphs. Speaking about tests, we like to see how well our code is covered by tests. There is plenty of services to test code, generate badges and show the data. But what if the budget is tight and the code should not leave the repository, for a legal reason, for example? 

## Let's test!
The simplest solution i have found is this one: 
```
go test ./... -coverprofile coverage.out
```
This will test all files and generate a file `coverage.out` with results. You can name the file as you wish.  


```
COVERAGE=`go tool cover -func=coverage.out | grep total: | grep -Eo '[0-9]+\.[0-9]+'`
echo $COVERAGE
```
Then parse the results and assign them to the new variable `COVERAGE`. Again, feel free to rename it as you wish.

## Let's build!

Good, now we have a number, and with pride or shame, we would like to show the number as a badge in the GitHub repo. 

```
COLOR=orange
if (( $(echo "$COVERAGE <= 50" | bc -l) )) ; then
    COLOR=red
    elif (( $(echo "$COVERAGE > 80" | bc -l) )); then
    COLOR=green
fi
```
and combine everything to an actial badge
```
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > badge.svg
```

It should all looks like somthing like this: 
```
go test ./... -coverprofile coverage.out
COVERAGE=`go tool cover -func=coverage.out | grep total: | grep -Eo '[0-9]+\.[0-9]+'`
echo $COVERAGE
COLOR=orange
if (( $(echo "$COVERAGE <= 50" | bc -l) )) ; then
    COLOR=red
    elif (( $(echo "$COVERAGE > 80" | bc -l) )); then
    COLOR=green
fi
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > badge.svg
```

## Let's show!
So, we have a badge generated and stored in the GitHub Actions workspace. The problem with this i can't replace the default GitHub Actions badge. It should be stored now somewhere. 
A good option would be to store it on AWS S3, GCP bucket or Cloudflare R2. But let's make it more complicated and store badge on GitHub and make our GitHub Actions  to commit the file back

```
git add badge.svg
git commit -m "added badge"
git push 
```
Sounds simple, right? Let's test!   
![ooops, main branch is protected](/2022/github-action-failed.jpg "ooops, main branch is protected") .  
Right, all commits to the main branch must be reviewed. The simplest solution would be to turn off protection... and let everyone merge to the main branch what they want? No bueno!  
Let's make a dedicated branch for the badge! 
```
git fetch
git checkout badge -f
git pull
git merge origin/main  
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > badge.svg
git add .
git commit -m "added badge"
git push 
 ```
and update README file to have a link to the badge
```
![Coverage](https://github.com/Prounckk/ZTM-DS-and-Algo-Golang/blob/badge/badge.svg?branch=badge)

```

On more push to the repo and ... It worked!  
Needless to say, don't copy-paste the link to your readme file unless you want to share my shaming results.