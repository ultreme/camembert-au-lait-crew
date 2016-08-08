package calcdashboard

import "math/rand"

type Entry interface {
	GetMinimalFields() (*MinimalFields, error)
}

type Entries []Entry

type MinimalFields struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image-url"`
	Kind        string `json:"kind"`
}

func (e *Entries) GetMinimalFields() ([]*MinimalFields, error) {
	var results []*MinimalFields

	for _, entry := range *e {
		result, err := entry.GetMinimalFields()
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (e *Entries) append(entry Entry) {
	*e = append(*e, entry)
}

type ManualEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image-url"`
	Kind        string `json:"kind"`
}

func (m *ManualEntry) GetMinimalFields() (*MinimalFields, error) {
	return &MinimalFields{
		Title:       m.Title,
		Description: m.Description,
		ImageURL:    m.ImageURL,
		Kind:        m.Kind,
	}, nil
}

func NewManualEntry(kind, title, description, imageURL string) *ManualEntry {
	return &ManualEntry{
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		Kind:        kind,
	}
}

func (e *Entries) shuffle() {
	for i := range *e {
		j := rand.Intn(i + 1)
		(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
	}
}
