# StackDB
A stackable database that provides extreme flexibility, self-management, high performance, low cost, and puts cloud native first

## Install and Setup
To install and setup StackDB, you first want to download the desired reslease version from [here](https://github.com/garrettlove8/StackDB/releases). After that's complete, run the following commands:

```bash
cd path/to/stackdb
go install
```

After you've run these two commands, StackDB will have been successfully installed on your machine. However, there is one final step before it is ready for use.

Last, you'll need to run the StackDB setup process. To do so, run the following:

```json
sdb setup
```

Now, you're ready to use StackDB!

## Creating a new database
To create a new database using the CLI (which is currently the only way of doing so) use the following command:

```bash
sdb database create [args]
```

### Args
Args are both positional and required.

Ordering:
1. Name
2. Type

An error will be returned if the command args don't abide by a few simple rules.

Rules:
1. The database name cannot be `system` since that is already in use by StackDB.
2. The only type of database currently available is `keyValue`, however more are coming (Time TBD)

### Example

```bash
sdb database create dbName keyValue
```