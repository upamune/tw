function peco-delete-tweet() {
tw tl 100 --with-id --user $1 | peco --prompt "TWEET> " | awk '{print $NF}' | tw del --pipe
}
