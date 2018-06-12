package main

import (
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func testStruct() {
	movies := []Movie{
		{Title: "Casblanca", Year: 1942, Color: false,
			Actors: []string{"A", "B"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"c"}},
	}

	// marshal
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("failed to marshal: %s\n", err)
	}
	log.Printf("[]Movie: marshal: %s\n", data)

	indentData, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal indent: %s\n", err)
	}
	log.Printf("[]Movie: marshal indent:\n%s\n", indentData)

	// unmarshal
	var titles []struct{ Title string }
	if err := json.Unmarshal(indentData, &titles); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}
	log.Printf("titles[]: unmarshal: %v\n", titles)

	var actors []struct{ Actors []string }
	if err := json.Unmarshal(indentData, &actors); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}
	log.Printf("actors[]: unmarshal: %v\n", actors)

	var rMovies []Movie

	if err := json.Unmarshal(data, &rMovies); err != nil {
		log.Fatalf("failed to unmarshl movies: %v", err)
	}

	log.Printf("rMovies[]: unmarshal: %v\n", rMovies)
}

func testMap() {
	m := make(map[int]string)

	m[97] = "a"
	m[98] = "b"
	m[99] = "c"

	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("failed to marshal: %s\n", err)
	}

	log.Printf("map[int]string: marshal: %s\n", string(data))

	rM := make(map[int]string)
	err = json.Unmarshal(data, &rM)
	log.Printf("unmarshal: %v\n", rM)
}

func main() {
	// testMap()
	testStruct()
}
