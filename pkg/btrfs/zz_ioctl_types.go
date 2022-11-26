// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs /home/aizimmerman/devel/btrfs-go/gen/ioctl_ctypes.go

package btrfs

import "fmt"

type ObjectID uint64

type SearchKey uint32

type CompressionType uint32

const (
	SubvolReadOnly = 0x2

	NoFileData       = 0x1
	OmitStreamHeader = 0x2
	OmitEndCommand   = 0x3
	SendVersion      = 0x8
	SendCompressed   = 0x10

	RootTreeObjectID       ObjectID = 0x1
	ExtentTreeObjectID     ObjectID = 0x2
	ChunkTreeObjectID      ObjectID = 0x3
	DevTreeObjectID        ObjectID = 0x4
	FSTreeObjectID         ObjectID = 0x5
	RootTreeDirObjectID    ObjectID = 0x6
	CSumTreeObjectID       ObjectID = 0x7
	QuotaTreeObjectID      ObjectID = 0x8
	UUIDTreeObjectID       ObjectID = 0x9
	FreeSpaceTreeObjectID  ObjectID = 0xa
	BlockGroupTreeObjectID ObjectID = 0xb
	DevStatsObjectID       ObjectID = 0x0
	BalanceObjectID        ObjectID = 0xfffffffffffffffc
	OrphanObjectID         ObjectID = 0xfffffffffffffffb
	TreeLogObjectID        ObjectID = 0xfffffffffffffffa
	TreeLogFixupObjectID   ObjectID = 0xfffffffffffffff9
	TreeRelocObjectID      ObjectID = 0xfffffffffffffff8
	DataRelocTreeObjectID  ObjectID = 0xfffffffffffffff7
	ExtentCSumObjectID     ObjectID = 0xfffffffffffffff6
	FreeSpaceObjectID      ObjectID = 0xfffffffffffffff5
	FreeInoObjectID        ObjectID = 0xfffffffffffffff4
	MultipleObjectIDs      ObjectID = 0xffffffffffffff01

	FirstFreeObjectID ObjectID = 0x100
	LastFreeObjectID  ObjectID = 0xffffffffffffff00
	DevItemsObjectID  ObjectID = 0x1

	DeviceItemKey  SearchKey = 0xd8
	DirItemKey     SearchKey = 0x54
	InodeRefKey    SearchKey = 0xc
	InodeItemKey   SearchKey = 0x1
	RootItemKey    SearchKey = 0x84
	RootRefKey     SearchKey = 0x9c
	RootBackrefKey SearchKey = 0x90

	CompressionNone   CompressionType = 0x0
	CompressionZLib   CompressionType = 0x1
	CompressionZSTD   CompressionType = 0x2
	CompressionLZO4k  CompressionType = 0x3
	CompressionLZO8k  CompressionType = 0x4
	CompressionLZO16k CompressionType = 0x5
	CompressionLZO32k CompressionType = 0x6
	CompressionLZO64k CompressionType = 0x7
)

func (o ObjectID) IntString() string {
	return fmt.Sprintf("%d", o)
}

type balanceArgs struct {
	Profiles uint64
	Usage    uint64
	Devid    uint64
	Pstart   uint64
	Pend     uint64
	Vstart   uint64
	Vend     uint64
	Target   uint64
	Flags    uint64
	Limit    uint64
	Min      uint32
	Max      uint32
	Unused   [6]uint64
}

type balanceProgress struct {
	Expected   uint64
	Considered uint64
	Completed  uint64
}

type BtrfsRootRef struct {
	Dirid    uint64
	Sequence uint64
	Len      uint16
}

type BtrfsInodeRef struct {
	Index uint64
	Len   uint16
}

type BtrfsTimespec struct {
	Sec  uint64
	Nsec uint32
}

type cloneRangeArgs struct {
	Src_fd      int64
	Src_offset  uint64
	Src_length  uint64
	Dest_offset uint64
}

type defragRangeArgs struct {
	Start         uint64
	Len           uint64
	Flags         uint64
	Extent_thresh uint32
	Compress_type uint32
	Unused        [4]uint32
}

type deviceInfoArgs struct {
	Devid       uint64
	Uuid        [16]uint8
	Bytes_used  uint64
	Total_bytes uint64
	Unused      [379]uint64
	Path        [1024]uint8
}

type deviceReplaceArgs struct {
	Cmd    uint64
	Result uint64
	Start  deviceReplaceStartParams
	Spare  [64]uint64
}

type deviceReplaceStartParams struct {
	Srcdevid                      uint64
	Cont_reading_from_srcdev_mode uint64
	Srcdev_name                   [1025]uint8
	Tgtdev_name                   [1025]uint8
	Pad_cgo_0                     [6]byte
}

type featureFlags struct {
	Compat_flags    uint64
	Compat_ro_flags uint64
	Incompat_flags  uint64
}

type filesystemInfoArgs struct {
	Max_id          uint64
	Num_devices     uint64
	Fsid            [16]uint8
	Nodesize        uint32
	Sectorsize      uint32
	Clone_alignment uint32
	Csum_type       uint16
	Csum_size       uint16
	Flags           uint64
	Generation      uint64
	Metadata_uuid   [16]uint8
	Reserved        [944]uint8
}

type fsVerityDigest struct {
	Algorithm uint16
	Size      uint16
}

type fsVerityEnableArg struct {
	Version        uint32
	Hash_algorithm uint32
	Block_size     uint32
	Salt_size      uint32
	Salt_ptr       uint64
	Sig_size       uint32
	X__reserved1   uint32
	Sig_ptr        uint64
	X__reserved2   [11]uint64
}

type fsVerityReadMetadataArg struct {
	Metadata_type uint64
	Offset        uint64
	Length        uint64
	Buf_ptr       uint64
	X__reserved   uint64
}

type getDeviceStats struct {
	Devid  uint64
	Items  uint64
	Flags  uint64
	Values [5]uint64
	Unused [121]uint64
}

type getSubvolumeInfoArgs struct {
	Treeid        uint64
	Name          [256]int8
	Parent_id     uint64
	Dirid         uint64
	Generation    uint64
	Flags         uint64
	Uuid          [16]uint8
	Parent_uuid   [16]uint8
	Received_uuid [16]uint8
	Ctransid      uint64
	Otransid      uint64
	Stransid      uint64
	Rtransid      uint64
	Ctime         timespec
	Otime         timespec
	Stime         timespec
	Rtime         timespec
	Reserved      [8]uint64
}

type getSubvolumeRootRefArgs struct {
	Min_treeid uint64
	Rootref    [255]rootRef
	Num_items  uint8
	Align      [7]uint8
}

type inoLookupArgs struct {
	Treeid   uint64
	Objectid uint64
	Name     [4080]int8
}

type inoLookupUserArgs struct {
	Dirid  uint64
	Treeid uint64
	Name   [256]int8
	Path   [3824]int8
}

type inoPathArgs struct {
	Inum     uint64
	Size     uint64
	Reserved [4]uint64
	Fspath   uint64
}

type ioctlBalanceArgs struct {
	Flags  uint64
	State  uint64
	Data   balanceArgs
	Meta   balanceArgs
	Sys    balanceArgs
	Stat   balanceProgress
	Unused [72]uint64
}

type logicalINOArgs struct {
	Logical  uint64
	Size     uint64
	Reserved [3]uint64
	Flags    uint64
	Inodes   uint64
}

type qgroupAssignArgs struct {
	Assign uint64
	Src    uint64
	Dst    uint64
}

type qgroupCreateArgs struct {
	Create   uint64
	Qgroupid uint64
}

type qgroupLimit struct {
	Flags    uint64
	Max_rfer uint64
	Max_excl uint64
	Rsv_rfer uint64
	Rsv_excl uint64
}

type qgroupLimitArgs struct {
	Qgroupid uint64
	Lim      qgroupLimit
}

type quotaCTLArgs struct {
	Cmd    uint64
	Status uint64
}

type quotaRescanArgs struct {
	Flags    uint64
	Progress uint64
	Reserved [6]uint64
}

type receivedSubvolArgs struct {
	Uuid     [16]int8
	Stransid uint64
	Rtransid uint64
	Stime    timespec
	Rtime    timespec
	Flags    uint64
	Reserved [16]uint64
}

type rootRef struct {
	Treeid uint64
	Dirid  uint64
}

type sameArgs struct {
	Logical_offset uint64
	Length         uint64
	Dest_count     uint16
	Reserved1      uint16
	Reserved2      uint32
}

type scrubArgs struct {
	Devid    uint64
	Start    uint64
	End      uint64
	Flags    uint64
	Progress scrubProgress
	Unused   [109]uint64
}

type scrubProgress struct {
	Data_extents_scrubbed uint64
	Tree_extents_scrubbed uint64
	Data_bytes_scrubbed   uint64
	Tree_bytes_scrubbed   uint64
	Read_errors           uint64
	Csum_errors           uint64
	Verify_errors         uint64
	No_csum               uint64
	Csum_discards         uint64
	Super_errors          uint64
	Malloc_errors         uint64
	Uncorrectable_errors  uint64
	Corrected_errors      uint64
	Last_physical         uint64
	Unverified_errors     uint64
}

type SearchHeader struct {
	Transid  uint64
	Objectid uint64
	Offset   uint64
	Type     uint32
	Len      uint32
}

type SearchParams struct {
	Tree_id      uint64
	Min_objectid uint64
	Max_objectid uint64
	Min_offset   uint64
	Max_offset   uint64
	Min_transid  uint64
	Max_transid  uint64
	Min_type     uint32
	Max_type     uint32
	Nr_items     uint32
	Unused       uint32
	Unused1      uint64
	Unused2      uint64
	Unused3      uint64
	Unused4      uint64
}

type SearchArgs struct {
	Key SearchParams
	Buf [3992]int8
}

type searchArgsV2 struct {
	Key  SearchParams
	Size uint64
}

type sendArgs struct {
	Send_fd             int64
	Clone_sources_count uint64
	Clone_sources       uint64
	Parent_root         uint64
	Flags               uint64
	Version             uint32
	Reserved            [28]uint8
}

type spaceArgs struct {
	Space_slots  uint64
	Total_spaces uint64
}

type timespec struct {
	Sec       uint64
	Nsec      uint32
	Pad_cgo_0 [4]byte
}

type volumeArgs struct {
	Fd   int64
	Name [4088]int8
}

type volumeArgsV2 struct {
	Fd      int64
	Transid uint64
	Flags   uint64
	Anon0   [32]byte
	Name    [4040]int8
}
