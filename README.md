# EVENTLY
This project is part of a tech challenge to become tech team lead.

It was done in a short period of time and made use of new technologies that I try to learn and implement as well as I can over the challenge.
# Complete Docs web-site
https://bright-figolla-5bb015.netlify.app

## Install and setup
* First, you have to clone the repository of [Evently](https://github.com/julian776/evently/tree/main).

> `git clone https://github.com/julian776/evently.git`

* Then you you have to create each .env file(Front, Events Manager and Notifier). 
For local development, you can copy the .env.template, and will be enough except for the email service variables in the notifier microservice.
Check how to generate a password for Gmail [here](https://support.google.com/mail/answer/185833?hl=en).

## Create .env Files:

* After cloning the repository, you need to create three separate .env files. These files are used to store environment variables that configure the application.
* The .env files are required for the Front-end, Events Manager, and Notifier microservices.
* For local development, you can start by copying the provided .env.template file and customizing it as needed.
* It's mentioned that you may need to modify the email service variables in the Notifier microservice .env file. The instructions provide a link to learn how to generate a password for Gmail if you plan to use Gmail as your email service.

## Running with Docker (optional):

* If you have Docker installed, you can use the Docker Compose file provided in the root directory of the project.
* It will create and run all the necessary services and databases for the Evently project.

## Running Services Individually:

* If you don't have Docker or choose not to use it, you can run each service individually.
* This means you would need to start each component of the project separately, such as the Front-end, Events Manager, and Notifier microservices.
