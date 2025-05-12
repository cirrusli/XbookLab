<template>
  <Header />
  <main class="profile-container">
    <div class="profile-header">
      <img
        :src="userInfo.avatar"
        data-src=" https://randomuser.me/api/portraits/men/33.jpg"
        alt="用户头像"
        class="profile-avatar"
      />
      <div class="profile-info">
        <h1 class="profile-name">{{ userInfo.name }}</h1>
        <p class="profile-bio">
          热爱阅读与分享，专注文学与哲学。通过阅读探索世界，在书籍中寻找智慧与共鸣。
        </p>
        <div class="profile-stats">
          <div class="profile-stat" @click="goToFollow">
            <strong>{{ userInfo.followCount }}</strong> 关注
          </div>
          <div class="profile-stat" @click="goToFollowMe">
            <strong>{{ userInfo.followMeCount }}</strong> 粉丝
          </div>
          <div class="profile-stat">
            <strong>{{ userInfo.topicCount }}</strong> 话题
          </div>
          <div class="profile-stat">
            <strong>{{ userInfo.bookCount }}</strong> 阅读
          </div>
        </div>
      </div>
    </div>

    <router-view></router-view>
  </main>
</template>

<script setup>
import { ref } from 'vue';
import Header from '@/components/Header/index.vue';
import { GetUserProfileApi } from '@/api/user.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const userInfo = ref({});
const handleGetProfile = async () => {
  const { data } = await GetUserProfileApi();
  userInfo.value = data;
  userInfo.value.name = JSON.parse(localStorage.getItem('userInfo')).name;
  userInfo.value.avatar = JSON.parse(localStorage.getItem('userInfo')).avatar;
};
handleGetProfile();

const goToUrl = (url) => {
  router.push({ path: url });
};

const goToFollow = () => {
  if (userInfo.value.followCount !== 0) {
    goToUrl('/user/follow');
  }
};

const goToFollowMe = () => {
  if (userInfo.value.followMeCount !== 0) {
    goToUrl('/user/followMe');
  }
};
</script>

<style lang="scss" scoped>
.profile-container {
  max-width: 1200px;
  margin: 2rem auto;
  padding: 0 2rem;
}

.profile-header {
  display: flex;
  align-items: center;
  margin-bottom: 3rem;
  padding: 2rem;
  background: var(--profile-card-bg);
  border-radius: 12px;
  box-shadow: var(--profile-shadow);
}

.profile-avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 2rem;
  border: 4px solid var(--profile-primary);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.profile-info {
  flex: 1;
}

.profile-name {
  font-size: 2rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: var(--profile-text);
}

.profile-bio {
  color: var(--profile-text-light);
  margin-bottom: 1.5rem;
  font-size: 1rem;
  line-height: 1.6;
  max-width: 80%;
}

.profile-stats {
  display: flex;
  gap: 2rem;
}

.profile-stat {
  cursor: pointer;
  font-size: 0.95rem;
  transition: all 0.3s ease;
  &:hover {
    color: var(--profile-primary);
  }
}

.profile-stat strong {
  font-size: 1.2rem;
  font-weight: 600;
  margin-right: 0.3rem;
  color: var(--profile-primary);
}
</style>
