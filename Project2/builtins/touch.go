package builtins

import (
    "fmt"
    "os"
)

func TouchFile(args ...string) error {
    switch len(args) {
    case 0: // no file path provided
        return fmt.Errorf("%w: expected one argument (file)", ErrInvalidArgCount)
    case 1: // create file at provided path
        _, err := os.OpenFile(args[0], os.O_RDONLY|os.O_CREATE, 0666)
        return err
    default: // more than one argument provided
        return fmt.Errorf("%w: expected one argument (file)", ErrInvalidArgCount)
    }
}