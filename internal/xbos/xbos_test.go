package xbos

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestRevocationList(t *testing.T) {
	rec := httptest.NewRecorder()
	(&Handler{}).ServeHTTP(rec, httptest.NewRequest("POST", PathRevocation, bytes.NewReader(make([]byte, 16))))
	if rec.Code != 200 || rec.Body.Len() != 276 || !bytes.Equal(rec.Body.Bytes(), make([]byte, 276)) {
		t.Fatalf("revocation reply %d %d bytes", rec.Code, rec.Body.Len())
	}
}
