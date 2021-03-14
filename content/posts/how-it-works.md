+++
title = "How it works"
date = "2021-03-13"
author = "prounckk"
tags = ["hugo", "golang", "cloudflare", "github actions"]
keywords = ["hugo", "golang","cloudflare", "github actions"]
description = "How to deploy a hugo website to cloudflare workers from github actions"
showFullContent = false
subtitle = "this is blog post "
+++

I like Simon Sinek's ["Start with Why"](https://www.goodreads.com/book/show/7108725-start-with-why) approach. I feel it fits well with the decision log for most of my life projects. In the previous blog, I've tried to explain WHY I started the project, so today, let me focus on HOW. Probably, it will help me to understand WHAT I'm trying to archive with this.

## Backend
This is a static website. I wanted to have something lightweight and easy to work with. I won't show you the massive list of static site generators. I will just say that Hugo[^1] is fast, written with golang and has clear documentation. I couldn't ask for more.

![alt text](/2021/php5torecicling.webp "Man throws a PHP guide book into the trash can")
### P.S.
I have nothing against PHP; it's a good language and powerful language. This book was just too old 😁

## Frontend
This blog post is a simple MD file that is parsed by Hudo and wrapped with love to a hugo-tania theme[^2]. So far, I just did some twigs at template files and didn't touch any JS or CSS files. Will I do it shortly? I hope not, but I will add some functionality with Rust code compiled into WebAssembly (wasm)


## Hosting
I had a choice between k8s and serverless. Yes, I know, this is so different! But it's where I wanted to improve my knowledge. K8s sounds cool, right? Cool, but expensive. Even with the traffic on the blog that closes to none, I would have to pay for always running containers.
Serverless, on another side, sounds like just the right solution. Free, exciting and modern.
Cough Cloudflare workers. Someone who knows me would understand the joke. CF is what I always recommend at work as a solution for every problem. To be fair, we do use CF workers, but the usage is so limited.  But it's an entirely different conversation.

## Deployment
This is the most exciting part of this post that can be explained in 1 file. But bear with me.
The deployment process is so simple and powerful at the same time.

I had two options: CircleCi or Github Actions. The first one, we use heavily at work. Why should I copy-paste a working solution? I'm not looking for an easy way. 😁 Needless to say, I use for this blog Github Actions.
What we do here.
Check the code from repo [here](https://github.com/Prounckk/eremeev/blob/master/.github/workflows/wrangler-action.yml
)
```YAML
      # Step 1 - Checks-out the repo
      - name: Checkout
        uses: actions/checkout@v2
```
install Hugo and delete old files, just in case 
```YAML
      # Step 2 - Sets up the latest version of Hugo
      - name: Setup Hugo
        uses: peaceiris/actions-hugo@v2
        with:
            hugo-version: 'latest'

        # Step 3 - Clean and don't fail
      - name: Clean public directory
        run: rm -rf public
```
generate static files, so I don't have to make it my local. I can even fix a type in GitHub, push to master and here you go! 
```YAML
        # Step 4 - Builds the site using the latest version of Hugo
      - name: Build
        run: hugo
```
And publishing at cloudflare. Here I'm using Wrangler, an official CLI.
```YAML
        # Step 5 - Publish the generated website to cloudflare Workers
      - name: Publish
        uses: cloudflare/wrangler-action@1.3.0
        env:
            USER: root
        with:
          apiToken: ${{ secrets.CF_API_TOKEN }}
          preCommands: echo "*** pre command ***"
          postCommands: |
            echo "*** post commands ***"
            wrangler publish --env production
            echo "******"
```
All of this is not rocket science but does its job.

Now, based on WHY and HOW I can say WHAT the blog is about. It's my learning platform, where I try things, improve my codding and English skills.

[^1]: [Hugo: A Fast and Flexible Static Site Generator](https://gohugo.io/)
[^2]: [This theme is created by Tania Rascia](https://github.com/taniarascia)