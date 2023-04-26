package download

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/irorikon/tools/command/flags"
	"github.com/irorikon/tools/util"
	"github.com/pkg/errors"
)

type Get struct {
	Trace  bool
	Output string
	Procs  int
	URLs   []string
	// args      []string
	timeout   int
	useragent string
	referer   string
}

// New for get package
func New() *Get {
	return &Get{
		Trace:   false,
		Procs:   runtime.NumCPU(), // default
		timeout: 10,
	}
}

// Run execute methods in get package
func (g *Get) Run(ctx context.Context, version string, args []string) error {
	if err := g.Ready(version, args); err != nil {
		return err
	}

	// TODO(codehex): calc maxIdleConnsPerHost
	client := newDownloadClinet(16)

	target, err := Check(ctx, &CheckConfig{
		URLs:    g.URLs,
		Timeout: time.Duration(g.timeout) * time.Second,
		Client:  client,
	})
	if err != nil {
		return err
	}

	filename := target.Filename

	var dir string
	if g.Output != "" {
		fi, err := os.Stat(g.Output)
		if err == nil && fi.IsDir() {
			dir = g.Output
		} else {
			dir, filename = filepath.Split(g.Output)
			if dir != "" {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return errors.Wrapf(err, "failed to create diretory at %s", dir)
				}
			}
		}
	}

	opts := []DownloadOption{
		WithUserAgent(g.useragent, version),
		WithReferer(g.referer),
	}

	return Download(ctx, &DownloadConfig{
		Filename:      filename,
		Dirname:       dir,
		ContentLength: target.ContentLength,
		Procs:         g.Procs,
		URLs:          target.URLs,
		Client:        client,
	}, opts...)
}

// Ready method define the variables required to Download.
func (g *Get) Ready(version string, args []string) error {
	if flags.Trace {
		g.Trace = flags.Trace
	}

	if flags.Timeout > 0 {
		g.timeout = flags.Timeout
	}

	if err := g.parseURLs(); err != nil {
		return errors.Wrap(err, "failed to parse of url")
	}

	g.Procs = flags.NumConnection * len(g.URLs)

	if flags.Output != "" {
		g.Output = flags.Output
	}

	if flags.UserAgent {
		g.useragent = util.GetRandomUserAgent()
	}

	if flags.Referer != "" {
		g.referer = flags.Referer
	}

	return nil
}

func (g *Get) parseURLs() error {

	if flags.URL == nil {
		fmt.Fprintf(stdout, "Please input url separate with space or newline\n")
		fmt.Fprintf(stdout, "Start download with ^D\n")

		// scanning url from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			scan := scanner.Text()
			urls := strings.Split(scan, " ")
			for _, url := range urls {
				if govalidator.IsURL(url) {
					g.URLs = append(g.URLs, url)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			return errors.Wrap(err, "failed to parse url from stdin")
		}

		if len(g.URLs) < 1 {
			return errors.New("urls not found in the arguments passed")
		}
	} else {
		for _, url := range flags.URL {
			if govalidator.IsURL(url) {
				g.URLs = append(g.URLs, url)
			}
		}
	}
	return nil
}
