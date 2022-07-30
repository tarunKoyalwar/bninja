# bninja

bninja is a collection of cross-platform command line tools with some extra features(clipboard support) and defaults to improve experience of bash/shell/terminal users.

Everything is written in Golang hence it can be used in Linux,Windows and MacOS.

All tools are subcommands of bninja command. These tools are -


* [bninja cut](./docs/bninja_cut.md)	 - Similar to GNU cut but better
* [bninja diff](./docs/bninja_diff.md)	 - Get unique and Common lines of files
* [bninja dummy](./docs/bninja_dummy.md)	 - Dummy Module to pass data from input source to output 
* [bninja getlen](./docs/bninja_getlen.md)	 - Returns length of the given string
* [bninja hcut](./docs/bninja_hcut.md)	 - Cut Data horizantally similar to GNU cut but horizantal
* [bninja html](./docs/bninja_html.md)	 - HTML Encoder/Decoder
* [bninja json](./docs/bninja_json.md)	 - Format and Colorize JSON
* [bninja replace](./docs/bninja_replace.md)	 - Match & Replace Data
* [bninja sort](./docs/bninja_sort.md)	 - Sort Input data in ascending/descending order
* [bninja uniq](./docs/bninja_uniq.md)	 - Print Unique lines
* [bninja url](./docs/bninja_url.md)	 - A Simple URL Encoder and Decoder


The Obvious advantage is that all these commands can read and write from clipboard & files .Other advantages include less options and default behaviour such as trim spaces from each line.

Some Commands have extra features than there counterpart UNIX Commands.

For Example-

1. Cut 
Same as GNU cut but delimeter can be more than 1 characters . Skip lines that does not contain delimeter using (--skip) option.

2. hcut
Similar to GNU cut but cut data horizantally i.e cut entire document based on delimeter and print rows based on user input.

3. diff
Compare two files and get unique lines that are common or belongs to a particular file.

4. uniq
Get unique lines from input . By default trims extra spaces on each line . (GNU uniq only detects duplicates only if they are adjacent)

5. replace
An Alternative to sed but less of an eyesore. Match and replace strings (NO Syntax uses options instead of subcommands like /s or /g) . options are similar to replace functions in go or python.

# Installation 

- Make Sure you have go installed properly or [install](https://go.dev/doc/install)

- Install Using go
```sh
go install github.com/tarunKoyalwar/bninja/cmd/bninja@latest
```

If you loved this tool . Star / Watch this repo to show support and to not miss any feature updates.


# Support

If you like `bninja` and want to see it improve furthur or want me to create other intresting projects , You can buy me a coffee 

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/B0B4CPU5V)

