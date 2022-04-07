# Remove Image Borders

Removes borders from images.

## Usage

```
make dir="./path/to/images/directory"
```

## Wait why not just imagemagick `trim`

I needed something a little more specific than just `trim` imagemagick, as trim will indiscriminately trim whitespace. This approach is probematic for assets that are intentionally use whitespace around content for layout purposes.

> This option removes any edges that are exactly the same color as the corner pixels.
> http://www.imagemagick.org/script/command-line-options.php#trim

All available options I could find adhere to this indiscriminate algorithm, so I had to write something to accomplish this.
