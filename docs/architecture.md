# StackDB Architecture

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
	"type": ""
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