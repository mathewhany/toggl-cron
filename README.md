# Toggl Cron

## Motivation
I've started this project because I thought it would be cool if I could answer the question: "What have I been doing in the past day? week? month? or even year?". So I tried to find a free app that I can use to do that, I found many options, most notably [Toggl](https://toggl.com/) and [Clockify](https://clockify.me/), but I've decided to go with Toggl because it felt more polished. 

I've tried it for a while but it didn't work out for me, I didn't like the fact that I have to start and stop the timer every time I switch tasks, I often forget to do that and spent a lot of time re-entering time entries manually. So I tried setting up a reminder that fires every 30 minutes or so and reminds me to fill in what I've been doing in the past 30 minutes, but the problem is that I often do multiple tasks in the same 30 minutes, so I have to fill in multiple time entries for the same time period, which is a pain, and if I entered them in the same time period this messed up the report at the end of the week, I often had entries like "TV & Coffee", "TV & Lunch", "Working on updating CV & Working on side project", "Side project". They were displayed as different tasks in the report, which is not what I wanted. Creating my own app was not an option because I didn't have the time to do that, so instead I decided to create a script that just runs every 15 minutes creating a time entry for the next 15 minutes, and I can just write what I'm doing in the same entry like "TV & Coffee @ProjectName & Working on a project @ProjectName", and when 15 minutes pass, the script run again, splitting them into separate entries, each with in it's project.

## How it works
It's just a lambda function hosted on AWS which is triggered every 15 minutes by a CloudWatch event.

It is written in Go, and uses the Toggl API to create the time entries.

Why Go and not Python? Because I wanted to learn Go, and this seemed like a good opportunity to do so. Why AWS and not a simpler option? Again because I wanted to learn AWS, and this seemed like a good opportunity to do so.

## Configuration
The script is configured using environment variables, here's a list of the variables and what they do:
| Variable | Description |
| - | - |
| TOGGL_EMAIL | The email address you use to login to Toggl |
| TOGGL_PASSWORD | The password you use to login to Toggl |
| TOGGL_WORKSPACE_ID | The ID of the workspace you want to create the time entries in |
| DURATION_MIN | The duration of the time entry in minutes |

## How to run locally
1. Must have Docker installed
2. Fill environment variables in `.env` file (Create one if it doesn't exist)
```
TOGGL_EMAIL=<email>
TOGGL_PASSWORD=<password>
TOGGL_WORKSPACE_ID=<workspace-id>
DURATION_MIN=15
```
2. Run `docker-compose up`
3. Test the function by running:
```sh
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
```
This should create a time entry in Toggl and the duration of 15 minutes.