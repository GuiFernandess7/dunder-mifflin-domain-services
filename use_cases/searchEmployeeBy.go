package use_cases

import (
	"context"
	"database/sql"
	"errors"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func GetEmployeesBy(min int32, max int32, sex *string, db_repository *database.Queries, ctx context.Context) ([]database.Employee, error) {
	if min == 0 || max == 0 {
		return nil, errors.New("min and max salary must be provided")
	}

	params := database.FilterEmployeeBySalaryAndSexParams{
		Min: sql.NullInt32{
			Int32: min,
			Valid: true,
		},
		Max: sql.NullInt32{
			Int32: max,
			Valid: true,
		},
	}

	if sex != nil {
		params.Sex = *sex
		result, err := db_repository.FilterEmployeeBySalaryAndSex(ctx, params)

		if err != nil {
			log.Fatalf("failed to filter employee by salary: %v", err)
			return nil, err
		}

		return result, nil

	} else {
		params, _ := buildSalaryFilterParams(min, max)
		result, err := db_repository.FilterEmployeeBySalary(ctx, params)

		if err != nil {
			log.Fatalf("failed to filter employee by salary and sex: %v", err)
			return nil, err
		}
		return result, nil
	}
}

func buildSalaryFilterParams(min int32, max int32)(database.FilterEmployeeBySalaryParams, error) {
	selectedRange := database.FilterEmployeeBySalaryParams{}
	if min == 0 || max == 0 {
		return selectedRange, errors.New("min and max salary must be provided")
	}

	selectedRange.Max = sql.NullInt32{
		Int32: max,
		Valid: true,
	}

	selectedRange.Min = sql.NullInt32{
		Int32: min,
		Valid: true,
	}

	return selectedRange, nil
}

