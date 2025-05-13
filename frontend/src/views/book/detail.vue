<template>
  <Header />
  <main class="detail-container">
    <div class="book-detail">
      <img :src="bookInfo.Cover" alt="书籍封面" class="detail-cover" />
      <div class="detail-info">
        <h1>{{ bookInfo.Title }}</h1>
        <p class="author">作者：{{ bookInfo.Author }}</p>
        <p class="rating">评分：{{ bookInfo.AverageRating }}</p>
        <p class="desc">{{ bookInfo.Description }}</p>
      </div>
    </div>

    <div>
      <h3 class="section-title">我要评分</h3>
      <a-rate v-model:value="rateVal" allow-half @change="handleSubmitRate" />
    </div>

    <!-- 评论区 -->
    <Comment :commentType="0" :bookId=state.bookId />
  </main>
</template>

<script setup>
import { reactive, ref } from 'vue';
import Header from '@/components/Header/index.vue';
import {
  GetBookDetailApi,
  RecordBookRatingApi,
  RecordBookViewApi,
} from '@/api/recommend.js';
import { useRoute } from 'vue-router';
import Comment from '@/components/Comments/index.vue';
import { CreateSuccessMessage } from '@/utils/alert.js';

const route = useRoute();

const state = reactive({
  isLogin: localStorage.getItem('token') ? true : false,
  bookId: 0,
});

const bookInfo = ref({}); // 书籍信息

const handleGetBookDetail = async () => {
  state.bookId = Number(route.params.id); // 获取书籍ID并转换为数字类型
  const { data } = await GetBookDetailApi(state.bookId);
  bookInfo.value = data;
};

handleGetBookDetail();

const rateVal = ref(1);
const handleSubmitRate = () => {
  let params = {
    book_id: Number(route.params.id),
    rating: rateVal.value,
  };
  RecordBookRatingApi(params)
    .then((result) => {
      CreateSuccessMessage('评分成功');
    })
    .catch((err) => {
      console.error(err);
    });
};

if (state.isLogin) {
  RecordBookViewApi({ book_id: state.bookId })
    .then((result) => {
      console.log('记录成功');
    })
    .catch((err) => {
      console.error('记录失败');
    });
}
</script>

<style lang="scss" scoped>
.detail-container {
  max-width: 1200px;
  margin: 40px auto;
  padding: 0 20px;
}

/* 书籍详情页专用样式 */
.book-detail {
  display: grid;
  grid-template-columns: 240px 1fr;
  gap: 40px;
  margin-bottom: 40px;
  background: white;
  padding: 32px;
  border-radius: 12px;
  box-shadow: var(--card-shadow);
}

.detail-cover {
  width: 100%;
  height: 320px;
  object-fit: cover;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
  &:hover {
    transform: scale(1.02);
  }
}

.detail-info {
  h1 {
    font-size: 28px;
    color: var(--text-dark);
    margin-bottom: 16px;
  }
  .author {
    font-size: 16px;
    color: var(--text-light);
    margin-bottom: 12px;
  }
  .rating {
    color: #f39c12;
    font-weight: 500;
    margin-bottom: 12px;
  }
  .desc {
    line-height: 1.8;
    color: var(--text-dark);
  }
}

.detail-header {
  margin-bottom: 32px;
}

.detail-title {
  font-size: 32px;
  color: #2c3e50;
  line-height: 1.3;
  margin-bottom: 16px;
}
</style>
