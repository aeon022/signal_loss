// Package signal_loss only carries the story files, so installed binaries
// (go install) bring their own and need no story.json next to them.
package signal_loss

import "embed"

//go:embed story.json story_de.json
var Stories embed.FS
