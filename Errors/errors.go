package main

import (
	"errors"
	"fmt"
)

func beyond30(age int) error {
	if age >= 30 {
		return errors.New("no beyond 30")
	}
	return nil
}

var ErrTooManyBoyFriend = errors.New("too many boy friend")
var ErrSpendTooManyMoney = errors.New("spend too many money")

type girl struct {
	boyfriend  int
	spendmoney int
	age        int
}

func checkGirl(g *girl) error {
	if g.boyfriend > 2 {
		return ErrTooManyBoyFriend
	} else if g.spendmoney > 1200 {
		return ErrSpendTooManyMoney
	}
	if err := beyond30(g.age); err != nil {
		return err
	}
	return nil
}

func main() {

	girls := []girl{
		{boyfriend: 2, spendmoney: 1200, age: 22}, // 凉宫美咲（虚构）
		{boyfriend: 4, spendmoney: 980, age: 36},  // 樱井奈绪（虚构）
		{boyfriend: 3, spendmoney: 1500, age: 28}, // 白石爱子（虚构）
		{boyfriend: 2, spendmoney: 760, age: 19},  // 高桥唯（虚构）
		{boyfriend: 1, spendmoney: 1340, age: 41}, // 中岛美琴（虚构）
	}

	for _, g := range girls {
		fmt.Println(g)
		if err := checkGirl(&g); err != nil {
			
			if errors.Is(err, ErrSpendTooManyMoney) {
				fmt.Println("太败家了，不能要")
			}

			if errors.Is(err, ErrTooManyBoyFriend) {
				fmt.Println("太花心了，不能要")
			}

			fmt.Println(err)
			continue
		}
		fmt.Println("好女人！")
	}

}
