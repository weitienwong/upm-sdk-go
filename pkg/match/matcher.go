package match

import "github.com/gobwas/glob"

func Match(path, pattern string) bool {
	var g glob.Glob

	// create simple glob
	g = glob.MustCompile(pattern)
	return g.Match(path) // true
}

// HitMatch 命中匹配，遍历patterns集合，有一个命中则返回true
func HitMatch(path string, patterns []string) bool {
	if len(patterns) == 0 {
		return false
	}
	var g glob.Glob
	for _, pattern := range patterns {
		g = glob.MustCompile(pattern)
		if g.Match(path) {
			return true
		}
	}
	return false
}
