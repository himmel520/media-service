// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Ad) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Image.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "image",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Priority.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "priority",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AdPost) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Post)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "post",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        false,
			Max:           0,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.ImagesID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "images_id",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        false,
			Max:           0,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.ColorsID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "colors_id",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        false,
			Max:           0,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.TgID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tg_id",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    40,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Title)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    5,
			MinLengthSet: true,
			MaxLength:    150,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Description)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "description",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Priority.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "priority",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s AdPostPriority) Validate() error {
	switch s {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s AdPriority) Validate() error {
	switch s {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *AdPut) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Post.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    3,
					MinLengthSet: true,
					MaxLength:    100,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "post",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.ImagesID.Get(); ok {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "images_id",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.ColorsID.Get(); ok {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "colors_id",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.TgID.Get(); ok {
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tg_id",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Title.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    3,
					MinLengthSet: true,
					MaxLength:    40,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Description.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    5,
					MinLengthSet: true,
					MaxLength:    150,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "description",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Priority.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "priority",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s AdPutPriority) Validate() error {
	switch s {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s Ads) Validate() error {
	alias := ([]Ad)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AdsResp) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ColorPost) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Title)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^#[0-9A-Fa-f]{6}$"],
		}).Validate(string(s.Hex)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ColorPut) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Title.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    3,
					MinLengthSet: true,
					MaxLength:    100,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Hex.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^#[0-9A-Fa-f]{6}$"],
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *Image) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ImagePost) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Title)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ImagePostType) Validate() error {
	switch s {
	case "adv":
		return nil
	case "logo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *ImagePut) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Title.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    3,
					MinLengthSet: true,
					MaxLength:    100,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Type.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ImagePutType) Validate() error {
	switch s {
	case "adv":
		return nil
	case "logo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ImageType) Validate() error {
	switch s {
	case "adv":
		return nil
	case "logo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *ImagesResp) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s LogosResp) Validate() error {
	var failures []validate.FieldError
	for key, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  key,
				Error: err,
			})
		}
	}

	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *LogosRespItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s LogosRespItemType) Validate() error {
	switch s {
	case "logo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TgPost) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Title)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *TgPut) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Title.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    3,
					MinLengthSet: true,
					MaxLength:    100,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s V1AdsGetPriorityItem) Validate() error {
	switch s {
	case 1:
		return nil
	case 2:
		return nil
	case 3:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
