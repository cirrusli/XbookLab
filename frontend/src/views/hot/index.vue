<template>
  <aside class="sidebar">
    <div class="sidebar-section">
      <h3 class="sidebar-title">榜单分类</h3>
      <div class="tag-cloud">
        <span
          class="tag"
          v-for="(item, index) in dateRange"
          :key="index"
          :class="{ active: item.value === curDateRange }"
          @click="handleSwitchDate(item.value)"
          >{{ item.name }}</span
        >
      </div>
    </div>
  </aside>

  <div class="content">
    <section class="content-section">
      <div class="flex aic jcb">
        <h2 class="section-title">话题热榜</h2>
        <span class="btn-create" @click="showCreateTopic">发起话题</span>
      </div>

      <div class="hot-list" id="hot-list">
        <div
          class="hot-item"
          v-for="(item, index) in hotTopicList"
          :key="index"
          @click="goToTopicDetail(item.id)"
        >
          <div class="hot-rank">{{ item.rank }}</div>
          <div class="hot-content">
            <h3 class="hot-title">{{ item.title }}</h3>
            <div class="hot-meta">
              <span>{{ item.author }}</span>
              <span>{{ item.views }}浏览</span>
              <span>{{ item.likes }}赞</span>
              <span>{{ item.comments }}评论</span>
              <span>{{ item.tag }}</span>
              <span>{{ item.date }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>

  <a-modal
    v-model:open="modalVisible"
    :closable="false"
    @ok="handleSubmit"
    cancelText="取消"
    okText="确认"
  >
    <a-form :model="formData" ref="formRef" :rules="rules">
      <a-form-item label="话题标题" name="title">
        <a-input v-model:value="formData.title" />
      </a-form-item>
      <a-form-item label="话题内容" name="content">
        <a-textarea v-model:value="formData.content" />
      </a-form-item>
      <a-form-item label="话题标签" name="tag">
        <a-select
          ref="select"
          v-model:value="formData.tag"
          style="width: 120px"
          @focus="focus"
          @change="handleChange"
        >
          <a-select-option
            v-for="(item, index) in tagList"
            :key="index"
            :value="item.tag_name"
            >{{ item.tag_name }}</a-select-option
          >
        </a-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { GetHotTopicListApi, CreateTopicApi, GetIndexHotTopicApi } from '@/api/hot.js';
import { GetManageTagListApi } from '@/api/tag.js';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const curDateRange = ref('today');
const dateRange = ref([
  {
    id: '1',
    name: '今日',
    value: 'today',
  },
  {
    id: '2',
    name: '一周内',
    value: 'week',
  },
  {
    id: '3',
    name: '一个月内',
    value: 'month',
  },
]);

const hotTopicList = ref([]);
const handleGetHotTopic = async () => {
  let params = {
    range: curDateRange.value,
  };
  // const { data } = await GetHotTopicListApi(params);
  const { data } = await GetIndexHotTopicApi(params);
  hotTopicList.value = data;
};

handleGetHotTopic();

const handleSwitchDate = (value) => {
  curDateRange.value = value;
  handleGetHotTopic();
};

const modalVisible = ref(false);
const formRef = ref({});
const formData = reactive({
  title: '',
  content: '',
  tag: '',
});

const showCreateTopic = () => {
  if (!localStorage.getItem('token')) {
    CreateErrorMessage('请先登录');
    return false;
  }
  modalVisible.value = true;
};

let tagList = ref([]);
const handleGetTagList = async () => {
  const { data } = await GetManageTagListApi();
  tagList.value = data;
};
handleGetTagList();

// 跳转话题详情
const goToTopicDetail = (topicId) => {
  router.push({ path: `/topic/${topicId}` });
};

const rules = {
  title: [
    { required: true, message: '请输入标题', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' },
  ],
  content: [
    { required: true, message: '请输入话题内容', trigger: 'blur' },
    { min: 10, max: 200, message: '长度在 10 到 200 个字符', trigger: 'blur' },
  ],
  tag: [
    { required: true, message: '请选择话题标签', trigger: 'change' },
    { type: 'string', message: '请选择话题标签', trigger: 'change' },
  ],
};

const handleSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      CreateTopicApi(formData)
        .then((res) => {
          if (res.code === 200) {
            CreateSuccessMessage('创建成功');
            modalVisible.value = false;
          } else {
            CreateErrorMessage(res.msg);
          }
        })
        .catch((error) => {
          console.log('error', error);
        });
    })
    .catch((error) => {
      console.log('error', error);
    });
};
</script>

<style lang="scss" scoped>
/* 新增分类图标 */
.sidebar-title::after {
  content: '▼';
  font-size: 12px;
  transition: transform 0.3s ease;
}

.sidebar-title.collapsed::after {
  transform: rotate(-90deg);
}

.hot-list {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: var(--card-shadow);
}

.hot-item {
  display: flex;
  padding: 16px;
  border-bottom: 1px solid var(--border-color);
  transition: var(--transition);
  cursor: pointer;
}

.hot-item:hover {
  background: var(--primary-light);
}

.hot-rank {
  font-size: 18px;
  font-weight: bold;
  color: var(--primary-color);
  width: 40px;
  text-align: center;
  margin-right: 16px;
}

.hot-content {
  flex: 1;
}

.hot-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 8px;
  margin-top: 0;
  color: var(--text-dark);
}

.hot-meta {
  display: flex;
  gap: 12px;
  color: var(--text-light);
  font-size: 13px;
}

.hot-meta span {
  display: flex;
  align-items: center;
}

.hot-meta span::before {
  content: '·';
  margin-right: 4px;
}

.hot-meta span:first-child::before {
  content: '';
  margin-right: 0;
}

.btn-create {
  cursor: pointer;
  transition: all 300ms ease-in-out;
  &:hover {
    color: var(--primary-color);
  }
}
</style>
