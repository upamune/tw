function peco-favorite-tweet() {
tw tl 100 --with-id | peco --prompt "TWEET> " | awk '{print $NF}' | tw fav --pipe
}
