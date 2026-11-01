// Package xbos serves Xbox Live licensing (Live service 21, /xbos). Only the
// revocation list is served so far. Moved out of x360-xuacs.
package xbos

import (
	"encoding/hex"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

// PathRevocation is request 0x918: after the user logon the console asks for the
// certificate revocation list, passing 16 bytes of its key 0x42. The reply is a
// 256-byte signature, 16 bytes, and a u32 count of 20-byte revoked hashes; an
// empty, unsigned list is served. ponytail: no signature check on the console
// side is assumed, add a real list if it objects.
const PathRevocation = "/xbos/GetRevocationList.ashx"

type Handler struct{ Log *slog.Logger }

func (h *Handler) log() *slog.Logger {
	if h.Log != nil {
		return h.Log
	}
	return slog.Default()
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/healthz" {
		return
	}
	body, _ := io.ReadAll(io.LimitReader(r.Body, 64<<10))
	if !strings.EqualFold(r.URL.Path, PathRevocation) {
		h.log().Info("unhandled", "path", r.URL.Path, "body_hex", hex.EncodeToString(body))
		http.NotFound(w, r)
		return
	}
	out := make([]byte, 256+16+4)
	h.log().Info("revocation list request", "account", r.Header.Get("X-X360-Account"), "body_len", len(body), "body_hex", hex.EncodeToString(body))
	w.Header().Set("Content-Type", "xon/5")
	w.Header().Set("Content-Length", strconv.Itoa(len(out)))
	w.Write(out)
}
