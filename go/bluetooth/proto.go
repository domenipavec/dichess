package bluetooth

import (
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func readProto(r io.Reader, msg proto.Message) error {
	var length uint64
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return errors.Wrap(err, "could not read length")
	}
	data := make([]byte, length)
	if _, err := io.ReadFull(r, data); err != nil {
		return errors.Wrapf(err, "couldn't read msg data of length %d", length)
	}
	if err := proto.Unmarshal(data, msg); err != nil {
		return errors.Wrap(err, "couldn't unmarshal msg")
	}
	return nil
}

func writeProto(w io.Writer, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "could not marshal proto message")
	}

	if err := binary.Write(w, binary.BigEndian, uint64(len(data))); err != nil {
		return errors.Wrap(err, "could not write data length")
	}

	_, err = w.Write(data)
	if err != nil {
		return errors.Wrap(err, "could not write data")
	}

	return nil
}
