<p align="center">
<img alt="Marshal logo" src="assets/logo.png" height="150"/>
<b>Marshal</b>
</p>

---

Marshal is a marshaller for watermill that just passes the message. To be used when interacting with Core

## Usage:

```golang

// With no logger 
logger := walrus.New()

// With a specific logger
logger := walrus.NewWithLogger(logrus.StandardLogger())

// With a specific log entry
logger := walrus.NewWithLogger(logrus.StandardLogger().WithField("foo", "bar")

// Example for NATS subscriber
subscriber, err := nats.NewStreamingSubscriber(
    nats.StreamingSubscriberConfig{
        ...
    },
    logger,
)

// Example for message router
router, err := message.NewRouter(message.RouterConfig{}, logger)
```
