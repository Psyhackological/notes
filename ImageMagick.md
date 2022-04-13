# ImageMagick

<img src="https://wikiless.org/media/wikipedia/commons/9/9a/ImageMagick_logo.svg" alt="ImageMagickLogo" width="256">

To merge all .png images vertically (y axis, up and down):

`magick convert -append *.png vertical.png`

To merge all .png images horizontally (x axis, left and right):

`magick convert +append *.png horizontal.png`
