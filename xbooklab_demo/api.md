``` protobuf
syntax = "proto3";

package xbooklab;
option go_package = "./;xbooklab";

service XbookLab {
  // User
  rpc Register(RegisterReq) returns (RegisterResp);
  // 用户主页
  rpc GetUserProfile(GetUserProfileReq) returns (GetUserProfileResp);
  rpc UpdateUserInfo(UpdateUserInfoReq) returns (UpdateUserInfoResp);

  // Auth
  rpc Login(LoginReq) returns (LoginResp);
  rpc Logout(LogoutReq) returns (LogoutResp);
  rpc ChangePassword(ChangePasswordReq) returns (ChangePasswordResp);

  // Book
  rpc CreateBook(CreateBookReq) returns (CreateBookResp);
  // 点击书籍进入详情页，获取书籍的详细信息，能否区分登录态，比如有token的情况会再调用RecordBookView
  rpc GetBookDetail(GetBookDetailReq) returns (GetBookDetailResp);
  // 会在interaction表中一并记录用户的评分行为
  rpc SetBookRating(SetBookRatingReq) returns (SetBookRatingResp);

  // Recommend 首页的推荐
  // 先只做书籍的个性化推荐，协同过滤推荐算法，还要考虑混合新鲜内容以及新用户的冷启动问题
  rpc GetRecommendedBooks(GetRecommendedBooksReq)
      returns (GetRecommendedBooksResp);

  // Interaction 
  // 用户点击书籍详情页，记录用户的浏览行为
  rpc RecordBookView(RecordBookViewReq) returns (RecordBookViewResp);
  // 记录用户评分行为
  rpc RecordBookRating(RecordBookRatingReq) returns (RecordBookRatingResp);
  // 记录用户评论行为
  rpc RecordBookComment(RecordCommentReq) returns (RecordCommentResp);

  // Topic
  rpc GetTopicList(GetTopicListReq) returns (GetTopicListResp);
  rpc CreateTopic(CreateTopicReq) returns (CreateTopicResp);
  rpc GetTopicDetail(GetTopicDetailReq) returns (GetTopicDetailResp);
  rpc LikeTopic(LikeTopicReq) returns (LikeTopicResp);
  rpc DeleteTopic(DeleteTopicReq) returns (DeleteTopicResp);

  // Comment
  rpc CreateComment(CreateCommentReq) returns (CreateCommentResp);
  rpc GetCommentList(GetCommentListReq) returns (GetCommentListResp);

  // Follow
  rpc FollowUser(FollowUserReq) returns (FollowResp);
  rpc UnfollowUser(UnfollowUserReq) returns (FollowResp);
  rpc GetFollowingList(GetFollowingReq) returns (FollowListResp);
  rpc GetFollowersList(GetFollowersReq) returns (FollowListResp);

  // Manager
  rpc GetBookList(GetBookListReq) returns (GetBookListResp);
  rpc GetUserList(GetUserListReq) returns (GetUserListResp);
  // 书籍、话题会有兴趣标签，在创建书籍、话题时添加，如科技、历史、心理、艺术、哲学等
  rpc GetTagList(GetTagListReq) returns (GetTagListResp);
  rpc AddTag(AddTagReq) returns (AddTagResp);
  rpc DeleteTag(DeleteTagReq) returns (DeleteTagResp);
  rpc DeleteUser(DeleteUserReq) returns (DeleteUserResp);
  rpc DeleteBook(DeleteBookReq) returns (DeleteBookResp);
}

message CreateCommentReq {
  Comment comment = 1;
}
message CreateCommentResp {
  BaseResp base_resp = 1;
}

message GetCommentListReq {
  uint32 comment_id = 1;  // 评论区会有唯一id，可以拉到评论列表
  uint32 comment_type = 2; // 评论类型，0是书籍，1是话题，区分是书籍还是话题的评论区
}

message GetCommentListResp {
  repeated Comment comments = 1; // 评论列表
}

message BaseResp {
  uint32 code = 1;
  string message = 2;
}
message UpdateBookReq { Book book = 1; }
message UpdateBookResp { BaseResp base_resp = 1; }

message DeleteUserReq { uint32 user_id = 1; }
message DeleteUserResp {BaseResp base_resp = 1;}

message GetUserListReq {
  uint32 offset = 1;
  uint32 limit = 2;
}

message GetUserListResp { repeated User users = 1; }
message RecordCommentReq {
  uint32 user_id = 1; // 带token，无需传user_id
  uint32 comment_id = 2;
}

message RecordCommentResp {BaseResp base_resp = 1;}

message AddTagReq { string tag_name = 1; }
message AddTagResp {
  BaseResp base_resp = 1;
}

message DeleteTagReq { uint32 tag_id = 1; }

message DeleteTagResp { BaseResp base_resp = 1;  }
message GetTagListReq {
}
message GetTagListResp { repeated Tag tags = 1; }

message Tag {
  uint32 tag_id = 1;
  string tag_name = 2;
}

message FollowUserReq {
  uint32 user_id = 1;
  uint32 follow_user_id = 2;
}

message UnfollowUserReq {
  uint32 user_id = 1;
  uint32 follow_user_id = 2;
}

message GetFollowingReq {
  uint32 user_id = 1;
}

message GetFollowersReq {
  uint32 user_id = 1;
}

message FollowListResp {
  repeated User user_list = 1;
}

message FollowResp { BaseResp base_resp = 1;  }

message RegisterReq { User user = 1; }
message RegisterResp { BaseResp base_resp = 1; }

message GetUserProfileReq { uint32 user_id = 1; }

message Event {
  uint32 event_id = 1;      // 事件id，例如阅读书籍、评论、评分等
  string event_content = 2; // 事件内容，例如阅读了《三体》
  uint32 event_time = 3;    // 事件时间timestamp
  uint32 user_id = 4;       // 事件的用户id
}
message GetUserProfileResp {
  User user = 1;
  Book recent_reading = 2;
  repeated User recent_top3_following_users = 3;
  repeated Event recent_events = 4; // 最近的10条事件
}

message UpdateUserInfoReq { User user = 1; }
message UpdateUserInfoResp { BaseResp base_resp = 1; }

message LoginReq {
  string username = 1;
  string password = 2;
}
message LoginResp { string token = 1; }

message LogoutReq { string token = 1; }
message LogoutResp { BaseResp base_resp = 1; }

message ChangePasswordReq {
  uint32 user_id = 1;
  string old_password = 2;
  string new_password = 3;
}

message ChangePasswordResp { BaseResp base_resp = 1; }

message GetBookListReq {
  uint32 limit = 1;
  uint32 offset = 2;
}
message GetBookListResp {
  repeated Book book_list = 1;
  uint32 total_num = 2;
}

message Book {
  uint32 book_id = 1;
  string title = 2; // 书名
  string author = 3;
  string cover = 4;
  string description = 5;
  string rating = 6;
  string tag_name = 7;
}
message Comment {
  uint32 comment_id = 1; // 用评论的书籍或者话题id作为comment_id，一个评论区中的所有评论都有相同的comment_id
  uint32 comment_type = 2; // 评论类型，0为书籍评论，1为话题评论
  // 直接把用户的username存到content，不考虑回复评论的情况
  uint32 content = 3; // 书友小李说：这本书很不错！
  uint32 user_id = 4; // 评论的用户id
}
message CreateBookReq { Book book = 1; }

message CreateBookResp { BaseResp base_resp = 1; }

message GetBookDetailReq { uint32 book_id = 1; }

message GetBookDetailResp { Book book = 1; }

message SetBookRatingReq {
  uint32 book_id = 1; // 书籍id
  uint32 rating = 2;  // 评分，0-10
  uint32 user_id = 3; // 用户id
}

message SetBookRatingResp { string message = 1; }

message DeleteBookReq { uint32 book_id = 1; }

message DeleteBookResp { BaseResp base_resp = 1;}

message GetRecommendedBooksReq {
  uint32 user_id = 1;
  uint32 limit = 2;
  uint32 offset = 3;
  uint32 tag_filter = 4; // 筛选指定分类的书籍，例如科技、历史、心理、艺术、哲学等
}
message GetRecommendedBooksResp { repeated Book books = 1; }

message RecordBookViewReq {
  uint32 user_id = 1; // 带token，无需传user_id
  uint32 book_id = 2;
}
message RecordBookViewResp { BaseResp base_resp = 1; }

message RecordBookRatingReq {
  uint32 user_id = 1; // 带token，无需传user_id
  uint32 book_id = 2;
  uint32 rating = 3; // 评分
}

message RecordBookRatingResp { BaseResp base_resp = 1; }

message GetTopicListReq {
  uint32 offset = 1;
  uint32 limit = 2;
  uint32 tag_filter = 3; // 筛选指定分类的话题，例如科技、历史、心理、艺术、哲学等
}

message GetTopicListResp {
  repeated Topic topics = 1;
  uint32 total = 2;
}

message Topic {
  uint32 topic_id = 1;
  string title = 2;
  string content = 3;
  uint32 author_user_id = 4;
  uint32 like_count = 5;
  string tag_name = 6; // 话题的分类，例如科技、历史、心理、艺术、哲学等
}

message CreateTopicReq {
  string title = 1;
  string content = 2;
  uint32 author_user_id = 3;
}

message CreateTopicResp {
  uint32 topic_id = 1;
  BaseResp base_resp = 2;
}

message GetTopicDetailReq { uint32 topic_id = 1; }

message GetTopicDetailResp { Topic topic = 1; }

message LikeTopicReq {
  uint32 user_id = 1; // 带token，无需传user_id
  uint32 topic_id = 2;
}
message LikeTopicResp { BaseResp base_resp = 1;}

message DeleteTopicReq { uint32 topic_id = 1; }
message DeleteTopicResp { BaseResp base_resp = 1; }

message User {
  uint32 id = 1;
  string username = 2;
  string password = 3;
  string avatar = 4;
  string bio = 5;
  repeated Tag tags = 6; // 兴趣标签，例如科技、历史、心理、艺术、哲学等
}
```

