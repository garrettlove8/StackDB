![License](https://img.shields.io/github/license/garrettlove8/StackDB)
[![Go Report Card](https://goreportcard.com/badge/github.com/garrettlove8/StackDB)](https://goreportcard.com/report/github.com/garrettlove8/StackDB)
![Workflow](https://img.shields.io/github/workflow/status/garrettlove8/StackDB/Test%20and%20Build)

# StackDB

A "stackable" database that prioritizes (in no particular order) ease of use, performance, adaptability, and scalability.

## Installation
To install and setup StackDB, you first want to download the desired reslease version from [here](https://github.com/garrettlove8/StackDB/releases). After that's complete, run the following commands:

```bash
cd path/to/stackdb/executable
go install
```

After you've run these two commands, StackDB will have been successfully installed on your machine. However, there is one final step before it is ready for use.

Last, you'll need to run the StackDB setup process. To do so, run the following:

```bash
sdb setup
```

Now, you're ready to use StackDB!

## Basic Concepts
Part of how StackDB realizes its priority of "Adaptability" is by limiting its core to a very simple and fundamental data structures - maps and sets. This essentially means that StackDB is a key/value store that allows you to add layers depending on the problem you are trying to solve.

In other words, StackDB isn't by itself a fully-feldged database management system such as Postgres or MongoDB, but rather provides a highly flexible storage machanism that can be used as a base layer to build out relational databases, document databases, graph database, etc..

The reason for going about it this way is that instead of having to decide between MySQL, Postgres, or MongoDB, DynamoDB, etc. is that now, you'll be able to use one database that can adapt to whatever the problem is that you're trying to solve.

## Usage

### Startup
StackDB is started by running

```bash
sdb start
```

Upon running this command, you'll have started and loaded up the database system as well as entered the shell. From the shell you can interact with everything StackDB has to offer via the SDBL (StackDB Language). Eventually, you'll also be able to connect to the database via the drivers and use it within your own applications.

### SDBL
SDBL is under development, more to come...

## Contributing
As this is a learning / passion project, I am not currently looking for contribution.

## To-Do
[] Remove Ginko and Gomega and use the standard testing library, along with Testify and Mockery where necessary

[] Running `go install` does seem to make the program usable via the `sdb` command