package validation

import (
    "errors"
    "fmt"
)

type Rule func(key string, value interface{}) error

type Rules []Rule

type Validator struct {
    rules Rules
}

func (v *Validator) Add(rule Rule) {
    v.rules = append(v.rules, rule)
}

func (v *Validator) Validate(data map[string]interface{}) []error {
    var errors []error
    for _, rule := range v.rules {
        for key, value := range data {
            if err := rule(key, value); err != nil {
                errors = append(errors, err)
            }
        }
    }
    return errors
}

func ValidateLength(maxLength int) Rule {
    return func(key string, value interface{}) error {
        str, ok := value.(string)
        if !ok {
            return fmt.Errorf("%s is not a string", key)
        }
        if len(str) > maxLength {
            return fmt.Errorf("%s must be less than or equal to %d characters", key, maxLength)
        }
        return nil
    }
}

func ValidatePresence(key string, value interface{}) error {
    if value == "" {
        return fmt.Errorf("%s can't be blank", key)
    }
    return nil
}

func ValidateRange(min, max int) Rule {
    return func(key string, value interface{}) error {
        num, ok := value.(int)
        if !ok {
            return fmt.Errorf("%s is not a number", key)
        }
        if num < min || num > max {
            return fmt.Errorf("%s must be between %d and %d", key, min, max)
        }
        return nil
    }
}

func ValidateEmail(key string, value interface{}) error {
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("%s is not a string", key)
    }
    // simplified email validation for example purposes
    if str == "" || str[:1] == "@" || str[len(str)-1:] == "@" {
        return fmt.Errorf("%s is not a valid email address", key)
    }
    return nil
}