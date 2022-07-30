## bninja diff

Get unique and Common lines of files

### Synopsis

Get Unique or Common Lines of Two Files.
It doesnot matter if it has extra spaces or lines are shuffled/proportional 
Implementation uses Golang Maps to get Common and unique lines.
	

```
bninja diff [flags]
```

### Examples

```

// Common Unique Lines
bninja diff file1 file2 -c

// Unique Lines of file1
bninja diff file1 file2 -1

//Unique Lines of file2
bninja diff file1 file2 -2

```

### Options

```
  -c, --common   Show Only Common Lines
  -1, --file1    Show Lines Unique to File1
  -2, --file2    Show Lines Unique to File2
  -h, --help     help for diff
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

