# Tgz

Small utility for packaging/extracting folders in tar.gz archives.


## Installation

Installing using go get is the easiest.

    go get github.com/walle/tgz

This requires you to have your `$GOPATH` setup and your `$GOPATH/bin` added to path as described here http://golang.org/doc/code.html#GOPATH.

## Usage

The tool tries to be smart, since it compresses directories to file and extracts files to directory so if you give a directory as input path it compresses it to the output file you supplied. And if you give a file as input it extracts it to the folder you supplied as output. If no output path is supplied when extracting the current directory is used.

    usage: tgz input_file [output_file]
      input_file: if input_file is a file tgz tries to extract it, if input_file is a directory tgz tries to compress it to output_file
      output_file: only required if input_file is a directory, otherwise the current directory is used

### Create an archive containing a folder

    $ tgz my_folder my_archive.tar.gz

### Extract an archive in the current directory

    $ tgz my_archive.tar.gz

## Contributing

All contributions are welcome! See [CONTRIBUTING](CONTRIBUTING.md) for more info.

## License

Licensed under MIT license. See [LICENSE](LICENSE) for more information.