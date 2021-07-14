/*
Copyright © 2021 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/blacktop/ipsw/pkg/dyld"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	dyldCmd.AddCommand(slideCmd)
	slideCmd.Flags().BoolP("auth", "a", false, "Print only slide info for mappings with auth flags")
	slideCmd.MarkZshCompPositionalArgumentFile(1, "dyld_shared_cache*")
}

// slideCmd represents the slide command
var slideCmd = &cobra.Command{
	Use:   "slide [options] <dyld_shared_cache>",
	Short: "Get slide info chained pointers",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		printAuthSlideInfo, _ := cmd.Flags().GetBool("auth")

		dscPath := filepath.Clean(args[0])

		fileInfo, err := os.Lstat(dscPath)
		if err != nil {
			return fmt.Errorf("file %s does not exist", dscPath)
		}

		// Check if file is a symlink
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			symlinkPath, err := os.Readlink(dscPath)
			if err != nil {
				return errors.Wrapf(err, "failed to read symlink %s", dscPath)
			}
			// TODO: this seems like it would break
			linkParent := filepath.Dir(dscPath)
			linkRoot := filepath.Dir(linkParent)

			dscPath = filepath.Join(linkRoot, symlinkPath)
		}

		f, err := dyld.Open(dscPath)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := os.Stat(dscPath + ".a2s"); os.IsNotExist(err) {
			log.Warn("parsing public symbols...")
			err = f.GetAllExportedSymbols(false)
			if err != nil {
				// return err
				log.Errorf("failed to parse all exported symbols: %v", err)
			}
			log.Warn("parsing private symbols...")
			err = f.ParseLocalSyms()
			if err != nil {
				return err
			}
			err = f.SaveAddrToSymMap(dscPath + ".a2s")
			if err != nil {
				return err
			}
		} else {
			a2sFile, err := os.Open(dscPath + ".a2s")
			if err != nil {
				return err
			}
			gzr, err := gzip.NewReader(a2sFile)
			if err != nil {
				return fmt.Errorf("failed to create gzip reader: %v", err)
			}
			// Decoding the serialized data
			err = gob.NewDecoder(gzr).Decode(&f.AddressToSymbol)
			if err != nil {
				return err
			}
			gzr.Close()
			a2sFile.Close()
		}

		if f.SlideInfoOffsetUnused > 0 {
			f.ParseSlideInfo(dyld.CacheMappingAndSlideInfo{
				Address:         f.Mappings[1].Address,
				Size:            f.Mappings[1].Size,
				FileOffset:      f.Mappings[1].FileOffset,
				SlideInfoOffset: f.SlideInfoOffsetUnused,
				SlideInfoSize:   f.SlideInfoSizeUnused,
			}, false)
		} else {
			for _, extMapping := range f.MappingsWithSlideInfo {
				if printAuthSlideInfo && !extMapping.Flags.IsAuthData() {
					continue
				}
				if extMapping.SlideInfoSize > 0 {
					f.ParseSlideInfo(extMapping.CacheMappingAndSlideInfo, true)
				}
			}
		}

		return nil
	},
}
