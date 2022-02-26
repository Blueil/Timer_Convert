package main;

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/sqweek/dialog"
	"strconv"
);

func main() {

	path, pathErr := os.Getwd();
	check(pathErr);

	var filename string;

	if len(os.Args) < 2 {
		//fmt.Println("Please specify a file to open!");
		//return;
		var dialogErr error;
		filename, dialogErr = dialog.File().Filter("Twisty Timer Export File (.txt)", "txt").Load();
		check(dialogErr);
	} else {
		filename = os.Args[1];
	}

	dat, datErr := os.ReadFile(filename);
	check(datErr);

	content := strings.Split(strings.ReplaceAll(string(dat), `"`, ""), "\n");

	solves := []string{"No.;Time;Comment;Scramble;Date;P.1"};

	var sovlesLimitStr string;
	fmt.Println("How much solves do you want to convert from " + fmt.Sprint(len(content) - 1) + " solves?");
	fmt.Scanln(&sovlesLimitStr);

	solvesLimit, solvesLimitErr := strconv.Atoi(sovlesLimitStr);
	check(solvesLimitErr);

	for i := len(content) - solvesLimit - 1; i < len(content); i++ {
		if content[i] == "" {
			continue;
		}

		data := strings.Split(content[i], ";");

		date, dateErr := time.Parse("2006-01-2T15:04:05.000", strings.Split(data[2], "+")[0]);
		check(dateErr);
		time := data[0]
		scramble := data[1];
		dateStr := date.Format("2006-01-2 15:04:05");
		dnf := len(data) == 4;

		if dnf {
			time = "DNF(" + time + ")"
		}

		solves = append(solves, fmt.Sprintf("%d;%s;;%s;%s;%s", i - len(content) + solvesLimit + 2, time, scramble, dateStr, data[0]));
	}

	err := os.WriteFile(path + "/" + "output.csv", []byte(strings.Join(solves, "\n")), 0222);
	check(err);
}

func check(err error) {
	if err != nil {
		panic(err);
	}
}
