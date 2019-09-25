# Mongo Migrator
The Mongo Migrator project is a very small Go-based project that provides a small and simplistic migration tool to make MongoDB migrations easier to prepare and automate.

![Stable builds](https://travis-ci.com/Dyescape/Mongo-Migrator.svg?branch=master)
[![Stable coverage](https://codecov.io/gh/Dyescape/Mongo-Migrator/branch/master/graph/badge.svg)](https://codecov.io/gh/Dyescape/Mongo-Migrator)



# Table of contents
- [Overview](#overview)
- [Commands](#commands)
- [Environment variables](#environment-variables)
- [Motivation](#motivation)
- [License](#license)

# Overview
Mongo Migrator is a small Go-based project which can help schedule and automate MongoDB based database migrations or updates. The tool itself is merely responsible for uploading the local files to a MongoDB instance and deleting non existing ones (if flagged to do so). However, when packaged in a Docker container and scheduled as a Kubernetes Job for example, one can use it to their advantage and automate database migrations. Continuous Integration can be used to create and publish new instances of the container, the migration, and Kubernetes will download the latest image whenever it schedules an automation.

# Commands
The Mongo Migrator has a few commands which can be used throughout your deployment. Below is a small list of these commands.
`Migrate` - The migrate command will perform the actual migration towards a MongoDB instance.
`Check` - Will perform a check, and summerizes what the `Migrate` command would do. This command does not do a migration; it merely calculates what it would do if you did a migration.

# Environment variables
The tool offers several environment variables. They are used throughout the application. All of these settings can also be used in the commands. Environment variables are upper case, while the command arguments are lowercase.
`PATH` - Specify the path of the directory to be migrated towards Mongo. Default `./migrate`.
`DELETE` - Delete files not part of this package. Default `false`.
`DEBUG` - Output detailed information during any activity. Default `false`.
`HOST` - The MongoDB host. Default `localhost`
`PORT` - The MongoDB port. Default `27017`
`USERNAME` - The username for MongoDB. Default `mongo`
`PASSWORD` - The password for MongoDB. Default `mongo`

# Motivation
We at Dyescape run a massive Minecraft MMORPG project. The majority of content in-game is all configured through JSON files, which are loaded from a MongoDB instance. With frequent updates, and a team constantly working hard to prepare new content updates and patches, we wanted some mechanism of easily being able to prepare for these updates and to them apply them to the network as easy and as automated as can be. Our systems automatically detect file changes in MongoDB, so deploying a (content) update for us merely involes migrating the JSON files. This project is created to make this process as easy and as automated as possible by preparing a Docker image from our content repositories. These repositories will constantly prepare update containers through Continuous Integration, until a scheduled Kubernetes Job decides to take the latest image and execute the migration.

# License
This project is licensed under the GPL-3.0. For the official license, see [LICENSE](LICENSE)