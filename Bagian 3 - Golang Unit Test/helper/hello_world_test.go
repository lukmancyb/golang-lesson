package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("L")
	}
}

func BenchmarkHelloWorldLukman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lukman Hakim")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Lukman",
			request:  "Lukman",
			expected: "Hello Lukman",
		},
		{
			name:     "Hakim",
			request:  "Hakim",
			expected: "Hello Hakim",
		},
		{
			name:     "HakimL",
			request:  "HakimL",
			expected: "Hello HakimL",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("TestLukman", func(t *testing.T) {
		result := HelloWorld("Lukman")
		require.Equal(t, "Hello Lukman", result, "Result must be hello lukman")
	})

	t.Run("TestHakim", func(t *testing.T) {
		result := HelloWorld("Lukman")
		require.Equal(t, "Hello Lukman", result, "Result must be hello lukman")
	})
}

// func TestMain(m *testing.M) {
// 	//before
// 	fmt.Println("Before unit test")
// 	m.Run()
// 	// After
// 	fmt.Println("After unit test")
// }

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak berjalan di windows")
	}
	result := HelloWorld("Lukman")
	assert.Equal(t, "Hello Lukman", result, "Result must be Hello Lukman")

	fmt.Println("Test Skip berjalan")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Lukmanh")
	assert.Equal(t, "Hello Lukman", result, "Result must be Hello Lukman")
	fmt.Println("Test assert selesai")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Lukmanh")
	require.Equal(t, "Hello Lukman", result, "Result must be Hello Lukman")
	fmt.Println("Test assert selesai")
}

func TestHelloWorldLukman(t *testing.T) {
	result := HelloWorld("Lukmanh")

	if result != "Hello Lukman" {
		t.Fail()
	}

	fmt.Println("Selesai")
}

func TestHelloWorldHakim(t *testing.T) {
	result := HelloWorld("Lukmanhakim")

	if result != "Hello Lukman" {
		t.FailNow()
	}
}
