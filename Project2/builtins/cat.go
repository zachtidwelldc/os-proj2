package builtins

import (
    "fmt"
    "io"
    "os"
)

func CatFile(w io.Writer, args ...string) error {
    if len(args) < 1 {
        return fmt.Errorf("%w: expected at least one argument (file)", ErrInvalidArgCount)
    }

    for _, filename := range args {
        content, err := os.ReadFile(filename)
        if err != nil {
            return err
		}
        fmt.Fprintln(w, string(content))
    }

    return nil
}