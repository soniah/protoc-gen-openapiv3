package generator

import (
	"fmt"
	"io"

	"github.com/pb33f/libopenapi/datamodel/high/v3"
)

// MergeDocuments merges multiple *v3.Document into one.
// It combines Paths, Components.Schemas, and Info from the first document.
func MergeDocuments(docs []*v3.Document) (*v3.Document, error) {
	if len(docs) == 0 {
		return nil, nil
	}

	merged := docs[0]
	bs, err := merged.Render()
	if err != nil {
		return nil, err
	}
	fmt.Fprint(io.Discard, bs)

	for i := 1; i < len(docs); i++ {
		doc := docs[i]

		bs, err := doc.Render()
		if err != nil {
			return nil, err
		}
		fmt.Fprint(io.Discard, bs)
		//for k, v := range doc.Paths.PathItems.FromOldest() {
		//	k
		//}
		//for k, v := range doc.Paths.PathItems.FromOldest() {
		//	fmt.Fprint(io.Discard, k, v)
		//}
		//bs, err := doc.Render()
		//if err != nil {
		//	return nil, nil
		//}
		//fmt.Fprint(io.Discard, bs)
	}
	//for _, doc := range docs {
	//	// Merge Paths
	//	if doc.Paths != nil && doc.Paths.PathItems != nil {
	//		for path, item := range doc.Paths.PathItems {
	//			merged.Paths.PathItems[path] = item
	//		}
	//	}
	// Merge Schemas
	//if doc.Components != nil && doc.Components.Schemas != nil {
	//	for name, schema := range doc.Components.Schemas {
	//		merged.Components.Schemas[name] = schema
	//	}
	//}
	//Merge Tags
	//if doc.Tags != nil {
	//	merged.Tags = append(merged.Tags, doc.Tags...)
	//}
	//}
	return merged, nil
}
