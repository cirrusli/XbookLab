package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

//	func main() {
//		// genInsertBookSQL()
//		// genInsertUserAndUserTagSQL()
//		// genInsertBookViewAndRatingSQL()
//		// genInsertFollowSQL()
//		genInsertRecommendSQL()
//	}
func genInsertBookSQL() {

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
	// coverTemplate := "https://example.com/covers/%d.jpg"

	var bookInserts []string
	var tagInserts []string

	// 记录每个tag已生成多少本书
	tagCount := make(map[int]int)
	var desc string
	for i := 1; i <= 100; i++ {
		tagID := rand.Intn(5) + 1 // 随机选择一个 tag
		tagName := tags[tagID]
		tagCount[tagID]++

		// 构造书名为：科技书籍1、历史书籍3 等等
		title := fmt.Sprintf("%s书籍%d", tagName, tagCount[tagID])

		author := authors[rand.Intn(len(authors))]
		// cover := fmt.Sprintf(coverTemplate, i)
		cover := "D:\\Gooo\\Dev_Projects\\XbookLab\\backend\\assets\\book_cover\\default.png"
		descTemplate := descriptions[rand.Intn(len(descriptions))]
		if strings.Contains(descTemplate, "%s") {
			desc = fmt.Sprintf(descTemplate, tagName)
		} else {
			desc = descTemplate
		}
		rating := float64(rand.Intn(51)+50) / 10.0 // 生成5.0到10.0的一位小数
		if rating > 8.0 {
			title = "高分好书" + title
		}
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

	// 写入到文件
	file, err := os.Create("../XbookLab/backend/sql/insert_book.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入批量书籍插入语句
	file.WriteString("-- 批量书籍插入语句\n\n")
	batchSQL := "INSERT INTO `book` (`book_id`, `title`, `author`, `cover`, `description`, `average_rating`) VALUES "
	for i, sql := range bookInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")

	// 写入批量标签关联插入语句
	file.WriteString("\n-- 批量标签关联插入语句\n\n")
	batchSQL = "INSERT INTO `book_tag` (`book_id`, `tag_id`) VALUES "
	for i, sql := range tagInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")
}
func genInsertUserAndUserTagSQL() {
	// 模拟用户名
	usernames := []string{"book_lover", "reading_fan", "literature_geek", "novel_enthusiast", "page_turner", "story_seeker", "word_devourer", "library_ghost", "ink_drinker", "bibliophile"}

	// 模拟个人简介
	bios := []string{
		"热爱阅读，喜欢分享读书心得",
		"每天都要读一本新书",
		"科幻小说爱好者",
		"历史书籍收藏家",
		"心理学研究者",
		"艺术与文学爱好者",
		"哲学思考者",
		"书是人类进步的阶梯",
		"在书中寻找人生的答案",
		"阅读让生活更美好",
	}

	// 密码哈希(123456)
	passwordHash := "$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW"

	var userInserts []string
	var userTagInserts []string

	for i := 1; i <= 100; i++ {
		username := fmt.Sprintf("%s%d", usernames[rand.Intn(len(usernames))], i)
		gender := "men"
		if rand.Intn(2) == 0 {
			gender = "women"
		}
		avatar := fmt.Sprintf("https://randomuser.me/api/portraits/%s/%d.jpg", gender, rand.Intn(100)+1)
		bio := bios[rand.Intn(len(bios))]

		// 构建用户插入语句
		userSQL := fmt.Sprintf(
			"INSERT INTO `user` (`username`, `password`, `avatar`, `bio`) "+
				"VALUES ('%s', '%s', '%s', '%s');",
			username, passwordHash, avatar, bio,
		)
		userInserts = append(userInserts, userSQL)

		// 每个用户随机关联1-3个标签
		tagCount := rand.Intn(3) + 1
		for j := 0; j < tagCount; j++ {
			tagID := rand.Intn(5) + 1
			userTagSQL := fmt.Sprintf("INSERT INTO `user_tag` (`user_id`, `tag_id`) VALUES (%d, %d);", i, tagID)
			userTagInserts = append(userTagInserts, userTagSQL)
		}
	}

	// 写入到文件
	file, err := os.Create("../XbookLab/backend/sql/insert_user.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入批量用户插入语句
	file.WriteString("-- 批量用户插入语句\n\n")
	batchSQL := "INSERT INTO `user` (`username`, `password`, `avatar`, `bio`) VALUES "
	for i, sql := range userInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")

	// 写入批量用户标签关联插入语句
	file.WriteString("\n-- 批量用户标签关联插入语句\n\n")
	batchSQL = "INSERT INTO `user_tag` (`user_id`, `tag_id`) VALUES "
	for i, sql := range userTagInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")
}
func genInsertBookViewAndRatingSQL() {
	var viewInserts []string
	var ratingInserts []string

	// 每个用户浏览3-10本书并随机评分
	for userID := 1; userID <= 100; userID++ {
		viewCount := rand.Intn(8) + 3 // 3-10次浏览
		viewedBooks := make(map[int]bool)

		for j := 0; j < viewCount; j++ {
			bookID := rand.Intn(100) + 1 // 1-100本书
			if _, exists := viewedBooks[bookID]; exists {
				continue // 避免重复浏览同一本书
			}
			viewedBooks[bookID] = true

			// 构建浏览记录插入语句
			viewSQL := fmt.Sprintf("INSERT INTO `book_view` (`user_id`, `book_id`) VALUES (%d, %d);", userID, bookID)
			viewInserts = append(viewInserts, viewSQL)

			// 50%概率给浏览过的书评分
			if rand.Intn(2) == 0 {
				rating := rand.Intn(10) + 1 // 1-10分
				ratingSQL := fmt.Sprintf("INSERT INTO `rating` (`user_id`, `book_id`, `rating_value`) VALUES (%d, %d, %d);", userID, bookID, rating)
				ratingInserts = append(ratingInserts, ratingSQL)
			}
		}
	}

	// 写入到文件
	file, err := os.Create("../XbookLab/backend/sql/insert_book_view_rating.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入批量浏览记录插入语句
	file.WriteString("-- 批量书籍浏览记录插入语句\n\n")
	batchSQL := "INSERT INTO `book_view` (`user_id`, `book_id`) VALUES "
	for i, sql := range viewInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")

	// 写入批量评分插入语句
	file.WriteString("\n-- 批量用户评分插入语句\n\n")
	batchSQL = "INSERT INTO `rating` (`user_id`, `book_id`, `rating_value`) VALUES "
	for i, sql := range ratingInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")
}
func genInsertFollowSQL() {
	var followInserts []string

	// 为每个用户(10-100)生成10-20个关注关系
	for userID := 10; userID <= 100; userID++ {
		followCount := rand.Intn(11) + 10 // 10-20个关注
		followedUsers := make(map[int]bool)

		for j := 0; j < followCount; j++ {
			followedID := rand.Intn(100) + 1 // 1-100用户
			if followedID == userID || followedUsers[followedID] {
				continue // 避免关注自己和重复关注
			}
			followedUsers[followedID] = true

			// 构建关注关系SQL
			followSQL := fmt.Sprintf("INSERT INTO `follow` (`follower_id`, `followed_id`) VALUES (%d, %d);", userID, followedID)
			followInserts = append(followInserts, followSQL)
		}
	}

	// 写入到文件
	file, err := os.Create("../XbookLab/backend/sql/insert_follow.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入批量关注关系插入语句
	file.WriteString("-- 批量用户关注关系插入语句\n\n")
	batchSQL := "INSERT INTO `follow` (`follower_id`, `followed_id`) VALUES "
	for i, sql := range followInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")
}

func genInsertRecommendSQL() {
	var recommendInserts []string

	// 科技类书籍ID范围:1-35
	// 艺术类书籍ID范围:36-45
	// 哲学类书籍ID范围:46-55

	// 为user_id=1生成推荐数据
	userID := 1

	// 生成推荐数据时按比例混合不同类型
	techCount := 14 // 科技类70%
	artCount := 2   // 艺术类10%
	philoCount := 4 // 哲学类20%

	// 生成科技类推荐
	for i := 0; i < techCount; i++ {
		bookID := rand.Intn(35) + 1
		score := 9.0 + rand.Float64() // 9.0-10.0分
		recommendSQL := fmt.Sprintf("INSERT INTO `recommend` (`user_id`, `book_id`, `score`, `update_at`) VALUES (%d, %d, %.1f, NOW());", userID, bookID, score)
		recommendInserts = append(recommendInserts, recommendSQL)
	}

	// 生成艺术类推荐
	for i := 0; i < artCount; i++ {
		bookID := rand.Intn(10) + 36
		score := 8.0 + rand.Float64() // 8.0-9.0分
		recommendSQL := fmt.Sprintf("INSERT INTO `recommend` (`user_id`, `book_id`, `score`, `update_at`) VALUES (%d, %d, %.1f, NOW());", userID, bookID, score)
		recommendInserts = append(recommendInserts, recommendSQL)
	}

	// 生成哲学类推荐
	for i := 0; i < philoCount; i++ {
		bookID := rand.Intn(10) + 46
		score := 7.5 + rand.Float64() // 7.5-8.5分
		recommendSQL := fmt.Sprintf("INSERT INTO `recommend` (`user_id`, `book_id`, `score`, `update_at`) VALUES (%d, %d, %.1f, NOW());", userID, bookID, score)
		recommendInserts = append(recommendInserts, recommendSQL)
	}

	// 打乱推荐顺序，确保每页10条数据中混合不同类型
	rand.Shuffle(len(recommendInserts), func(i, j int) {
		recommendInserts[i], recommendInserts[j] = recommendInserts[j], recommendInserts[i]
	})

	// 写入到文件
	file, err := os.Create("../XbookLab/backend/sql/insert_recommend.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入批量推荐插入语句
	file.WriteString("-- 批量推荐数据插入语句\n\n")
	batchSQL := "INSERT INTO `recommend` (`user_id`, `book_id`, `score`, `update_at`) VALUES "
	for i, sql := range recommendInserts {
		if i > 0 {
			batchSQL += ", "
		}
		batchSQL += sql[strings.Index(sql, "VALUES ")+7 : len(sql)-1]
	}
	file.WriteString(batchSQL + ";\n")
}
