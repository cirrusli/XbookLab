<template>
  <Header />
  <main class="detail-container">
    <article class="topic-detail">
      <h1>{{ topicInfo.title }}</h1>
      <div class="meta">
        <span>{{ topicInfo.tag }}</span>
        <span>{{ topicInfo.likes }}赞</span>
        <span
          class="topic-likes"
          :class="{ active: topicInfo.isLiked === 1 }"
          @click="handleLikeTopic(topicInfo.isLiked)"
        ></span>
      </div>
      <div class="content">{{ topicInfo.desc }}</div>
    </article>
    <!-- 评论区 -->
    <Comment :commentType="1" :topicId="state.topicId" />
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue';
import Header from '@/components/Header/index.vue';
import { useRoute } from 'vue-router';
import { GetTopicDetailApi, LikeTopicApi } from '@/api/hot';
import Comment from '@/components/Comments/index.vue';
import { CreateSuccessMessage, CreateErrorMessage } from '../../utils/alert';

const route = useRoute();

const topicInfo = ref({}); // 话题信息
const state = reactive({
  topicId: null,
});
const handleGetTopicDetail = async () => {
  state.topicId = route.params.id; // 获取话题ID
  const { data } = await GetTopicDetailApi(state.topicId);
  topicInfo.value = data;
};

handleGetTopicDetail();

const handleLikeTopic = (curStatus) => {
  if (!localStorage.getItem('token')) {
    CreateErrorMessage('请先登录');
    return false;
  }
  let newStatus = curStatus === 1 ? 2 : 1;
  LikeTopicApi({ is_liked: newStatus })
    .then((result) => {
      CreateSuccessMessage(result.message);
      topicInfo.value.isLiked = newStatus;
    })
    .catch((err) => {
      console.error(err);
    });
};
</script>

<style lang="scss" scoped>
.detail-container {
  max-width: 1200px;
  margin: 40px auto;
  padding: 0 20px;
}

.topic-detail {
  background: white;
  padding: 32px;
  border-radius: 12px;
  box-shadow: var(--card-shadow);
  margin-bottom: 40px;
}

.topic-detail h1 {
  font-size: 28px;
  color: var(--text-dark);
  margin-bottom: 16px;
}

.topic-detail .meta {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  color: var(--text-light);
}

.topic-detail .content {
  line-height: 1.8;
  color: var(--text-dark);
  font-size: 16px;
}

/* 交互效果 */
.topic-detail {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.topic-detail:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
}

.topic-likes {
  display: flex;
  align-items: center;
  margin-right: 12px;
  cursor: pointer;
  &.active {
    &::before {
      content: '♥';
      color: #ff4757;
      margin-right: 4px;
    }
  }
  &::before {
    content: '♥';
    color: #ccc;
    margin-right: 4px;
  }
}
</style>
