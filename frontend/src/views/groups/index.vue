<template>
  <aside class="sidebar">
    <div class="sidebar-section">
      <h3 class="sidebar-title">圈子分类</h3>
      <div class="tag-cloud">
        <span
          v-for="(item, index) in menus"
          :key="index"
          class="tag"
          :class="{ active: curMenuId === item.id }"
          @click="handleChangeCategory(item.id, item.tagVal)"
          >{{ item.tag }}</span
        >
      </div>
    </div>
  </aside>

  <div class="content">
    <section class="content-section">
      <h2 class="section-title">热门书友圈</h2>
      <div class="groups-container" id="groups-container">
        <div v-for="(item, index) in groupList" :key="index" class="group-card">
          <img :src="item.cover" :alt="item.title" class="group-cover" />
          <div class="group-info">
            <h3 class="group-title">{{ item.title }}</h3>
            <p class="group-desc">{{ item.desc }}</p>
            <div class="group-meta">
              <span>{{ item.members }}成员 · {{ item.topics }}话题</span>
              <button class="join-btn" @click="handleJoinGroup(item.id)">加入</button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GetGroupListApi } from '@/api/group.js';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';

const menus = ref([
  { id: 0, tag: '全部', tagVal: '' },
  { id: 1, tag: '文学', tagVal: 'literature' },
  { id: 2, tag: '科技', tagVal: 'technology' },
  { id: 3, tag: '历史', tagVal: 'history' },
  { id: 4, tag: '心理', tagVal: 'psychology' },
  { id: 5, tag: '艺术', tagVal: 'art' },
  { id: 6, tag: '商业', tagVal: 'business' },
  { id: 7, tag: '哲学', tagVal: 'philosophy' },
]);
const curMenuId = ref(0);
const groupList = ref([]);
const handleGetGroups = async (tagVal) => {
  let params = {
    tag: tagVal || '',
  };
  const { data } = await GetGroupListApi(params);
  groupList.value = data;
};

const handleJoinGroup = async (groupId) => {
  try {
    // 这里需要调用加入API
    // const { data } = await JoinGroupApi({ groupId });
    CreateSuccessMessage('加入成功');
  } catch (error) {
    console.error(error);
    CreateErrorMessage('加入失败');
  }
};

const handleChangeCategory = (id, tagVal) => {
  curMenuId.value = id;
  handleGetGroups(tagVal);
};
handleGetGroups();
</script>

<style lang="scss" scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  gap: 12px;
  background: white;
  border-radius: 8px;
  padding: 16px;
  box-shadow: var(--card-shadow);
  width: 200px;
  height: fit-content;
  flex-shrink: 0;
}

/* 侧边栏标题可点击样式 */
.sidebar-title {
  cursor: pointer;
  user-select: none;
  position: relative;
  padding-right: 20px;
  &::after {
    content: '▼';
    position: absolute;
    right: 0;
    top: 50%;
    transform: translateY(-50%);
    font-size: 10px;
    transition: transform 0.3s ease;
  }
  &.collapsed::after {
    transform: translateY(-50%) rotate(-90deg);
  }
}

.tag-cloud {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.tag {
  display: block;
  padding: 10px 16px;
  margin: 6px 0;
  background-color: white;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  font-size: 14px;
  color: var(--text-dark);
  cursor: pointer;
  transition: var(--transition);
  width: 100%;
  position: relative;
  overflow: hidden;
  box-sizing: border-box;
  &::after {
    content: '';
    position: absolute;
    left: 0;
    top: 0;
    width: 4px;
    height: 100%;
    background: var(--primary-light);
    transition: var(--transition);
  }
  &:hover {
    border-color: var(--primary-color);
    color: var(--primary-color);
    transform: translateX(4px);
    box-shadow: 0 4px 12px rgba(66, 185, 131, 0.15);
    &::after {
      width: 6px;
      background: var(--primary-color);
    }
  }
  &.active {
    background-color: var(--primary-color);
    color: white;
    border-color: var(--primary-color);
    transform: translateX(4px);
    &::after {
      width: 6px;
      background: white;
    }
  }
}

.content {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.content-section {
  margin-bottom: 40px;
  background: white;
  padding: 32px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-dark);
  margin-bottom: 20px;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--border-color);
}

.groups-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
  margin-top: 24px;
}

.group-card {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: var(--card-shadow);
  transition: var(--transition);
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.group-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(66, 185, 131, 0.15);
}

.group-cover {
  height: 160px;
  width: 100%;
  object-fit: cover;
  border-bottom: 1px solid var(--border-color);
}

.group-info {
  padding: 16px;
}

.group-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--text-dark);
}

.group-desc {
  color: var(--text-light);
  font-size: 14px;
  margin-bottom: 12px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.group-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: var(--text-light);
  font-size: 13px;
}

.join-btn {
  background-color: var(--primary-color);
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  font-weight: 500;
  cursor: pointer;
  transition: var(--transition);
  font-size: 13px;
}

.join-btn:hover {
  background-color: #3aa876;
}
</style>
