package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	for i := 1; i <= 9; i++ {
		inputFileName := fmt.Sprintf("testdata/input_tests/test_case_%d_input.txt", i)
		outputFileName := fmt.Sprintf("testdata/output_tests/test_case_%d_output.txt", i)

		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			input, err := os.ReadFile(inputFileName)
			if err != nil {
				t.Fatalf("Error reading input file: %v", err)
			}
			expectedOutput, err := os.ReadFile(outputFileName)
			if err != nil {
				t.Fatalf("Error reading expected output file: %v", err)
			}

			oldStdout := os.Stdout
			_, w, _ := os.Pipe()

			os.Stdout = w

			os.Args = []string{"main", string(input)}
			main()

			w.Close()
			os.Stdout = oldStdout

			got, _ := os.ReadFile(outputFileName)
			result := strings.TrimSpace(string(got))
			expected := strings.TrimSpace(string(expectedOutput))

			if result != expected {
				t.Errorf("Test failed. Expected:\n%s\nGot:\n%s", expected, result)
			}

			indicator := fmt.Sprintf("Passed!\nInput:%s\nExpected:\n%s\nGot:\n%s", input, expected, result)
			t.Log(indicator)
		})
	}
}

