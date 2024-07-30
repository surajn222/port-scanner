package scanner

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock logrus for testing
type mockLogger struct {
	messages []string
}

func (m *mockLogger) Write(p []byte) (n int, err error) {

}

func (m *mockLogger) Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	m.messages = append(m.messages, message)
}

func (m *mockLogger) GetMessages() []string {
	return m.messages
}

func TestScanPort(t *testing.T) {
	// Initialize mock logger
	mockLog := &mockLogger{}
	originalLog := logrus.StandardLogger()
	logrus.SetOutput(mockLog)

	// Define test cases
	tests := []struct {
		target      string
		timeout     time.Duration
		expectError bool
		expectedMsg string
	}{
		{"localhost:8080", 1 * time.Second, false, "Port 8080 is open"},
		{"localhost:12345", 1 * time.Second, true, "Port 12345 is closed"}, // Assuming port 12345 is closed
	}

	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)

			// Call the ScanPort function
			ScanPort(tt.target, tt.timeout, &wg)

			// Wait for ScanPort to finish
			wg.Wait()

			// Check if an error was expected
			if tt.expectError {
				require.NotEmpty(t, mockLog.GetMessages())
				assert.Contains(t, mockLog.GetMessages()[0], tt.expectedMsg)
			} else {
				assert.Empty(t, mockLog.GetMessages())
				assert.Contains(t, mockLog.GetMessages()[0], tt.expectedMsg)
			}
		})
	}

	// Restore original logger
	logrus.SetOutput(originalLog)
}
