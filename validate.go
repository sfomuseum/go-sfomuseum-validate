package validate

import (
	"context"
	"fmt"
)

type Options struct {
	ValidatePlacetype bool
	ValidateSFOLevel  bool
}

func DefaultValidateOptions() *Options {

	return &Options{
		ValidatePlacetype: true,
		ValidateSFOLevel:  true,
	}
}

func Validate(ctx context.Context, body []byte) error {

	opts := DefaultValidateOptions()
	return ValidateWithOptions(ctx, body, opts)
}

func ValidateWithOptions(ctx context.Context, body []byte, options *Options) error {

	if options.ValidatePlacetype {

		err := ValidatePlacetype(body)

		if err != nil {
			return fmt.Errorf("Failed to validate sfomuseum:placetype, %w", err)
		}
	}

	if options.ValidateSFOLevel {

		err := ValidateSFOLevel(body)

		if err != nil {
			return fmt.Errorf("Failed to validate sfo:level, %w", err)
		}
	}

	return nil
}
