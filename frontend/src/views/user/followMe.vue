<template>
  <div class="profile-content section-box">
    <h3>关注我的人</h3>
    <div class="following-list">
      <div class="following-item" v-for="(item, index) in list" :key="index">
        <img :src="item.avatar" alt="用户头像" class="following-avatar" />
        <span>{{ item.name }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetFollowMeListApi } from '@/api/user.js';

let list = ref([]);
const handleGetList = async () => {
  const { data } = await GetFollowMeListApi();
  list.value = data.list;
};
handleGetList();
</script>

<style lang="scss" scoped>
/* Following List */
.following-list {
  display: flex;
  flex-wrap: wrap;
}

.following-item {
  width: 25%;
  display: flex;
  align-items: center;
  padding: 0.8rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  cursor: pointer;
  box-sizing: border-box;
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
</style>
