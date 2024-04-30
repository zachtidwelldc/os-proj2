package builtins

import (
    //"errors"
    "fmt"
    "os"
)

func MakeDirectory(args ...string) error {
    switch len(args) {
    case 0: // no directory path provided
        return fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgCount)
    case 1: // create directory at provided path
        return os.Mkdir(args[0], 0755)
    default: // more than one argument provided
        return fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgCount)
    }
}