+++
title = "GitHub Copilot vs Tabnine vs ChatGPT: What is it and how to use it"
date = "2022-12-22"
author = "prounckk"
tags = ["Copilot","github","autocompleation"]
keywords = ["github", "coding", "autocompleation", "copilot", "chatGPT", "Tabnine"]
description = "What is GitHub Copilot? How can it help, and what are the alternatives? What is the funcy chatGPT? Can AI take my job?"
showFullContent = false
+++

I like to write code, and like many SREs and SWEs, i don't really like to do repetitive things. Copy-pasting sucks, and line-by-line editing is not any penny better. Many tools can help you to write code faster and more efficiently. Some of them are GitHub Copilot, Tabnine and ChatGPC. In this article, i will try to show how to use AI in the code writing and share my experience.

## Tools for developers
There are some tools i've tried:
- [GitHub Copilot](https://github.com/features/copilot)
- [Tabnine](https://www.tabnine.com/)
- [ChatGPT](https://chat.openai.com/chat)
- [Vim Autocompleate](https://lual.dev/blog/how-to-use-autocompletion-in-vim/)
- Don't forget about default autocompleation in your IDE.

## What is GitHub Copilot?
Copilot is an AI from GitHub that can help you to write code faster. It can write code based on your comments or just trying to guess what do you have in mind. It is not perfect, but it can help you to save time and be more efficient. Copilot is not a replacement for a human, yet... (evil laugh). Let me show you how it works.

## Let's try it!
Here is an example of how it can help you to write faster:
![Here is an example of how it can help you to write code faster](/2022/copilot-wrote-article-for-me.png "Here is an example of how it can help you to write code faster")

Look how it finished my sentence! I had to make some changes to make it sounds like my accent ;) But in general, it was a good start.

What about writing code?
![](/2022/github-copilot-bubble-sorting-go.gif )
I've tried to run the function that was written by Copilot, and [it worked](https://play.golang.com/p/Z9AJUEOhud6)  
Unfortunately, it failed to write a binary search algorithm. I think it is because of the algorithm's complexity or because we do not usually write it, importing already created packages instead.


## What people say about GitHub Copilot
[Samuel Tissot - Senior developer, entrepreneur](https://www.linkedin.com/in/samueltissot/):
>  I gave GitHub Copilot a good try but I had to remove it for three main reasons. First of all, Copilot takes me out of the "flow". As I write code, I build a mental map of the code and its data flow. However, Copilot will often write / suggest a fairly good chunk of code that I need to stop and analyze to make sure it conforms to what I want to create. The code is often very similar to what I would write with the right variable names and all, but accomplishes something different that I originally intended. On this regards if I'm not diligent at analyzing the code that Copilot propose, I might inadvertently introduce bugs or errors in the code. The other reason is that I prefer researching for an answer in the code documentation which will often yield a better understanding of a feature and result in better code down the line.  The third main reason is a code-style preference, I have an opinionated view on how the code should be structured, and it mostly follows [the "Clean Coder" way](https://www.amazon.ca/Clean-Coder-Conduct-Professional-Programmers/dp/0137081073/ref=sr_1_1?crid=2XDP6XRZS58XH&qid=1671729626&sr=8-1) but with Copilot I will find myself having to rewrite some part of the suggestion anyway to conform to my coding preferences.  
So overall I think despite Copilot having some good suggestions I'm definitely faster without it and write better code without it.

[A friend of mine, lead developer and a shy person](https://www.linkedin.com/in/zarif-safiullin/):
> We should not be afraid of using Copilot or other AI tools. 
It's just an assistant. Developers still need to combine everything written together, then can see the big picture of how it should be interconnected. How to translate business needs to the code.  I will be concerned when Copilot could write a migration from baobab to redux in an 8-year-old project, test it and push it to production without errors.   
In fact, developers, with the help of node modules, Composer or Kubernetes, are already doing this. Outsourcing a lot of tasks to known libraries and tools.


## How much does GitHub Copilot cost
For the day of publishing the post, the price for an individual contributor is $10 per month or $100 per year.

## Tabnine
I really like Tabnine, it has a free plan and what is important, it promises not to use my code for training the public model; as you can see - I help Tabnine be better ;) I feel it is less annoying and more predictable. I see it as an advanced auto-completion.   
In the same test case of bubble sort, Tabnine failed to write from scratch but held with obvious and repetitive things:
![](/2022/tabnine-bubble-sorting-go.gif )


## Chat GPT
This is WOW! Honestly! I'm impressed with how AI has advanced in the last few years. ChatGPC is the most advanced (as i know and might be mistaken) publicly available AI chat.   
It helped us to find a big issue we had with the Logsstash config file and showed us how to fix it. What i think was interesting in this experience, chatGPT took the problem and oversaw it with fresh eyes. No emotions, no background knowledge of why this was designed this way. It checked what the settings are now, what kind of problem it creates, and what the documentation said. I wish I could act like this!
But let's go back to our bubble sorting algo! 
![](/2022/chatgpt-bubble-sorting-go.png )
It's still O(n^2), but looks slightly better than other AIs. Did you notice the variable `n` that helped not to check the length of the array every time? Did you notice my typo in the question? The chatGPT was ok with it.  ¯\\_(ツ)_/¯ 


## Concerns
Right after finishing writing the article, my subscription to Github Copilot will be over. I will not renew it, not because it was a terrible experience, It wasn't. I won't, but I have some concerns.
1. I like what i'm doing, and i really don't want to become just a comment editor. I want to write code, by myself. I don't want to be a code monkey. I am probably not good yet (who is perfect, huh?), but practicing is the only way to improve. 
2. Legally, all code i wrote from the laptop provided by my employer belongs to the hiring me organization. The rules might be different and vary depending on the company, in some companies, even working on weekends from a personal computer can be considered as writing code for the company. So, using Copilot can be a problem here. Imagine you wrote a fantastic function or piece of code that is pure gold. The copilot can pick it up and show it to someone else. We all do copy-paste from StackOverflow, but it's copied from an open source, not from a private company. So pasting code from the Copilot can lead to legal issues: stealing code, plagiarism, etc and a developer without knowing it can put the company in dire straits. I might be paranoid, but I don't want to risk the reputation of the company i work for and my own career. How could someone prove that the code was written by Copilot and not stealed from the competitors company's codebase? I don't know, and this is my concern.
By the way, it can be disabled in the [settings](https://github.com/settings/copilot), but it is not a default option.

![Disallow GitHub to use my code snippets for product improvements](/2022/github-copilot-settings.png "Disallow GitHub to use my code snippets for product improvements")


## Conclusion
I've asked chatGPT about the article: 
> It looks like a well-written article that discusses the use of AI tools such as GitHub Copilot, Tabnine, and ChatGPT in code writing and how they can help developers save time and be more efficient. The author also shares their personal experience with using these tools and provides some insights into their limitations and potential drawbacks. Overall, the article provides a balanced and informative perspective on the use of AI in code writing and can be useful for developers who are interested in exploring these tools.

Who are we to argue about this article with AI ;) Share with your friends! 
