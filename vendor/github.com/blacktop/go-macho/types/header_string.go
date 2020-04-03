// Code generated by "stringer -type=HeaderType,HeaderFlag -output header_string.go"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Obj-1]
	_ = x[Exec-2]
	_ = x[FVMLib-3]
	_ = x[Core-4]
	_ = x[Preload-5]
	_ = x[Dylib-6]
	_ = x[Dylinker-7]
	_ = x[Bundle-8]
	_ = x[DylibStub-9]
	_ = x[Dsym-10]
	_ = x[KextBundle-11]
}

const _HeaderType_name = "ObjExecFVMLibCorePreloadDylibDylinkerBundleDylibStubDsymKextBundle"

var _HeaderType_index = [...]uint8{0, 3, 7, 13, 17, 24, 29, 37, 43, 52, 56, 66}

func (i HeaderType) String() string {
	i -= 1
	if i >= HeaderType(len(_HeaderType_index)-1) {
		return "HeaderType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _HeaderType_name[_HeaderType_index[i]:_HeaderType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoUndefs-1]
	_ = x[IncrLink-2]
	_ = x[DyldLink-4]
	_ = x[BindAtLoad-8]
	_ = x[Prebound-16]
	_ = x[SplitSegs-32]
	_ = x[LazyInit-64]
	_ = x[TwoLevel-128]
	_ = x[ForceFlat-256]
	_ = x[NoMultiDefs-512]
	_ = x[NoFixPrebinding-1024]
	_ = x[Prebindable-2048]
	_ = x[AllModsBound-4096]
	_ = x[SubsectionsViaSymbols-8192]
	_ = x[Canonical-16384]
	_ = x[WeakDefines-32768]
	_ = x[BindsToWeak-65536]
	_ = x[AllowStackExecution-131072]
	_ = x[RootSafe-262144]
	_ = x[SetuidSafe-524288]
	_ = x[NoReexportedDylibs-1048576]
	_ = x[PIE-2097152]
	_ = x[DeadStrippableDylib-4194304]
	_ = x[HasTLVDescriptors-8388608]
	_ = x[NoHeapExecution-16777216]
	_ = x[AppExtensionSafe-33554432]
	_ = x[NlistOutofsyncWithDyldinfo-67108864]
	_ = x[SimSupport-134217728]
	_ = x[DylibInCache-2147483648]
}

const _HeaderFlag_name = "NoUndefsIncrLinkDyldLinkBindAtLoadPreboundSplitSegsLazyInitTwoLevelForceFlatNoMultiDefsNoFixPrebindingPrebindableAllModsBoundSubsectionsViaSymbolsCanonicalWeakDefinesBindsToWeakAllowStackExecutionRootSafeSetuidSafeNoReexportedDylibsPIEDeadStrippableDylibHasTLVDescriptorsNoHeapExecutionAppExtensionSafeNlistOutofsyncWithDyldinfoSimSupportDylibInCache"

var _HeaderFlag_map = map[HeaderFlag]string{
	1:          _HeaderFlag_name[0:8],
	2:          _HeaderFlag_name[8:16],
	4:          _HeaderFlag_name[16:24],
	8:          _HeaderFlag_name[24:34],
	16:         _HeaderFlag_name[34:42],
	32:         _HeaderFlag_name[42:51],
	64:         _HeaderFlag_name[51:59],
	128:        _HeaderFlag_name[59:67],
	256:        _HeaderFlag_name[67:76],
	512:        _HeaderFlag_name[76:87],
	1024:       _HeaderFlag_name[87:102],
	2048:       _HeaderFlag_name[102:113],
	4096:       _HeaderFlag_name[113:125],
	8192:       _HeaderFlag_name[125:146],
	16384:      _HeaderFlag_name[146:155],
	32768:      _HeaderFlag_name[155:166],
	65536:      _HeaderFlag_name[166:177],
	131072:     _HeaderFlag_name[177:196],
	262144:     _HeaderFlag_name[196:204],
	524288:     _HeaderFlag_name[204:214],
	1048576:    _HeaderFlag_name[214:232],
	2097152:    _HeaderFlag_name[232:235],
	4194304:    _HeaderFlag_name[235:254],
	8388608:    _HeaderFlag_name[254:271],
	16777216:   _HeaderFlag_name[271:286],
	33554432:   _HeaderFlag_name[286:302],
	67108864:   _HeaderFlag_name[302:328],
	134217728:  _HeaderFlag_name[328:338],
	2147483648: _HeaderFlag_name[338:350],
}

func (i HeaderFlag) String() string {
	if str, ok := _HeaderFlag_map[i]; ok {
		return str
	}
	return "HeaderFlag(" + strconv.FormatInt(int64(i), 10) + ")"
}
