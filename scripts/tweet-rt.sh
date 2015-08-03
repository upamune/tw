function peco-retweet-tweet() {
tw tl --with-id | peco --prompt "TWEET> " | awk '{print $NF}' | tw rt --pipe
}
