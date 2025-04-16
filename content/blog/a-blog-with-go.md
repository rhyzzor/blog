---
ogImage: 
  component: BlogPost
  props:
    title: "A blog with Go"
    description: "I've created a blog with Go. This is not a code tutorial."
sitemap:
  lastmod: 2025-04-15
robots: index, follow
schemaOrg:
  - "@type": "BlogPosting"
    headline: "A blog with Go"
    author: 
      type: "Person"
      name: "Rhyzzor"
    datePublished: "2025-01-02"
description: I've created a blog with Go. This is not a code tutorial.
date: "2025-01-02T18:00:00Z"
title: A blog with Go
---
I remember the first time I saw a blog, it was a blog with Tibia content and I was really impressed. After that, I knew I wanted to make my own blog, but I didn't know what the content I should write, I thought about having a gaming blog, a Tibia and Habbo blog, but I always left that "dream" for later. Now, many years later, I am finally creating my blog to share my journey.

And the design of the blog is deliberate. It reminds me of the old Tibia blog I used to follow. That's the part I like best, it reminds me of my childhood

## Why Go?

Well, my first decision it's made a blog with Nuxt (honestly, I really like Vue), but I was already spending hours using JavaScript in my job, so I should learn something new outside of work and for this year, I decided to improve my knowledge of Go. I also use Go for some things in my work, but not as much as JavaScript and I'm also sick of so many JavaScript frameworks (nothing against React, it's just my opinion).

And there's one more reason to learning Go. Today, I'm a Full-Stack Developer, but I don't really like Front-end, I just do it because it's my job. This year, I've decided to change that and improve my skills as a Back-end. I'm revisiting a lot of topics that I saw in my college, but I've forgotten over time and also several things about DSA.

## Nginx + Certbot vs Fly.io

I was talking to a friend about setting up a Lightsail with Nginx + Cerbot and Docker, so... my first decision was to use Lightsail with these tools, but I needed to find something easier to host a blog, because doesn't need a "hard work" to be launched (and I needed something cheaper too, the dollar in Brazil is expensive and Fly.io costs me only 3 dollars per month).

And I found the Fly.io, it's a very good option to host my blog and it's very simple to set up, the only thing that I need to do is to run the Fly CLI and remove the backslashes from Dockerfile, because the Github Actions throws an error when the backslash is not removed, and generate a certificate to my domain and blog. It’s far easier than configuring Nginx and Certbot, which, the first time I tried, took me hours. I’d rather not go through that experience again.

## Next steps

I’ve also decided to improve my English, so all my posts will now be written in English. It’s a great way to expand my vocabulary and practice writing better texts. My plan is to publish a new post every week, whether it’s about a current idea or some technical knowledge I’ve gained. My next post will be about **“The Basics of System Design”**, a crucial topic for everyone in the field.