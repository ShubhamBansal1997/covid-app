package utils

import "errors"

func GetStateName(code string) (string, error) {
	stateCodeMap := map[string]string{
		"AN": "Andaman and Nicobar Islands",
		"AP": "Andhra Pradesh",
		"AR": "Arunachal Pradesh",
		"AS": "Assam",
		"BR": "Bihar",
		"CH": "Chandigarh",
		"CT": "Chhattisgarh",
		"DL": "Delhi",
		"DN": "Dadra and Nagar Haveli and Daman and Diu",
		"GA": "Goa",
		"GJ": "Gujarat",
		"HP": "Himachal Pradesh",
		"HR": "Haryana",
		"JH": "Jharkhand",
		"JK": "Jammu and Kashmir",
		"KA": "Karnataka",
		"KL": "Kerala",
		"LA": "Ladakh",
		"LD": "Lakshadweep",
		"MH": "Maharashtra",
		"ML": "Meghalaya",
		"MN": "Manipur",
		"MP": "Madhya Pradesh",
		"MZ": "Mizoram",
		"NL": "Nagaland",
		"OR": "Odisha",
		"PB": "Punjab",
		"PY": "Puducherry",
		"RJ": "Rajasthan",
		"SK": "Sikkim",
		"TG": "Telangana",
		"TN": "Tamil Nadu",
		"TR": "Tripura",
		"UP": "Uttar Pradesh",
		"UT": "Uttarakhand",
		"WB": "West Bengal",
	}
	state, ok := stateCodeMap[code]
	if !ok {
		return "", errors.New("invalid state code")
	}
	return state, nil
}
