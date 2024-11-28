# Food Delivery APP Notification Service

### Design a notification system for a food delivery app.

#### Scenario:
- Send real-time notifications (e.g., order updates) to millions of users.
- Support multiple notification channels (e.g., SMS, email, push notifications).
- Handle retries for failed notifications across all channels with exponential backoff.
- Allow dynamic prioritization of notifications (e.g., critical updates must be sent before promotional ones).
- Scale horizontally for traffic spikes, such as during lunch/dinner hours.
- Provide observability, including metrics for sent, failed, and retried notifications.

### Task:

#### Design the backend architecture for this notification system.

- How would you design the system to handle the above requirements, ensuring scalability and reliability?
- How would you handle dynamic prioritization of notifications across channels?
- Provide a Go code snippet to demonstrate the worker service for handling retries with exponential backoff.

#### Reply:

Publisher
```
1. Read the order Data
    1.1 Detect the priority of data/notification.
        set priority_status low
        if priority is >
            Set priority_status High
    1.2 Prepare the JSON payload with required information
        based on order details
            set order_id
            status
            etc...
    1.3 Check the MQ status
        Check que status
        implement the functionality like ping/pong to check a status
    1.4 Sent the message in the que
    1.5 Acknoledge the message Sent
    1.6 Update the status
```

Listener

```
    1. Deploy an multiple listener to listen Message que (priority/promotional)
    2. Trigger the worker once data/message received on specific message que
        worker to listen que data length
            if priority_que > 0
                if data.length > 0
                    validate data
                        check the notification triggered for this user
                            if already triggered with same payload or template id same notification channel
                                check the status of notification
                                    success
                                        mark in db as success
                                    fail
                                        retry n(specified number) number of times
                                            success
                                                mark status in db
                                            fail
                                                log it into system & notify administrator                                            
                            else
                                extract the requried data
                                    fetch the additional data if any
                                prepare the payload
                                    get template if any
                                connect to the service provider
                                    sent the payload the service provider (email/push/sms)
                                        success
                                            get the acknoledgement
                                                if success
                                                    update the status in db
                                                if false
                                                    put back status to ready so listener can pull data again and send
                                                if unknown
                                                    put back status to unknown so listener can pull data again and send
                                        fail
                                            record in db
                                                retry & get the acknoledgement
                                                    record status in db
                else
                    log the instance
                    wait for the next message
            else
                promotional que will trigger & send notification
```

## Environment

```
MQ_USER=admin
MQ_PASS=admin123
MQ_HOST=0.0.0.0
MQ_PORT=5672

APP_HOST=0.0.0.0
APP_PORT=8000

PRIORITY_QUEUE=
PRAMOTIONAL_QUEUE=
```