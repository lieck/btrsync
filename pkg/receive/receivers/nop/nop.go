package nop

import (
	"io/fs"
	"time"

	"github.com/google/uuid"

	"github.com/tinyzimmer/btrsync/pkg/btrfs"
	"github.com/tinyzimmer/btrsync/pkg/receive/receivers"
)

type nopReceiver struct{}

func New() receivers.Receiver {
	return &nopReceiver{}
}

func (n *nopReceiver) Subvol(ctx receivers.ReceiveContext, path string, uuid uuid.UUID, ctransid uint64) error {
	return nil
}

func (n *nopReceiver) Snapshot(ctx receivers.ReceiveContext, path string, uuid uuid.UUID, ctransid uint64, cloneUUID uuid.UUID, cloneCtransid uint64) error {
	return nil
}

func (n *nopReceiver) Mkfile(ctx receivers.ReceiveContext, path string, ino uint64) error {
	return nil
}

func (n *nopReceiver) Mkdir(ctx receivers.ReceiveContext, path string, ino uint64) error {
	return nil
}

func (n *nopReceiver) Mknod(ctx receivers.ReceiveContext, path string, ino uint64, mode fs.FileMode, rdev uint64) error {
	return nil
}

func (n *nopReceiver) Mkfifo(ctx receivers.ReceiveContext, path string, ino uint64) error {
	return nil
}

func (n *nopReceiver) Mksock(ctx receivers.ReceiveContext, path string, ino uint64) error {
	return nil
}

func (n *nopReceiver) Symlink(ctx receivers.ReceiveContext, path string, ino uint64, linkTo string) error {
	return nil
}

func (n *nopReceiver) Rename(ctx receivers.ReceiveContext, oldPath string, newPath string) error {
	return nil
}

func (n *nopReceiver) Link(ctx receivers.ReceiveContext, path string, linkTo string) error {
	return nil
}

func (n *nopReceiver) Unlink(ctx receivers.ReceiveContext, path string) error {
	return nil
}

func (n *nopReceiver) Rmdir(ctx receivers.ReceiveContext, path string) error {
	return nil
}

func (n *nopReceiver) Write(ctx receivers.ReceiveContext, path string, offset uint64, data []byte) error {
	return nil
}

func (n *nopReceiver) EncodedWrite(ctx receivers.ReceiveContext, path string, op *btrfs.EncodedWriteOp, forceDecompress bool) error {
	return nil
}

func (n *nopReceiver) Clone(ctx receivers.ReceiveContext, path string, offset uint64, len uint64, cloneUUID uuid.UUID, cloneCtransid uint64, clonePath string, cloneOffset uint64) error {
	return nil
}

func (n *nopReceiver) SetXattr(ctx receivers.ReceiveContext, path string, name string, data []byte) error {
	return nil
}

func (n *nopReceiver) RemoveXattr(ctx receivers.ReceiveContext, path string, name string) error {
	return nil
}

func (n *nopReceiver) Truncate(ctx receivers.ReceiveContext, path string, size uint64) error {
	return nil
}

func (n *nopReceiver) Chmod(ctx receivers.ReceiveContext, path string, mode fs.FileMode) error {
	return nil
}

func (n *nopReceiver) Chown(ctx receivers.ReceiveContext, path string, uid uint64, gid uint64) error {
	return nil
}

func (n *nopReceiver) Utimes(ctx receivers.ReceiveContext, path string, atime, mtime, ctime time.Time) error {
	return nil
}

func (n *nopReceiver) UpdateExtent(ctx receivers.ReceiveContext, path string, fileOffset uint64, tmpSize uint64) error {
	return nil
}

func (n *nopReceiver) EnableVerity(ctx receivers.ReceiveContext, path string, algorithm uint8, blockSize uint32, salt []byte, sig []byte) error {
	return nil
}

func (n *nopReceiver) Fallocate(ctx receivers.ReceiveContext, path string, mode fs.FileMode, offset uint64, len uint64) error {
	return nil
}

func (n *nopReceiver) Fileattr(ctx receivers.ReceiveContext, path string, attr uint32) error {
	return nil
}

func (n *nopReceiver) FinishSubvolume(ctx receivers.ReceiveContext) error {
	return nil
}
