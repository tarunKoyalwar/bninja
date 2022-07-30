## bninja url

A Simple URL Encoder and Decoder

### Synopsis

A Simple URL Encoder & Decoder with Similar syntax to base64 unix command

```
bninja url [flags]
```

### Examples

```

// URL Encode data
echo "some Url encoded_data" | bninja url

// URL Decode data
echo "some+Url+encoded%5\fdata" | bninja url -d

```

### Options

```
  -d, --decode   URL Decode
  -h, --help     help for url
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

