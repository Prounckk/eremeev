+++
title = "GCP App Engine: The Good, the Bad and the Ugly"
date = "2021-03-20"
author = "prounckk"
tags = ["appengine", "gcp"]
keywords = ["app engine", "gcp","Google Cloud App Engine"]
description = "App Engine is an excellent service available at Google Cloud, but is it as good as it's described?"
showFullContent = false
+++

As a small team that wanted to be more focused on code than on infrastructure, we tried several hosting platforms for our main website and microservices. We went from WPEngine, stopped by GKE and ended up with Google App Engine. Here I'd like to summarize my experience with App Engine.

## The main evaluation criteria
### Scalable
The marketing website should be scalable. We can hope, but we can't predict how successful will be one of our campaigns, we can't predict how many visitors we get after announcements of big news. But at the same time, we don't want to pay for underused resources.
### Easy to use
It should be easy to deploy, rollback to previous versions, straightforward rules to make the canary deployment. We want to have control and ssh, but we don't want to be in charge of hardware, security updates and maintenance.
### Reliable
Uptime and response time our main SLIs at marketing. We love our visitors and want to make sure they are happy with the website's speed and availability.
![alt text](/2021/app-engine.jpeg "Google Cloud as a service started from App Engine")

## App Engine in a nutshell
Google Cloud as a service started from App Engine, which has pros and cons.
* App Engine has excellent documentation, but sometimes you can find an out-of-date piece of information.
* App Engine runs on VMs, but you have limited access to their settings as a developer.

When we talk about AppEngine, important to mention that there are two main types of service: Standard and Flexible.
### Standard
Standard is good to go option for non-tech-savvy people. It's running with prebuild docker images, has a limited range of options for programming languages. For example, [PHP7.4  recently become available](https://cloud.google.com/appengine/docs/standard/php7/runtime#php-7.4)

It's limited for customization, but it has all that you need for work.

An important thing to mention, Standard AppEngine might be scaled down to 0 instances, so you pay nothing if the service is not getting any requests.
Sounds good, right?
### Flexible
On the other hand, Flexible App Engine is really customizable; write your dockerfile with your favourite language version and all dependencies you'd like to have.
Keep in mind; it will always be running with at least one instance. The freedom costs money ;)

[Read more about the difference between them](https://cloud.google.com/appengine/docs/the-appengine-environments)
## Pros and Cons
### Be aware
This all sounds amazing, but nobody is perfect, so App Engine is not an exception to the rule. Here are some points to consider before migration to App Engine
* Be aware when you create your first App Engine service. Even if you delete the service, you won't be able to change the App Engine services' location. It sets once and forever. No way to change it. No way to have two App Engine services running at different locations. Need to spin up a new app Europe, for example, create a new GCP project. Sad, but true.
* Your Cloud SQL must be at the same zone as your App Engine instance.
* Cloud Tasks must be at the same zone as your App Engine instance.
* Deployment is slower than Compute Engine even if you prebuild docker image at your Ci pipeline.
* There is a limit of 10k files per application. Do you think it is a good idea to deploy your node_modules folder? Think twice.
* Your secrets are not secrets anymore. Everyone with access to GCP can preview environmental variable sets for App Engine. Just click on view configs.
  ![alt text](/2021/app-engine-view-secrets.png "Everyone with access to GCP can preview environmental variable sets for App Engine. Just click on view configs")

* One instance per service will have regular downtime. Flexible environment VM instances are restarted weekly. During restarts, Google's management services apply any necessary operating system and security updates.
* App Engine won't let you use 99% of CPU or memory as GKE; it's not bad, it's not good. It's just a way how App Engine scales up and scales down. By default, target_utilization has a default value of 0.5. It is 50% of CPU usage; you might change it to 0.9.
* The most annoying limitation, you can define up to 20 routing rules. That is it. Need more? Spean up a new project. Arghhh!

### Be happy
There is some good news:
* App Engine support Unix domain socket to cloud SQL, so the Unix domain sockets are often twice as fast as a TCP socket when both peers are on the same host
* CloudTask works amazingly well with App Engine and even can autoscale your instance based on a load of Cloud Tasks!
  ![alt text](/2021/google-tasks-and-app-engine.png "Google Cloud Tasks works amazingly well with App Engine and even can autoscale your instance based on a load of Cloud Tasks!")


## Our evaluation
* The deployment process is simple and predictable. An app.yaml file has all settings for the app and [can de deployed to GCP from local with one command](https://cloud.google.com/sdk/gcloud/reference/app/deploy) or from CI pipeline as we do. 
* A database connection is faster than from VM. Sorry, I'm lazy to provide you with the actual number, but it's significantly quicker than VM connected to CloudSQL with proxy.
* Switch from K8S to App Engine reduced bus factor and let team sleep better. I don't remember to be pages cause of infrastructure failure.


##  Final thoughts
* AppEngine is a good choice if you don't need to scale up your application across regions;
* It's a great option and probably the best option if you don't want to spend time/money writing and maintain a bunch of bash scripts, HELM config and be lost in yaml files.
* And this is a perfect solution for a small budget startup, where traffic unpredictable and, most of the time, close to none  (like my blog haha) 