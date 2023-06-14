---
title: "Events Manager"
linkTitle: "Events Manager"
weight: 2
---

The Events Manager is a crucial component of the application and holds control over events and users. It consists of two main layers: the domain layer and the infrastructure layer.

## Domain Layer
This layer handles the business logic. It manages events and users. Additionally, it emits events when specific actions occur, such as when an event is created, updated, or deleted. When a user registers for an event, an event is emitted as well.
You can check [use cases](https://github.com/julian776/evently/tree/main/events-manager/domain/events/usecases) to make a better idea.

## Infrastructure Layer
This layer deals with the implementation details of the Events Manager.


