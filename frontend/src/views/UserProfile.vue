<template>
  <div class="profile-container">
    <h2 class="section-title">我的资料</h2>
    <div class="user-card">
      <div class="avatar-wrapper">
        <img src="https://picsum.photos/100/100" class="user-avatar" alt="头像">
      </div>
      <div class="user-info">
        <h3 class="username">示例用户</h3>
        <p class="bio">这个人很懒，什么都没有写～</p>
      </div>
    </div>
    
    <h3 class="subsection-title">我的收藏</h3>
    <div class="collection-list">
      <div v-for="n in 3" :key="n" class="collection-item">
        <img src="https://picsum.photos/80/120" class="book-cover" alt="封面">
        <div class="collection-info">
          <h4 class="book-title">收藏图书{{ n }}</h4>
          <p class="collect-time">收藏时间：2023-07-15</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  padding: 20px 0;
}

.user-card {
  display: flex;
  align-items: center;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  margin-bottom: 30px;
}

.user-avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  margin-right: 30px;
}

.username {
  font-size: 20px;
  margin-bottom: 8px;
}

.bio {
  color: #666;
  font-size: 14px;
}

.subsection-title {
  font-size: 16px;
  color: #333;
  margin-bottom: 20px;
}

.collection-list {
  display: grid;
  gap: 20px;
}

.collection-item {
  display: flex;
  align-items: center;
  padding: 15px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.1);
}

.book-cover {
  width: 80px;
  height: 120px;
  object-fit: cover;
  margin-right: 20px;
}

.collection-info {
  flex: 1;
}

.book-title {
  font-size: 16px;
  margin-bottom: 8px;
}

.collect-time {
  color: #666;
  font-size: 14px;
}
</style>