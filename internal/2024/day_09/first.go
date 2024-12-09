package day_09

func (d Day) firstPart() any {
	indices := d.readFileSystem()

	var left = -1
	advanceLeft(&left, indices)

	var right = len(indices)
	advanceRight(&right, indices)

	for left < right {
		indices[left] = indices[right]
		indices[right] = -1
		advanceLeft(&left, indices)
		advanceRight(&right, indices)
	}

	return d.computeChecksum(indices)
}

func advanceLeft(left *int, fileSystem []int) {
	*left++
	for *left < len(fileSystem) && fileSystem[*left] != -1 {
		*left++
	}
}

func advanceRight(right *int, fileSystem []int) {
	*right--
	for *right >= 0 && fileSystem[*right] == -1 {
		*right--
	}
}
