package mud

import (
	"bufio"
	crypto "crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"
)

// Calling getRandomInt(100) will return a random number 0 to 100 inclusive
// If max is less then zero, then zero will be returned
func getRandomInt(max int) int {
	if max >= 0 {
		return rand.Intn(max)
	}
	return 0
}

func getRandomIntRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Generates a randomly secure integer (int64) from 0 to maxInclusive
// Takes in an int64
// Returns 0 for the integer if there was an error as well as the error
func secureRandomInt(max int64) (int64, error) {

	if max < 0 {
		return 0, errors.New("maxInclusive is less then zero")
	}
	maxInt := big.NewInt(max)

	result, err := crypto.Int(crypto.Reader, maxInt)
	if err != nil {
		return 0, err
	}

	return result.Int64(), nil
}

// Generates a randomly secure integer (int64) from minInclusive to maxInclusive
// Takes in an int64
// Returns 0 for the integer if there was an error as well as the error
func secureRandomIntRange(min, max int64) (int64, error) {

	if max < 0 {
		return 0, errors.New("maxInclusive is less then zero")
	}

	maxInt := big.NewInt(max - min)

	// Sets maxInt = maxInt + min
	maxInt.Add(maxInt, big.NewInt(min))
	result, err := crypto.Int(crypto.Reader, maxInt)

	if err != nil {
		return 0, err
	}
	return result.Int64(), nil
}

// Returns a radom direction. Returns north if there was an error
func getRandomDirection() Direction {
	result, err := secureRandomInt(3)
	if err != nil {
		fmt.Println("Can't get random integer for direction", err)
	}
	switch result {
	case 0:
		return NORTH
	case 1:
		return EAST
	case 2:
		return SOUTH
	case 3:
		return WEST
	default:
		fmt.Println("Invalid direction, this should be impossible")
	}
	return NORTH
}

func (floor *Floor) getRandomRoomOnFloor() Room {
	log := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	randNum, err := secureRandomInt(int64(len(floor.Rooms) - 1))
	if err != nil {
		log.Println(err)
	}
	return floor.Rooms[randNum]
}

// Selects a random tile on the wall of a room
func (floor *Floor) getRandomTileOnWall() Tile {
	log := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	room := floor.getRandomRoomOnFloor()
	randNum, err := secureRandomIntRange(0, int64(len(room.Tiles)-1))
	if err != nil {
		log.Println(err)
	}
	return room.Wall[randNum]
}

func getRandomItemFromPath(file *os.File) string {

	scanner := bufio.NewScanner(file)
	var counter int64
	counter = 0

	for scanner.Scan() {
		counter++
	}
	maxNum, err := secureRandomInt(counter)

	if err != nil {
		fmt.Println("random.go getRandomFromPath 0", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	scanner = bufio.NewScanner(file)
	counter = 0
	item := ""

	for scanner.Scan() {
		counter++
		if counter == maxNum {

			item = scanner.Text()
		}
	}
	return item
}

// Returns a random dagger name
func GetRandomDaggerName() string {

	const daggerPath = "mud/equipment/generated/weapons/daggers.txt"
	dagger, err := os.Open(daggerPath)
	defer dagger.Close()

	if err != nil {
		fmt.Println("random.go getRandomDaggerName 0", err)
	}
	return getRandomItemFromPath(dagger)
}

func getRandomBeltsName() string {

	const beltsPath = "mud/equipment/generated/armor/belts.txt"
	belts, err := os.Open(beltsPath)
	defer belts.Close()

	if err != nil {
		fmt.Println("random.go getRandomBeltsName 0", err)
	}
	return getRandomItemFromPath(belts)
}

func getRandomBootsName() string {

	const bootsPath = "mud/equipment/generated/armor/boots.txt"
	boots, err := os.Open(bootsPath)
	defer boots.Close()

	if err != nil {
		fmt.Println("random.go getRandomBeltsName 0", err)
	}
	return getRandomItemFromPath(boots)
}

func getRandomLegsName() string {

	const legsPath = "mud/equipment/generated/armor/legs.txt"
	legs, err := os.Open(legsPath)
	defer legs.Close()

	if err != nil {
		fmt.Println("random.go getRandomLegsName 0", err)
	}
	return getRandomItemFromPath(legs)
}

func getRandomShieldsName() string {

	const shieldsPath = "mud/equipment/generated/armor/shields.txt"
	shields, err := os.Open(shieldsPath)
	defer shields.Close()

	if err != nil {
		fmt.Println("random.go getRandomShieldsName 0", err)
	}
	return getRandomItemFromPath(shields)
}

func getRandomTorsosName() string {

	const torsoPath = "mud/equipment/generated/armor/torso.txt"
	torso, err := os.Open(torsoPath)
	defer torso.Close()

	if err != nil {
		fmt.Println("random.go getRandomTorsosName 0", err)
	}
	return getRandomItemFromPath(torso)
}