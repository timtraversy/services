+++
title = "Creating a Cloud Development Environment"
date = "2021-06-20"
draft = true
+++

## Background

For the past half-year I've been [Dockerizing](https://www.docker.com) most of the development I do for work and personal projects. It's satisfying to be able to version control my dependencies, easily share projects with others, and experiment with packages without polluting my host machine, 

Then I came across Mark O'Connor's [blog post](https://yieldthought.com/post/31857050698/ipad-linode-1-year-later) about setting up a development server that he SSH's into from his iPad. By moving computation-heavy tasks to the cloud, he can work from any machine that has a text editor and an internet connection. His description of the freedoms this gave him inspired me to attempt something similar.

## What I Wanted
* Docker running applications as processes like it was meant to, instead of inside a VM as it has to do on the Mac.
* Be able to program from any number of clients with different operating systems, memory, and CPU.
* Be able to work even from lightweight clients like tablets or even smartphones, when in a pinch.
* Have a single source of truth for version control history. Git allows you to sync progress across machines pretty well, but you still sometimes forget to publish a branch or push your latest change to the central server.

### Single Source of Truth


### Reproducible Environment

Checking Dockerfiles into version control gets you pretty close to this. But 
NixOS 
- ARM 

### Lightweight Clients

I don't imagine I'd transition to an iPad for full time 
    
## How I Did It