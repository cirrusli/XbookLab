<template>
  <aside class="sidebar">
    <div class="sidebar-section">
      <h3 class="sidebar-title">好友分类</h3>
      <div class="tag-cloud">
        <span
          class="tag"
          v-for="(item, index) in categoryList"
          :key="index"
          :class="{ active: item.id === curCategoryId }"
          @click="handleSwitchtype(item.id)"
          >{{ item.name }}</span
        >
      </div>
    </div>
  </aside>

  <div class="content">
    <section class="content-section">
      <h2 class="section-title">好友在读</h2>
      <div class="friends-container" id="friends-container">
        <div
          class="friend-item"
          v-for="(item, index) in friendsBookList"
          :key="index"
        >
          <img :src="item.avatar" :alt="item.name" class="friend-avatar" />
          <div class="friend-info">
            <div class="friend-name">{{ item.name }}</div>
            <div class="friend-reading">{{ item.reading }}</div>
            <div class="friend-book">
              <img
                :src="item.bookCover"
                :alt="item.bookTitle"
                class="friend-book-cover"
              />
              <div class="friend-book-info">
                <div class="friend-book-title">{{ item.bookTitle }}</div>
                <p>{{ item.bookComment }}</p>
                <div class="friend-book-progress">
                  <div
                    class="friend-book-progress-bar"
                    :style="{ width: item.bookProgress + '%' }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetFriendsBookListApi } from '@/api/friends.js';

const curCategoryId = ref('0');
const categoryList = ref([
  {
    id: '0',
    name: '全部',
    value: 'all',
  },
  {
    id: '1',
    name: '最近活跃',
    value: 'active',
  },
  {
    id: '2',
    name: '同城',
    value: 'local',
  },
]);
const friendsBookList = ref([]);
const handleGetFriendsBook = async () => {
  let params = {
    type: curCategoryId.value,
  };
  const { data } = await GetFriendsBookListApi(params);
  friendsBookList.value = data;
};
handleGetFriendsBook();

const handleSwitchtype = (type) => {
  curCategoryId.value = type;
  handleGetFriendsBook();
};
</script>

<style lang="scss" scoped>
.friend-item {
  display: flex;
  align-items: flex-start;
  padding: 20px;
  background: white;
  border-radius: 12px;
  margin-bottom: 20px;
  box-shadow: var(--card-shadow);
  transition: var(--transition);
  position: relative;
  overflow: hidden;
}

.friend-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(66, 185, 131, 0.15);
}

.friend-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  background: var(--primary-color);
  opacity: 0;
  transition: var(--transition);
}

.friend-item:hover::before {
  opacity: 1;
}

.friend-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 20px;
  border: 3px solid white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.friend-info {
  flex: 1;
}

.friend-name {
  font-weight: 600;
  margin-bottom: 6px;
  font-size: 16px;
  color: var(--text-dark);
}

.friend-reading {
  font-size: 14px;
  color: var(--primary-color);
  margin-bottom: 12px;
  font-weight: 500;
}

.friend-book {
  display: flex;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--primary-light);
}

.friend-book-cover {
  width: 70px;
  height: 95px;
  object-fit: cover;
  border-radius: 6px;
  margin-right: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: var(--transition);
}

.friend-book-cover:hover {
  transform: scale(1.03);
}

.friend-book-info {
  flex: 1;
}

.friend-book-title {
  font-size: 15px;
  margin-bottom: 8px;
  font-weight: 600;
}

.friend-book-comment {
  font-size: 14px;
  color: var(--text-light);
  line-height: 1.5;
  margin-bottom: 12px;
}

.friend-book-progress {
  height: 6px;
  background: #f0f0f0;
  border-radius: 3px;
  margin-top: 12px;
  overflow: hidden;
}

.friend-book-progress-bar {
  height: 100%;
  background: linear-gradient(90deg, var(--primary-color), #5cd6a9);
  border-radius: 3px;
  transition: width 0.6s ease;
  position: relative;
}

.friend-book-progress-bar::after {
  content: attr(data-progress);
  position: absolute;
  right: 8px;
  top: -22px;
  font-size: 12px;
  color: var(--primary-color);
  font-weight: 600;
}
</style>
