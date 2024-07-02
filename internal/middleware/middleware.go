package middleware

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaginationStruc struct {
	PerPage int64
	Page    int64
}

// This is a struct for filter
type FilterStruc struct {
	Property    string //"int"/"string"
	StringValue string
	IntValue    *IntValuesStruc
}

type IntValuesStruc struct {
	Gt  *int64
	Gte *int64
	Lt  *int64
	Lte *int64
	Eq  *int64
}

func checkProperty(str string) bool {
	return slices.Contains([]string{"type_1", "type_2", "gen", "att", "def", "spatt", "spdef", "spd", "total"}, str)
}

func checkOperator(str string) bool {
	return slices.Contains([]string{"gt", "gte", "lt", "lte", "eq"}, str)
}

func parseQuery(inputStr string) (string, string, error) {
	property := ""
	operator := ""
	regex, err := regexp.Compile(`^(type_1|type_2)=(\w+)|(att|def|spatt|spdef|spd|total)(?:\[(gt|gte|lt|lte|eq)\])?=(\d+)$`)
	if err != nil {
		if !checkProperty(inputStr) {
			return property, operator, err // case3
		}

		return property, operator, nil // case1

	}

	checkValue := regex.ReplaceAllString(inputStr, "$1")

	if checkValue != "" {
		if checkValue == inputStr {
			return "", "", nil
		}
		property = regex.ReplaceAllString(inputStr, "$1")
		operator = regex.ReplaceAllString(inputStr, "$2")
	} else {
		property = regex.ReplaceAllString(inputStr, "$3")
		operator = regex.ReplaceAllString(inputStr, "$4")
	}

	return property, operator, nil // case 1 and 3???

}

func FilterMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			filterMap := map[string]FilterStruc{}
			queryParams := c.QueryParams()
			fmt.Println(queryParams)
			// key = type_1 value = fire
			for key, value := range queryParams {
				property, operator, err := parseQuery(key)

				// testing value
				fmt.Println("property:", property, " operator:", operator)
				// testing value

				if err != nil {
					continue
				}
				if operator == "" {
					if slices.Contains([]string{"type_1", "type_2"}, property) {
						filterMap[property] = FilterStruc{
							Property:    "string",
							StringValue: value[0],
							IntValue:    nil,
						}
					} else {
						// for atk=xx
						valueInt, err := strconv.ParseInt(value[0], 10, 64)
						fmt.Println("else statement ", valueInt)
						if err != nil {
							continue
						}
						if slices.Contains([]string{"gen", "att", "def", "spatt", "spdef", "spd", "total"}, property) {
							filter := FilterStruc{
								Property:    "integer",
								StringValue: "",
								IntValue:    &IntValuesStruc{},
							}
							filterMap[property] = filter
							// switch oper

						}
					}
				} else {
					valueInt, err := strconv.ParseInt(value[0], 10, 64)
					if err != nil {
						continue
					}
					if existingFilter, ok := filterMap[property]; ok {

						switch operator {
						case "gt":
							existingFilter.IntValue.Gt = &valueInt
						case "gte":
							existingFilter.IntValue.Gte = &valueInt
						case "lt":
							existingFilter.IntValue.Lt = &valueInt
						case "lte":
							existingFilter.IntValue.Lte = &valueInt
						case "eq":
							existingFilter.IntValue.Eq = &valueInt
						}

						filterMap[property] = existingFilter
					} else {

						filter := FilterStruc{
							Property: "int",
							IntValue: &IntValuesStruc{},
						}

						switch operator {
						case "gt":
							filter.IntValue.Gt = &valueInt
						case "gte":
							filter.IntValue.Gte = &valueInt
						case "lt":
							filter.IntValue.Lt = &valueInt
						case "lte":
							filter.IntValue.Lte = &valueInt
						case "eq":
							filter.IntValue.Eq = &valueInt
						}

						filterMap[property] = filter

					}
				}

			}
			c.Set("filter", filterMap)
			return next(c)
		}
	}
}

func PaginationMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pagination := &PaginationStruc{
				PerPage: 10,
				Page:    1,
			}

			per_page := c.QueryParam("per_page")
			page := c.QueryParam("page")

			if per_page != "" {
				if per_page_int, err := strconv.ParseInt(per_page, 10, 64); err == nil {
					pagination.PerPage = per_page_int
				}
			}

			if page != "" {
				if page_int, err := strconv.ParseInt(page, 10, 64); err == nil {
					pagination.Page = page_int
				}
			}

			c.Set("pagination", pagination)

			return next(c)
		}
	}
}
