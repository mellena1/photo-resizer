# photo-resizer

Quick Go program to resize images

## Usage

### Plan
Plan allows you to see which files the glob finds before applying. `-o` allows you to create a file from that plan so you can ensure that the planned files are the only ones affected.
```
photo-resizer plan -g 'pictures/*.jpg' -o plan.json
```

### Apply
One of plan (-f) or fileglob (-g) must be set as the input for the apply. Width and height sets the width and height of the new photos. Note that it will overwrite the current images unless you use `--suffix`.
```
photo-resizer apply -f plan.json -w 1920 -h 1080
```

Check the help for more options. `photo-resizer [command] --help`


## File Glob
Uses [github.com/bmatcuk/doublestar](https://github.com/bmatcuk/doublestar) for the globbing. Check that repo for info on the understood glob syntax.
