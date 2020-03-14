# Notifier

The `notifier` provides a minimal interface for sending simple
notifications.

The `SlackNotifier` is the only practical implementation at present.

A `MockNotifier` exists to make it easy when needing to stub out the
notifier in tests.

## Slack

You will need to configure your Slack account to allow notifications to a
channel.

```
// your slack endpoint
url := "https://hooks.slack.com/services/xxx/yyy/zzzzzz"

n := NewSlackNotifier(url)

if err := n.Notify("Some text"); err != nil {
    return err
}
```

[Rate Limits](https://api.slack.com/docs/rate-limits) apply to Slack notifications.

## License

The MIT License (MIT)

Copyright (c) 2020 Scott Barr

See [LICENSE](LICENSE)
