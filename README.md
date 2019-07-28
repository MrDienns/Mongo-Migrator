# Mongo Migrator
The Mongo Migrator project is a very small Go-based project that provides a small and simplistic migration tool to make MongoDB migrations easier to prepare and automate.

![Stable builds](https://travis-ci.com/Dyescape/Mongo-Migrator.svg?branch=master)

# Table of contents
- [Overview](#overview)
- [Motivation](#motivation)
- [License](#license)

# Overview
Mongo Migrator is a small Go-based project which can help schedule and automate MongoDB based database migrations or updates. The tool itself is merely responsible for uploading the local files to a MongoDB instance and deleting non existing ones (if flagged to do so). However, when packaged in a Docker container and scheduled as a Kubernetes Job for example, one can use it to their advantage and automate database migrations. Continuous Integration can be used to create and publish new instances of the container, the migration, and Kubernetes will download the latest image whenever it schedules an automation.

# Motivation
We at Dyescape run a massive Minecraft MMORPG project. The majority of content in-game is all configured through JSON files, which are loaded from a MongoDB instance. With frequent updates, and a team constantly working hard to prepare new content updates and patches, we wanted some mechanism of easily being able to prepare for these updates and to them apply them to the network as easy and as automated as can be. Our systems automatically detect file changes in MongoDB, so deploying a (content) update for us merely involes migrating the JSON files. This project is created to make this process as easy and as automated as possible by preparing a Docker image from our content repositories. These repositories will constantly prepare update containers through Continuous Integration, until a scheduled Kubernetes Job decides to take the latest image and execute the migration.

# License
This project is licensed under the GPL-3.0. For the official license, see [LICENSE](LICENSE)