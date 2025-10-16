// nolint
package externalsort

import (
	"bytes"
	"fmt"
	"io"
)

func ReadLine(reader io.ReadSeeker) ([]byte, error) {
	startPos, err := reader.Seek(0, io.SeekCurrent)
	var readOffset int64
	var readBuffer bytes.Buffer

	for {
		tmp := make([]byte, 256)
		n, err := reader.Read(tmp)
		if n > 0 {
			if pos := bytes.IndexByte(tmp, '\n'); pos > -1 {
				readOffset += int64(pos) + 1
				if _, err = readBuffer.Write(append(tmp[:pos], []byte("\n")...)); err != nil {
					return nil, err
				}
				break
			}

			readOffset += int64(n)
			if _, err = readBuffer.Write(tmp); err != nil {
				return nil, err
			}
		}

		if err != nil {
			if err == io.EOF && readBuffer.Len() > 0 {
				if _, err = readBuffer.Write([]byte("\n")); err != nil {
					return nil, err
				}
				break
			}
			return nil, err
		}

	}

	if _, err = reader.Seek(startPos+readOffset, io.SeekStart); err != nil {
		return nil, fmt.Errorf("on Seek: %w", err)
	}

	return readBuffer.Bytes(), nil
}
