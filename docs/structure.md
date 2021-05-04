# Resource Structure

## System
The overall system will maintain resources to manage its state and knowledge of that state. The first level of which is the `system.json` file, which keeps a record of every database, as well as their acossiated collections. This resource is not to be confused with `config/stackdb.json` which is where system settings are located.

```json
{
	"databases": [
		{
			"name": "databaseOne",
			"type": "keyValue"
		}
	]
}
```

## Database
Behind the scenes, on the local file system, a database is stored in a single directory. Within a given database's directory, things are broken down as follows:

	/databaseOne
		/databaseOne.json
		/collections
			/collectionOne.json
			/collectionTwo.json
			/collectionThree.json
			...

Within each database file is the following:
```json
{
	"uuid": "",
	"name": "",
	"type": "",
	"collections": ["", ""]
}
```

Within each collection file is the following:
```json
{
	"uuid": "",
	"name": "",
	"data": {
		"docOne": {
			...
		},
		"docTwo": {
			...
		},
		"docThree": {
			...
		},
		...
	},
}
```