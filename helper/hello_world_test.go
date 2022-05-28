package helper

import (
	"fmt"
	"testing"
	_"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	} {
		{
			name: "Rikaz",
			request: "Rikaz",
		},
		{
			name: "Fawaiz",
			request: "Fawaiz",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i:=0; i<b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})

	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Rikaz", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			HelloWorld("Rikaz")
		}
	})
	b.Run("Fawaiz", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			HelloWorld("Fawaiz")
		}
	})
}

func BenchmarkHelloWorld1(b *testing.B) {
	for i:=0; i<b.N; i++ {
		HelloWorld("Rikaz")
	}
}

func BenchmarkHelloWorld2(b *testing.B) {
	for i:=0; i<=b.N; i++ {
		HelloWorld("Fawaiz")
	}
}

func TestTableTest(t *testing.T) { //rikaz ver
	slice1 := []string{"rikaz","fawaiz","haerul","afgani"}
	for _, val := range slice1 {
		t.Run(val, func(t *testing.T){
			result := HelloWorld(val)
			assert.Equal(t,"Hello " + val, result, "Result must be 'Hello " + val + "'")
		})
	}
}

func TestHelloWorldTableTest(t *testing.T) { //Konsep table test menggunakan slice & perulangan untuk test
	tests := []struct{
		name string
		request string
		expected string
	} {
		{
			name : "Rikaz",
			request : "Rikaz",
			expected : "Hello Rikaz",
		},
		{
			name : "Fawaiz",
			request : "Fawaiz",
			expected : "Hello Fawaiz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Rikaz", func(t *testing.T){
		result := HelloWorld("Rikaz")
		assert.Equal(t,"Hello Rikaz", result, "Result must be 'Hello Rikaz'")
	})
	t.Run("Fawaiz", func(t *testing.T){
		result := HelloWorld("Fawaiz")
		assert.Equal(t,"Hello Fawaiz", result, "Result must be 'Hello Fawaiz'")
	})
	t.Run("Haerul", func(t *testing.T){
		result := HelloWorld("Haerul")
		require.Equal(t,"Hello Haerul", result)
	})

}

func TestMain(m *testing.M) {
	//before test
	fmt.Println("BEFORE UNIT TEST")
	
	m.Run()

	//after test
	fmt.Println("AFTER UNIT TEST")
}

//Test pake lib testify
func TestSkip(t *testing.T) { //skip test
	if runtime.GOOS == "windows" {
		t.Skip("cannot run windows")
	}
	result := HelloWorld("Rikaz")
	assert.Equal(t,"Hello Rikaz", result, "Result must be 'Hello Rikaz'")
}

func TestHelloWorldAssert(t *testing.T) { //Assert jika program error panggil Fail()
	result := HelloWorld("Rikaz")
	assert.Equal(t,"Hello Rikaz", result, "Result must be 'Hello Rikaz'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) { //Require jika program error panggil FailNow()
	result := HelloWorld("Rikaz")
	require.Equal(t,"Hello Rikaz", result, "Result must be 'Hello Rikaz'")
	fmt.Println("TestHelloWorldRequire Done")
}

//Test Manual
func TestHelloWorldRikaz(t *testing.T) {
	result := HelloWorld("Rikaz")
	if result != "Hello Rikaz" {
		//error
		// t.Fail() //jika error program masih di run hingga selesai
		t.Error("Result must be 'Hello Rikaz'") //jika error panggil Fail()
	}
	fmt.Println("TestHelloWordRikaz Done")
}

func TestHelloWorldFawaiz(t *testing.T) {
	result := HelloWorld("Fawaiz")
	if result != "Hello Fawaiz" {
		//error
		// t.FailNow() //jika error program langsung berhenti & tidak dilanjutkan eksekusi
		t.Fatal("Result must be 'Hello Fawaiz'") //jika error panggil FailNow()
	}
	fmt.Println("TestHelloWordFawaiz Done")
}