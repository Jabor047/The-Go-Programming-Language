package main

import (
	"bufio"
	"strings"
	"encoding/json"
	"fmt"
	"github.com/mozillazg/go-slugify"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const APIKey = "3361cedd" //Enter your APIkey here
const APIURL = "http://www.omdbapi.com/?apikey=" + APIKey + "&"

type Movie struct {
	Title string
	Year string
	Runtime string `json:"length_in_minutes"`
	Director string
	Actors string
	Genre string
	Plot string
	Poster string
}

func (m Movie) posterFilename()string{
	ext := filepath.Ext(m.Poster)
	title := slugify.Slugify(m.Title)
	return fmt.Sprintf("%s_(%s)%s", title, m.Year, ext)
}

func getMovie(title string)(movie Movie, err error){
	OmdbURL := fmt.Sprintf("%st=%s", APIURL, url.QueryEscape(title))
	resp, err := http.Get(OmdbURL)
	if err != nil{
		return 
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, OmdbURL)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return
	}
	return
}

func (m Movie) writePoster() error {
	posterURL := m.Poster
	resp, err := http.Get(posterURL) 
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, posterURL)
		return err
	}
	file, err := os.Create(m.posterFilename())
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main(){
	// if len(os.Args) != 2 {
	// 	fmt.Fprintln(os.Stderr, "Usage: Poster Movie_title")
	// 	os.Exit(1)
	// }
	title := os.Args[1:]
	fullTitle := strings.Join(title, " ")
	movie, err := getMovie(fullTitle)
	if err != nil {
		log.Fatal(err)
	}
	if Zero := new(Movie); movie == *Zero {
		fmt.Fprintf(os.Stderr, "No results for '%s' \n", title)
		os.Exit(2)
	}
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
			   movie.Title, movie.Year, movie.Runtime, movie.Director,
			   movie.Actors, movie.Genre, movie.Plot)
	fmt.Println("The movie's poster is saved as: \t", movie.posterFilename())
	err = movie.writePoster()
	if err != nil{
		log.Fatal(err)
	}

}