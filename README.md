# lab2spl

This command generates a spline (.spl) for DaVinci Resolve or Fusion from a .lab file.

## Download

[Latest release](https://github.com/1l0/lab2spl/releases/latest)

## Usage

1. Run the following command with the .lab file path.
    - By default, `<.lab file path>.srt` should be generated.

```sh
lab2spl <file name>.lab
```

### Options

Specify the output file path:

```sh
lab2spl -o example.spl <file name>.lab
```

Specify the fps:

```sh
lab2spl -fps 30 <file name>.lab
```
