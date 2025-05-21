package utils

import (
	"testing"
)


func TestGetCapital (t *testing.T) {
	tests := []struct {
		name     string 
		province     string
		expectedCapital string
	}{
		{"real province", "Chaco", "Resistencia"},
		{"case and accent marks", "sÁnTá fÉ", "Santa Fe"},
		{"another word", "Chiavenato", ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receivedCapital, err := GetCapital(tt.province)
			
			if err != nil {
				t.Errorf("GetCapital(%s) = %s, %v; want %s, nil", tt.province, receivedCapital, err, tt.expectedCapital)
			}
			if receivedCapital != tt.expectedCapital {
				t.Errorf("GetCapital(%s) = %s; want %s", tt.province, receivedCapital, tt.expectedCapital)
			}
 		})
	}
}

func TestRemoveAccentMarks (t *testing.T) {
	tests := []struct {
		name     string
		entryString     string
		expectedString string
	}{
		{"string with accent marks", "ríáchúéló", "riachuelo"},
		{"string without accent marks", "Totalidad", "Totalidad"},
		{"string with other marks", "pingüino", "pingüino"},
	}
	
	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receivedString := RemoveAccentMarks(tt.entryString)
			
			if receivedString != tt.entryString {
				t.Errorf("RemoveActionMarks(%s) = %s; want %s", tt.entryString, receivedString, tt.expectedString)
			}
		})
	}
} 