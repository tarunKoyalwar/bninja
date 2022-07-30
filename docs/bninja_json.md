## bninja json

Format JSON

### Synopsis

Format JSON and Color(Linux Only) returns given data if invalid json

```
bninja json [flags]
```

### Examples

```

// Format JSON and colorize it
curl -v http://somewhere/data.json | bninja json --color

// Sort Keys In Ascending Order
cat data.json | bninja json --sort

```

### Options

```
      --color   Colorize JSON (Only Linux)
  -h, --help    help for json
      --sort    Sort JSON Keys
```

### Options inherited from parent commands

```
  -i, --clipin        Import data from clipboard
  -o, --clipout       Export data to clipboard
      --file string   Import data from File
      --out string    Export data to File
      --stdout        Print to Stdout even if data if data is exported to clipboard or file
```

### SEE ALSO

* [bninja](bninja.md)	 - A Handy Toolkit for bash ninja's

