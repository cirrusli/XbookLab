<template>
  <header>
    <div class="nav-container">
      <ul class="nav-links">
        <li
          v-for="(item, index) in menuItems"
          :key="index"
          :class="{ active: item.path == curPath }"
        >
          <router-link :to="item.path">{{ item.name }}</router-link>
        </li>
      </ul>
      <div class="user-section">
        <button v-if="!isLogin" class="login-btn" @click="goTo('/login')">
          登录
        </button>
        <template v-else>
          <div class="user-profile flex aic">
            <img :src="userInfo.avatar" alt="用户头像" class="user-avatar" />
            <span class="username">{{ userInfo.name }}</span>
          </div>
          <ul class="user-dropdown">
            <li><router-link to="/user">个人主页</router-link></li>
            <li><router-link to="/user/account">后台管理</router-link></li>
            <li>
              <a class="logout-btn" @click="handleLogout">退出登录</a>
            </li>
          </ul>
        </template>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { Modal } from 'ant-design-vue';
import { LogoutApi } from '@/api/user.js';

const route = useRoute();
const router = useRouter();

const isLogin = ref(localStorage.getItem('token') ? true : false);
// 用户信息，直接存储在本地了
let userInfo = ref({});
if (isLogin.value) {
  userInfo.value = JSON.parse(localStorage.getItem('userInfo'));
}

const curPath = ref(route.path);
const menuItems = [
  { name: '推荐', path: '/index' },
  { name: '书友圈', path: '/groups' },
  { name: '话题热榜', path: '/hot' },
  { name: '好友在读', path: '/friends' },
  // { name: '书友匹配', path: '#', id: 'match-trigger' },
];

const goTo = (url) => {
  router.push({ path: url });
};

watch(
  () => route.path,
  (val) => {
    curPath.value = val;
  }
);

const handleLogout = () => {
  Modal.confirm({
    centered: true,
    content: `您确定要退出登录吗？`,
    okText: '确认',
    onOk() {
      LogoutApi()
        .then((result) => {
          localStorage.removeItem('token');
          localStorage.removeItem('userInfo');
          isLogin.value = false;
          userInfo.value = {};
          if (route.meta.needAuth) {
            router.push('/');
          }
          window.location.reload();
        })
        .catch((err) => {
          console.error(err);
        });
    },
    cancelText: '取消',
    onCancel() {},
  });
};
</script>

<style lang="scss" scoped>
header {
  background-color: var(--primary-light);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  position: sticky;
  top: 0;
  z-index: 100;
  padding: 0 5%;
}

.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
}

.nav-links {
  display: flex;
  list-style: none;
  li {
    margin-right: 24px;
    position: relative;
    padding: 20px 0;
  }
  a {
    font-weight: 500;
    color: var(--text-dark);
    transition: var(--transition);
    &:hover {
      color: var(--primary-color);
    }
  }
  .active {
    color: var(--primary-color);
    &::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 3px;
      background-color: var(--primary-color);
      border-radius: 3px 3px 0 0;
    }
  }
}

.user-section {
  display: flex;
  align-items: center;
  height: 100%;
  &:hover {
    .user-dropdown {
      display: block;
    }
  }
}

.login-btn {
  background-color: var(--primary-color);
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
  transition: var(--transition);
  &:hover {
    background-color: #3aa876;
    transform: translateY(-1px);
  }
}

.user-profile {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  margin-right: 8px;
  object-fit: cover;
}

.username {
  font-weight: 500;
}

.user-dropdown {
  display: none;
  position: absolute;
  top: 60px;
  right: 5%;
  background-color: white;
  box-shadow: var(--shadow);
  border-radius: 8px;
  padding: 16px;
  z-index: 1000;
  text-align: center;
  li {
    padding: 8px 16px;
    cursor: pointer;
    &:hover {
      background-color: var(--primary-light);
    }
    a {
      color: var(--text-light);
      font-size: 14px;
    }
  }
}
</style>
