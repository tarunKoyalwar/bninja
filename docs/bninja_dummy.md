## bninja dummy

Dummy Module to pass data from input source to output 

### Synopsis

Dummy Moudle to pass data from input source(clipboard,file,stdin,args) to output(file,clipboard)

```
bninja dummy [flags]
```

### Examples

```

// export stdin data to clipboard
echo "somestuff" | bninja dummy -o

//import data from clipboard
bninja dummy -i
	
```

### Options

```
  -h, --help   help for dummy
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

