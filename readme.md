# GO-REPO-DISPATCH-CONTROLLER

Since githubs own cron schedule is not that precicse\
I've decided to create a small go controller that will trigger\
workflow dispatches using crontab on one of my servers.

Keep your github bearer token to the repo you want to trigger\
in the bearer.token file. Support for only one repo atm. 

To be improved.

## Example usage

```ps1
.\dispatch-controller.exe -userName eddeT -eventType book_test -repoName book-feelgood
```