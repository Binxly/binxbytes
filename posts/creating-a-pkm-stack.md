---

title: "Creating a PKM Stack - My Approach After 2 Years"
description: "Personal Knowledge Management: What worked, what didn't, and how to stop wasting time."
date: "2025-04-04"
category: "thoughts"
author: "zac"

---

Years ago, I spent an embarrassing amount of time trying to come up with my own method for personal knowledge management (PKM) and tinkering with various applications advertised to streamline productivity. It all started while browsing documentation and reading articles one day, and I realized many of the things that I'd read while self-studying often faded from memory. I'd come across a problem in a project that I'd previously spent a lot of time reading about, and only vague details would surface. Something had to change, and keeping dozens of physical notebooks (even though most of my work lives on my computer) simply didn't scale well. Not to mention, I rarely referred back to my notebooks because they would eventually become lost in a box of all the others from the last 10+ years. Meanwhile, any digital notes I took at the time were lost inside nested folders, and good luck finding them months later amongst some ~30+ TB of hard drive space.

### The Productivity Influencer Trap: A Warning

![influencers](https://www.vice.com/wp-content/uploads/sites/2/2021/01/1611226953994-screenshot-2021-01-21-at-105850.png)

If, like me, you've ever googled something like “best practices for taking notes”, you've likely encountered this endless rabbit hole. You will see things like [Building a Second Brain](https://www.buildingasecondbrain.com), [The PARA Method](https://fortelabs.com/blog/para/), [Zettelkasten Workflows](https://zettelkasten.de/overview/), [Creating Notion Dashboards](https://www.notion.com/help/guides/personal-work-dashboard), [Mind-Mapping](https://www.lucidchart.com/pages/tutorial/how-to-make-a-mind-map), etc. Depending on where you start, it can be all too easy for you to watch a video, or read an article (oops), and become convinced that the author's approach to tracking and organizing information is the end-all. The truth is ***different people require different systems***. While these methods may work for some, you may feel differently when trying to put them into practice.

I was certainly not immune to this rabbit hole when I first tried to tackle the digital productivity problem for myself. Far too often, I would find myself getting excited about a new method of note taking, new bullet journal ideas, or a new system of digital organization, only to find myself abandoning it after the better part of a week. Of course, this was all because it didn't *actually* mesh well with the way that I learned— much less the way I preferred to use the tools available to me. All these methods were either too granular, too broad, or required too much effort to maintain for me to actually *want* to use them.

## What Actually Worked: KISS (keep it Simple, stupid)

![simplicity](https://focusandaction.com/wp-content/uploads/2023/08/applenotesmeme.jpeg)

Instead of spreading my notes out across more physical notebooks, various apps on my phone, and different software on my desktop, I decided to consolidate everything into two key places, and have kept it that way for the last couple years: [Obsidian](https://obsidian.md) for digital notes and books, and a folder in the root directory of each of my devices where I put all of my coding projects and [GitHub](https://github.com) repositories.

### Obsidian - My Digital Library and Notes

![obsidian graph](https://forum.obsidian.md/uploads/default/original/3X/3/3/338d6f9cb03a40154eb4ada379ab725de934e678.png)

[Obsidian](https://obsidian.md) has become invaluable for me, and I highly recommend it. It's not only where I store my digital notes, but also all of the PDFs on my [reading list](https://github.com/binxly/reading-list). As far as the construction of my [vault](https://help.obsidian.md/vault) (i.e. the folder where your notes are stored), I wanted it to have as few roadblocks as possible when I wanted to save something for the future. Be aware that if you're new to the app, it's best to start slow and only add complexity once your life and interests call for it!

Since I already had hundreds of notes (digital and physical) across a wide range of subjects, I decided to organize my vault using a system which is fairly common: [Maps of Content](https://obsidian.rocks/maps-of-content-effortless-organization-for-notes/). All this means is that I link notes to indexes (or their parent notes) instead of placing them in folders, where they are likely to get lost over time. This takes full advantage of the [note-linking](https://help.obsidian.md/link-notes) offered by Obsidian, which was what attracted me to the app in the first place. In the future, I may cover my approach in more detail (especially given how many people come to me for advice on this subject), but for now, here is a general overview of my indexes, and how I use them:

- 000 Index — primary index, where I list all other indexes and keep a list of the most relevant Active notes.

- 005 Active — an index of “active” notes, i.e. ones that require more attention. New notes are automatically added to this.

- 010 Library — index of notes based on reading material (the PDFs which are saved in my vault).

- 020 Writing & Creativity — an index of my creative writing, notes on art projects, and my blog posts.

- 030 Personal Development — this index links to notes on things like my monthly budget, gift ideas for family, and any introspective thoughts, as well as subjects like philosophy and occasionally theology.

- 040 Learning & Knowledge — index of notes on subjects I tend to study.

- 050 Quotes & Snippets — any quotes/snippets from reading material that I found particularly resonant.

- 060 Projects — an index containing links to notes about the development of various projects.

As far as folders are concerned, I only use **two folders** in Obsidian. They are prefixed with underscores to ensure they appear at the very top of my directory. For my ebooks and PDF documents, I created a `_library/` folder. My new notes are automatically placed in another folder titled `_inbox/`, and get linked to the `005 Active` index by default, thanks to Obsidian's [template system](https://help.obsidian.md/plugins/templates). After I feel a note has summarized the idea I was trying to capture well enough, I simply replace the link to `005 Active` with the relevant index (or parent note), and drag it out of this folder, placing it in the root of my vault's directory.

### Creating a Dedicated Workspace Folder

For projects in development, I create a `workspace/` folder in the root directory of my devices. This helps me quickly find projects, regardless of which system I'm working on. Inside of this folder lives a home for repositories, labeled `github.com/`. I name it this way such that the folder's working directory appears the same way links to repositories do, online— with a user, followed by the cloned repo of interest (i.e. `workspace/github.com/binxly/monkey-interpreter`).

### Syncing Between Devices

![synology drive](https://image.coolblue.be/840x473/content/3d1017622dc1d75ad0292e4ed76a0659)

For my Obsidian vault, I use a community plugin named [Remotely Save](https://github.com/remotely-save/remotely-save). With this, I'm able to backup my vault to [Dropbox](https://www.dropbox.com/) and sync it between all of my devices. Backups are scheduled at regular intervals, and I've turned on auto-syncing to push/pull any changes as soon as I open the app. This keeps all of my notes and books updated across all my devices.

I sync my `workspace/` folder (and my Obsidian vault) to my NAS via [Synology Drive](https://www.synology.com/en-global/dsm/feature/drive). There are a few headaches I encountered with this method (looking at you, `node_modules/`), but I created rules to ignore anything outside of the core files for each project. An added benefit of using Synology Drive is it allows me to access my files and projects from anywhere— which means I can look at them even from my phone. Creating a two-way sync will immediately push/pull any changes in the directory to all of my devices, which comes in handy given my homelab environment.

### Backups: Redundancy for Disaster Prevention

![backups](https://missfreddy.com/wp-content/uploads/2016/03/digitalbackups.jpg)

In addition to my NAS, I backup part of my `workspace/` folder and my entire Obsidian vault to iCloud, Google Drive, and an external drive I keep plugged into my primary workstation. This may not be necessary for you, but given the volume of data (and unfortunate mistakes I've made in the past) I figured this would give me enough room for error without having to worry about weeks worth of work suddenly disappearing. I can't tell you the number of times this has saved me from disaster, and I always recommend keeping multiple copies of any data you don't want to lose.

## Closing Thoughts

Ultimately, most PKM systems fall short not because they're inherently flawed, but because they're not tailored to *you*. After years of trial, error, and a few too many hours lost to the productivity content vortex, I've landed on a simple system that fits *my* brain and workflow. It's not flashy, and it's definitely not optimized to impress anyone raking in that sweet ad revenue that comes with viewer retention on YouTube—but it works. If there's one takeaway here, it's this: focus less on building the “perfect system” and more on creating one that you'll *actually start to use*. Keep it simple, make it yours, and let it evolve naturally over time.