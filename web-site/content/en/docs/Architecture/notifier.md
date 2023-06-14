---
title: "Notifier"
linkTitle: "Notifier"
weight: 2
---

The Notifier acts as a receiver for messages sent by the Events Manager. It registers handlers for specific types of messages and listens to specific queues within a message broker. When a message arrives, the Notifier executes the appropriate handler for that particular type of message. This allows the Notifier to perform actions based on the events received from the Events Manager, such as sending notifications to users.

## Domain Layer
This layer handles the business logic.

You can check [use cases](https://github.com/julian776/evently/tree/main/notifier/domain/events/usecases) to make a better idea.

## Infrastructure Layer
This layer deals with the implementation details of the Events Manager.

The most interesting is the [listener](https://github.com/julian776/evently/blob/main/notifier/infrastructure/rabbit/listener.go).


