package day_09

func (d Day) secondPart() any {
	indices := d.readFileSystem()
	right := len(indices) - 1

	// Locate the last valid element from the end
	for right >= 0 && indices[right] == -1 {
		right--
	}

	for right >= 0 {
		// Find the block of identical elements
		tmp := right
		for tmp >= 0 && indices[tmp] == indices[right] {
			tmp--
		}

		blockSize := right - tmp
		freeStart := findSuitablePlace(indices, blockSize)
		if freeStart != -1 && freeStart <= tmp {
			// Move block to free space
			for i := 0; i < blockSize; i++ {
				indices[freeStart+i] = indices[right-i]
				indices[right-i] = -1
			}
		}

		// Update `right` to the end of the next block
		right = tmp
		for right >= 0 && indices[right] == -1 {
			right--
		}
	}

	return d.computeChecksum(indices)
}

func findSuitablePlace(indices []int, requiredSpace int) int {
	left := 0
	n := len(indices)

	for left < n {
		// Find the start of a free block
		for left < n && indices[left] != -1 {
			left++
		}
		// Measure the size of the free block
		freeBlockStart := left
		for left < n && indices[left] == -1 {
			left++
		}
		if left-freeBlockStart >= requiredSpace {
			return freeBlockStart
		}
	}

	return -1
}
