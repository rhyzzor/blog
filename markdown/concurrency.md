---
title: "Subroutines, coroutines, concurrency, paralelism and race condition."
slug: concurrency
description: A simple way to learn about concurrency and a little bit about processors.
date: 2025-01-16 18:00:00
---

I could very well deliver a generic answer about Concurrency x Paralelism, but I don't think I'd like that very much, so... last weekend, I studied more about how  processes used to work on our computer, so to start talking about concurrency, I think it's interesting to start with the topic of "subroutines a.k.a procedures"

## Subroutines

As you can see in the drawing below, a subroutine (also known as a procedure) is basically an executable unit that can be called in our main function, for example, let's say we have a main() function that executes a series of processes and these processes are subroutines of our main() function, we have a bob() subroutine and it is called inside the main() function, they will be executed sequentially, that is, main() will call bob() and then return to main(). In other words, main() will call bob(), bob() will be called and then return to main(), but it's also worth remembering that once a subroutine has finished, it will return from the beginning again; if it is called again, it will start again from its initial step, which is very different from coroutines.

![Subroutine Example](/static/images/subroutine.png)

## Coroutines

Coroutines have been around for a long time and back in the day, processors weren't multicore, if I'm not mistaken, the first "popular" multicore processor was launched on the market in 2009, the Intel Core 2 Duo (there was also the IBM POWER4, but it was the Core 2 Duo that popularized this technology among consumers) and knowing this, let me ask you a question: how did our computers manage to run, "simultaneously", our OS, Counter Strike and other processes as well? Basically, coroutines. These processes weren't simultaneous, but their contexts were changed so quickly that they were running in parallel. Coroutines have the ability, unlike subroutines, to "pause", i.e yield. They paused, save their context and other processes could run, which gave us the impression that they were "parallel", but they were always an alternation of tasks, all controlled by our Scheduler. However, our OS didn't always work in this way, there are other ways in which it executes this "task list", this is just one of them, but they are quite similar.

![Coroutine and Subroutine Example](/static/images/coroutine.png)


