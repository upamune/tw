function peco-favorite-tweet() {
tw tl --with-id | peco --prompt "TWEET> " | awk '{print $NF}' | tw fav --pipe
}
