---
title: "App Flow"
linkTitle: "App Flow"
weight: 1
resources:
  - src: "**.{png,jpg}"
    title: "Image #:counter"
---

{{% pageinfo %}}
A Brief App Overview.
{{% /pageinfo %}}

{{< imgproc tech_challenge Fill "950x550" >}}
{{< /imgproc >}}
[Full Resolution](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=tech_challenge.drawio#R7Vldb6M4FP01eQwCzOdjPtqutDPajirNTp8iBxzwDGDWOE3SX7%2FXYAcoadJq0kl2tJES4eMLNufce33tjNAs395xXKafWUyykW3G2xGaj2w7DEL4lcCuAVzHb4CE07iBrBZ4oM9EgaZC1zQmVc9QMJYJWvbBiBUFiUQPw5yzTd9sxbL%2BqCVOyAB4iHA2RP%2BmsUgbNHDNFv%2BD0CTVI1um6lni6EfC2bpQ441sNKs%2FTXeO9bOUfZXimG06ELoZoRlnTDRX%2BXZGMkmtpq257%2FaV3v28OSnEW26429lVUo2L9JFR%2FOWfcPe4nIxdJVYldpoQEgM%2Fqsm4SFnCCpzdtOi0fmkiH2tCq7X5xFgJoAXgdyLETomN14IBlIo8U70wY777pu6vG4%2ByYbi6Od92O%2Bc71aoEZz%2FIjGWM13NFZv2BniEZip%2BKrXlEjjDgBMrpME%2BIOEaV1xhKfjpDKLLvCMsJzBUMOMmwoE99%2F8LKTZO9nbp1wjnedQxKRgtRdZ58LwEwUAHneMqfVLwBCX3Z32cPF80MdKvzKi1Uu9I73EqPeSm3MpDv9l0rRMEJ56pb94RTeHnCr8LjwnN7XEf194uq5%2F2Es7UW9QnevhrnuIA0yQeat4pKeTYpFeShxDU5G1hM%2BuqtaJZ1iPbiwPccibNCdPBV%2FTkkDZq4DrKOSfNEuCDbo1SqXjt4ETU662%2FaNcLSmT3trA%2Fa88%2FOvmtfOKR6AWVdTTCFbw0m67qCKRwE04ozmbh%2Fnxiy%2BzGEgovHUHDRGLJGF692dFl7gWrnp5Tz3P%2Bz38%2FICVngmuTU8%2B5kv4IJuqK%2FVRHhOVdXRFgD3gd8s7XIaAGc6P22dNEYV%2Bk%2BkiQDFHbQn%2FCSZPesooKyAvqWTAiWdwwmGU1kh5DxNsWqFQGjMiq6usHmuJTj59tEnjIYeFMhIxWiXJScCRaxbCiua4fODNiYwi0xJa3ABSvImZYwJzDcvobOUEMPHZDQ%2BygJvdMSajZpXh9%2BnBYskx3T%2FYGGJjImK7zOxOuKVmXjIiu6ld4xrQecaNTUCFxLLYGKiXx5%2BzZGArHnlcjjPEXUiDK2juvyxyggjdm3IHq8jupZQoPFEa7EYl1mDEMmWJCSViwm0CVjyrP2F2PL9X03MJFlhp45dpzYBMQjzpI4xvcyOY9TOC%2F2BlbgDlwiPBDV4YdF9bCWHLiEdoWBZJ0gVEcPYO9OR%2B78UMjKPKqWS8vuyEvzBKae0SX84uc1J%2FKVYizwElekkocMElzMFbIAJxQJJw9fPi0eCAfujerpTPJYL%2BTRR1wddYJgqI5l%2BQZCH1W%2BOG8I2nop0meQ9vH8uMrYJkoxF4Ym%2BZW18QyE%2Bn6fUMfbE9VdyKwDLm%2BhwPD8jyJ1eB7x38uE5bIyxAbCx4hg%2FDrzwTJHFrW1DBzLcWGB8T3Tc90Q2Z5tAZ%2B3d7fPi%2FuxSxaOaW7ha5TFmcLHD01DL146wZnWIcW9Vtveynceyb%2Bi6D72vq7uPo%2Bn5l%2FbP7%2BRYj5GA4GveBdwquB%2FUYuf3AEcZES5%2FMkNQHPe8Avq%2FWOzfEuk4khIvk4V%2Bq%2B49QESjxR3L%2FLagfOJ4EBS89%2B%2FjkOz%2FaOnOWRv%2F0xDN%2F8C)

The Events Manager is responsible for exposing HTTP endpoints that the Frontend can communicate with. This allows the Frontend to interact with the Events Manager by making HTTP requests.

Events Manager publishes events to the Notifier using a message broker (In this case RabbitMQ).

------
# Flow Overview
## Frontend communicates with the Events Manager

The Frontend talk with the Events Manager by sending HTTP requests.
These requests perform operations like creating new events, updating existing events, or retrieving events information.

## Events Manager processes the requests

The Events Manager receives the HTTP requests from the Frontend and processes them accordingly.
It performs the necessary action.
Then Events Manager publishes events to the Notifier through RabbitMQ:

After processing the requests, the Events Manager publishes relevant events to the Notifier microservice using RabbitMQ.
RabbitMQ acts as a message broker, facilitating the communication between the Events Manager and the Notifier.
The Events Manager sends messages containing event data to a specific RabbitMQ exchange, which routes the messages to the appropriate queues.

## Notifier consumes events

The Notifier consumes the events by receiving messages from the queue.
The Notifier then processes these events.

