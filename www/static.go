package www

import "embed"

//go:embed *.html
//go:embed *.js
//go:embed *.css
//go:embed css/*.css
//go:embed fonts/*.ttf
//go:embed images
var Static embed.FS
