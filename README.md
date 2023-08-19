# s3clt
A lightweight AWS S3 client with a focus on minimizing binary size.

## Installation

```shell
go install github.com/knaka/s3clt/cmd/s3clt
go install github.com/knaka/s3clt/cmd/s3get
go install github.com/knaka/s3clt/cmd/s3put
```

## Synopsis

Retrieve an S3 object and output it to stdout. `s3get ...` is equivalent to `s3clt get ...`.


```shell
s3clt get us-east-1 DOC-EXAMPLE-BUCKET1 foo.tar | tar xvf -
s3get us-east-1 DOC-EXAMPLE-BUCKET1 foo.tar | tar xvf -
```

Read data from stdin and upload it as an S3 object. `s3put ...` is equivalent to `s3clt put ...`.


```shell
tar cvf * | s3clt put us-east-1 DOC-EXAMPLE-BUCKET1 bar.tar
tar cvf * | s3put us-east-1 DOC-EXAMPLE-BUCKET1 bar.tar
```

You can omit the region name if it is obtainable from the environment variables `$AWS_REGION` or `$AWS_DEFAULT_REGION`, the shared config file `~/.aws/config`, or EC2 instance metadata.

```shell
s3get DOC-EXAMPLE-BUCKET1 foo.tar | tar xvf -
```

## Todo

* Set up GitHub Actions to build and place binaries in the "Releases" section.
* Further reduce binary size using UPX.

