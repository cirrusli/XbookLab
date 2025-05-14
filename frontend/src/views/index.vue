<template>
  <!-- Left Sidebar -->
  <aside class="sidebar">
    <div class="sidebar-section">
      <h3 class="sidebar-title">领域探索</h3>
      <div class="tag-cloud">
        <span
          v-for="(item, index) in menus"
          :key="index"
          class="tag"
          :class="{ active: curMenuVal === item.value }"
          @click="handleChangeCategory(item.name)"
          >{{ item.name }}</span
        >
      </div>
    </div>
    <!-- <div class="sidebar-section">
      <button class="settings-btn">个性化设置</button>
    </div> -->
  </aside>

  <!-- Right Content -->
  <div class="content">
    <section class="content-section">
      <h2 class="section-title">为你推荐</h2>
      <div class="topic-grid" id="topic-grid">
        <div
          class="topic-card"
          v-for="(item, index) in hotTopicList"
          :key="index"
          @click="goToTopicDetail(item.id)"
        >
          <div class="topic-content">
            <h3 class="topic-title">{{ item.title }}</h3>
            <p class="topic-excerpt">{{ item.excerpt }}</p>
            <div class="topic-meta">
              <span class="topic-likes">{{ item.likes }}</span>
              <span class="topic-category">{{ item.tag }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="content-section">
      <h2 class="section-title">精选好书</h2>
      <div class="book-list">
        <div
          v-for="(item, index) in bookList"
          :key="index"
          class="book-card"
          @click="goToBookDetail(item.id)"
        >
          <div class="book-cover-container">
            <img :src="item.cover" :alt="item.title" class="book-cover" />
          </div>
          <div class="book-info">
            <h3 class="book-title">{{ item.title }}</h3>
            <p class="book-author">{{ item.author }}</p>
            <p class="book-rating">{{ item.rating }}</p>
            <p class="book-desc">{{ item.desc }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetIndexHotTopicApi } from '@/api/hot.js';
import { GetBooksListApi } from '@/api/recommend.js';
import { useRouter } from 'vue-router';
import { GetRecommendedBooksApi } from '@/api/recommend.js';

const router = useRouter();

const menus = ref([
  { id: 0, name: '全部', value: '' },
  { id: 1, name: '科技', value: 'technology' },
  { id: 2, name: '历史', value: 'history' },
  { id: 3, name: '心理', value: 'psychology' },
  { id: 4, name: '艺术', value: 'art' },
  { id: 5, name: '哲学', value: 'philosophy' },
  { id: 6, name: '商业', value: 'business' },
  { id: 7, name: '文学', value: 'literature' },
]);
const curMenuVal = ref('');

const handleChangeCategory = (name) => {
  const menuItem = menus.value.find(item => item.name === name);
  curMenuVal.value = menuItem ? menuItem.id : '';
  handleGetHotTopic();
  handleGetRecommendBookList();
};

// 获取热门话题
const hotTopicList = ref([]);
const handleGetHotTopic = async () => {
  let params = {};
  if (curMenuVal.value) {
    params.tag = curMenuVal.value;
  }
  const { data } = await GetIndexHotTopicApi(params);
  hotTopicList.value = data.slice(0, 6);
};

handleGetHotTopic();

// 获取书单
const bookList = ref([]);
const handleGetBookList = async () => {
  let params = {};
  if (curMenuVal.value) {
    params.tag = curMenuVal.value;
  }
  const { data } = await GetBooksListApi(params);
  bookList.value = data;
};
// 获取推荐书单
const handleGetRecommendBookList = async () => {
  let params = {
    Limit: 10,
    Offset: 0,
    TagFilter: curMenuVal.value || 0
  };
  try {
    const { data } = await GetRecommendedBooksApi(params);
    bookList.value = data.Books.map(book => ({
      id: book.BookID,
      title: book.Title,
      author: book.Author,
      cover: book.Cover,
      rating: book.AverageRating,
      desc: book.Description,
      tag: book.Category
    }));
  } catch (error) {
    router.push('/login');
  }
};
handleGetRecommendBookList();
// 跳转书籍详情
const goToBookDetail = (bookId) => {
  router.push({ path: `/book/${Number(bookId)}` });
};

// 跳转话题详情
const goToTopicDetail = (topicId) => {
  router.push({ path: `/topic/${topicId}` });
};
</script>

<style lang="scss" scoped>
/* Left Sidebar */
.sidebar {
  background: white;
  border-radius: 8px;
  padding: 16px;
  box-shadow: var(--card-shadow);
  height: fit-content;
}

.sidebar-section {
  margin-bottom: 24px;
}

.sidebar-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
  color: var(--primary-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  &::after {
    content: '▼';
    font-size: 12px;
    transition: var(--transition);
  }
  &.collapsed::after {
    transform: rotate(-90deg);
  }
}

.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}

.tag {
  background-color: var(--primary-light);
  color: var(--primary-color);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 14px;
  transition: var(--transition);
  cursor: pointer;
  &:hover,
  &.active {
    background-color: var(--primary-color);
    color: white;
  }
}

.settings-btn {
  display: block;
  width: 100%;
  padding: 8px;
  background-color: var(--primary-light);
  color: var(--primary-color);
  border: none;
  border-radius: 4px;
  font-weight: 500;
  cursor: pointer;
  transition: var(--transition);
  text-align: center;
  &:hover {
    background-color: var(--primary-color);
    color: white;
  }
}
/* Right Content */
.content-section {
  margin-bottom: 32px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 16px;
  color: var(--text-dark);
  display: flex;
  align-items: center;
  &::before {
    content: '';
    display: inline-block;
    width: 4px;
    height: 16px;
    background-color: var(--primary-color);
    margin-right: 8px;
    border-radius: 2px;
  }
}

.topic-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.topic-card {
  background: white;
  border-radius: 8px;
  padding: 28px;
  padding-top: 6px;
  box-shadow: var(--card-shadow);
  transition: var(--transition);
  cursor: pointer;
  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  }
}

.topic-title {
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--text-dark);
}

.topic-excerpt {
  color: var(--text-light);
  font-size: 14px;
  margin-bottom: 12px;
  display: -webkit-box;
-webkit-line-clamp: 3;
line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.topic-meta {
  display: flex;
  align-items: center;
  font-size: 13px;
  color: var(--text-light);
}

.topic-likes {
  display: flex;
  align-items: center;
  margin-right: 12px;
  &::before {
    content: '♥';
    color: #ff4757;
    margin-right: 4px;
  }
}

.book-list {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 24px;
}

.book-card {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: var(--card-shadow);
  transition: var(--transition);
  cursor: pointer;
  &:hover {
    transform: scale(1.05);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  }
}

.book-cover {
  width: 100%;
  height: 160px;
  object-fit: cover;
}

.book-info {
  padding: 12px;
}

.book-title {
  font-weight: 600;
  font-size: 14px;
  margin-bottom: 4px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.book-author {
  color: var(--text-light);
  font-size: 12px;
  margin-bottom: 6px;
}

.book-rating {
  color: #f39c12;
  font-size: 12px;
  margin-bottom: 6px;
}

.book-desc {
  font-size: 12px;
  color: var(--text-light);
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
