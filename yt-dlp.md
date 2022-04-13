[![yt-dlp](https://raw.githubusercontent.com/yt-dlp/yt-dlp/master/.github/banner.svg)](https://github.com/yt-dlp/yt-dlp)
# [Making it work](https://pypi.org/project/yt-dlp/)
If Python Pip is not installed (ArchLinux):

`paru -S python-pip`

`yay -S python-pip`

`sudo pacman -S python-pip`

To install

`pip install yt-dlp`

Adding to a PATH

`echo -e 'export PATH="~/.local/bin:$PATH"\n' >> ~/.bashrc`

# Understanding how it works
Default `-f`, `--format`:

`yt-dlp -f "bv*+ba/b" "https://youtu.be/iik25wqIuFo"`

`yt-dlp --format "bestvideo*+bestaudio/best" "https://youtu.be/iik25wqIuFo"`

By default, yt-dlp tries to download the best available quality if you __don't__ pass any options.

The default format selector is `bv*+ba/b`. This means that if a combined video + audio format that is better than the best video-only format is found, the former will be preferred.

Change the output template name

`yt-dlp -o "%(title)s.%(ext)s" "https://youtu.be/iik25wqIuFo"`

Result: __title.ext__

Example: __Rick roll, but with different link.webm__

To embed subtitles:

`yt-dlp --embed-subs "https://youtu.be/iik25wqIuFo"`

To embed metadata and chapters:

`yt-dlp --embed-metadata "https://youtu.be/iik25wqIuFo"`

