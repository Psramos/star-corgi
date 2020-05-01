**Star Corgi**

![Corgi](./corgi.jpg)

Star corgi is just a small application that retrieves your daily horoscope and posts it to
the desired slack channel.

**How it works:**
Create an application on slack and install it to any channel, retrieve
the link, and assign it to the following var in main.go:
```
var slackUrl string = "YOUR_SLACK_URL"
```

Then you can fill the desired horoscopes to retrieve on the main code:

```
func main() {
	guidance("aries")
}
```