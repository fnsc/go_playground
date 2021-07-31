package main

import "fmt"

type Grade struct {
	score1, score2, score3, score4, score5, average, finalAverage float64
}

func calculate(grade Grade) Grade {
	const WEIGHT1 = 2
	const WEIGHT2 = 3
	const WEIGHT3 = 4
	const WEIGTH4 = 1

	grade.average = ((grade.score1 * WEIGHT1) + (grade.score2 * WEIGHT2) + (grade.score3 * WEIGHT3) + (grade.score4 * WEIGTH4)) / (WEIGHT1 + WEIGHT2 + WEIGHT3 + WEIGTH4)

	return grade
}

func inExam(grade Grade) bool {
	return grade.average >= 5.0 && grade.average <= 6.9
}

func calculateExam(grade Grade) Grade {
	grade.finalAverage = (grade.average + grade.score5) / 2

	return grade
}

func printExamResult(grade Grade) {
	if grade.finalAverage >= 5.0 {
		fmt.Printf("Media: %.1f\nAluno em exame.\nNota do exame: %.1f\nAluno aprovado.\nMedia final: %.1f\n", grade.average, grade.score5, grade.finalAverage)

		return
	}

	fmt.Printf("Media: %.1f\nAluno em exame.\nNota do exame: %.1f\nAluno reprovado.\nMedia final: %.1f\n", grade.average, grade.score5, grade.finalAverage)
}

func isAproved(grade Grade) bool {
	return grade.average >= 7.0
}

func main() {
	var grade Grade

	fmt.Scanf("%f %f %f %f", &grade.score1, &grade.score2, &grade.score3, &grade.score4)

	grade = calculate(grade)

	if inExam(grade) {
		fmt.Scanf("%f", &grade.score5)
		grade = calculateExam(grade)
		printExamResult(grade)

		return
	}

	if isAproved(grade) {
		fmt.Printf("Media: %.1f\nAluno aprovado.\n", grade.average)

		return
	}

	fmt.Printf("Media: %.1f\nAluno reprovado.\n", grade.average)
}
