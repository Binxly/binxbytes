---
title: "How I Stopped Worrying and Learned to Ditch node_modules"
date: "2025-06-27"
category: "notes"
excerpt: "While staring down the barrel of a stack trace after updating my portfolio dependencies, readying myself for an afternoon of treading through hell, I took a step back. Why do I need any of this? The web should be simple."
slug: "the-simple-web"
---

There it was—like a bolt of lightning while I was knee-deep in npm packages, third-party components, Vercel deployment settings, and everyone’s favorite JavaScript library. I was staring down the barrel of a stack trace after updating my portfolio's dependencies, readying myself for an afternoon of treading through hell, when I took a step back. Why do I need any of this? The web should be simple.

![Node Modules Singularity](http://i.imgur.com/lrgCHVu.jpg)

And it wasn't the first time this thought had occurred to me. Before I'd ever heard of Next.js, I'd read articles such as [The Bloated State of Web Development](https://hackernoon.com/the-bloated-state-of-web-development) and [Just Fucking Use HTML](https://justfuckingusehtml.com/). Further back, I remember seeing a [Casey Muratori tweet](https://x.com/cmuratori/status/1426299131270615040) about “dependency culture” in software development, and what it means in the long-term for projects when you add more dependencies to them. He posits “the probability that (a project) build remains working after x years is p^xn, where n is the number of tools used in the build.” If we assume a 90% chance that the tools used in a codebase still work after one year (which is very generous in today's world, as Casey points out), with just 3 tools there is very little chance that things will work as expected after five years.

![90% probability of a tool still operational over time](https://pbs.twimg.com/media/E8s8li8VUAMaT_b?format=png&name=4096x4096)

Nevertheless, Next.js appealed to me at some point, and all I knew at the time was that I wanted to make myself a website. In addition to this shiny framework, I wound up using Vercel to host it for some time, which happens to be the company that created Next.js. Although, I was on edge due to seemingly frequent price hikes in their subscription tiers, and the slew of [questionable billing practices](https://www.linkedin.com/posts/theburningmonk_another-vercel-billing-surprise-unfortunately-activity-7204470386976026624-LZTe/) that had started plaguing popular web applications hosted on their platform. While it's true I'd (likely) never encounter these issues on a personal webpage, it still didn't sit well with me.

So I took a leap from the warm embrace of a managed platform, said my goodbyes to all the chic frameworks, and decided to create a website using a serverless, compiled Go binary running on AWS Lambda. To my surprise (and barring finding my way around the interface for AWS), it wasn't so bad.

## Simplifying

![Node Project vs Go Project](https://i.imgur.com/1TGARUp.png)

Just by looking at these two project directories side-by-side, the difference is obvious. My old Next.js project was dense—tons of files, folders, and layers of abstraction. And that's before adding any serious functionality or getting too heavy into the design. From 62 files down to 18, with room for improvement.

The biggest shift? Cutting out dependency bloat. With Node, I was dragging along a massive `node_modules` directory that had packages for everything: routing, fonts, components, framework features I wasn’t using. Even the stuff I _did_ rely on came with five layers of abstraction. Half the time, it felt like I wasn’t sure which dependency was actually responsible for rendering a button.

In contrast, [my Go project](https://www.github.com/binxly/gofolio) is lightweight, and uses just four dependencies— each with a clear purpose:

- Templ – an HTML templating engine for Go
- algnhsa – plugs Go's net/http directly into AWS Lambda
- gorilla/mux – a popular HTTP router
- goldmark – a markdown parser with frontmatter support

That's it. Just Go code (with some simple HTML for each page), which I compile into a single binary. That binary gets uploaded as my Lambda payload, and whenever an HTTPS request comes in through CloudFront, AWS spins up the binary to handle the request.

## Switching Hosts

Self-hosting was out of the question for me due to limited resources and my power flickering during the mildest of storms, so I reached for the next best thing: AWS. The free tier options offered by Amazon are good enough for my purposes—the hardest part was just figuring out how to navigate their portal and get their available resources running together the right way. Fortunately, there were many online resources for exactly my use case, and with the help of some articles and videos from YouTubers such as [Micah Walter](https://www.youtube.com/watch?v=bAwg-9Fpy4E&pp=0gcJCc4JAYcqIYzv), [codeHeim](https://www.youtube.com/watch?v=aAA4tgkv2cI&t=52s), and [AWS with Jaymit](https://www.youtube.com/watch?v=lw6lG9tV0ZM&pp=0gcJCc4JAYcqIYzv), I was able to grok _most_ of what I needed to know to get a simple page up and running.

![AWS HTTPS Request Handling](https://i.imgur.com/Q2YSaXw.png)

From a high level, whenever someone visits my website, the request goes through Route 53, which handles the DNS for my custom domain. After configuring my domain's nameservers to the ones provided by AWS, and requesting a public certificate for it, this whole setup didn't take long to propagate. Once in place, my domain points to a CloudFront distribution, which is AWS's content delivery network (CDN) that handles HTTPS requests, caching, and routing traffic.

CloudFront receives the request, and if there's no static asset from S3 that it can cache or serve from the edge, it forwards it along to AWS Lambda, where my Go binary lives. When Lambda receives the request, it spins up a container (if one isn't already warm), runs the binary, and pipes the request in using algnhsa, which makes it all work like any typical net/http Go server.

Thanks to Templ and Goldmark, the Go binary handles the request, renders the requested page, and sends back a raw HTML response. That response flows back out through CloudFront, and presto, everything's automagically sent to the user's browser.

There's no node process watching for changes, no runtime React stuff booting up, no Vercel deployment pipeline I need to watch like a hawk — just a binary, waiting to be invoked.

## Notes

I can safely say that migrating to AWS taught me a lot more about the moving parts behind a modern website than any managed hosting platform ever did—I had to configure things that platforms like Vercel intentionally abstract away. With how much I've enjoyed creating my own [interpreter in Go](https://www.github.com/binxly/monkey-interpreter), and some one-off TUI applications using [bubbletea](https://github.com/charmbracelet/bubbletea), I was itching to try recreating my website without all the fuss.

There's still plenty of room for improvement: many things I'd simplify further, scuffed workarounds I did just to get a minimally viable page going, and I'll be redoing my tailwind implementation at some point. For now, it works, it's mine, and I'm happy with the progress I made. Throw in a script that runs on a cronjob that looks for changes to the directory to push changes through to production, and whatever other quality of life improvements I can think of, I'm sure it'll be even better in the future.

I'm not saying everyone should rewrite their site in Go and run with their own hosting setup. But for me, this was about more than just picking a host, or using a certain language. I wanted something simple, something smaller, that just _works_.

Turns out, building for the web can still be simple. Who knew?
