// Code generated by "stringer -type IoctlCmd,ObjectID,SearchKey,CompressionType -output zz_stringers.go"; DO NOT EDIT.

package btrfs

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FS_IOC_ENABLE_VERITY-1082156677]
	_ = x[FS_IOC_MEASURE_VERITY-3221513862]
	_ = x[FS_IOC_READ_VERITY_METADATA-3223873159]
	_ = x[BTRFS_IOC_SNAP_CREATE-1342215169]
	_ = x[BTRFS_IOC_DEFRAG-1342215170]
	_ = x[BTRFS_IOC_RESIZE-1342215171]
	_ = x[BTRFS_IOC_SCAN_DEV-1342215172]
	_ = x[BTRFS_IOC_FORGET_DEV-1342215173]
	_ = x[BTRFS_IOC_TRANS_START-37894]
	_ = x[BTRFS_IOC_TRANS_END-37895]
	_ = x[BTRFS_IOC_SYNC-37896]
	_ = x[BTRFS_IOC_CLONE-1074041865]
	_ = x[BTRFS_IOC_ADD_DEV-1342215178]
	_ = x[BTRFS_IOC_RM_DEV-1342215179]
	_ = x[BTRFS_IOC_BALANCE-1342215180]
	_ = x[BTRFS_IOC_CLONE_RANGE-1075876877]
	_ = x[BTRFS_IOC_SUBVOL_CREATE-1342215182]
	_ = x[BTRFS_IOC_SNAP_DESTROY-1342215183]
	_ = x[BTRFS_IOC_DEFRAG_RANGE-1076925456]
	_ = x[BTRFS_IOC_TREE_SEARCH-3489698833]
	_ = x[BTRFS_IOC_TREE_SEARCH_V2-3228603409]
	_ = x[BTRFS_IOC_INO_LOOKUP-3489698834]
	_ = x[BTRFS_IOC_DEFAULT_SUBVOL-1074304019]
	_ = x[BTRFS_IOC_SPACE_INFO-3222311956]
	_ = x[BTRFS_IOC_START_SYNC-2148045848]
	_ = x[BTRFS_IOC_WAIT_SYNC-1074304022]
	_ = x[BTRFS_IOC_SNAP_CREATE_V2-1342215191]
	_ = x[BTRFS_IOC_SUBVOL_CREATE_V2-1342215192]
	_ = x[BTRFS_IOC_SUBVOL_GETFLAGS-2148045849]
	_ = x[BTRFS_IOC_SUBVOL_SETFLAGS-1074304026]
	_ = x[BTRFS_IOC_SCRUB-3288372251]
	_ = x[BTRFS_IOC_SCRUB_CANCEL-37916]
	_ = x[BTRFS_IOC_SCRUB_PROGRESS-3288372253]
	_ = x[BTRFS_IOC_DEV_INFO-3489698846]
	_ = x[BTRFS_IOC_FS_INFO-2214630431]
	_ = x[BTRFS_IOC_BALANCE_V2-3288372256]
	_ = x[BTRFS_IOC_BALANCE_CTL-1074041889]
	_ = x[BTRFS_IOC_BALANCE_PROGRESS-2214630434]
	_ = x[BTRFS_IOC_INO_PATHS-3224933411]
	_ = x[BTRFS_IOC_LOGICAL_INO-3224933412]
	_ = x[BTRFS_IOC_SET_RECEIVED_SUBVOL-3234370597]
	_ = x[BTRFS_IOC_SEND-1078498342]
	_ = x[BTRFS_IOC_DEVICES_READY-2415957031]
	_ = x[BTRFS_IOC_QUOTA_CTL-3222311976]
	_ = x[BTRFS_IOC_QGROUP_ASSIGN-1075352617]
	_ = x[BTRFS_IOC_QGROUP_CREATE-1074828330]
	_ = x[BTRFS_IOC_QGROUP_LIMIT-2150667307]
	_ = x[BTRFS_IOC_QUOTA_RESCAN-1077974060]
	_ = x[BTRFS_IOC_QUOTA_RESCAN_STATUS-2151715885]
	_ = x[BTRFS_IOC_QUOTA_RESCAN_WAIT-37934]
	_ = x[BTRFS_IOC_GET_DEV_STATS-3288896564]
	_ = x[BTRFS_IOC_DEV_REPLACE-3391657013]
	_ = x[BTRFS_IOC_FILE_EXTENT_SAME-3222836278]
	_ = x[BTRFS_IOC_RM_DEV_V2-1342215226]
	_ = x[BTRFS_IOC_LOGICAL_INO_V2-3224933435]
	_ = x[BTRFS_IOC_GET_SUBVOL_INFO-2180551740]
	_ = x[BTRFS_IOC_GET_SUBVOL_ROOTREF-3489698877]
	_ = x[BTRFS_IOC_INO_LOOKUP_USER-3489698878]
	_ = x[BTRFS_IOC_SNAP_DESTROY_V2-1342215231]
	_ = x[BTRFS_IOC_ENCODED_READ-2155910208]
	_ = x[BTRFS_IOC_ENCODED_WRITE-1082168384]
}

const _IoctlCmd_name = "BTRFS_IOC_TRANS_STARTBTRFS_IOC_TRANS_ENDBTRFS_IOC_SYNCBTRFS_IOC_SCRUB_CANCELBTRFS_IOC_QUOTA_RESCAN_WAITBTRFS_IOC_CLONEBTRFS_IOC_BALANCE_CTLBTRFS_IOC_DEFAULT_SUBVOLBTRFS_IOC_WAIT_SYNCBTRFS_IOC_SUBVOL_SETFLAGSBTRFS_IOC_QGROUP_CREATEBTRFS_IOC_QGROUP_ASSIGNBTRFS_IOC_CLONE_RANGEBTRFS_IOC_DEFRAG_RANGEBTRFS_IOC_QUOTA_RESCANBTRFS_IOC_SENDFS_IOC_ENABLE_VERITYBTRFS_IOC_ENCODED_WRITEBTRFS_IOC_SNAP_CREATEBTRFS_IOC_DEFRAGBTRFS_IOC_RESIZEBTRFS_IOC_SCAN_DEVBTRFS_IOC_FORGET_DEVBTRFS_IOC_ADD_DEVBTRFS_IOC_RM_DEVBTRFS_IOC_BALANCEBTRFS_IOC_SUBVOL_CREATEBTRFS_IOC_SNAP_DESTROYBTRFS_IOC_SNAP_CREATE_V2BTRFS_IOC_SUBVOL_CREATE_V2BTRFS_IOC_RM_DEV_V2BTRFS_IOC_SNAP_DESTROY_V2BTRFS_IOC_START_SYNCBTRFS_IOC_SUBVOL_GETFLAGSBTRFS_IOC_QGROUP_LIMITBTRFS_IOC_QUOTA_RESCAN_STATUSBTRFS_IOC_ENCODED_READBTRFS_IOC_GET_SUBVOL_INFOBTRFS_IOC_FS_INFOBTRFS_IOC_BALANCE_PROGRESSBTRFS_IOC_DEVICES_READYFS_IOC_MEASURE_VERITYBTRFS_IOC_SPACE_INFOBTRFS_IOC_QUOTA_CTLBTRFS_IOC_FILE_EXTENT_SAMEFS_IOC_READ_VERITY_METADATABTRFS_IOC_INO_PATHSBTRFS_IOC_LOGICAL_INOBTRFS_IOC_LOGICAL_INO_V2BTRFS_IOC_TREE_SEARCH_V2BTRFS_IOC_SET_RECEIVED_SUBVOLBTRFS_IOC_SCRUBBTRFS_IOC_SCRUB_PROGRESSBTRFS_IOC_BALANCE_V2BTRFS_IOC_GET_DEV_STATSBTRFS_IOC_DEV_REPLACEBTRFS_IOC_TREE_SEARCHBTRFS_IOC_INO_LOOKUPBTRFS_IOC_DEV_INFOBTRFS_IOC_GET_SUBVOL_ROOTREFBTRFS_IOC_INO_LOOKUP_USER"

var _IoctlCmd_map = map[IoctlCmd]string{
	37894:      _IoctlCmd_name[0:21],
	37895:      _IoctlCmd_name[21:40],
	37896:      _IoctlCmd_name[40:54],
	37916:      _IoctlCmd_name[54:76],
	37934:      _IoctlCmd_name[76:103],
	1074041865: _IoctlCmd_name[103:118],
	1074041889: _IoctlCmd_name[118:139],
	1074304019: _IoctlCmd_name[139:163],
	1074304022: _IoctlCmd_name[163:182],
	1074304026: _IoctlCmd_name[182:207],
	1074828330: _IoctlCmd_name[207:230],
	1075352617: _IoctlCmd_name[230:253],
	1075876877: _IoctlCmd_name[253:274],
	1076925456: _IoctlCmd_name[274:296],
	1077974060: _IoctlCmd_name[296:318],
	1078498342: _IoctlCmd_name[318:332],
	1082156677: _IoctlCmd_name[332:352],
	1082168384: _IoctlCmd_name[352:375],
	1342215169: _IoctlCmd_name[375:396],
	1342215170: _IoctlCmd_name[396:412],
	1342215171: _IoctlCmd_name[412:428],
	1342215172: _IoctlCmd_name[428:446],
	1342215173: _IoctlCmd_name[446:466],
	1342215178: _IoctlCmd_name[466:483],
	1342215179: _IoctlCmd_name[483:499],
	1342215180: _IoctlCmd_name[499:516],
	1342215182: _IoctlCmd_name[516:539],
	1342215183: _IoctlCmd_name[539:561],
	1342215191: _IoctlCmd_name[561:585],
	1342215192: _IoctlCmd_name[585:611],
	1342215226: _IoctlCmd_name[611:630],
	1342215231: _IoctlCmd_name[630:655],
	2148045848: _IoctlCmd_name[655:675],
	2148045849: _IoctlCmd_name[675:700],
	2150667307: _IoctlCmd_name[700:722],
	2151715885: _IoctlCmd_name[722:751],
	2155910208: _IoctlCmd_name[751:773],
	2180551740: _IoctlCmd_name[773:798],
	2214630431: _IoctlCmd_name[798:815],
	2214630434: _IoctlCmd_name[815:841],
	2415957031: _IoctlCmd_name[841:864],
	3221513862: _IoctlCmd_name[864:885],
	3222311956: _IoctlCmd_name[885:905],
	3222311976: _IoctlCmd_name[905:924],
	3222836278: _IoctlCmd_name[924:950],
	3223873159: _IoctlCmd_name[950:977],
	3224933411: _IoctlCmd_name[977:996],
	3224933412: _IoctlCmd_name[996:1017],
	3224933435: _IoctlCmd_name[1017:1041],
	3228603409: _IoctlCmd_name[1041:1065],
	3234370597: _IoctlCmd_name[1065:1094],
	3288372251: _IoctlCmd_name[1094:1109],
	3288372253: _IoctlCmd_name[1109:1133],
	3288372256: _IoctlCmd_name[1133:1153],
	3288896564: _IoctlCmd_name[1153:1176],
	3391657013: _IoctlCmd_name[1176:1197],
	3489698833: _IoctlCmd_name[1197:1218],
	3489698834: _IoctlCmd_name[1218:1238],
	3489698846: _IoctlCmd_name[1238:1256],
	3489698877: _IoctlCmd_name[1256:1284],
	3489698878: _IoctlCmd_name[1284:1309],
}

func (i IoctlCmd) String() string {
	if str, ok := _IoctlCmd_map[i]; ok {
		return str
	}
	return "IoctlCmd(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RootTreeObjectID-1]
	_ = x[ExtentTreeObjectID-2]
	_ = x[ChunkTreeObjectID-3]
	_ = x[DevTreeObjectID-4]
	_ = x[FSTreeObjectID-5]
	_ = x[RootTreeDirObjectID-6]
	_ = x[CSumTreeObjectID-7]
	_ = x[QuotaTreeObjectID-8]
	_ = x[UUIDTreeObjectID-9]
	_ = x[FreeSpaceTreeObjectID-10]
	_ = x[BlockGroupTreeObjectID-11]
	_ = x[DevStatsObjectID-0]
	_ = x[BalanceObjectID-18446744073709551612]
	_ = x[OrphanObjectID-18446744073709551611]
	_ = x[TreeLogObjectID-18446744073709551610]
	_ = x[TreeLogFixupObjectID-18446744073709551609]
	_ = x[TreeRelocObjectID-18446744073709551608]
	_ = x[DataRelocTreeObjectID-18446744073709551607]
	_ = x[ExtentCSumObjectID-18446744073709551606]
	_ = x[FreeSpaceObjectID-18446744073709551605]
	_ = x[FreeInoObjectID-18446744073709551604]
	_ = x[MultipleObjectIDs-18446744073709551361]
	_ = x[FirstFreeObjectID-256]
	_ = x[LastFreeObjectID-18446744073709551360]
	_ = x[DevItemsObjectID-1]
}

const (
	_ObjectID_name_0 = "DevStatsObjectIDRootTreeObjectIDExtentTreeObjectIDChunkTreeObjectIDDevTreeObjectIDFSTreeObjectIDRootTreeDirObjectIDCSumTreeObjectIDQuotaTreeObjectIDUUIDTreeObjectIDFreeSpaceTreeObjectIDBlockGroupTreeObjectID"
	_ObjectID_name_1 = "FirstFreeObjectID"
	_ObjectID_name_2 = "LastFreeObjectIDMultipleObjectIDs"
	_ObjectID_name_3 = "FreeInoObjectIDFreeSpaceObjectIDExtentCSumObjectIDDataRelocTreeObjectIDTreeRelocObjectIDTreeLogFixupObjectIDTreeLogObjectIDOrphanObjectIDBalanceObjectID"
)

var (
	_ObjectID_index_0 = [...]uint8{0, 16, 32, 50, 67, 82, 96, 115, 131, 148, 164, 185, 207}
	_ObjectID_index_2 = [...]uint8{0, 16, 33}
	_ObjectID_index_3 = [...]uint8{0, 15, 32, 50, 71, 88, 108, 123, 137, 152}
)

func (i ObjectID) String() string {
	switch {
	case i <= 11:
		return _ObjectID_name_0[_ObjectID_index_0[i]:_ObjectID_index_0[i+1]]
	case i == 256:
		return _ObjectID_name_1
	case 18446744073709551360 <= i && i <= 18446744073709551361:
		i -= 18446744073709551360
		return _ObjectID_name_2[_ObjectID_index_2[i]:_ObjectID_index_2[i+1]]
	case 18446744073709551604 <= i && i <= 18446744073709551612:
		i -= 18446744073709551604
		return _ObjectID_name_3[_ObjectID_index_3[i]:_ObjectID_index_3[i+1]]
	default:
		return "ObjectID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeviceItemKey-216]
	_ = x[DirItemKey-84]
	_ = x[InodeRefKey-12]
	_ = x[InodeItemKey-1]
	_ = x[RootItemKey-132]
	_ = x[RootRefKey-156]
	_ = x[RootBackrefKey-144]
}

const (
	_SearchKey_name_0 = "InodeItemKey"
	_SearchKey_name_1 = "InodeRefKey"
	_SearchKey_name_2 = "DirItemKey"
	_SearchKey_name_3 = "RootItemKey"
	_SearchKey_name_4 = "RootBackrefKey"
	_SearchKey_name_5 = "RootRefKey"
	_SearchKey_name_6 = "DeviceItemKey"
)

func (i SearchKey) String() string {
	switch {
	case i == 1:
		return _SearchKey_name_0
	case i == 12:
		return _SearchKey_name_1
	case i == 84:
		return _SearchKey_name_2
	case i == 132:
		return _SearchKey_name_3
	case i == 144:
		return _SearchKey_name_4
	case i == 156:
		return _SearchKey_name_5
	case i == 216:
		return _SearchKey_name_6
	default:
		return "SearchKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompressionNone-0]
	_ = x[CompressionZLib-1]
	_ = x[CompressionZSTD-2]
	_ = x[CompressionLZO4k-3]
	_ = x[CompressionLZO8k-4]
	_ = x[CompressionLZO16k-5]
	_ = x[CompressionLZO32k-6]
	_ = x[CompressionLZO64k-7]
}

const _CompressionType_name = "CompressionNoneCompressionZLibCompressionZSTDCompressionLZO4kCompressionLZO8kCompressionLZO16kCompressionLZO32kCompressionLZO64k"

var _CompressionType_index = [...]uint8{0, 15, 30, 45, 61, 77, 94, 111, 128}

func (i CompressionType) String() string {
	if i >= CompressionType(len(_CompressionType_index)-1) {
		return "CompressionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompressionType_name[_CompressionType_index[i]:_CompressionType_index[i+1]]
}
