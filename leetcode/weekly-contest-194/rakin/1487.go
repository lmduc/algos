package main

import "strconv"

func getFolderNames(names []string) []string {
	tracker := map[string]int{}
	res := []string{}
	for _, name := range names {
		count := tracker[name]
		pick := name
		for {
			if count > 0 {
				pick = name + "(" + strconv.Itoa(count) + ")"
			}

			if tracker[pick] == 0 {
				tracker[pick] = 1
				tracker[name] = count + 1
				break
			}

			count++
		}
		res = append(res, pick)
	}

	return res
