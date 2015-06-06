#/bin/bash

# Install		: gem install compass sass bootstrap-sass
# Prepare		: cd public && compass create assets -x sass -r bootstrap-sass --using bootstrap --sass-dir "sass" --css-dir "css" --javascripts-dir "js" --images-dir "img"
# Expanded		: compass w sass/main.sass -s expanded
# Compressed	: compass w sass/main.sass -s compressed

compass w sass/main.sass -s compressed
