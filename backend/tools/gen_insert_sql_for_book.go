package main

import (
	"fmt"
	"math/rand"
)

func genBooks() {

	// 已有标签
	tags := map[int]string{
		1: "科技",
		2: "历史",
		3: "心理",
		4: "艺术",
		5: "哲学",
	}

	// 模拟作者
	authors := []string{"张三", "李四", "王五", "赵六", "陈七", "钱八", "孙九", "周十"}

	// 描述模板
	descriptions := []string{
		"这是一本关于%s的精彩书籍。",
		"深入探讨%s，适合所有对这个领域感兴趣的人。",
		"内容丰富，见解深刻，值得一读。",
		"%s方面的权威之作。",
		"结合实际案例，理论与实践相结合。",
	}

	// 封面图片链接模板
	coverTemplate := "https://example.com/covers/%d.jpg"

	var bookInserts []string
	var tagInserts []string

	// 记录每个tag已生成多少本书
	tagCount := make(map[int]int)

	for i := 1; i <= 100; i++ {
		tagID := rand.Intn(5) + 1 // 随机选择一个 tag
		tagName := tags[tagID]
		tagCount[tagID]++

		// 构造书名为：科技书籍1、历史书籍3 等等
		title := fmt.Sprintf("%s书籍%d", tagName, tagCount[tagID])

		author := authors[rand.Intn(len(authors))]
		cover := fmt.Sprintf(coverTemplate, i)
		desc := fmt.Sprintf(descriptions[rand.Intn(len(descriptions))], tagName)
		rating := float64(rand.Intn(101)) / 10.0 // 修改为 0.0 到 10.0 的一位小数

		// 构建书籍插入语句
		bookSQL := fmt.Sprintf(
			"INSERT INTO `book` (`book_id`, `title`, `author`, `cover`, `description`, `average_rating`) "+
				"VALUES (%d, '%s', '%s', '%s', '%s', %.1f);",
			i, title, author, cover, desc, rating,
		)
		bookInserts = append(bookInserts, bookSQL)

		// 每本书只关联一个 tag
		tagSQL := fmt.Sprintf("INSERT INTO `book_tag` (`book_id`, `tag_id`) VALUES (%d, %d);", i, tagID)
		tagInserts = append(tagInserts, tagSQL)
	}

	// 输出结果
	fmt.Println("-- 书籍插入语句\n")
	for _, sql := range bookInserts {
		fmt.Println(sql)
	}

	fmt.Println("\n-- 标签关联插入语句\n")
	for _, sql := range tagInserts {
		fmt.Println(sql)
	}
}
