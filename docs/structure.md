# Resource Structure

## StackDB database
StackDB will maintain its own database (called stackdb) to support managing itself. This database is no different than any other database other than it being managed internally.

## Database
Behind the scenes, on the local file system, a database is stored in a single directory. Within a given database's directory, things are broken down as follows:

```
/data
	/databaseOne.json
	/databaseTwo.json
	...
```

Within each database file is the following:
```json
{
	"uuid": "",
	"name": "",
	"type": "",
	"collections": [
		{
			"uuid": "",
			"name": "",
			"data": {
				"docOne": {
					...
				},
				...
			},
		}
	]
}
```