package files

import (
	"os"
)

func ReadFile() {
	// path := filepath.Join(os.TempDir, "")
	// io.ReadAll("/welcome.txt")

	os.Open("welcome.txt")

}
