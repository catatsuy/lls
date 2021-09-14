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

If you want to reduce the memory usage, you can specify the `-buf-size` option. If the size you specify is smaller than the actual size needed, you will not get the full file list.

## FAQs

### I can't run lls

If the size of the directory is larger than 2GB, you may get the following error.

```
$ lls
invalid argument
```

You can specify the size of the buffer by passing the `-buf-size` option. You can specify `2147483647`, which is 1 byte smaller than 2GB.

```
$ lls -buf-size 2147483647
```

### I can't get a list of all files

Maybe the size of the buffer is too small.

For example, you run `lls -debug` and get the following output.

```
$ lls -debug
bufSize: 4096; getdents ret: 4080
```

The `bufSize` is the actual size of the buffer passed to the system call, and `getdents ret` is the return value of the getdents system call, which is the number of bytes actually used in the buffer.

It consumes `(20+filename) bytes` per file. If the difference between these two numbers is less than one file, it is possible that there are still other files, but we are not able to list all of them.

If you give a larger number to the `-buf-size` option, you should get a list of all files.
