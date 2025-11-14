package rpmver

func CompareEVR(epoch1, version1, release1, epoch2, version2, release2 string) int {
	e1 := "0"
	if epoch1 != "" {
		e1 = epoch1
	}
	e2 := "0"
	if epoch2 != "" {
		e2 = epoch2
	}

	if e1 < e2 {
		return -1
	}
	if e1 > e2 {
		return 1
	}
	vcmp := RpmVersionCompare(version1, version2)
	if vcmp != 0 {
		return vcmp
	}
	return RpmVersionCompare(release1, release2)
}
