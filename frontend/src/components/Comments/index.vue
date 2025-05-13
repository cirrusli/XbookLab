<template>
  <div class="comment-container">
    <h3 class="section-title">读者讨论</h3>
    <section class="comment-section">
      <div class="comment-list" id="comment-list">
        <div
          class="comment-card"
          v-for="(item, index) in commentList"
          :key="index"
        >
          <div class="comment-header">
            <img src="@/assets/icon-avatar.png" class="comment-avatar" />
            <span class="comment-author">{{ item.author }}</span>
            <span class="comment-time">{{ item.time }}</span>
          </div>
          <div class="comment-content">{{ item.content }}</div>
        </div>
      </div>
      <div class="comment-editor">
        <textarea
          v-model="commentContent"
          class="comment-input"
          placeholder="写下你的思考和评论..."
        ></textarea>
        <div class="editor-actions">
          <button class="submit-btn" @click="handleSubmitComment">
            发表评论
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetBookCommentsApi, RecordBookCommentApi } from '@/api/recommend.js';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';

const props = defineProps({
  commentType: { type: Number }, //评论类型，0是书籍，1是话题
  bookId: { type: Number }, //书籍id
  topicId: { type: Number }, //话题id
});

const commentContent = ref('');

const commentList = ref([]); // 评论列表
const handleGetCommentList = async () => {
  const { data } = await GetBookCommentsApi();
  commentList.value = data;
};
handleGetCommentList();

const handleSubmitComment = () => {
  if (!localStorage.getItem('token')) {
    CreateErrorMessage('请先登录');
    return false;
  }
  let params = {
    content: commentContent.value,
    comment_type: props.commentType,
  };

  if (props.commentType === 0) {
    params.comment_id = props.bookId;
  } else if (props.commentType === 1) {
    params.comment_id = props.topicId;
  }
  if (!commentContent.value.trim()) {
    CreateErrorMessage('评论内容不能为空');

    return false;
  }
  RecordBookCommentApi(params)
    .then((result) => {
      CreateSuccessMessage('评论成功');
      commentContent.value = '';
      window.location.reload();
    })
    .catch((err) => {
      console.error(err);
    });
};
</script>

<style lang="scss" scoped>
.comment-container {
  width: 100%;
  margin-top: 40px;
  min-height: 400px;
  padding: 0;
  background: none;
  border: none;
  box-shadow: none;
}

/* 统一评论区样式 */
.comment-section {
  border-radius: 12px;
  box-shadow: var(--card-shadow);
  background-color: #ffffff;
  .comment-card {
    padding: 24px;
    border-bottom: 1px dashed var(--border-color);
    box-shadow: none;
  }
  .comment-header {
    display: flex;
    align-items: center;
    // &:hover {
    //   transform: translateY(-2px);
    //   box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    // }
  }
  .comment-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    object-fit: cover;
    margin-right: 12px;
    border: 2px solid rgba(66, 185, 131, 0.1);
    transition: all 0.3s ease;
    &:hover {
      transform: scale(1.05);
    }
  }

  .comment-content {
    padding-left: 56px;
    font-size: 14px;
    line-height: 1.8;
  }
}

.comment-list {
  display: grid;
  gap: 10px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-dark);
  margin-bottom: 20px;
  padding-bottom: 8px;
}

.comment-time {
  color: #64748b;
  font-size: 12px;
  margin-left: 10px;
}

.comment-editor {
  padding: 24px;
}

.comment-input {
  width: 100%;
  height: 200px;
  padding: 16px;
  border: 1px solid #ecf0f1;
  border-radius: 8px;
  resize: vertical;
  font-size: 15px;
  line-height: 1.6;
  resize: none;
  transition: border-color 0.3s ease;
  box-sizing: border-box;
  &:focus {
    outline: none;
    border-color: var(--primary-color);
    box-shadow: 0 0 0 3px rgba(66, 185, 131, 0.2);
  }
}

.editor-actions {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}

.submit-btn {
  background: #3498db;
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.3s ease;
  &:hover {
    background: #2980b9;
  }
}
</style>
