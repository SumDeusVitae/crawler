package main

import "testing"

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one/chewawa">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one/chewawa">
					<span>Boot.dev</span>
				</a>
				<a href="http://example.com/path/one/chewawa">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one/chewawa", "https://other.com/path/one/chewawa", "http://example.com/path/one/chewawa"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/">
					<span>Boot.dev</span>
				</a>
				<a href="http://example.com/path/">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/", "https://other.com/path/", "http://example.com/path/"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v URLs, got %v", i, tc.name, len(tc.expected), len(actual))
			}
			for j := range tc.expected {
				// Check if the actual URL matches the expected one
				if actual[j] != tc.expected[j] {
					t.Errorf("Test %v - '%s' FAIL: expected URL %v, but got %v", i, tc.name, tc.expected[j], actual[j])
				}
			}
		})
	}
}
