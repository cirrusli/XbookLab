<template>
  <div class="profile-content">
    <aside class="profile-sidebar">
      <h3>最近在读</h3>
      <div class="reading-book">
        <img
          src="http://localhost:8000/static/book_cover/san_ti.jpg"
          alt="三体"
          class="book-cover"
        />
        <div class="book-info">
          <h4>三体</h4>
          <p>刘慈欣 · 科幻</p>
        </div>
      </div>

      <h3>关注的人</h3>
      <div class="following-list">
        <div
          class="following-item"
          v-for="(item, index) in followList"
          :key="index"
        >
          <img :src="item.avatar" alt="用户头像" class="following-avatar" />
          <span>{{ item.name }}</span>
        </div>
      </div>
    </aside>

    <div class="profile-activity">
      <div class="activity-header">
        <h2>动态</h2>
      </div>

      <div class="activity-item">
        <p>发表了话题 <a href="#">"如何理解《存在与时间》中的此在概念"</a></p>
        <p class="activity-time">2小时前</p>
      </div>

      <div class="activity-item">
        <p>开始阅读 <a href="#">《三体》</a> 并添加了书评</p>
        <p class="activity-time">昨天</p>
      </div>

      <div class="activity-item">
        <p>关注了 <a href="#">黑格尔</a> 并收藏了他的书单</p>
        <p class="activity-time">3天前</p>
      </div>

      <div class="activity-item">
        <p>加入了书友圈 <a href="#">"哲学与思辨"</a></p>
        <p class="activity-time">5天前</p>
      </div>

      <div class="activity-item">
        <p>完成了 <a href="#">《人类简史》</a> 的阅读</p>
        <p class="activity-time">1周前</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetFollowersListApi } from '@/api/user.js';

let followList = ref([]);
const handleGetFollowersList = async () => {
  const { data } = await GetFollowersListApi();
  followList.value = data.list.length > 6 ? data.list.slice(0, 6) : data.list;
};
handleGetFollowersList();
</script>

<style lang="scss" scoped>
.profile-content {
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 2rem;
}

.profile-sidebar {
  background: var(--profile-card-bg);
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: var(--profile-shadow);
  position: sticky;
  top: 6rem;
  align-self: flex-start;
  h3 {
    font-size: 1.1rem;
    font-weight: 600;
    margin-bottom: 1rem;
    color: var(--profile-primary);
    padding-bottom: 0.5rem;
    border-bottom: 1px solid var(--profile-border);
  }
}

.profile-activity {
  background: var(--profile-card-bg);
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: var(--profile-shadow);
}

.activity-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid var(--profile-border);
  h2 {
    font-size: 1.3rem;
    font-weight: 600;
    color: var(--profile-text);
  }
}

.activity-item {
  padding: 1.2rem 0;
  border-bottom: 1px solid var(--profile-border);
  transition: all 0.3s ease;
  &:hover {
    background-color: rgba(66, 185, 131, 0.05);
  }
  &:last-child {
    border-bottom: none;
  }
}

.activity-item {
  p {
    margin-bottom: 0.5rem;
    color: var(--profile-text);
    line-height: 1.6;
  }
  a {
    color: var(--profile-primary);
    font-weight: 500;
    transition: all 0.2s ease;
    &:hover {
      text-decoration: underline;
    }
  }
}

.activity-time {
  color: var(--profile-text-light);
  font-size: 0.85rem;
  display: flex;
  align-items: center;
  &::before {
    content: '🕒';
    margin-right: 0.3rem;
    font-size: 0.8rem;
  }
}

/* Reading Book */
.reading-book {
  display: flex;
  align-items: center;
  margin-bottom: 1.5rem;
  padding: 0.8rem;
  border-radius: 8px;
  background-color: rgba(66, 185, 131, 0.05);
  transition: all 0.3s ease;
  &:hover {
    background-color: rgba(66, 185, 131, 0.1);
  }
}

.book-cover {
  width: 60px;
  height: 80px;
  border-radius: 4px;
  object-fit: cover;
  margin-right: 1rem;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.book-info {
  h4 {
    font-size: 1rem;
    font-weight: 600;
    margin-bottom: 0.3rem;
    color: var(--profile-text);
  }

  p {
    font-size: 0.85rem;
    color: var(--profile-text-light);
  }
}

/* Following List */
.following-list {
  display: flex;
  flex-direction: column;
  gap: 0.8rem;
}

.following-item {
  display: flex;
  align-items: center;
  padding: 0.8rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  cursor: pointer;
  &:hover {
    background-color: rgba(66, 185, 131, 0.05);
  }
}

.following-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 0.8rem;
}

/* Responsive */
@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    text-align: center;
    padding: 2rem 1rem;
  }

  .profile-avatar {
    margin-right: 0;
    margin-bottom: 1.5rem;
  }

  .profile-bio {
    max-width: 100%;
  }

  .profile-content {
    grid-template-columns: 1fr;
  }

  .profile-sidebar {
    position: static;
  }
}
</style>
