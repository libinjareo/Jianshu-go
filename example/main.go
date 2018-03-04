package main

import (
	"fmt"
	"jianshu-go"
)

func main() {

	//fmt.Println("**** Test one ****")
	//user := jianshu.NewUser("https://www.jianshu.com/u/b7d7dfdc14fb", "", "")
	//fmt.Println(user.GetUserID())
	//fmt.Println(user.GetUserGender())
	//fmt.Println(user.GetUserLink())
	//fmt.Println(user.GetFollowNumber())
	//fmt.Println(user.GetFollowerNumber())
	//fmt.Println(user.GetPassageNumber())
	//fmt.Println(user.GetWriteNumber())
	//fmt.Println(user.GetLikeNumber())
	//fmt.Println(user.GetHomepagePassage())
	//fmt.Println(user.GetPersonalDetail())
	//fmt.Println(user.GetTwitterInfo())
	//fmt.Sprint(user.GetLikedNotes())
	//fmt.Println(user.GetSubscription())
	////fmt.Println(jianshu.GetUuidFromLink(user))
	//
	//fmt.Println("**** Test two ****")
	//user2 := jianshu.NewUser("https://www.jianshu.com/u/3c1c79ba19f3", "", "")
	//fmt.Println(user2.GetUserID())
	//fmt.Println(user2.GetUserGender())
	//fmt.Println(user2.GetUserLink())
	//fmt.Println(user2.GetFollowNumber())
	//fmt.Println(user2.GetFollowerNumber())
	//fmt.Println(user2.GetPassageNumber())
	//fmt.Println(user2.GetWriteNumber())
	//fmt.Println(user2.GetLikeNumber())
	//fmt.Println(user2.GetHomepagePassage())
	//fmt.Println(user2.GetPersonalDetail())
	//fmt.Println(user2.GetTwitterInfo())
	//fmt.Sprint(user2.GetLikedNotes())
	//fmt.Println(user2.GetSubscription())
	//fmt.Println("**** Test three ****")
	//user3 := jianshu.NewUser("https://www.jianshu.com/users/fefb7dffd893/timeline", "", "")
	//fmt.Println(user3.GetUserID())
	//fmt.Println(user3.GetUserGender())
	//fmt.Println(user3.GetUserLink())
	//fmt.Println(user3.GetFollowNumber())
	//fmt.Println(user3.GetFollowerNumber())
	//fmt.Println(user3.GetPassageNumber())
	//fmt.Println(user3.GetWriteNumber())
	//fmt.Println(user3.GetLikeNumber())
	//fmt.Println(user3.GetHomepagePassage())
	//fmt.Println(user3.GetPersonalDetail())
	//fmt.Println(user3.GetTwitterInfo())
	//fmt.Sprint(user3.GetLikedNotes())
	//fmt.Println(user3.GetSubscription())
	//fmt.Println(jianshu.GetUuidFromLink(user3))
	//fmt.Println("**** Test Four ****")
	//user4 := jianshu.NewUser("https://www.jianshu.com/u/1441f4ae075d", "", "")
	//fmt.Println(user4.GetUserID())
	//fmt.Println(user4.GetUserGender())
	//fmt.Println(user4.GetUserLink())
	//fmt.Println(user4.GetFollowNumber())
	//fmt.Println(user4.GetFollowerNumber())
	//fmt.Println(user4.GetPassageNumber())
	//fmt.Println(user4.GetWriteNumber())
	//fmt.Println(user4.GetLikeNumber())
	//fmt.Println(user4.GetHomepagePassage())
	//fmt.Println(user4.GetPersonalDetail())
	//fmt.Println(user4.GetTwitterInfo())
	//fmt.Sprint(user4.GetLikedNotes())
	//fmt.Println(user4.GetSubscription())
	//fmt.Println("**** Test five ****")
	//user5 := jianshu.NewUser("https://www.jianshu.com/u/4062aaeba322", "", "")
	//fmt.Println(user5.GetUserID())
	//fmt.Println(user5.GetUserGender())
	//fmt.Println(user5.GetUserLink())
	//fmt.Println(user5.GetFollowNumber())
	//fmt.Println(user5.GetFollowerNumber())
	//fmt.Println(user5.GetPassageNumber())
	//fmt.Println(user5.GetWriteNumber())
	//fmt.Println(user5.GetLikeNumber())
	//fmt.Println(user5.GetHomepagePassage())
	//fmt.Println(user5.GetPersonalDetail())
	//fmt.Println(user5.GetTwitterInfo())
	//fmt.Println(user5.GetLikedNotes())
	//fmt.Println(user5.GetSubscription())
	//fmt.Println(user5.GetLatestActive())
	//fmt.Println(user5.GetLatestCommented())
	//fmt.Println(user5.GetHotPassage())
	//
	//newTopic := jianshu.NewTopic(2)
	//fmt.Println(newTopic.GetTopicCollectionRecommend())
	//fmt.Println(newTopic.GetTopicCollectionHot())
	//fmt.Println(newTopic.GetTopicCollectionCity())
	//fmt.Println(newTopic.GetTopicCollectionSchoolyard())
	//
	//newAuthor := jianshu.NewRecommendAuthor(7)
	//fmt.Println(newAuthor.GetListRecommendAuthor())
	//
	//newHomePage := jianshu.NewHomePage(2, 1)
	//fmt.Println(newHomePage.GetHomePagePassages())
	//fmt.Println("新上榜")
	//newHomePage.GetNewList()
	//fmt.Println("七日热门")
	//newHomePage.GetHotSeven()
	//fmt.Println("30日热门")
	//newHomePage.GetHotMonth()
	//fmt.Println("简书大学堂")
	//newHomePage.GetJianShuSchool()

	//newPublication := jianshu.NewPublication()
	//// 已出版图书
	//fmt.Println("已出版图书")
	//newPublication.GetPublicizedBook()
	//// 小说
	//fmt.Println("小说")
	//newPublication.GetNovelBooks()
	//// IT
	//fmt.Println("IT")
	//newPublication.GetITAndJobMarket()
	//// culture
	//fmt.Println("culture")
	//newPublication.GetCultureAndHistory()
	//// magazine
	//fmt.Println("magazine")
	//newPublication.GetMonthlyMagazine()
	//newSearch := jianshu.NewSearch("谢小路")
	//newSearch.GetAuthor(2)

	newArticle := jianshu.NewArticle("https://www.jianshu.com/p/a043b5008190")
	fmt.Println(newArticle.GetAuthor())
	fmt.Println(newArticle.GetTitle())
	fmt.Println(newArticle.GetDescription())
	fmt.Println(newArticle.GetContent())

	newTrendSearch := jianshu.NewTrendSearch()
	newTrendSearch.GetTrendSearchList()

}
