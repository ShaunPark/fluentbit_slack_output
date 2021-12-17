# fluentbit slack output plugin (prettyslack)

Fluentbit Slack Outputn Plugin

Original slack output plugin can just send whole data and does not support formatting.

![sample slack message](https://github.com/ShaunPark/fluentbit_slack_output/blob/main/images/fl_output.png)

```
    [OUTPUT]
        Name             prettyslack
        Match            kernel.slack
        Webhook          https://hooks.slack.com/services/TXXXXXXXXXX/BXXXXXXXXX/XXXXXXXXXXXXXX
        Channel          myChannel
        Message_Field    msg
        Title            logs from ${NODE_NAME}
        Max_Attachments  7
```

- Webhook : slack incomming webhook url
- Channel : slack channel name
- Title : slack message title
- Message_Field : key for first line of slack message
- Max_Attachments : number of max attachments of 1 slack message. ( Defaut : 5, Max : 20)

- Color value of attachment from data with key 'color' (Default : #A9AAAA)

# Reference

"slack.go" : copied from https://github.com/ashwanthkumar/slack-go-webhook and modified to add block feature
