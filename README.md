# lls

The 'lls' tool is a lightweight version of the 'ls' command, which allows you to generate a list of files within a directory that has a large number of files (e.g. 30M).

## How?

To use the 'lls' tool, you need to allocate a buffer in memory for the size of the directory you want to list. You can use the 'getdents64' system call to retrieve the list of files.

The larger the directory size, the more memory will be required for the buffer. To determine the necessary amount of memory in advance, you can run the `ls -dl` command.

For additional information on using the 'lls' tool, you can refer to [this article](http://be-n.com/spw/you-can-list-a-million-files-in-a-directory-but-not-with-ls.html).

## Usage

```sh
$ lls > output.txt

$ lls / > output.txt
```

To adjust the memory usage of the 'lls' tool, you can use the `-buf-size` option (with a default value of 5MB). This allows you to specify the size of the buffer used to hold the list of files.
