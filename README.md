# lls

lls is lightweight ls. Using lls, you can get a list of files in a directory that contains a large number of files.

## How?

You allocate a buffer for the size of the directory, and then call the getdents64 system call directly.

The larger the size of the directory, the more memory you will need. By running `ls -dl`, you will know how much memory to allocate in advance.

See [this article](http://be-n.com/spw/you-can-list-a-million-files-in-a-directory-but-not-with-ls.html) for more information.

## Usage

```sh
$ lls > output.txt

$ lls / > output.txt
```

If you want to change the memory usage, you can specify the `-buf-size` option (default: 5MB).
