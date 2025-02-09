# evo

Smooth out command execution times, Just it.

## How To

the follows:

```bash
$ evo --cron "* * * * *" --run "echo 'Hello, World!'"
```

To see this is useful, imagine you have a system that calls a command every minutes as like a cron job.
Let's say you want a command to run every 5 minutes. However, the parent process runs the command every minute, so it gets called 4 extra times. So the command returns the cached standard output for those 4 times. And that's it.

## License

MIT by [@6jz](https://twitter.com/6jz)
