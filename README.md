# evo

Smooth out command execution times, Just it.

## How To

the follows:

```bash
$ evo --cron "* * * * *" --run "echo 'Hello, World!'"
```

To see this is useful, imagine you have a system that calls a command every minutes as like a cron job.
Let's say you want a command to run every 5 minutes. However, the parent process runs the command every minute, so it gets called 4 extra times. So the command returns the cached standard output for those 4 times. And that's it.

## Advanced Usage

Let's say that for some reason, the command doesn't return the correct value when executed via a scheduled agent (is that possible? Yes, it is). In that case, you can use the following trick to always return a cached value, and update the value using crontab.

```bash
# agent (always returns cache value)
$ evo --cron "* * * * *" --delay 25 --run "echo 'Hello, World!'"

# crontab (update cache value)
* * * * * /usr/local/bin/evo --cron "* * * * *" --delay 10 --random --run "echo 'Hello, World!'" > /var/log/hello-world.log 2>&1
```

## License

MIT by [@6jz](https://twitter.com/6jz)
