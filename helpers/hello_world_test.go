package helpers

/**
How to RUN tets at Terminal

if all file
> go test -v ./...

if by Name Func
> go test -v -run={name of func}

if by Name Func and sub test
> go test -v -run={name of func}/{name of sub func}
> go test -v -run=/{name of sub func} // run all test func include name of sub func


How to Run Benchmark
> go test -v -bench=. ./...
> go test -v -run=notest -bench=. ./...
*/

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M)  {
	// Before Unit Tets run
	fmt.Println("Start Unit Test")

	m.Run()

	// After Unit Tets run
	fmt.Println("End The Unit Test")
}

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Naim")
	
	if result != "Hello Naim" {
		t.Error("Result is not Hello Naim")
	}
	
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Naim")
	assert.Equal(t, "Hello Naim", result, "Result is not Hello Naim")
	/**
		assert.Equal = t.Error / t.Fail yang mana code selanjutnya akan tetap di jalankan
	*/
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Naim")
	require.Equal(t, "Hello Naim", result, "Result is not Hello Naim")
	/**
		require.Equal = t.Fatal/ t.FailNow yang mana code selanjutnya tidak akan tetap di jalankan
	*/
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Naim")
	require.Equal(t, "Hello Naim", result, "Result is not Hello Naim")
}

func TestSubTest(t *testing.T)  {
	t.Run("Name is Naim", func(t *testing.T) {
		result := HelloWorld("Naim")
		require.Equal(t, "Hello Naim", result, "Result is not Hello Naim")
	})
	t.Run("Name is Abrar", func(t *testing.T) {
		result := HelloWorld("Abrar")
		require.Equal(t, "Hello Abrar", result, "Result is not Hello Abrar")
	})
}

// ***
func TestHelloworldTable(t *testing.T)  {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "HelloWorld(Abrar)",
			request: "Abrar",
			expected: "Hello Abrar",
		},
		{
			name: "HelloWorld(Naim)",
			request: "Naim",
			expected: "Hello Naim",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Naim")
	}
}

func BenchmarkHelloWorldSub(b *testing.B)  {
	b.Run("Abrar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Abrar")
		}
	})
	b.Run("Naim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Naim")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B)  {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "HelloWorld(Abrar)",
			request: "Abrar",
		},
		{
			name: "Helloworld(Naim)",
			request: "Naim",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}