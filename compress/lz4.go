// Copyright © 2018 Nikolas Sepos <nikolas.sepos@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compress

import (
	"io"
	"io/ioutil"

	"github.com/pierrec/lz4"
)

func init() {
	RegisterCompressor("lz4", lz4Compressor{})
}

type lz4Compressor struct{}

func (comp lz4Compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	return lz4.NewWriter(w), nil
}

func (comp lz4Compressor) Decompress(r io.Reader) (io.ReadCloser, error) {
	return ioutil.NopCloser(lz4.NewReader(r)), nil
}
