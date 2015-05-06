SlackMS
========

1) Chat on slack with your team via SMS

2) Send SMS to a team member from any Slack channel
## Usage

1) To chat on Slack via SMS, first word in text message must be the destination channel of your Slack team

    #general Hello team!!!
    
2) To send an SMS to a team member from a slack channel, make sure the team member's phone number is available in the their profile. Then one can send text like below:

    /sms <slack_username_of_team_member> <message body>


## Installation

### Setup your own server

Make sure to change the **Slash Command** URL to whatever your URL is.

##### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/ueg1990/SlackMS/tree/master)

And then:

```bash
$ heroku config:set SLACK_WEARHACKS_WEBHOOK_URL=<URL>
$ heroku config:set TWILIO_ACCOUNT_SID=<ACCOUNT_SID> 
$ heroku config:set TWILIO_AUTH_TOKEN=<AUTH_TOKEN>
$ heroku config:set TWILIO_NUMBER=<NUMBER>
```

### Setup Integration

- Go to your channel
- Click on **Configure Integrations**.
- Scroll all the way down to **DIY Integrations & Customizations** section.

#### Add a new slash command with the following settings:

- Click on **Add** next to **Slash Commands**.

  - Command: `/sms`
  - URL: `http://YOUR-URL.com/sms`
  - Method: `POST`

  ![](http://i.imgur.com/zLrHkf5.png)

All other settings can be set on your own discretion.

#### Set up a new incoming webhook

Click on **Add** next to **Incoming WebHooks**.

  - Choose a channel to integrate with (this doesn't matter -- it'll always respond to the channel you called it from)
  - Note the new Webhook URL.

  ![](http://i.imgur.com/tgiTLdj.png)
  
### Setup Twilio

- Create a Twilio Account
- Go to your [Twilio Account](https://www.twilio.com/user/account/settings) to retrieve the Twilio Account SID and Auth Token associated with your account
- Update the Messaging Request URL to your URL with route `/twiml`:
    
    ![](http://i.imgur.com/Mkf7HGa.png)

## Contributing

- Please use the [issue tracker]() to report any bugs or file feature requests.

- PRs to add new sources are welcome.
