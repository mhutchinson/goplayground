package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/transparency-dev/armored-witness-common/release/firmware/ftlog"
	"github.com/transparency-dev/serverless-log/client"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	tlogURL = flag.String("url", "https://api.transparency.dev/armored-witness-firmware/ci/log/1/", "Base URL of a serverless log")
	origin  = flag.String("origin", "transparency.dev/armored-witness/firmware_transparency/ci/1", "Origin string in checkpoint")
	vkey    = flag.String("vkey", "transparency.dev-aw-ftlog-ci+f5479c1e+AR6gW0mycDtL17iM2uvQUThJsoiuSRirstEj9a5AdCCu", "Public key for log")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	f := func(ctx context.Context, path string) ([]byte, error) {
		u, err := url.Parse(*tlogURL + path)
		if err != nil {
			klog.Exitf("Failed to parse URL: %v", err)
		}
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req.WithContext(ctx))
		if err != nil {
			return nil, err
		}
		switch resp.StatusCode {
		case 404:
			klog.Infof("Not found: %q", u.String())
			return nil, os.ErrNotExist
		case 200:
			break
		default:
			return nil, fmt.Errorf("unexpected http status %q", resp.Status)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				klog.Errorf("resp.Body.Close(): %v", err)
			}
		}()
		return io.ReadAll(resp.Body)
	}
	v, err := note.NewVerifier(*vkey)
	if err != nil {
		klog.Exitf("Failed to create verifier: %v", err)
	}
	cp, _, _, err := client.FetchCheckpoint(ctx, f, v, *origin)
	if err != nil {
		klog.Exitf("Failed to fetch checkpoint: %v", err)
	}
	klog.Infof("Got checkpoint for log of size %d", cp.Size)

	for i := uint64(0); i < cp.Size; i++ {
		leaf, err := client.GetLeaf(ctx, f, i)
		if err != nil {
			klog.Errorf("Failed to get leaf %d: %v", i, err)
			continue
		}
		releaseNote, err := note.Open([]byte(leaf), note.VerifierList())
		if err != nil {
			if e, ok := err.(*note.UnverifiedNoteError); ok {
				releaseNote = e.Note
				klog.V(2).Infof("Note at leaf %d was not verified: %v", i, err)
			} else {
				klog.Errorf("Failed to open note at leaf %d: %v", i, err)
				continue
			}
		}
		var release ftlog.FirmwareRelease
		if err := json.Unmarshal([]byte(releaseNote.Text), &release); err != nil {
			klog.Errorf("Failed to unmarshal release at index %d: %v", i, err)
		}
		klog.Infof("Leaf %d: %s", i, release.Component)
	}
}
