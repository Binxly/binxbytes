---
title: The Vibe-Coder's Dilemma – How AI Is Dangerous for Juniors and What To Do About It
date: 2025-03-18
category: thoughts
excerpt: The dangers of a recent trend for newcomers in software development.
slug: vibe-coders-dilemma

---

> In response to a post by Geoffrey Huntley: [Dear Student: Yes, AI is here, you're screwed unless you take action...](https://ghuntley.com/screwed/)

## The Rise of AI and "Vibe-Coding"

In 2025, the impact of AI on software development is undeniable. Many knowledgeable people are using it to accelerate the development of their projects, usually by having it do grunt-work such as creating unit tests or writing small functions for simple operations. However, SOTA models have reached a point where they have enabled a new fad in development called "vibe-coding". This is where a user tells an AI what they want to build, blindly accepts generated code suggestions, copy-pastes traceback errors into the context window whenever they need to debug, and development continues in this way— fully "embracing the vibes" until their application either comes to fruition, or their codebase is so tangled in knots that they `rm -rf` the whole thing and move on.

### "Expert Beginners" - The Danger

While this new paradigm of programming may seem interesting to seasoned developers, I believe it's disastrous for junior-level developers and those who are just beginning to explore the field. Many times, newcomers lack the experience required to understand fundamentals, sometimes computers in general, causing them to make nonsensical requests to AI agents so they can feel like "10x engineers". Anytime I see threads online where beginners dig themselves into holes trying to use ChatGPT for their *ambitious* projects, I'm reminded of my first week in college; I was surprised that much of my class was unfamiliar with common hotkeys and operations which were second-nature to anyone who grew up around a computer. I can't imagine how lost those same people would feel after they've confidently copy-pasted hundreds of lines of slop from their chat interfaces to an IDE.

![localhost](https://preview.redd.it/ai-is-taking-over-v0-l2pzxryu5gxa1.jpg?auto=webp&s=5698e400ee5dd1282ca511d4baee9acfb430800b)

Worse, folks who were junior-level developers prior to AI's growth are admittedly creating much more complex projects, yet *fail to understand the higher-level flow of information within it*. Either that, or their codebases are riddled with unused imports, variables, libraries which they can't tell you anything about, and sometimes entire functions that are left completely isolated. Of course, bring this up and there are always things you can do to improve the output of these tools—

- The latest forks of VS Code with bigger and better sidebars covering your screen.
- New Ultra-GPT-5o-o4-blueberry-mini-high.SFT powered search engines.
- New prompt engineering practices released every 30 minutes on X.

—but seriously, there's a real trend happening, and more users heavily adopting and relying upon this technology every day.

> "We were promised flying cars but got comment spam bots."
> 	- u/[kpetrie77](https://www.reddit.com/r/GenX/comments/1g8bumd/comment/lsxadcr/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button)

## The Nuances of Learning Software Development

What makes programming difficult is simple: it challenges the way you learn. Libraries never quite do exactly what you want them to, computers are often times less predictable than they ought to be, and you have to continually solve problems that you create for yourself every other day. When in the middle of large projects, you frequently get confronted with the fact that your week-old code actually sucks, forcing you to rewrite it, and this process is how you learn and absorb the most to become an invaluable developer.

### Proposing Solutions for Juniors Adopting AI

Juniors must approach AI carefully to ensure they aren't at risk of becoming reliant on it. Some people outright tell everyone not to use AI at all, but I think that could hinder your growth in the modern day as AI becomes more capable and widely adopted, yet there's also something to be said about others encouraging everyone to use AI everything. These tools are beginning to appear in seemingly *every corner of the internet*, and most people will try them out, at some stage.

I believe the solution for healthy adoption lies in the degrees of separation you place between yourself, any specific problem, and the AI model. Newcomers must shift their perception of modern AI to align with their learning journey, rather than obsess over rapid development of prototypes. Specifically, *preventing AI from giving you direct solutions*, and instead treating it as more of a *scratchpad* to help you think through problems as you run into them. This makes more sense once you gain a basic understanding of how LLMs work, and begin to *pair their objective strengths with your own unique approach to learning*.

## Using AI to Learn - Rubber Duck Debugging

![Rubber Duck Debugging](https://storage.googleapis.com/duckly-blog/2019/10/IMG_7540.jpeg)

Whenever I open up a new chat with my favorite models (either local or on the cloud), I like imagining it as a new-age form of [rubber duck debugging](https://en.wikipedia.org/wiki/Rubber_duck_debugging). This is an evolution of the classic debugging technique—one where the "duck" not only listens, but also responds, questions, and occasionally surprises me with insights I hadn't considered. It transforms problem-solving from a solitary, internal process into a free-flowing conversation. In a way, it's not just about fixing bugs anymore—it's about refining your thinking, reinforcing core concepts, and overcoming particular challenges.

### Preventing Yourself From Falling Behind

For someone just beginning to learn about software development in 2025, it takes mindful practice, intentional studying, and a desire to keep pushing yourself. Juniors *must* tamp their AI usage down into more specific scopes.

- **DO NOT** use it to provide code snippets.
- **DO** have it provide direct links to documentation.
- **CONSIDER** using it to quickly access specific areas of documentation, in a *proper RAG application setting*.
- **DO NOT** take everything AI provides you at face-value.
- **DO** carefully check any summaries, asking for direct links to pages where it sourced information from.
- **DO NOT** have it design your application's architecture from the ground-up.
- **DO** use it to talk about the viability of ideas and solutions **you** come up with.

This goes back to the rubber duck analogy. Your goal should be to use AI in ways that will encourage you to come to your own conclusions, while taking advantage of the vast amounts of information readily available. Do not take shortcuts, but feel free to ask questions if you come across something that makes you feel stumped. Don't forget that you should be building projects for the purpose of learning more about your stack, rather than automating a million versions of the same Top 10 Portfolio Projects we've all seen before.

### Deciding to Generate Code - Things to Avoid

If you consider yourself skilled when it comes to programming, you may feel inclined to use AI for code generation anyways. However, instead of using AI generated code everywhere, think carefully about *where* in your codebase you use it. Under no circumstances should you allow changes to things such as:

- Security-critical code
- Core proprietary business logic
- Database schemas
- API contracts and public-facing interfaces
- Performance-critical code

You should likely not use AI generated code in these areas **at all** if you're using public APIs that generate text for you. There's a high chance that your messages are not kept private and could be used for training in the future, which could wind up putting your company, project, or private information at risk.

![Dangers of Vibe-Coding](https://i.redd.it/acp1k9csabpe1.png)

## Closing Thoughts

While I find deep learning fascinating, and there are many capable models available that can help people complete certain tasks, I think that new students and graduates alike must reconsider how they approach AI text generation. Using an AI model to shortcut your own learning process doesn't benefit you, and no tool is worth using if it means sacrificing your own autonomy and diminishing industry-critical skills.

