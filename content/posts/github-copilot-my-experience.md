+++
title = "What is GitHub Copilot and how to use it"
date = "2022-12-10"
author = "prounckk"
tags = ["Copilot","github","autocompleation"]
keywords = ["github", "coding", "autocompleation", "copilot"]
description = "What is GitHub Copilot? How can it help, and what are the alternatives?"
showFullContent = false
+++

I like to write code, and like many SREs and SWEs, i don't really like to do repetitive things. Copy-pasting sucks, and line-by-line editing is not any penny better. Many tools can help you to write code faster and more efficiently. One of them is GitHub Copilot. In this article, i will try to show how to use it and share my experience.

## What is GitHub Copilot?
Copilot is an AI from GitHub that can help you to write code faster. It can write code based on your comments or just trying to guess what do you have in mind. It is not perfect, but it can help you to save time and be more efficient. Copilot is not a replacement for a human, yet... (evil laugh). Let me show you how it works.

## Let's try it!
Here is an example of how it can help you to write faster:
![Here is an example of how it can help you to write code faster](/2022/copilot-wrote-article-for-me.png "Here is an example of how it can help you to write code faster")

Look how it finished my sentence! I had to make some changes to make it sounds like my accent ;) But in general, it was a good start.

What about writing code?
![](/2022/github-copilot-bubble-sorting-go.gif )
I've tried to run the function that was written by Copilot, and [it worked](https://play.golang.com/p/Z9AJUEOhud6)
Unfortunately, it failed to write a binary search algorithm. I think it is because of the algorithm's complexity, or maybe it is because we do not usually write it, importing already created packages instead.


## What people say about GitHub Copilot
[Samuel Tissot, Senior developer, entrepreneur](https://www.linkedin.com/in/samueltissot/):
>  I removed Copilot for this main reason: It takes me out of the flow. I know what I want to write, and, Copilot writes something very similar but not quite what I want… Thus I need to stop and understand and correct the code. 
Also, sometimes it looks good, but it is not what I want. I'm afraid it's dangerous and might introduce bugs.  I think there is more value in searching for an answer in docs than trusting Copilot.  I think overall, I'm faster without it since I am building a mental map of the code in my head as I write it.

[A friend of mine, lead developer and a shy person](https://www.linkedin.com/in/zarif-safiullin/):
> We should not be afraid of using Copilot or other AI tools. 
It's just an assistant. Developers still need to combine everything written together, then can see the big picture of how it should be interconnected. How to translate business needs to the code.  I will be concerned when Copilot could write a migration from baobab to redux in an 8-year-old project, test it and push it to production without errors.   
In fact, developers, with the help of node modules, Composer or Kubernetes, are already doing this. Outsourcing a lot of tasks to known libraries and tools.


## How much does GitHub Copilot cost
For the day of publishing the post, the price for an individual contributor is $10 per month or $100 per year.


## Alternatives
There are many alternatives to GitHub Copilot. Here are some of them that i've tried:
 - [Vim Autocompleate](https://lual.dev/blog/how-to-use-autocompletion-in-vim/)
 - [Tabnine](https://www.tabnine.com/)
 - Don't forget about default autocompleation in your IDE.

I really liked Tabnine, it has a free plan and what is important, it promised not to use my code for training the public model; as you can see - I help Tabnine be better ;) 

## Concerns and Conclusion
Right after finishing writing the article, my subscription will be over. I will not renew it, not because it was a terrible experience, It wasn't, but I have some concerns.
1. I like what i'm doing, and i really don't want to become just a comment editor. I want to write code, by myself. I don't want to be a code monkey. I am probably not good yet (who is perfect, huh?), but practicing is the only way to improve. 
2. Legally, all code i wrote from the laptop provided by my employer belongs to the hiring me organization. The rules might be different and vary depending on the company, in some companies, even working on weekends from a personal computer can be considered as writing code for the company. So, using Copilot can be a problem here. Imagine you wrote a fantastic function or piece of code that is pure gold. The copilot can pick it up and show it to someone else. We all do copy-paste from StackOverflow, but it's copied from an open source, not from a private company. So pasting code from the Copilot can lead to legal issues: stealing code, plagiarism, etc and a developer without knowing it can put the company in dire straits. I might be paranoid, but I don't want to risk the reputation of the company i work for and my own career. How could soneone prove that the code was written by Copilot and not stealed from the competitors company's codebase? I don't know, but it is a concern.
By the way, it can be disabled in the [settings](https://github.com/settings/copilot), but it is not a default option.

![Disallow GitHub to use my code snippets for product improvements](/2022/github-copilot-settings.png "Disallow GitHub to use my code snippets for product improvements")




