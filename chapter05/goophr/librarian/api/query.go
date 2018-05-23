package api

import "sort"

type docResult struct {
	DocID   string   `json:"doc_id"`
	Score   int      `json:"doc_score"`
	Indices tIndices `json:"token_indices"`
}

type result struct {
	Count int         `json:"count"`
	Data  []docResult `json:"data"`
}

// getResult returns unsorted search results & a map of documents containing tokens.
func getResult(out chan tcMsg, count int) tCatalog {
	tc := tCatalog{}

	for i := 0; i < count; i++ {
		dc := <-out
		tc[dc.Token] = dc.DC
	}
	close(out)

	return tc
}

func getFScores(docIDScore map[string]int) (map[int][]string, []int) {
	// fScore maps frequency score to set of documents.
	fScore := map[int][]string{}

	fSorted := []int{}

	for dID, score := range docIDScore {
		fs := fScore[score]
		fScore[score] = append(fs, dID)
		fSorted = append(fSorted, score)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(fSorted)))

	return fScore, fSorted
}

func getDocMaps(tc tCatalog) (map[string]int, map[string]tIndices) {
	// docIDScore maps DocIDs to occurences of all tokens.
	// key: DocID
	// val: Sum of all occurences of tokens so far.
	docIDScore := map[string]int{}
	docIndices := map[string]tIndices{}

	// for each token's catalog
	for _, dc := range tc {
		// for each document registered under the token
		for dID, doc := range dc {
			// add to docID score
			var tokIndices tIndices
			for _, tList := range doc.Indices {
				tokIndices = append(tokIndices, tList...)
			}
			docIDScore[dID] += doc.Count

			dti := docIndices[dID]
			docIndices[dID] = append(dti, tokIndices...)
		}
	}

	return docIDScore, docIndices
}

func sortResults(tc tCatalog) []docResult {
	docIDScore, docIndices := getDocMaps(tc)
	fScore, fSorted := getFScores(docIDScore)

	results := []docResult{}
	addedDocs := map[string]bool{}

	for _, score := range fSorted {
		for _, docID := range fScore[score] {
			if _, exists := addedDocs[docID]; exists {
				continue
			}
			results = append(results, docResult{
				DocID:   docID,
				Score:   score,
				Indices: docIndices[docID],
			})
			addedDocs[docID] = false
		}
	}
	return results
}

// getSearchResults returns a list of documents.
// They are listed in decending order of occurences.
func getSearchResults(sts []string) []docResult {
	callback := make(chan tcMsg)

	for _, st := range sts {
		go func(term string) {
			tcGet <- tcCallback{
				Token: term,
				Ch:    callback,
			}
		}(st)
	}
	cts := getResult(callback, len(sts))
	results := sortResults(cts)
	return results
}
