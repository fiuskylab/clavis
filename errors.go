package clavis

import "fmt"

const (
	templateNotFound = "%s not found"
	templateMissing  = "Missing %s"
	templateExists   = "%s already exists"
	templateExpired  = "%s is expired"
)

type errat struct {
	message string
	isNil   bool
}

func NilErrat() errat {
	return errat{
		isNil: true,
	}
}

func ErratNotFound(key string) errat {
	return errat{
		isNil:   false,
		message: fmt.Sprintf(templateNotFound, key),
	}
}

func ErratUnknown(msg string) errat {
	return errat{
		isNil:   false,
		message: msg,
	}
}

func ErratMissing(key string) errat {
	return errat{
		isNil:   false,
		message: fmt.Sprintf(templateMissing, key),
	}
}

func ErratExists(key string) errat {
	return errat{
		isNil:   false,
		message: fmt.Sprintf(templateExists, key),
	}
}

func ErratExpired(key string) errat {
	return errat{
		isNil:   false,
		message: fmt.Sprintf(templateExpired, key),
	}
}

func (e *errat) Nil() bool {
	return e.isNil
}
