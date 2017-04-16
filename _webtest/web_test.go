package webtest

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jonpchin/gochess/gostuff"
	"github.com/sclevine/agouti"
)

var db *sql.DB

//localhost testing
func TestLoginDev(t *testing.T) {

	var dbFile string
	if gostuff.IsEnvironmentTravis() {
		dbFile = "../_travis/data/dbtravis.txt"
	} else {
		dbFile = "../secret/checkdb.txt"
	}

	// make sure MySQL connection is alive before proceeding
	if gostuff.CheckDBConnection(dbFile) == false {
		t.Fatal("Failed to connect to MySQL in Travis CI in localhost")
	}

	var err error
	dbString, _ := gostuff.ReadFile(dbFile)
	db, err = sql.Open("mysql", dbString)
	//defer db.Close()

	if err != nil {
		t.Fatal("Can't open MySQL")
	}

	gostuff.SetDb(db)

	//if database ping fails here that means connection is alive but database is missing
	if db.Ping() != nil {
		t.Fatal("Can't ping MySQL")
	}

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		t.Fatal("Failed to start Chrome Driver:", err)
	}
	page1, err := driver.NewPage(agouti.Browser("Chrome"))
	if err != nil {
		t.Fatal("Failed to open page:", err)
	}

	if err := page1.Navigate("https://localhost:443"); err != nil {
		t.Fatal("Failed to navigate index at localhost:", err)
	}

	if err := page1.Navigate("https://localhost/login"); err != nil {
		t.Fatal("Failed to navigate login at localhost:", err)
	}

	loginURL, err := page1.URL()
	if err != nil {
		t.Fatal("Failed to get page URL:", err)
	}

	expectedLoginURL := "https://localhost/login"
	if loginURL != expectedLoginURL {
		t.Fatal("Expected URL to be", expectedLoginURL, "but got", loginURL)
	}

	time.Sleep(time.Second)
	user1 := "can"

	// Player should have zero games on Travis
	if gostuff.IsEnvironmentTravis() {
		storage := gostuff.GetGames(user1)
		if len(storage) > 0 {
			t.Fatal("There are more then zero games for ", user1, " when there shouldn't be")
		}
	}

	err = page1.FindByID("user").Fill(user1)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}
	pass := readPass(user1)
	err = page1.FindByID("password").Fill(pass)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}

	err = page1.FindByID("login").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}

	time.Sleep(time.Second)
	if err := page1.Navigate("https://localhost/server/lobby"); err != nil {
		t.Fatal("Failed to navigate lobby at localhost:", err)
	}

	user2 := "ben"

	success := testResumeGame(user1, user2, page1)

	if success != "false" {
		t.Fatal("resumeGame Ajax failed:", success)
	}

	time.Sleep(time.Second)

	err = page1.FindByID("sendSeek").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}

	// start second browser
	page2, err := driver.NewPage(agouti.Browser("Chrome"))
	if err != nil {
		t.Fatal("Failed to open page:", err)
	}

	if err := page2.Navigate("https://localhost:443"); err != nil {
		t.Fatal("Failed to navigate index at localhost:", err)
	}

	if err := page2.Navigate("https://localhost/login"); err != nil {
		t.Fatal("Failed to navigate login at localhost:", err)
	}

	err = page2.FindByID("user").Fill(user2)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}
	pass = readPass(user2)
	err = page2.FindByID("password").Fill(pass)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}

	err = page2.FindByID("login").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}
	time.Sleep(time.Second * 1)

	if err := page2.Navigate("https://localhost/server/lobby"); err != nil {
		t.Fatal("Failed to navigate lobby at localhost:", err)
	}

	err = page2.FindByID("sendSeek").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}
	time.Sleep(time.Second)
	var whitePlayer string
	page2.RunScript("return WhiteSide;", map[string]interface{}{}, &whitePlayer)

	if user1 == whitePlayer {
		executeGame(page1, page2, user1, user2, t)
	} else if user2 == whitePlayer {
		executeGame(page2, page1, user2, user1, t)

	} else {
		// then navigate to chess page and try to terminate any possible games that are left over
		if err := page2.Navigate("https://localhost/chess/memberChess"); err != nil {
			t.Fatal("Failed to navigate login to chess page:", err)
		}
		err = page2.FindByID("abortButton").Click()
		if err != nil {
			t.Fatal("Couldn't find abort button  user 2:", err)
		}
		t.Fatal("No user matched as whitePlayer")
	}

	page1.Destroy()
	page2.Destroy()
	time.Sleep(time.Second)
	if err := driver.Stop(); err != nil {
		t.Error("Failed to close pages and stop WebDriver:", err)
	}
}

func testResumeGame(user1 string, user2 string, page1 *agouti.Page) string {

	success := "true"

	err := page1.RunScript(`
		var result = "true";
		$.ajax({
			url: '../resumeGame',
			type: 'post',
			dataType: 'html',
			data : { 'id': "1", 'white': "`+user1+`", 'black': "`+user2+`"},
			async: false,
			success:function(data) {
				result = data; 
			}
		});
		return result;`, map[string]interface{}{}, &success)
	if err != nil {
		fmt.Println("Can't run ajax script", err)
	}
	return success
}

func executeGame(page1 *agouti.Page, page2 *agouti.Page, user1 string, user2 string, t *testing.T) {

	var jsResult string

	page1.RunScript("sendMove('e2', 'e4');", map[string]interface{}{}, &jsResult)
	page2.RunScript("sendMove('c7', 'c5');", map[string]interface{}{}, &jsResult)
	page1.RunScript("sendMove('g1', 'f3');", map[string]interface{}{}, &jsResult)
	page1.RunScript("return board.fen();", map[string]interface{}{}, &jsResult)

	// check to make sure the position is what it should be
	if jsResult != "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R" {
		t.Error("board does not match user1", jsResult)
	}

	// now try to resign the game
	err := page1.FindByID("resignButton").Click()
	if err != nil {
		t.Fatal("Couldn't resign:", user1, err)
	}
	err = page1.ConfirmPopup()
	if err != nil {
		t.Fatal("Couldn't confirm resign popup:", user1, err)
	}

	// Gives time to save to database
	time.Sleep(2 * time.Second)
	// Player should have one games on Travis
	if gostuff.IsEnvironmentTravis() {
		storage := gostuff.GetGames(user1)
		if len(storage) != 1 || storage[0].Status != "White Resigned" ||
			storage[0].Rated != "Yes" || storage[0].TimeControl != 5 {
			if len(storage) > 0 {
				t.Fatal("GetGames does not match the expected output for ", user1, storage[0].Status,
					storage[0].Rated, storage[0].TimeControl)
			} else {
				t.Fatal("GetGames does not match the expected output for ", user1, len(storage))
			}
		}
	}

	history, result, err := gostuff.GetRatingHistory(user1, "blitz")
	if result == false || history == "" {
		t.Fatal("Could not GetRatingHistory for blitz", user1, err)
	}
	history, result, err = gostuff.GetRatingHistory(user2, "blitz")
	if result == false || history == "" {
		t.Fatal("Could not GetRatingHistory 2 for blitz", user2, err)
	}

	err = page1.FindByID("rematchButton").Click()
	if err != nil {
		t.Fatal("Couldn't find rematch button :", user1, err)
	}
	err = page2.FindByID("rematchButton").Click()
	if err != nil {
		t.Fatal("Couldn't find rematch button :", user1, err)
	}
	err = page1.FindByID("abortButton").Click()
	if err != nil {
		t.Fatal("Couldn't find abort button :", user1, err)
	}
	err = page2.FindByID("rematchButton").Click()
	if err != nil {
		t.Fatal("Couldn't find rematch button  :", user1, err)
	}
	err = page1.FindByID("rematchButton").Click()
	if err != nil {
		t.Fatal("Couldn't find rematch button :", user2, err)
	}
	err = page1.FindByID("drawButton").Click()
	if err != nil {
		t.Fatal("Couldn't find draw button  :", user1, err)
	}
	err = page2.FindByID("drawButton").Click()
	if err != nil {
		t.Fatal("Couldn't find draw button :", user2, err)
	}
	// TODO: Check if game really ended and check if the other player really won
	// Still need to test abort failure, abort sucess, draw, and checkmate

	time.Sleep(2 * time.Second)

	// Player should have two games on Travis
	if gostuff.IsEnvironmentTravis() {
		storage := gostuff.GetGames(user1)
		if len(storage) != 2 || storage[1].Status != "Agreed Draw" ||
			storage[1].Rated != "Yes" || storage[1].TimeControl != 5 {
			if len(storage) > 1 {
				t.Fatal("GetGames 2 does not match the expected output for ", user1, storage[1].Status,
					storage[1].Rated, storage[1].TimeControl)
			} else {
				t.Fatal("GetGames 2 does not match the expected output for ", user1)
			}
		}
	}

	page1.RunScript("sendMove('e2', 'e4');", map[string]interface{}{}, &jsResult)
	page2.RunScript("sendMove('c7', 'c6');", map[string]interface{}{}, &jsResult)
	gostuff.Cleanup()

	if err := page1.Navigate("https://localhost/server/lobby"); err != nil {
		t.Fatal("Failed to navigate lobby at localhost:", err)
	}
	if err := page2.Navigate("https://localhost/saved?user=" + user2); err != nil {
		t.Fatal("Failed to navigate saved at localhost:", user2, err)
	}

	success := testResumeGame(user1, user2, page1)

	if success != "false" {
		t.Fatal("resumeGame Ajax failed:", user1, user2, success)
	}
	page1.RunScript("sendMove('e4', 'e5');", map[string]interface{}{}, &jsResult)
	page2.RunScript("sendMove('c6', 'c5');", map[string]interface{}{}, &jsResult)

	if jsResult != "rnbqkbnr/pp1ppppp/8/2p1P3/8/8/PPPP1PPP/RNBQKBNR" {
		t.Error("board does not match", user1, user2, jsResult)
	}

	err = page1.FindByID("resignButton").Click()
	if err != nil {
		t.Fatal("Couldn't resign:", user1, err)
	}
	err = page1.ConfirmPopup()
	if err != nil {
		t.Fatal("Couldn't confirm resign popup:", user1, err)
	}
}

/*
func TestLoginProduction(t *testing.T) {

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		t.Fatal("Failed to start Chrome Driver:", err)
	}
	page1, err := driver.NewPage(agouti.Browser("Chrome"))
	if err != nil {
		t.Fatal("Failed to open page:", err)
	}

	if err := page1.Navigate("https://goplaychess.com:443"); err != nil {
		t.Fatal("Failed to navigate index:", err)
	}

	if err := page1.Navigate("https://goplaychess.com/login"); err != nil {
		t.Fatal("Failed to navigate login:", err)
	}

	loginURL, err := page1.URL()
	if err != nil {
		t.Fatal("Failed to get page URL:", err)
	}

	expectedLoginURL := "https://goplaychess.com/login"
	if loginURL != expectedLoginURL {
		t.Fatal("Expected URL to be", expectedLoginURL, "but got", loginURL)
	}

	user1 := "foo"

	err = page1.FindByID("user").Fill(user1)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}
	pass := readPass(user1)
	err = page1.FindByID("password").Fill(pass)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}

	err = page1.FindByID("login").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}
	time.Sleep(time.Second)
	if err := page1.Navigate("https://goplaychess.com/server/lobby"); err != nil {
		t.Fatal("Failed to navigate lobby:", err)
	}
	time.Sleep(time.Second)
	err = page1.FindByID("sendSeek").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}

	// start second browser
	page2, err := driver.NewPage(agouti.Browser("Chrome"))
	if err != nil {
		t.Fatal("Failed to open page:", err)
	}

	if err := page2.Navigate("https://goplaychess.com:443"); err != nil {
		t.Fatal("Failed to navigate index at localhost:", err)
	}

	if err := page2.Navigate("https://goplaychess.com/login"); err != nil {
		t.Fatal("Failed to navigate login at localhost:", err)
	}

	user2 := "Carl"
	err = page2.FindByID("user").Fill(user2)
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}

	err = page2.FindByID("password").Fill(readPass(user2))
	if err != nil {
		t.Fatal("Couldn't fill login info:", err)
	}

	err = page2.FindByID("login").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}
	time.Sleep(time.Second)

	if err := page2.Navigate("https://goplaychess.com/runtest"); err != nil {
		t.Fatal("Failed to navigate to runtest at goplaychess.com:", err)
	}

	time.Sleep(time.Second * 3)
	htmlContent, err := page2.HTML()

	if err != nil {
		t.Fatal("Failed to get html of runtest webpage", err)
	}

	if strings.Contains(htmlContent, "not-available.png") || strings.Contains(htmlContent, "loading.gif") {
		t.Fatal("AJAX unit tests failed on goplaychess.com", htmlContent)
	}

	if err := page2.Navigate("https://goplaychess.com/server/lobby"); err != nil {
		t.Fatal("Failed to navigate lobby at localhost:", err)
	}
	time.Sleep(time.Second)
	err = page2.FindByID("sendSeek").Click()
	if err != nil {
		t.Fatal("Couldn't submit:", err)
	}

	time.Sleep(2 * time.Second)
	var whitePlayer string
	page2.RunScript("return WhiteSide;", map[string]interface{}{}, &whitePlayer)
	var jsResult string
	time.Sleep(2 * time.Second)
	if user1 == whitePlayer {
		page1.RunScript("sendMove('e2', 'e4');", map[string]interface{}{}, &jsResult)
		page2.RunScript("sendMove('c7', 'c5');", map[string]interface{}{}, &jsResult)
		page1.RunScript("sendMove('g1', 'f3');", map[string]interface{}{}, &jsResult)
		time.Sleep(time.Second)
		page1.RunScript("return board.fen();", map[string]interface{}{}, &jsResult)

		// check to make sure the position is what it should be
		if jsResult != "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R" {
			t.Fatal("board does not match user1", jsResult)
		}

		// now try to resign the game
		err = page1.FindByID("resignButton").Click()
		if err != nil {
			t.Fatal("Couldn't resign user1:", err)
		}
		err = page1.ConfirmPopup()
		if err != nil {
			t.Fatal("Couldn't confirm resign popup user1:", err)
		}
		err = page1.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 1:", err)
		}
		err = page2.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 1:", err)
		}
		err = page1.FindByID("abortButton").Click()
		if err != nil {
			t.Fatal("Couldn't find abort button  user 1:", err)
		}
		err = page2.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 1:", err)
		}
		err = page1.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 2:", err)
		}
		err = page1.FindByID("drawButton").Click()
		if err != nil {
			t.Fatal("Couldn't find draw button  user 1:", err)
		}
		err = page2.FindByID("drawButton").Click()
		if err != nil {
			t.Fatal("Couldn't find draw button  user 2:", err)
		}
		// TODO: Check if game really ended and check if the other player really won
		// Still need to test abort failure, abort sucess, draw, and checkmate

	} else if user2 == whitePlayer {
		page2.RunScript("sendMove('e2', 'e4');", map[string]interface{}{}, &jsResult)
		page1.RunScript("sendMove('c7', 'c5');", map[string]interface{}{}, &jsResult)
		page2.RunScript("sendMove('g1', 'f3');", map[string]interface{}{}, &jsResult)
		time.Sleep(time.Second)
		page2.RunScript("return board.fen();", map[string]interface{}{}, &jsResult)

		if jsResult != "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R" {
			t.Error("board does not match user2")
		}
		err = page2.FindByID("resignButton").Click()
		if err != nil {
			t.Fatal("Couldn't resign user2:", err)
		}
		err = page2.ConfirmPopup()
		if err != nil {
			t.Fatal("Couldn't confirm resign popup user2:", err)
		}
		err = page2.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 2:", err)
		}
		err = page1.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 1:", err)
		}
		err = page1.FindByID("abortButton").Click()
		if err != nil {
			t.Fatal("Couldn't find abort button  user 1:", err)
		}
		err = page1.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 1:", err)
		}
		err = page2.FindByID("rematchButton").Click()
		if err != nil {
			t.Fatal("Couldn't find rematch button  user 2:", err)
		}
		err = page1.FindByID("drawButton").Click()
		if err != nil {
			t.Fatal("Couldn't find draw button  user 1:", err)
		}
		err = page2.FindByID("drawButton").Click()
		if err != nil {
			t.Fatal("Couldn't find draw button  user 2:", err)
		}

	} else {
		// then navigate to chess page and try to terminate any possible games that are left over
		if err := page2.Navigate("https://goplaychess.com/chess/memberChess"); err != nil {
			t.Fatal("Failed to navigate login to chess page:", err)
		}
		err = page2.FindByID("abortButton").Click()
		if err != nil {
			t.Fatal("Couldn't find abort button  user 2:", err)
		}
		t.Fatal("No user matched as whitePlayer", whitePlayer)
	}
	page1.Destroy()
	page2.Destroy()
	time.Sleep(time.Second)
	if err := driver.Stop(); err != nil {
		t.Error("Failed to close pages and stop WebDriver:", err)
	}
}
*/
// returns pass of user's account
func readPass(user string) string {
	config, err := os.Open("data/" + user + ".txt")
	defer config.Close()
	if err != nil {
		log.Println("web_Test.go readAccount 1 ", err)
	}
	scanner := bufio.NewScanner(config)
	scanner.Scan()

	return scanner.Text()
}
