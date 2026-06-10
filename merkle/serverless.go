package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/transparency-dev/armored-witness-common/release/firmware/ftlog"
	"github.com/transparency-dev/serverless-log/client"
	"golang.org/x/mod/sumdb/note"
)

var (
	tlogURL = flag.String("url", "https://api.transparency.dev/armored-witness-firmware/ci/log/3/", "Base URL of a serverless log")
	origin  = flag.String("origin", "transparency.dev/armored-witness/firmware_transparency/ci/3", "Origin string in checkpoint")
	vkey    = flag.String("vkey", "transparency.dev-aw-ftlog-ci-3+3f689522+Aa1Eifq6rRC8qiK+bya07yV1fXyP156pEMsX7CFBC6gg", "Public key for log")
	start   = flag.Uint64("start", 0, "First index to start outputting details")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	f := func(ctx context.Context, path string) ([]byte, error) {
		u, err := url.Parse(*tlogURL + path)
		if err != nil {
			slog.Error("Failed to parse URL", "error", err)
			os.Exit(1)
		}
		req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		switch resp.StatusCode {
		case 404:
			slog.Info("Not found", "url", u.String())
			return nil, os.ErrNotExist
		case 200:
			break
		default:
			return nil, fmt.Errorf("unexpected http status %q", resp.Status)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				slog.Error("resp.Body.Close() failed", "error", err)
			}
		}()
		return io.ReadAll(resp.Body)
	}
	v, err := note.NewVerifier(*vkey)
	if err != nil {
		slog.Error("Failed to create verifier", "error", err)
		os.Exit(1)
	}
	cp, _, _, err := client.FetchCheckpoint(ctx, f, v, *origin)
	if err != nil {
		slog.Error("Failed to fetch checkpoint", "error", err)
		os.Exit(1)
	}
	slog.Info("Got checkpoint", "size", cp.Size)

	for i := *start; i < cp.Size; i++ {
		leaf, err := client.GetLeaf(ctx, f, i)
		if err != nil {
			slog.Error("Failed to get leaf", "index", i, "error", err)
			continue
		}
		releaseNote, err := note.Open([]byte(leaf), note.VerifierList())
		if err != nil {
			if e, ok := err.(*note.UnverifiedNoteError); ok {
				releaseNote = e.Note
				slog.Debug("Note was not verified", "index", i, "error", err)
			} else {
				slog.Error("Failed to open note", "index", i, "error", err)
				continue
			}
		}
		var release ftlog.FirmwareRelease
		if err := json.Unmarshal([]byte(releaseNote.Text), &release); err != nil {
			slog.Error("Failed to unmarshal release", "index", i, "error", err)
		}
		slog.Info("Leaf info",
			"index", i,
			"hex", fmt.Sprintf("%x", i),
			"component", release.Component,
			"tag", release.Git.TagName,
			"commit", fmt.Sprintf("%.6s", release.Git.CommitFingerprint),
		)
	}
}
