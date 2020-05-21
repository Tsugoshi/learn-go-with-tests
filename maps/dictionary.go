package maps

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("Could not find the word")
	ErrWordExist        = DictionaryErr("Word already exists")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	word, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return word, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
